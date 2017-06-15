package gander

import (
	"encoding/csv"
	"io"
	"os"

	"github.com/tophatsteve/urlreader"
)

// A Summary describes the statisical properties of a Series.
type Summary struct {
	Name     string
	Mean     float64
	Median   float64
	Mode     []float64
	Min      float64
	Max      float64
	StdDev   float64
	Variance float64
}

// LoadCSVFromURL creates a DataFrame by loading a csv file
// from a specific url. Note: at the moment this does not
// support https.
func LoadCSVFromURL(url string) (*DataFrame, error) {
	u := urlreader.NewReader(url)
	return loadCSVFromReader(u)
}

// LoadCSVFromPath creates a DataFrame by loading a csv file
// from a specific file system path.
func LoadCSVFromPath(path string) (*DataFrame, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	return loadCSVFromReader(f)
}

func loadCSVFromReader(reader io.Reader) (*DataFrame, error) {
	r := csv.NewReader(reader)
	data, err := r.ReadAll()

	if err != nil {
		return nil, err
	}

	return NewDataFrame(data)
}
