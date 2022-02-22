package wharfapi

import (
	"context"
	"errors"
	"math"

	v5 "github.com/iver-wharf/wharf-api-client-go/v2/api/wharfapi/v5"
	"github.com/iver-wharf/wharf-api-client-go/v2/pkg/model/request"
	"github.com/iver-wharf/wharf-api-client-go/v2/pkg/model/response"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// CreateBuildLogStream contains methods for sending log creation requests in
// a streamed fashion.
type CreateBuildLogStream interface {
	Send([]request.Log) error
	Close() error
	CloseAndRecv() (response.CreatedLogsSummary, error)
}

type createBuildLogStream struct {
	stream v5.Builds_CreateLogStreamClient
}

func (s createBuildLogStream) Send(logs []request.Log) error {
	grpcLogLines := make([]*v5.NewLogLine, len(logs))
	for i, log := range logs {
		grpcLogLines[i] = &v5.NewLogLine{
			BuildId:   uint64(log.BuildID),
			Message:   log.Message,
			Timestamp: timestamppb.New(log.Timestamp),
		}
	}
	return s.stream.Send(&v5.CreateLogStreamRequest{
		Lines: grpcLogLines,
	})
}

func (s createBuildLogStream) Close() error {
	return s.stream.CloseSend()
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
func (c *Client) CreateBuildLogStream(ctx context.Context) (CreateBuildLogStream, error) {
	conn, err := grpc.Dial(c.APIURL)
	if err != nil {
		return nil, err
	}
	builds := v5.NewBuildsClient(conn)
	stream, err := builds.CreateLogStream(ctx)
	if err != nil {
		return nil, err
	}
	return createBuildLogStream{stream}, nil
}
