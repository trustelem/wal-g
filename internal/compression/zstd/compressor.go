package zstd

import (
	"io"

	"github.com/klauspost/compress/zstd"
)

const (
	AlgorithmName = "zstd"
	FileExtension = "zst"
)

type Compressor struct{}

func (compressor Compressor) NewWriter(writer io.Writer) io.WriteCloser {
	zw, _ := zstd.NewWriter(writer, zstd.WithEncoderLevel(zstd.SpeedDefault))

	return zw
}

func (compressor Compressor) FileExtension() string {
	return FileExtension
}
