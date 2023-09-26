package cbz

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
)

// Compressor is the interface that wraps the basic methods for compressing files
type Compressor interface {
	AddFile(filename string) (io.Writer, error)
	AddFileFromPath(filename, path string) error
	GetFile(filename string) (io.Reader, error)
	Close() error
}

// ZIPCompressor is a Compressor that uses the ZIP format
type ZIPCompressor struct {
	zipWriter *zip.Writer
	zipReader *zip.Reader
}

func (z *ZIPCompressor) AddFile(filename string) (io.Writer, error) {
	return z.zipWriter.Create(filename)
}

func (z *ZIPCompressor) AddFileFromPath(filename, path string) error {
	file, err := z.zipWriter.Create(filename)
	if err != nil {
		return fmt.Errorf("error creating page file in cbz file: %w", err)
	}

	f, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("error opening file: %w", err)
	}
	defer f.Close()

	if _, err := io.Copy(file, f); err != nil {
		return fmt.Errorf("error copying file to cbz file: %w", err)
	}

	return nil
}

func (z *ZIPCompressor) GetFile(filename string) (io.Reader, error) {
	file, err := z.zipReader.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}

	return file, nil
}

func (z *ZIPCompressor) Close() error {
	return z.zipWriter.Close()
}

// NewZIPCompressor creates a new ZIPCompressor
func NewZIPCompressor(w io.Writer) Compressor {
	return &ZIPCompressor{
		zipWriter: zip.NewWriter(w),
	}
}
