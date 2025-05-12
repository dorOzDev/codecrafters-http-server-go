package main

import (
	"bytes"
	"compress/gzip"
	"strings"
	"sync"
)

type Compresser interface {
	compress(input []byte) ([]byte, error)
	compressorName() string
}

type GzipCompresser struct {
}

func (GzipCompresser) compress(input []byte) ([]byte, error) {
	var buf bytes.Buffer
	gzipWriter := gzip.NewWriter(&buf)
	_, err := gzipWriter.Write([]byte(input))
	if err != nil {
		return nil, err
	}
	gzipWriter.Close()
	return buf.Bytes(), nil
}

func (GzipCompresser) compressorName() string {
	return GZIP
}

const (
	GZIP = "gzip"
)

var (
	supportedEncodingMap map[string]struct{}
	once                 sync.Once
)

func isSupportedEncoding(encodingArray []string) (string, bool) {
	once.Do(func() {
		supportedEncodingMap = map[string]struct{}{
			GZIP: {},
		}
	})

	for _, encoding := range encodingArray {
		_, exists := supportedEncodingMap[strings.ToLower(encoding)]
		if exists {
			return encoding, true
		}
	}

	return "", false
}

func parseAcceptEncoding(headerValue string) []string {
	encodings := strings.Split(headerValue, ",")
	for i := range encodings {
		encodings[i] = strings.TrimSpace(encodings[i])
	}
	return encodings
}
