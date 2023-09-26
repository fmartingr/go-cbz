package cbz

import (
	"fmt"
	"os"

	"github.com/fmartingr/go-comicinfo/v2"
)

// Save saves the comic book to the path specified
func (c *CBZ) Save(path string) error {
	if c.compressor != nil {
		comicInfoFile, err := c.compressor.AddFile("ComicInfo.xml")
		if err != nil {
			return fmt.Errorf("error creating page file in cbz file: %w", err)
		}

		if err := comicinfo.Write(c.ComicInfo(), comicInfoFile); err != nil {
			return fmt.Errorf("error writing comic info: %w", err)
		}

		if err := c.compressor.Close(); err != nil {
			return fmt.Errorf("error closing zip file: %w", err)
		}

		if err := os.Rename(c.path, path); err != nil {
			return fmt.Errorf("error moving file to destination: %w", err)
		}

		return nil
	}

	return fmt.Errorf("no compressor to use for saving")
}

// New creates a new ComicBook object
func New() (file *CBZ, err error) {
	f, err := os.CreateTemp("", "*.cbz")
	if err != nil {
		return nil, fmt.Errorf("error creating temporary comic file: %w", err)
	}

	file = &CBZ{
		path:       f.Name(),
		compressor: NewZIPCompressor(f),
	}

	return
}

// open creates a new ComicBook object from the path specified
// leaving private since this is not the intended way of creating a comic book and we need to test
// it properly.
func open(path string) (file *CBZ, err error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}
	defer f.Close()

	file = &CBZ{
		path:       path,
		compressor: NewZIPCompressor(f),
		Pages:      []Page{},
	}

	// Read `comicinfo.xml` if it exists
	comicInfoFile, err := file.compressor.GetFile("ComicInfo.xml")
	if err != nil && !os.IsNotExist(err) {
		return file, nil
	}

	if comicInfoFile != nil {
		file.Info, err = comicinfo.Read(comicInfoFile)
		if err != nil {
			return nil, fmt.Errorf("error opening comic info: %w", err)
		}
	}

	return file, nil
}
