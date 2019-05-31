# compresshttpheader [![GoDoc](http://godoc.org/github.com/paulvollmer/go-compresshttpheader?status.svg)](http://godoc.org/github.com/paulvollmer/go-compresshttpheader)

this package compress a string of http headers with focus on the smalles size.

## Installation

```
go get github.com/paulvollmer/go-compresshttpheader
```

## Benchmark

```
GZIP SIZE:	 338
FLATE SIZE:	 243 <- compresshttpheader
```

## License

Released under the [MIT license](LICENSE).
