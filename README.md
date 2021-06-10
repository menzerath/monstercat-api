# Monstercat API

[![GoDoc](https://img.shields.io/static/v1?label=godoc&message=reference&color=blue)](https://pkg.go.dev/github.com/menzerath/monstercat-api/monstercat?tab=doc)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/menzerath/monstercat-api)

Go-based wrapper to access Monstercat's API for releases and tracks.

## Supported Features

As there is no up-to-date API documentation, this repository is based on the reverse engineering of their website.
Sadly this also means that the API may break at any time.  
For additional features open an issue or feel free to create a pull request.

* login with your Monstercat account
* load and traverse release list
* download releases as FLAC or mp3 (login required)

## Usage

```bash
go get -u github.com/menzerath/monstercat-api
```

```go
package main

import (
	"fmt"
	"os"

	"github.com/menzerath/monstercat-api/monstercat"
)

func main() {
	client := monstercat.NewClient()
	releases, err := client.ReleaseList()
	if err != nil {
		fmt.Printf("error: %s", err)
		os.Exit(1)
	}
	fmt.Printf("releases: %+v", releases)
}
```
