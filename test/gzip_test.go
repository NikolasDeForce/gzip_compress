package gzip_test

import (
	"gzip"
	"os"
	"testing"
)

func TestGzip(t *testing.T) {
	got := gzip.Compress("exemple1.txt")

	_, err := os.ReadFile("exemple1.txt.gz")
	if err != nil {
		t.Errorf("Compress() not working -> %v err --> %v", got, err)
	}
	defer os.Remove("exemple1.txt.gz")
}
