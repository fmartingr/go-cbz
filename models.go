package cbz

import (
	"github.com/fmartingr/go-comicinfo/v2"
)

// CBZ represents a comic book
type CBZ struct {
	path       string
	compressor Compressor
	Info       *comicinfo.ComicInfo
	Pages      []Page
}

// Page represents a page in a comic book
// Keep in mind that the pages we store and the pages present in the ComicInfo may be different
// when opening files, since they may not be present in the ComicInfo file or the ComicInfo file
// may not be present at all.
type Page struct {
	Number int
	Path   string
}
