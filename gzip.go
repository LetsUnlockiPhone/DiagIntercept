package main

import (
	"bytes"
	"io"
	"compress/gzip"
)

// Taken from https://gist.github.com/alex-ant/aeaaf497055590dacba760af24839b8d
func GunzipData(data []byte) (resData []byte, err error) {
	b := bytes.NewBuffer(data)

	var r io.Reader
	r, err = gzip.NewReader(b)
	if err != nil {
		return
	}

	var resB bytes.Buffer
	_, err = resB.ReadFrom(r)
	if err != nil {
		return
	}

	resData = resB.Bytes()

	return
}
