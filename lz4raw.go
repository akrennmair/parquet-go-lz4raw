package lz4raw

import (
	"bytes"
	"errors"
	"io"

	goparquet "github.com/fraugster/parquet-go"
	"github.com/fraugster/parquet-go/parquet"
	lz4 "github.com/pierrec/lz4/v4"
)

type LZ4RawBlockCompressor struct{}

func NewLZ4RawBlockCompressor() *LZ4RawBlockCompressor {
	return &LZ4RawBlockCompressor{}
}

func (c *LZ4RawBlockCompressor) CompressBlock(data []byte) ([]byte, error) {
	buf := &bytes.Buffer{}

	w := lz4.NewWriter(buf)
	n, err := w.Write(data)
	if err != nil {
		return nil, err
	}
	if n < len(data) {
		return nil, errors.New("short write")
	}
	if err := w.Close(); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (c *LZ4RawBlockCompressor) DecompressBlock(data []byte) ([]byte, error) {
	r := lz4.NewReader(bytes.NewReader(data))
	return io.ReadAll(r)
}

func init() {
	goparquet.RegisterBlockCompressor(parquet.CompressionCodec_LZ4_RAW, NewLZ4RawBlockCompressor())
}
