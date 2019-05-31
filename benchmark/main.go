package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"

	"github.com/paulvollmer/go-compresshttpheader"
)

func main() {
	textContentType := []byte(`{
	  "Cache-Control":["no-cache","post-check=0, pre-check=0"],
	  "Content-Type":["text/html; charset=UTF-8"],
	  "Date":["Fri, 24 May 2019 22:06:06 GMT"],
	  "Expires":["Mon, 26 Jul 1997 05:00:00 GMT"],
	  "Last-Modified":["Fri, 24 May 2019 22:06:06 GMT"],
	  "Pragma":["no-cache"],
	  "Server":["Apache"],
	  "Set-Cookie":["Session=1234567890; path=/; secure; HttpOnly"],
	  "Strict-Transport-Security":["max-age=63072000; preload"],
	  "Vary":["Accept-Encoding"],
	  "X-Content-Type-Options":["nosniff"]
	}`)

	compressGzip(textContentType)
	compressHeader(bytes.NewReader(textContentType))
}

func compressGzip(source []byte) {
	var b bytes.Buffer
	gz, err := gzip.NewWriterLevel(&b, gzip.BestCompression)
	if err != nil {
		panic(err)
	}
	if _, err := gz.Write(source); err != nil {
		panic(err)
	}
	if err := gz.Flush(); err != nil {
		panic(err)
	}
	if err := gz.Close(); err != nil {
		panic(err)
	}
	fmt.Println("GZIP SIZE:\t", len(b.Bytes()))
}

func compressHeader(source io.Reader) {
	result, err := compresshttpheader.Encode(source)
	if err != nil {
		panic(err)
	}
	fmt.Println("FLATE SIZE:\t", len(result))
}
