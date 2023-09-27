# CBZ

[![Go Reference](https://pkg.go.dev/badge/github.com/fmartingr/go-cbz.svg)](https://pkg.go.dev/github.com/fmartingr/go-cbz)
[![Go Report Card](https://goreportcard.com/badge/github.com/fmartingr/go-cbz)](https://goreportcard.com/report/github.com/fmartingr/go-cbz)
[![codecov](https://codecov.io/gh/fmartingr/go-cbz/graph/badge.svg?token=AFCNOW5C2A)](https://codecov.io/gh/fmartingr/go-cbz)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Manage your CBZ files with Go.

## Features

- 📘 Create a CBZ file from scratch.

## Roadmap

- (TODO) 📙 Modify a CBZ file metadata (`comicinfo.xml`)
- (TODO) 📗 Automatically create metadata from contents (`comicinfo.xml`)
- (TODO) 📕 Extract a CBZ file (or just use `unzip`...?)
- (TODO) 🖥️ Use as CLI tool. `cbz`

## Usage

### Create a CBZ file

```go
package main

import (
	"fmt"
	"log"

	"github.com/fmartingr/go-cbz"
)

func main() {
	comic, err := cbz.New()
	if err != nil {
		log.Fatal(err)
	}

	// Set some metadata
	comic.ComicInfo().Series = "My Comic"
	comic.ComicInfo().Volume = 1

	// Add some pages
	if err := comic.AppendPage("testdata/page01.jpg"); err != nil {
		log.Fatal(err)
	}
	if err := comic.AppendPage("testdata/page02.jpg"); err != nil {
		log.Fatal(err)
	}

	// Save the comic
	if err := comic.Save("my-comic v01.cbz"); err != nil {
		log.Fatal(err)
	}
}
```

## License

[MIT License](LICENSE)
