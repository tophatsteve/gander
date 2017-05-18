package gander

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"sync"
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

// DropRows removes the rows specified by the provided row numbers.
// Row numbers are zero based.
func (d *DataFrame) DropRows(r ...int) error {
	if maxInt(r) > d.Rows()-1 {
		return errors.New("a specified row is out of range")
	}

	if sort.IntsAreSorted(r) == false {
		sort.Ints(r)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(r)))

	for _, v := range r {
		d.dropRow(v)
	}

	return nil
}

// DropRowsWhere removes all the rows where the provided function
// evaluates to true.
func (d *DataFrame) DropRowsWhere(fn func([]float64) bool) error {
	for i := d.Rows() - 1; i >= 0; i-- {
		r := d.toRow(i)
		if fn(r) == true {
			d.dropRow(i)
		}
	}

	return nil
}

// DropColumns removes the columns specified by the provided column numbers.
// Column numbers are zero based.
func (d *DataFrame) DropColumns(r ...int) error {
	if maxInt(r) > d.Columns()-1 {
		return errors.New("a specified column is out of range")
	}

	df := DataFrame{}
	for i, v := range *d {
		if containsInt(i, r) == false {
			df = append(df, v)
		}
	}

	*d = df

	return nil
}

// DropColumnsByName removes the columns specified by the provided column names.
// Row indexes are zero based.
func (d *DataFrame) DropColumnsByName(n ...string) error {
	for _, c := range n {
		if containsString(c, d.ColumnNames()) == false {
			return fmt.Errorf("column '%s' does not exist in the DataFrame", c)
		}
	}

	df := DataFrame{}
	for _, v := range *d {
		if containsString(v.Name, n) == false {
			df = append(df, v)
		}
	}

	*d = df

	return nil
}

// DropColumnsWhere removes all the columns where the provided function
// evaluates to true.
func (d *DataFrame) DropColumnsWhere(fn func(*Series) bool) error {
	df := DataFrame{}
	for _, v := range *d {
		if fn(v) == false {
			df = append(df, v)
		}
	}

	*d = df

	return nil
}

// ColumnNames returns a slice of the column names in the DataFrame.
func (d *DataFrame) ColumnNames() []string {
	cols := []string{}

	for _, c := range *d {
		cols = append(cols, c.Name)
	}

	return cols
}

// Columns returns the number of columns in the DataFrame.
func (d *DataFrame) Columns() int {
	return len(*d)
}

// Rows returns the number of rows in the DataFrame.
func (d *DataFrame) Rows() int {
	return len((*d)[0].Values)
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
			if df[c].IsCategorical() == true {
				output += fmt.Sprintf(" %"+strconv.Itoa(colWidths[c]-3)+"s  ", df[c].CategoricalLabels[df[c].Values[r]])
			} else {
				output += fmt.Sprintf(" %"+strconv.Itoa(colWidths[c]-3)+".2f  ", df[c].Values[r])
			}
		}
		output += fmt.Sprintf("\n")
	}

	return output
}

// Standardize scales the values in all non-categorical Series
// to standard form.
func (d *DataFrame) Standardize() {
	for _, v := range *d {
		if v.IsCategorical() == false {
			v.Standardize()
		}
	}
}

func (d *DataFrame) toRow(i int) []float64 {
	r := []float64{}

	for _, v := range *d {
		r = append(r, v.Values[i])
	}

	return r
}

func (d *DataFrame) dropRow(r int) {
	var wg sync.WaitGroup
	wg.Add(d.Columns())

	for i, _ := range *d {
		go func(s *Series) {
			defer wg.Done()
			s.dropRow(r)
		}((*d)[i])
	}
	wg.Wait()
}

func maxInt(l []int) int {
	m := -1

	for _, v := range l {
		if v > m {
			m = v
		}
	}

	return m
}

func containsInt(i int, l []int) bool {
	// TODO: make more efficient
	for _, v := range l {
		if i == v {
			return true
		}
	}

	return false
}

func containsString(i string, l []string) bool {
	// TODO: make more efficient
	for _, v := range l {
		if i == v {
			return true
		}
	}

	return false
}
