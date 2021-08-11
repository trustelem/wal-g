package zstd

import (
	"io"

	"github.com/klauspost/compress/zstd"
	"github.com/wal-g/wal-g/internal/compression/computils"
)

type Decompressor struct{}

func (decompressor Decompressor) Decompress(src io.Reader) (io.ReadCloser, error) {
	zstdReader, err := zstd.NewReader(computils.NewUntilEOFReader(src))
	return zstdReader.IOReadCloser(), err
}

func (decompressor Decompressor) FileExtension() string {
	return FileExtension
}
