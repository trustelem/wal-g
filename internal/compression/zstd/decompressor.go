package zstd

import (
	"io"

	"github.com/klauspost/compress/zstd"
	"github.com/pkg/errors"
	"github.com/wal-g/wal-g/internal/compression/computils"
	"github.com/wal-g/wal-g/utility"
)

type Decompressor struct{}

func (decompressor Decompressor) Decompress(dst io.Writer, src io.Reader) error {
	zstdReader, err := zstd.NewReader(computils.NewUntilEOFReader(src))
	if err != nil {
		return err
	}
	_, err = utility.FastCopy(dst, zstdReader)
	if err != nil {
		return errors.Wrap(err, "DecompressZstd: zstd copy failed")
	}
	zstdReader.Close()
	return errors.Wrap(err, "DecompressZstd: zstd reader close failed")
}

func (decompressor Decompressor) FileExtension() string {
	return FileExtension
}
