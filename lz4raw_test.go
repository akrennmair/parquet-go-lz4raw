package lz4raw_test

import (
	"bytes"
	"testing"

	goparquet "github.com/fraugster/parquet-go"
	"github.com/fraugster/parquet-go/parquet"
	"github.com/fraugster/parquet-go/parquetschema"
	"github.com/stretchr/testify/require"
)

func TestWriteThenReadLZ4Compressed(t *testing.T) {
	sd, err := parquetschema.ParseSchemaDefinition(`message msg {
		required binary foo (STRING);
		required int64 bar;
	}`)
	require.NoError(t, err)

	buf := &bytes.Buffer{}

	fw := goparquet.NewFileWriter(buf, goparquet.WithSchemaDefinition(sd), goparquet.WithCompressionCodec(parquet.CompressionCodec_LZ4_RAW))

	err = fw.AddData(map[string]interface{}{
		"foo": []byte("hello world"),
		"bar": int64(4200023),
	})
	require.NoError(t, err)
	require.NoError(t, fw.Close())

	r, err := goparquet.NewFileReader(bytes.NewReader(buf.Bytes()))
	require.NoError(t, err)

	record, err := r.NextRow()
	require.NoError(t, err)

	require.Equal(t, []byte("hello world"), record["foo"].([]byte))
	require.Equal(t, int64(4200023), record["bar"].(int64))
}
