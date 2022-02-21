# parquet-go-lz4raw

<p align="center">
<a href="https://pkg.go.dev/github.com/akrennmair/parquet-go-lz4raw"><img src="https://pkg.go.dev/badge/github.com/akrennmair/parquet-go-lz4raw.svg" alt="Go Reference"></a>
<a href="https://github.com/akrennmair/parquet-go-lz4raw/blob/main/LICENSE"><img src="https://img.shields.io/badge/license-Apache%202-blue"></a>
</p>

This library implements the LZ4\_RAW compression algorithm for [github.com/fraugster/parquet-go](github.com/fraugster/parquet-go). By default,
`parquet-go` library only supports `GZIP` and `SNAPPY` as compression algorithms to minimize the list
of dependencies.

All you need to do is import this package into your program and the compression
algorithm will be automatically available in `parquet-go`.

```go
import (
    _ "github.com/akrennmair/parquet-go-block-compressors/lz4raw" // registers the LZ4 block compressor with the LZ4_RAW compression type with parquet-go
)
```

## License

See the file `LICENSE` for further license information.

## Author

Andreas Krennmair <ak@synflood.at>
