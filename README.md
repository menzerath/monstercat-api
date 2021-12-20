# Monstercat API

![GitHub Workflow Status](https://img.shields.io/github/workflow/status/menzerath/monstercat-api/go)
![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/menzerath/monstercat-api)
[![GoDoc](https://img.shields.io/static/v1?label=godoc&message=reference&color=blue)](https://pkg.go.dev/github.com/menzerath/monstercat-api/v2/monstercat)

Go-based wrapper and CLI to access Monstercat's API for releases and tracks.

## Supported Features

As there is no up-to-date API documentation, this repository is based on the reverse engineering of their website.
Sadly this also means that the API may break at any time.  
For additional features open an issue or feel free to create a pull request.

* login with your Monstercat account
* list, search and filter all releases
* download releases as FLAC or mp3 (login required)

## Usage

There are two ways you can utilize this project.

### CLI

Run a simple cli based on this project by downloading a [release](https://github.com/menzerath/monstercat-api/releases) or compiling it yourself.

Afterwards usage is as simple as this:

```bash
$ ./monstercat catalog --search="mix contest"
+------------+-------------------------------------------------+------------+---------+--------------+
| CATALOG ID |                      TITLE                      |   ARTIST   |  TYPE   | RELEASE DATE |
+------------+-------------------------------------------------+------------+---------+--------------+
| MMC604     | S6E4 - The Mix Contest - "You & Me"             | Monstercat | Podcast | 2021-08-11   |
| MMC603     | S6E3 - The Mix Contest - "Orbit"                | Monstercat | Podcast | 2021-08-04   |
| MMC602     | S6E2 - The Mix Contest - "There and Back"       | Monstercat | Podcast | 2021-07-28   |
| MMC601     | S6E1 - The Mix Contest - "Opening Ceremonies"   | Monstercat | Podcast | 2021-07-21   |
| MMCS600    | The Mix Contest 2021 - Submissions Open Now!    | Monstercat | Podcast | 2021-05-19   |
| MMC508     | The Mix Contest 2020 - Winner’s Showcase        | Monstercat | Podcast | 2020-09-23   |
| MMC507     | S5E7 - The Mix Contest - "Showdown"             | Monstercat | Podcast | 2020-09-02   |
| MMC506     | S5E6 - The Mix Contest - "Unity"                | Monstercat | Podcast | 2020-08-26   |
| MMC505     | S5E5 - The Mix Contest - "Bittersweet Horizons" | Monstercat | Podcast | 2020-08-19   |
| MMC504     | S5E4 - The Mix Contest - "How We Win, Together" | Monstercat | Podcast | 2020-08-12   |
+------------+-------------------------------------------------+------------+---------+--------------+
10 of 25 results
```

A list of all configurable options can be obtained by adding the `--help` flag to any command.

### API

To use this project in your own work, follow these steps:

```bash
go get -u github.com/menzerath/monstercat-api/v2
```

```go
package main

import (
	"fmt"
	"os"

	"github.com/menzerath/monstercat-api/v2/monstercat"
)

func main() {
	client := monstercat.NewClient()
	catalog, err := client.BrowseCatalog(WithSearch("mix contest"))
	if err != nil {
		fmt.Printf("error: %s", err)
		os.Exit(1)
	}
	fmt.Printf("catalog: %+v", catalog)
}
```

A list of all `BrowseOption`s is available in our [API documentation](https://pkg.go.dev/github.com/menzerath/monstercat-api/v2/monstercat#BrowseOption).
