package zstd

import (
	"io"
	"os"

	"github.com/klauspost/compress/zstd"
	"github.com/wal-g/tracelog"
)

const (
	AlgorithmName = "zstd"
	FileExtension = "zst"
)

type Compressor struct{}

func (compressor Compressor) NewWriter(writer io.Writer) io.WriteCloser {
	speed := zstd.SpeedBestCompression

	val, ok := os.LookupEnv("WALG_COMPRESSION_LEVEL")
	if ok {
		switch val {
		case "fastest", "1":
			speed = zstd.SpeedFastest
		case "default", "fast", "2", "":
			speed = zstd.SpeedDefault
		case "better", "3":
			speed = zstd.SpeedBetterCompression
		case "best", "4":
			speed = zstd.SpeedBestCompression

		default:
			tracelog.WarningLogger.Printf("zstd compressor: unknown compression level '%s', using 'default'", val)
			tracelog.DebugLogger.Print("    possible values are : fastest|fast|better|best")
		}
	}

	zw, err := zstd.NewWriter(writer, zstd.WithEncoderLevel(speed))
	if err != nil {
		panic(err)
	}

	return zw
}

func (compressor Compressor) FileExtension() string {
	return FileExtension
}
