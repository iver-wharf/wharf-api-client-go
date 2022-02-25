package wharfapi

import (
	"context"
	"crypto/x509"
	"errors"
	"fmt"
	"math"
	"strings"

	v5 "github.com/iver-wharf/wharf-api-client-go/v2/api/wharfapi/v5"
	"github.com/iver-wharf/wharf-api-client-go/v2/pkg/model/request"
	"github.com/iver-wharf/wharf-api-client-go/v2/pkg/model/response"
	"golang.org/x/oauth2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/credentials/oauth"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// CreateBuildLogStream contains methods for sending log creation requests in
// a streamed fashion.
type CreateBuildLogStream interface {
	Send([]request.Log) error
	CloseAndRecv() (response.CreatedLogsSummary, error)
}

type createBuildLogStream struct {
	stream v5.Builds_CreateLogStreamClient
}

func (s createBuildLogStream) Send(logs []request.Log) error {
	grpcLogLines := make([]*v5.NewLogLine, len(logs))
	for i, log := range logs {
		grpcLogLines[i] = &v5.NewLogLine{
			BuildID:   uint64(log.BuildID),
			Message:   log.Message,
			Timestamp: timestamppb.New(log.Timestamp),
		}
	}
	return s.stream.Send(&v5.CreateLogStreamRequest{
		Lines: grpcLogLines,
	})
}

func (s createBuildLogStream) CloseAndRecv() (response.CreatedLogsSummary, error) {
	res, err := s.stream.CloseAndRecv()
	if err != nil {
		return response.CreatedLogsSummary{}, err
	}
	if res.LinesInserted > math.MaxUint {
		return response.CreatedLogsSummary{}, errors.New("inserted logs count is bigger than maximum uint size")
	}
	return response.CreatedLogsSummary{
		LogsInserted: uint(res.LinesInserted),
	}, nil
}

// CreateBuildLogStream creates a log creation stream used to sending log
// creation requests in a streamed fashion by reusing the same TCP connection
// for higher throughput during log injection.
//
// Added in wharf-api v5.1.0.
func (c *Client) CreateBuildLogStream(ctx context.Context) (CreateBuildLogStream, error) {
	conn, err := c.grpcDial()
	if err != nil {
		return nil, fmt.Errorf("dial grpc: %w", err)
	}
	builds := v5.NewBuildsClient(conn)
	stream, err := builds.CreateLogStream(ctx)
	if err != nil {
		return nil, fmt.Errorf("open log creation stream: %w", err)
	}
	return createBuildLogStream{stream}, nil
}

func (c *Client) grpcDial() (*grpc.ClientConn, error) {
	transportCred, err := c.grpcTransportCred()
	if err != nil {
		return nil, fmt.Errorf("get transport credentials: %w", err)
	}
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(transportCred),
	}
	if c.AuthHeader != "" {
		typ, token, ok := cutString(c.AuthHeader, ' ')
		if !ok {
			return nil, errors.New("invalid auth header format, expected 'Bearer abc123'")
		}
		perRPC := oauth.NewOauthAccess(&oauth2.Token{
			TokenType:   typ,
			AccessToken: token,
		})
		opts = append(opts, grpc.WithPerRPCCredentials(perRPC))
	}
	return grpc.Dial(trimProtocol(c.APIURL), opts...)
}

func (c *Client) grpcTransportCred() (credentials.TransportCredentials, error) {
	if !strings.HasPrefix(c.APIURL, "https://") {
		return insecure.NewCredentials(), nil
	}
	certPool, err := x509.SystemCertPool()
	if err != nil {
		return nil, fmt.Errorf("load system cert pool: %w", err)
	}
	return credentials.NewClientTLSFromCert(certPool, ""), nil
}
