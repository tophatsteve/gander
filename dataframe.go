package gander

import (
	"errors"
	"fmt"
	"strconv"
)

type DataFrame []*Series

// NewDataFrame creates a DataFrame from a 2 dimensional string slice.
func NewDataFrame(data [][]string) (*DataFrame, error) {
	if !columnCountsMatch(data) {
		return nil, errors.New("not all rows have the same number of columns")
	}

	var headers []string
	if hasHeaderRow(data) {
		headers = data[0]
		data = data[1:]
	} else {
		headers = []string{}
		for x := 0; x < len(data[0]); x++ {
			headers = append(headers, fmt.Sprintf("Column %v", x+1))
		}
	}

	d := DataFrame{}

	for x := 0; x < len(data[0]); x++ {
		s := createSeries(headers[x], data, x)
		d = append(d, s)
	}

	return &d, nil
}

// Columns returns a slice of the column names in the DataFrame.
func (d *DataFrame) Columns() []string {
	cols := []string{}

	for _, c := range *d {
		cols = append(cols, c.Name)
	}

	return cols
}

// String returns a tabular representation of the DataFrame.
func (d *DataFrame) String() string {
	df := *d
	columns := len(df)
	colWidths := []int{}
	output := ""

	for c := 0; c < columns; c++ {
		cl := len(df[c].Name) + 2 // add 2 for padding
		if cl < 12 {
			cl = 12
		}
		colWidths = append(colWidths, cl)
		output += fmt.Sprintf("%"+strconv.Itoa(cl)+"v", fmt.Sprintf("%v  ", df[c].Name))
	}

	output += fmt.Sprintf("\n")

	rows := len(df[0].Values)
	if rows > 10 {
		rows = 10 // only print 1st 10 rows
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < columns; c++ {
			output += fmt.Sprintf(" %"+strconv.Itoa(colWidths[c]-3)+".2f  ", df[c].Values[r])
		}
		output += fmt.Sprintf("\n")
	}

	return output
}
