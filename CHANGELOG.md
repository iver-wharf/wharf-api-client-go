# Wharf API Go client changelog

This project tries to follow [SemVer 2.0.0](https://semver.org/).

<!--
	When composing new changes to this list, try to follow convention.

	The WIP release shall be updated just before adding the Git tag.
	From (WIP) to (YYYY-MM-DD), ex: (2021-02-09) for 9th of February, 2021

	A good source on conventions can be found here:
	https://changelog.md/
-->

## v2.2.1 (WIP)

- Fixed gRPC logs streaming failing if no port is specified in the `APIURL`
  field of `wharfapi.Client`. Now defaults to `80` for `HTTP` URLs, and `443`
  for `HTTPS`. (#48)

## v2.2.0 (2022-05-02)

- Added methods for execution engine: (#45)

  - `Client.GetEngineList() EngineList`: `GET /api/engine`

- Added query parameter for specifying engine in new builds: (#45)

  - `ProjectStartBuild.Engine`: `POST /api/project/{projectId}/build?engine=...`

## v2.1.0 (2021-03-08)

- Added `Client.CreateBuildLogStream` that uses wharf-api's gRPC API that was
  introduced in wharf-api v5.1.0 to allow streaming log creation. (#35)

- Added numerous dependencies:

  - `github.com/alta/protopatch` v0.5.0 (#35)
  - `github.com/blang/semver/v4` v4.0.0. (#39)
  - `golang.org/x/oauth2` v0.0.0-20200107190931-bf48bf16ab8d (#35)
  - `google.golang.org/grpc` v1.44.0 (#35)
  - `google.golang.org/protobuf` v1.27.1 (#35)

- Changed minimum version of Go from 1.16 to 1.17. (#36)

- Changed `BuildSearch.Status` to a slice to support searching for builds of any
  matching status. Support was added to wharf-api in v5.1.0 with [wharf-api#150](https://github.com/iver-wharf/wharf-api/pull/150).
  (#37)

- Added `WorkerID` to `response.Build` model, as added in
  [wharf-api#156](https://github.com/iver-wharf/wharf-api/pull/156). (#38)

- Added methods for API health and metadata: (#39)

  - `Client.GetHealth() HealthStatus`: `GET /api/health`
  - `Client.GetVersion() app.Version`: `GET /api/version`
  - `Client.Ping() Ping`: `GET /api/ping`

- Added methods for project overrides: (#40)

  - `Client.GetProjectOverrides(uint) ProjectOverrides`:
    `GET /api/project/{projectId}/override`

  - `Client.UpdateProjectOverrides(uint, ProjectOverridesUpdate) ProjectOverrides`:
    `PUT /api/project/{projectId}/override`

  - `Client.DeleteProjectOverrides(uint)`:
    `DELETE /api/project/{projectId}/override`

- Added method to delete a project (USE WITH CAUTION!): (#40)

  - `Client.DeleteProject(uint)`: `DELETE /api/project/{projectId}`

- Added methods for uploading artifacts and test results: (#40)

  - `Client.CreateBuildArtifact(uint, string, io.Reader)`:
    `POST /api/build/{buildId}/artifact`

  - `Client.CreateBuildTestResult(uint, string, io.Reader) []ArtifactMetadata`:
    `POST /api/build/{buildId}/test-result`

- Added client and server version validation on all endpoints. This is disabled
  by default, but can be enabled with the following new flags: (#39)

  ```go
  client := wharfapi.Client{
      ErrIfOutdatedClient: true,
      ErrIfOutdatedServer: true,
	}
  ```

- Changed `Client` method receiver to a pointer on all methods. This will not
  break regular usage, but may break edge case usages during compilation. (#39)

- Changed all endpoints that sends JSON payload in the request to stream the
  writing instead of buffering it. This will lower the memory footprint for
  larger payloads. (#40)

- Fixed `GetBuildArtifact` having invalid implementation where it expected an
  artifact JSON response, while in reality that's the endpoint to download an
  artifact. This will break compilation of existing code, but the endpoint
  never worked before to begin with. (#40)

## v2.0.0 (2022-01-25)

- BREAKING: Changed module path from `github.com/iver-wharf/wharf-api-client-go`
  to `github.com/iver-wharf/wharf-api-client-go/v2`. (#30)

- BREAKING: Removed support for wharf-api v4.2.0 and below. (#29)

- Changed minimum version of Go from 1.13 to 1.16. (#29)

- Added support for `github.com/iver-wharf/wharf-api` v5.0.0. (#29)

- Changed version of `github.com/iver-wharf/wharf-core` from v1.2.0 to v1.3.0.
  (#29)

- Added dependency `github.com/google/go-querystring` v1.1.0. (#29)

## v1.4.0 (2021-09-07)

- Added methods for search endpoints: (#18)

  - `Client.SearchProject(Project) []Project`: `POST /api/projects/search`
  - `Client.SearchProvider(Provider) []Provider`: `POST /api/providers/search`
  - `Client.SearchToken(Token) []Token`: `POST /api/tokens/search`

## v1.3.1 (2021-08-30)

- Added documentation to all exported methods. (#11)

- Added dependency on `github.com/iver-wharf/wharf-core` v1.1.0.
  (#12, #13, #16)

- Changed all logging via `fmt.Print` and sirupsen/logrus to instead use the new
  `github.com/iver-wharf/wharf-core/pkg/logger`. (#13)

- Changed to use `problem.Response` from `wharf-core` instead of the
  `wharfapi.Problem` that was added in v1.3.0/#4. (#12, #14)

- Removed `wharfapi.ProblemError` as the `problem.Response` type from the
  `wharf-core` module conforms to the `error` interface. (#14)

- Removed dependency on `github.com/sirupsen/logrus`. (#13)

## v1.3.0 (2021-06-10)

- Fixed implementation of the refactor for ApiUrl to APIURL. (#8)

- Added endpoint `PUT /api/provider`. (#8)

- Added endpoint `PUT /api/token`. (#8)

- Added problem response handling so better and more informative errors are
  returned as `wharfapi.ProblemError` instead. (#4)

- Fixed "non-2xx status code" checks. (#4)

- Fixed invalid initialisms in method and field names. Affected names: (#5)

  | Before                       | After                        |
  | ------                       | -----                        |
  | `Client.ApiUrl`              | `Client.APIURL`              |
  | `Client.GetProjectById`      | `Client.GetProjectByID`      |
  | `Client.GetProviderById`     | `Client.GetProviderByID`     |
  | `Client.GetTokenById`        | `Client.GetTokenByID`        |
  | `Project.AvatarUrl`          | `Project.AvatarURL`          |
  | `ProjectRunResponse.BuildId` | `ProjectRunResponse.BuildID` |

- Remove ProviderID from Token to match the corresponding change in the
  DB datastructure. (#6)

## v1.2.0 (2021-02-25)

- Added CHANGELOG.md to repository. (!10)

- Added support for endpoint `PUT /api/build/{buildid}?status={status}`
  (via function `PutStatus`). (!11)

- Added support for endpoint `POST /api/build/{buildid}/log`
  (via function `PostLog`). (!11)

## v1.1.0 (2021-01-07)

- Fixed all endpoint functions not checking status code on HTTP responses. (!9)

## v1.0.0 (2020-12-02)

- Changed package name from `client` to `wharf` and changed the type name
  `WharfClient` to just `Client`, resulting in usage changing from
  `client.WharfClient` to `wharf.Client`. (!7, !8)

## v0.1.5 (2020-11-24)

- Removed group endpoints support, a reflection of the changes from the API
  v1.0.0. (!6)

  - `GET /api/group` (via `GetGroupById`)
  - `GET /api/groups/search` (via `GetGroup`)
  - `PUT /api/group` (via `PutGroup`)

## v0.1.4 (2020-10-23)

- Added support for endpoint `PUT /api/branches` (via function `PutBranches`)
  that was introduced in the API v1.0.0. (!5)

- Added `README.md` to document that this package is only meant to be used to
  talk to the main API and not to the provider/plugin APIs just to avoid
  circular dependencies. (!3)

## v0.1.3 (2020-04-30)

- Added support for endpoint `POST /api/project/{projectid}/{stage}/run`
  (via `PostProjectRun` function). (!1)

## v0.1.2 (2019-11-21)

- Fixed `/api/providers/search` function (`GetProvider`) to only filter by
  token ID `if tokenID > 0`. (a57a10d2)

## v0.1.1 (2019-11-21)

- Changed module name and moved repository from `/tools/wharf-project/client`
  to `/tools/wharf-client` due to restrictions in how Go handles Git paths.
  (755a2fc4)

## v0.1.0 (2019-11-21)

- Added initial commit. Endpoints supported: (df25308e)

  - `POST /api/branch` (via `PutBranch`)
  - `GET /api/group/{groupId}` (via `GetGroupById`)
  - `GET /api/groups/search` (via `GetGroup`)
  - `PUT /api/group` (via `PutGroup`)
  - `GET /api/project/{projectId}` (via `GetProjectById`)
  - `PUT /api/project` (via `PutProject`)
  - `GET /api/provider/{providerId}` (via `GetProviderById`)
  - `POST /api/providers/search` (via `GetProvider`)
  - `POST /api/provider` (via `PostProvider`)
  - `GET /api/token/{tokenId}` (via `GetTokenById`)
  - `POST /api/tokens/search` (via `GetToken`)
  - `POST /api/token` (via `PostToken`)
