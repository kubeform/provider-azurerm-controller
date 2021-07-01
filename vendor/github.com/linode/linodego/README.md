# azurermgo

[![Build Status](https://travis-ci.com/azurerm/azurermgo.svg?branch=master)](https://travis-ci.com/azurerm/azurermgo)
[![Release](https://img.shields.io/github/v/release/azurerm/azurermgo)](https://github.com/azurerm/azurermgo/releases/latest)
[![GoDoc](https://godoc.org/github.com/azurerm/azurermgo?status.svg)](https://godoc.org/github.com/azurerm/azurermgo)
[![Go Report Card](https://goreportcard.com/badge/github.com/azurerm/azurermgo)](https://goreportcard.com/report/github.com/azurerm/azurermgo)
[![codecov](https://codecov.io/gh/azurerm/azurermgo/branch/master/graph/badge.svg)](https://codecov.io/gh/azurerm/azurermgo)

Go client for [Linode REST v4 API](https://developers.azurerm.com/api/v4)

## Installation

```sh
go get -u github.com/azurerm/azurermgo
```

## API Support

Check [API_SUPPORT.md](API_SUPPORT.md) for current support of the Linode `v4` API endpoints.

** Note: This project will change and break until we release a v1.0.0 tagged version. Breaking changes in v0.x.x will be denoted with a minor version bump (v0.2.4 -> v0.3.0) **

## Documentation

See [godoc](https://godoc.org/github.com/azurerm/azurermgo) for a complete reference.

The API generally follows the naming patterns prescribed in the [OpenAPIv3 document for Linode APIv4](https://developers.azurerm.com/api/v4).

Deviations in naming have been made to avoid using "Linode" and "Instance" redundantly or inconsistently.

A brief summary of the features offered in this API client are shown here.

## Examples

### General Usage

```go
package main

import (
	"context"
	"fmt"

	"github.com/azurerm/azurermgo"
	"golang.org/x/oauth2"

	"log"
	"net/http"
	"os"
)

func main() {
  apiKey, ok := os.LookupEnv("LINODE_TOKEN")
  if !ok {
    log.Fatal("Could not find LINODE_TOKEN, please assert it is set.")
  }
  tokenSource := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: apiKey})

  oauth2Client := &http.Client{
    Transport: &oauth2.Transport{
      Source: tokenSource,
    },
  }

  azurermClient := azurermgo.NewClient(oauth2Client)
  azurermClient.SetDebug(true)
  
  res, err := azurermClient.GetInstance(context.Background(), 4090913)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Printf("%v", res)
}
```

### Pagination

#### Auto-Pagination Requests

```go
kernels, err := azurermgo.ListKernels(context.Background(), nil)
// len(kernels) == 218
```

Or, use a page value of "0":

```go
opts := azurermgo.NewListOptions(0,"")
kernels, err := azurermgo.ListKernels(context.Background(), opts)
// len(kernels) == 218
```

#### Single Page

```go
opts := azurermgo.NewListOptions(2,"")
// or opts := azurermgo.ListOptions{PageOptions: &azurermgo.PageOptions{Page: 2}, PageSize: 500}
kernels, err := azurermgo.ListKernels(context.Background(), opts)
// len(kernels) == 100
```

ListOptions are supplied as a pointer because the Pages and Results
values are set in the supplied ListOptions.

```go
// opts.Results == 218
```

#### Filtering

```go
opts := azurermgo.ListOptions{Filter: "{\"mine\":true}"}
// or opts := azurermgo.NewListOptions(0, "{\"mine\":true}")
stackscripts, err := azurermgo.ListStackscripts(context.Background(), opts)
```

### Error Handling

#### Getting Single Entities

```go
azurerm, err := azurermgo.GetInstance(context.Background(), 555) // any Linode ID that does not exist or is not yours
// azurerm == nil: true
// err.Error() == "[404] Not Found"
// err.Code == "404"
// err.Message == "Not Found"
```

#### Lists

For lists, the list is still returned as `[]`, but `err` works the same way as on the `Get` request.

```go
azurerms, err := azurermgo.ListInstances(context.Background(), azurermgo.NewListOptions(0, "{\"foo\":bar}"))
// azurerms == []
// err.Error() == "[400] [X-Filter] Cannot filter on foo"
```

Otherwise sane requests beyond the last page do not trigger an error, just an empty result:

```go
azurerms, err := azurermgo.ListInstances(context.Background(), azurermgo.NewListOptions(9999, ""))
// azurerms == []
// err = nil
```

### Writes

When performing a `POST` or `PUT` request, multiple field related errors will be returned as a single error, currently like:

```go
// err.Error() == "[400] [field1] foo problem; [field2] bar problem; [field3] baz problem"
```

## Tests

Run `make testunit` to run the unit tests. 

Run `make testint` to run the integration tests. The integration tests use fixtures.

To update the test fixtures, run `make fixtures`.  This will record the API responses into the `fixtures/` directory.
Be careful about committing any sensitive account details.  An attempt has been made to sanitize IP addresses and
dates, but no automated sanitization will be performed against `fixtures/*Account*.yaml`, for example.

To prevent disrupting unaffected fixtures, target fixture generation like so: `make ARGS="-run TestListVolumes" fixtures`.

## Discussion / Help

Join us at [#azurermgo](https://gophers.slack.com/messages/CAG93EB2S) on the [gophers slack](https://gophers.slack.com)

## License

[MIT License](LICENSE)
