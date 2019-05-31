package compresshttpheader

import (
	"bytes"
	"compress/flate"
	"io"
	"strings"
)

func Encode(source string) ([]byte, error) {
	var b bytes.Buffer
	zw, err := flate.NewWriterDict(&b, flate.BestCompression, dictionary)
	if err != nil {
		return nil, err
	}
	if _, err := io.Copy(zw, strings.NewReader(source)); err != nil {
		return nil, err
	}
	if err := zw.Close(); err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

func Decode(source []byte) ([]byte, error) {
	zr := flate.NewReaderDict(bytes.NewReader(source), dictionary)
	var b bytes.Buffer
	if _, err := io.Copy(&b, zr); err != nil {
		return nil, err
	}
	if err := zr.Close(); err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
