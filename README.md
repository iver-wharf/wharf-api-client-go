# Go-lang RPC client for Wharf

[![Go Reference](https://pkg.go.dev/badge/github.com/iver-wharf/wharf-api-client-go.svg)](https://pkg.go.dev/github.com/iver-wharf/wharf-api-client-go)

A library to talk to Wharf via Wharf's main API written in Go.

Uses `net/http` to send HTTP requests and `encoding/json` to
serialize/deserialize each message back and forth.

This project is for example used inside the providers to create projects
into the database when importing from GitLab, GitHub, or Azure DevOps.

## Usage

```go
package main

import (
	"fmt"
	"github.com/iver-wharf/wharf-api-client-go/pkg/wharfapi"
)

func main() {
	client := wharfapi.Client{
		ApiUrl:     "https://example.wharf.com",
		AuthHeader: "Bearer some-auth-token",
	}

	project,err := client.GetProjectById(125)

	if err != nil {
		fmt.Printf("Unable to find project\n")
	} else {
		fmt.Printf("Project #%d: %s\n", project.ProjectID, project.Name)
	}
}
```

### Sample output

```
GET | PROJECT | 125
Project #125: MyProject
```

---

Maintained by [Iver](https://www.iver.com/en).
Licensed under the [MIT license](./LICENSE).
