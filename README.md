# compresshttpheader [![GoDoc](http://godoc.org/github.com/paulvollmer/go-compresshttpheader?status.svg)](http://godoc.org/github.com/paulvollmer/go-compresshttpheader) [![Build Status](https://travis-ci.org/paulvollmer/go-compresshttpheader.svg?branch=master)](https://travis-ci.org/paulvollmer/go-compresshttpheader) [![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/paulvollmer/go-compresshttpheader/blob/master/LICENSE)

this package compress a string of http headers with focus on the smalles size.

## Installation

```
go get github.com/paulvollmer/go-compresshttpheader
```

## Benchmark

```
GZIP SIZE:   338
FLATE SIZE:  243 <- compresshttpheader
```

## License

Released under the [MIT license](LICENSE).
