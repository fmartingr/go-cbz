# CBZ

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
	if err := comic.AppendPage(fmt.Sprintf("testdata/page01.jpg", i)); err != nil {
		log.Fatal(err)
	}
	if err := comic.AppendPage(fmt.Sprintf("testdata/page02.jpg", i)); err != nil {
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
