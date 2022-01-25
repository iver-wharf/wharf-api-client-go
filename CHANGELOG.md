# Wharf API Go client changelog

This project tries to follow [SemVer 2.0.0](https://semver.org/).

<!--
	When composing new changes to this list, try to follow convention.

	The WIP release shall be updated just before adding the Git tag.
	From (WIP) to (YYYY-MM-DD), ex: (2021-02-09) for 9th of February, 2021

	A good source on conventions can be found here:
	https://changelog.md/
-->

## v2.0.0 (2022-01-25)

- BREAKING: Changed module path from `github.com/iver-wharf/wharf-api-client-go`
  to `github.com/iver-wharf/wharf-api-client-go/v2`. (#30)

- BREAKING: Changed minimum version of Go from 1.13 to 1.16. (#29)

- BREAKING: Removed support for wharf-api v4.2.0 and below. (#29)

- Added support for wharf-api v5. (#29)

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
