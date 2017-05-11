package gander

import (
	"encoding/csv"
	"github.com/tophatsteve/urlreader"
	"io"
	"os"
)

func LoadCSVFromURL(url string) (*DataFrame, error) {
	u := urlreader.NewReader(url)
	return LoadCSVFromReader(u)
}

func LoadCSVFromPath(path string) (*DataFrame, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	return LoadCSVFromReader(f)
}

func LoadCSVFromReader(reader io.Reader) (*DataFrame, error) {
	r := csv.NewReader(reader)
	data, err := r.ReadAll()
	if err != nil {
		return nil, err
	}

	return NewDataFrame(data)
}
