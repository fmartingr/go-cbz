package cbz

import (
	"fmt"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateComic(t *testing.T) {
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "comic.cbz")

	comic, err := New()
	if err != nil {
		t.Errorf("New() error = %v", err)
	}

	// Set metadata
	comic.ComicInfo().Series = "My Comic"
	comic.ComicInfo().Volume = 1

	// Add pages
	for i := 1; i < 3; i++ {
		if err := comic.AppendPage(fmt.Sprintf("testdata/page%02d.jpg", i)); err != nil {
			t.Errorf("AppendPage() error = %v", err)
		}
	}

	if err := comic.Save(filePath); err != nil {
		t.Errorf("Save() error = %v", err)
	}

	assert.NotNil(t, comic)
	assert.FileExists(t, filePath)
	assert.Equal(t, "My Comic", comic.ComicInfo().Series)
	assert.Equal(t, 1, comic.ComicInfo().Volume)
	assert.Equal(t, 2, comic.Info.Pages.Len())
	assert.Equal(t, 2, len(comic.Pages))
}
