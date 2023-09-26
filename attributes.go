package cbz

import (
	"fmt"
	"path/filepath"

	"github.com/fmartingr/go-comicinfo/v2"
)

// SetComicInfo sets the comic info for the comic book
func (c *CBZ) SetComicInfo(comicInfo *comicinfo.ComicInfo) {
	c.Info = comicInfo
}

// ComicInfo returns the comic info for the comic book
func (c *CBZ) ComicInfo() *comicinfo.ComicInfo {
	if c.Info == nil {
		c.Info = comicinfo.NewComicInfo()
	}

	return c.Info
}

// appendPage adds a page to the comic book
func (c *CBZ) appendPage(page *Page) error {
	if err := c.compressor.AddFileFromPath(fmt.Sprintf("%4d.%s", page.Number, filepath.Ext(page.Path)), page.Path); err != nil {
		return fmt.Errorf("error adding page file to cbz file: %w", err)
	}

	c.Pages = append(c.Pages, *page)

	if c.ComicInfo() != nil {
		c.ComicInfo().Pages.Append(comicinfo.ComicPageInfo{
			Image: page.Number,
		})
	}

	return nil
}

// AppendPage adds a page to the comic book
func (c *CBZ) AppendPage(path string) error {
	return c.appendPage(&Page{
		Path:   path,
		Number: len(c.Pages) + 1,
	})
}
