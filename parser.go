package gander

// Load data and convert to a DataFrame

import (
	"encoding/csv"
	"io"
	"os"
	"strconv"
)

func LoadCSVFromURL(url string) (*DataFrame, error) {
	return &DataFrame{}, nil
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

func isNumeric(v string) bool {
	_, err := strconv.ParseFloat(v, 64)
	if err != nil {
		return false
	}
	return true
}

func hasHeaderRow(data [][]string) bool {
	// is every cell on the 1st row not a number
	n := 0
	for x := 0; x < len(data[0]); x++ {
		if isNumeric(data[0][x]) == false {
			n += 1
		}
	}

	if n == len(data[0]) {
		return true
	}

	return false
}

func hasCategoricalData(data [][]string, column int) bool {
	// is any value in the column not numeric
	startRow := 0

	if hasHeaderRow(data) == true {
		startRow = 1
	}

	for r := startRow; r < len(data); r++ {
		if isNumeric(data[r][column]) == false {
			return true
		}
	}

	return false
}

func columnCountsMatch(data [][]string) bool {
	// do all rows have the same number of columns
	if len(data) == 0 {
		return false
	}

	c := len(data[0])

	for _, v := range data {
		if len(v) != c {
			return false
		}
	}

	return true
}

func createSeries(name string, data [][]string, column int) *Series {
	if hasCategoricalData(data, column) {
		values := []string{}
		for _, v := range data {
			values = append(values, v[column])
		}
		return NewCategoricalSeries(name, values)
	}

	values := []float64{}

	for _, v := range data {
		value, _ := strconv.ParseFloat(v[column], 64)
		values = append(values, value)
	}

	return NewSeries(name, values)
}
