package compresshttpheader

import (
	"testing"
)

func TestEncodeDecode(t *testing.T) {
	testSource := `{
		"Cache-Control":["no-cache","post-check=0, pre-check=0"],
		"Content-Type":["text/html; charset=UTF-8"],
		"Date":["Fri, 24 May 2019 22:06:06 GMT"],
		"Expires":["Mon, 26 Jul 1997 05:00:00 GMT"],
		"Last-Modified":["Fri, 24 May 2019 22:06:06 GMT"],
		"Pragma":["no-cache"],
		"Server":["Apache"],
		"Set-Cookie":["Session=55307921331a1ff5104a583f1a7b7f79; path=/; secure; HttpOnly"],
		"Strict-Transport-Security":["max-age=63072000; preload"],
		"Vary":["Accept-Encoding"],
		"X-Content-Type-Options":["nosniff"]
		}`

	compressedResult, err := Encode(testSource)
	if err != nil {
		t.Error(err)
	}

	decompressedResult, err := Decode(compressedResult)
	if err != nil {
		t.Error(err)
	}
	if testSource != string(decompressedResult) {
		t.Error("decompressedResult not equal")
	}
}
