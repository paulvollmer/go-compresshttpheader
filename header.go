package compresshttpheader

import (
	"bytes"
	"compress/flate"
	"io"
)

func Encode(source io.Reader) ([]byte, error) {
	var b bytes.Buffer
	zw, err := flate.NewWriterDict(&b, flate.BestCompression, dictionary)
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(zw, source);
	if err != nil {
		return nil, err
	}
	err = zw.Close();
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

func Decode(source io.Reader) ([]byte, error) {
	zr := flate.NewReaderDict(source, dictionary)
	var b bytes.Buffer
	_, err := io.Copy(&b, zr);
	if err != nil {
		return nil, err
	}
	err = zr.Close();
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
