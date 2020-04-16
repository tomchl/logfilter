package logfilter

import (
	"log"
	"strings"
)

// IgnoreHTTPWriter  struct with write
type IgnoreHTTPWriter struct{}

func (*IgnoreHTTPWriter) Write(p []byte) (n int, err error) {
	if !strings.HasPrefix(string(p), "http:") {
		log.Println(string(p))
		return len(p), nil
	}
	return len(p), nil
}
