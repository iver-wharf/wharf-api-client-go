# Go-lang RPC client for Wharf

[![Codacy Badge](https://app.codacy.com/project/badge/Grade/3ab0d0c67ee642bfa1952dae4d99f55d)](https://www.codacy.com/gh/iver-wharf/wharf-api-client-go/dashboard?utm_source=github.com\&utm_medium=referral\&utm_content=iver-wharf/wharf-api-client-go\&utm_campaign=Badge_Grade)
[![Go Reference](https://pkg.go.dev/badge/github.com/iver-wharf/wharf-api-client-go.svg)](https://pkg.go.dev/github.com/iver-wharf/wharf-api-client-go)

A library to talk to Wharf via Wharf's main API written in Go.

Uses `net/http` to send HTTP requests and `encoding/json` to
serialize/deserialize each message back and forth.

This project is for example used inside the providers to create projects
into the database when importing from GitLab, GitHub, or Azure DevOps.

## Compatibility

| wharf-api-client-go | wharf-api |
| :---: | :---: |
| v1 | v0, v1, v2, v3, v4, *and* v5<sup>1</sup> |
| v2 | v5, *and beyond* |

<sup>1</sup>: The `PUT` endpoints `/api/project`, `/api/token`, and `/api/provider` were removed in `wharf-api v5.0.0`.

## Usage

```console
$ go get github.com/iver-wharf/wharf-api-client-go/v2/pkg/wharfapi
```

```go
package main

import (
	"fmt"
	"github.com/iver-wharf/wharf-api-client-go/v2/pkg/wharfapi"
)

func main() {
	client := wharfapi.Client{
		APIURL:     "https://example.wharf.com",
		AuthHeader: "Bearer some-auth-token",
	}

	project,err := client.GetProjectByID(125)

	if err != nil {
		fmt.Printf("Unable to find project\n")
	} else {
		fmt.Printf("Project #%d: %s\n", project.ProjectID, project.Name)
	}
}
```

### Sample output

```log
GET | PROJECT | 125
Project #125: MyProject
```

## Linting Golang

- Requires Node.js (npm) to be installed: <https://nodejs.org/en/download/>
- Requires Revive to be installed: <https://revive.run/>

```sh
go get -u github.com/mgechev/revive
```

```sh
npm run lint-go
```

## Linting markdown

- Requires Node.js (npm) to be installed: <https://nodejs.org/en/download/>

```sh
npm install

npm run lint-md

# Some errors can be fixed automatically. Keep in mind that this updates the
# files in place.
npm run lint-md-fix
```

## Linting

```sh
make deps # download linting dependencies

make lint

make lint-go # only lint Go code
make lint-md # only lint Markdown files
```

Some errors can be fixed automatically. Keep in mind that this updates the
files in place.

```sh
make lint-fix

make lint-fix-go # only lint and fix Go files
make lint-fix-md # only lint and fix Markdown files
```

---

Maintained by [Iver](https://www.iver.com/en).
Licensed under the [MIT license](./LICENSE).
