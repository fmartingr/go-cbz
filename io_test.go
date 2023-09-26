package cbz

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	comic, err := New()
	if err != nil {
		t.Errorf("New() error = %v", err)
	}

	assert.NotNil(t, comic)
}

func TestSave(t *testing.T) {
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "comic.cbz")

	comic, err := New()
	if err != nil {
		t.Errorf("New() error = %v", err)
	}

	if err := comic.Save(filePath); err != nil {
		t.Errorf("Save() error = %v", err)
	}

	assert.FileExists(t, filePath)
}
