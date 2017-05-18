package gander

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func cellValuesMatch(df *DataFrame, i [][]float64) bool {
	for k, v := range *df {
		for j, x := range v.Values {
			if i[j][k] != x {
				return false
			}
		}
	}
	return true
}

func TestCreateDataFrame(t *testing.T) {
	df, err := NewDataFrame(createSampleDataWithHeaders())
	assert.Equal(t, nil, err, "error is not nil")
	assert.Equal(t, 5, len(*df), "dataframe does not have the correct number of columns")
}

func TestCreateDataFrameWithCategoricalData(t *testing.T) {
	df, err := NewDataFrame(createSampleDataWithCategoricalData())
	assert.Equal(t, nil, err, "error is not nil")
	s := (*df)[3]
	assert.Equal(t, true, s.IsCategorical(), "series does not contain categorical data")
	assert.Equal(t, 2, len(s.CategoricalLabels), "wrong number of category labels created")
	assert.Equal(t, 2, len(s.CategoricalValues), "wrong number of category values created")
}

func TestStringFullFrame(t *testing.T) {
	s := `         a           b           c           d           e  
      1.00        2.00        3.00        4.00        5.00  
      3.00        5.00        2.00        2.00        4.00  
      7.00        6.00        1.00        3.00        3.00  
      4.00        2.00        4.00        7.00        6.00  
`
	df, err := NewDataFrame(createSampleDataWithHeaders())
	assert.Equal(t, nil, err, "error is not nil")
	r := df.String()
	assert.Equal(t, s, r, "string not returned in correct format")
}

func TestStringHeadOnly(t *testing.T) {
	s := `         a           b           c           d           e  
      1.00        2.00        3.00        4.00        5.00  
      3.00        5.00        2.00        2.00        4.00  
      7.00        6.00        1.00        3.00        3.00  
      4.00        2.00        4.00        7.00        6.00  
      1.00        2.00        3.00        4.00        5.00  
      3.00        5.00        2.00        2.00        4.00  
      7.00        6.00        1.00        3.00        3.00  
      4.00        2.00        4.00        7.00        6.00  
      1.00        2.00        3.00        4.00        5.00  
      3.00        5.00        2.00        2.00        4.00  
`
	df, err := NewDataFrame(createLargerSampleData())
	assert.Equal(t, nil, err, "error is not nil")
	r := df.String()
	assert.Equal(t, s, r, "string not returned in correct format")
}

func TestListColumnNames(t *testing.T) {
	expected := []string{"a", "b", "c", "d", "e"}
	df, err := NewDataFrame(createSampleDataWithHeaders())
	assert.Equal(t, nil, err, "error is not nil")
	cols := df.ColumnNames()
	assert.Equal(t, 5, len(cols), "wrong number of columns returned")

	for i, v := range expected {
		assert.Equal(t, v, cols[i], "column header is not correct")
	}
}

func TestDropRowsSingleRow(t *testing.T) {
	df, err := NewDataFrame(createSampleDataWithHeaders())
	assert.Equal(t, nil, err, "error is not nil")
	df.DropRows(2)
	assert.Equal(t, 3, df.Rows(), "wrong number of rows remaining")
	matchValues := cellValuesMatch(df, [][]float64{
		{1, 2, 3, 4, 5},
		{3, 5, 2, 2, 4},
		{4, 2, 4, 7, 6},
	})

	assert.Equal(t, true, matchValues, "values in dataframe do not match expected")
}

func TestDropRowsMultipleRows(t *testing.T) {
	df, err := NewDataFrame(createSampleDataWithHeaders())
	assert.Equal(t, nil, err, "error is not nil")
	df.DropRows(0, 2)
	assert.Equal(t, 2, df.Rows(), "wrong number of rows remaining")
	matchValues := cellValuesMatch(df, [][]float64{
		{3, 5, 2, 2, 4},
		{4, 2, 4, 7, 6},
	})

	assert.Equal(t, true, matchValues, "values in dataframe do not match expected")
}

func TestDropRowsDropAllRows(t *testing.T) {
	df, err := NewDataFrame(createSampleDataWithHeaders())
	assert.Equal(t, nil, err, "error is not nil")
	df.DropRows(0, 2, 1, 3)
	assert.Equal(t, 0, df.Rows(), "wrong number of rows remaining")
}

func TestDropRowsOutOfRange(t *testing.T) {
	df, err := NewDataFrame(createSampleDataWithHeaders())
	assert.Equal(t, nil, err, "error is not nil")
	err = df.DropRows(4)
	assert.Equal(t, "a specified row is out of range", err.Error(), "error message is not correct")
}

func TestDropRowsWhereSingleRow(t *testing.T) {
	df, err := NewDataFrame(createSampleDataWithHeaders())
	assert.Equal(t, nil, err, "error is not nil")
	df.DropRowsWhere(func(i []float64) bool {
		return i[0] == 3
	})
	assert.Equal(t, 3, df.Rows(), "wrong number of rows remaining")
	matchValues := cellValuesMatch(df, [][]float64{
		{1, 2, 3, 4, 5},
		{7, 6, 1, 3, 3},
		{4, 2, 4, 7, 6},
	})

	assert.Equal(t, true, matchValues, "values in dataframe do not match expected")
}

func TestDropRowsWhereMulitpleRows(t *testing.T) {
	df, err := NewDataFrame(createSampleDataWithHeaders())
	assert.Equal(t, nil, err, "error is not nil")
	df.DropRowsWhere(func(i []float64) bool {
		return i[0] == 3 || i[0] == 4
	})
	assert.Equal(t, 2, df.Rows(), "wrong number of rows remaining")
	matchValues := cellValuesMatch(df, [][]float64{
		{1, 2, 3, 4, 5},
		{7, 6, 1, 3, 3},
	})

	assert.Equal(t, true, matchValues, "values in dataframe do not match expected")
}

func TestDropRowsWhereDropAllRows(t *testing.T) {
	df, err := NewDataFrame(createSampleDataWithHeaders())
	assert.Equal(t, nil, err, "error is not nil")
	df.DropRowsWhere(func(i []float64) bool {
		return true
	})
	assert.Equal(t, 0, df.Rows(), "wrong number of rows remaining")
}

func TestDropColumnsSingleColumn(t *testing.T) {
	df, err := NewDataFrame(createSampleDataWithHeaders())
	assert.Equal(t, nil, err, "error is not nil")
	df.DropColumns(2)
	assert.Equal(t, 4, df.Columns(), "wrong number of columns remaining")
	matchValues := cellValuesMatch(df, [][]float64{
		{1, 2, 4, 5},
		{3, 5, 2, 4},
		{7, 6, 3, 3},
		{4, 2, 7, 6},
	})

	assert.Equal(t, true, matchValues, "values in dataframe do not match expected")
}

func TestDropColumnsMultipleColumns(t *testing.T) {
	df, err := NewDataFrame(createSampleDataWithHeaders())
	assert.Equal(t, nil, err, "error is not nil")
	df.DropColumns(2, 4, 0)
	assert.Equal(t, 2, df.Columns(), "wrong number of columns remaining")
	matchValues := cellValuesMatch(df, [][]float64{
		{2, 4},
		{5, 2},
		{6, 3},
		{2, 7},
	})

	assert.Equal(t, true, matchValues, "values in dataframe do not match expected")
}

func TestDropColumnsOutOfRange(t *testing.T) {
	df, err := NewDataFrame(createSampleDataWithHeaders())
	assert.Equal(t, nil, err, "error is not nil")
	err = df.DropColumns(5)
	assert.Equal(t, "a specified column is out of range", err.Error(), "error message is not correct")
}

func TestDropColumnsDropAllColumns(t *testing.T) {
	df, err := NewDataFrame(createSampleDataWithHeaders())
	assert.Equal(t, nil, err, "error is not nil")
	df.DropColumns(2, 4, 0, 1, 3)
	assert.Equal(t, 0, df.Columns(), "wrong number of columns remaining")
}

func TestDropColumnsByNameSingleColumn(t *testing.T) {
	df, err := NewDataFrame(createSampleDataWithHeaders())
	assert.Equal(t, nil, err, "error is not nil")
	df.DropColumnsByName("c")
	assert.Equal(t, 4, df.Columns(), "wrong number of columns remaining")
	matchValues := cellValuesMatch(df, [][]float64{
		{1, 2, 4, 5},
		{3, 5, 2, 4},
		{7, 6, 3, 3},
		{4, 2, 7, 6},
	})

	assert.Equal(t, true, matchValues, "values in dataframe do not match expected")
}

func TestDropColumnsByNameMultipleColumns(t *testing.T) {
	df, err := NewDataFrame(createSampleDataWithHeaders())
	assert.Equal(t, nil, err, "error is not nil")
	df.DropColumnsByName("c", "e", "a")
	assert.Equal(t, 2, df.Columns(), "wrong number of columns remaining")
	matchValues := cellValuesMatch(df, [][]float64{
		{2, 4},
		{5, 2},
		{6, 3},
		{2, 7},
	})

	assert.Equal(t, true, matchValues, "values in dataframe do not match expected")
	assert.Equal(t, "b", (*df)[0].Name, "column name is not correct for column 0")
	assert.Equal(t, "d", (*df)[1].Name, "column name is not correct for column 0")
}

func TestDropColumnsByNameDropAllColumns(t *testing.T) {
	df, err := NewDataFrame(createSampleDataWithHeaders())
	assert.Equal(t, nil, err, "error is not nil")
	df.DropColumnsByName("c", "e", "a", "b", "d")
	assert.Equal(t, 0, df.Columns(), "wrong number of columns remaining")
}

func TestDropColumnsByNameInvalidName(t *testing.T) {
	df, err := NewDataFrame(createSampleDataWithHeaders())
	assert.Equal(t, nil, err, "error is not nil")
	err = df.DropColumnsByName("f")
	assert.Equal(t, "column 'f' does not exist in the DataFrame", err.Error(), "")
	assert.Equal(t, 5, df.Columns(), "wrong number of columns remaining")
}

func TestDropColumnsWhereSingleColumn(t *testing.T) {
	df, err := NewDataFrame(createSampleDataWithHeaders())
	assert.Equal(t, nil, err, "error is not nil")
	df.DropColumnsWhere(func(s *Series) bool {
		return s.Name == "c"
	})
	assert.Equal(t, 4, df.Columns(), "wrong number of columns remaining")
	matchValues := cellValuesMatch(df, [][]float64{
		{1, 2, 4, 5},
		{3, 5, 2, 4},
		{7, 6, 3, 3},
		{4, 2, 7, 6},
	})

	assert.Equal(t, true, matchValues, "values in dataframe do not match expected")
}

func TestDropColumnsWhereMultipleColumns(t *testing.T) {
	df, err := NewDataFrame(createSampleDataWithHeaders())
	assert.Equal(t, nil, err, "error is not nil")
	df.DropColumnsWhere(func(s *Series) bool {
		return s.Name == "c" || s.Name == "e"
	})
	assert.Equal(t, 3, df.Columns(), "wrong number of columns remaining")
	matchValues := cellValuesMatch(df, [][]float64{
		{1, 2, 4},
		{3, 5, 2},
		{7, 6, 3},
		{4, 2, 7},
	})

	assert.Equal(t, true, matchValues, "values in dataframe do not match expected")
}

func TestDropColumnsWhereDropAllColumns(t *testing.T) {
	df, err := NewDataFrame(createSampleDataWithHeaders())
	assert.Equal(t, nil, err, "error is not nil")
	df.DropColumnsWhere(func(s *Series) bool {
		return true
	})
	assert.Equal(t, 0, df.Columns(), "wrong number of columns remaining")
}

func TestStandardizeDataFrame(t *testing.T) {
	df, err := NewDataFrame(createSampleDataWithCategoricalData())
	assert.Equal(t, nil, err, "error is not nil")
	df.Standardize()

	dfe, err := NewDataFrame(createSampleDataWithCategoricalData())

	for c, s := range *df {
		if s.IsCategorical() == false {
			assert.Equal(t, true, toleratedError(0.0, s.Mean()), "mean is not zero(ish)")
			assert.Equal(t, true, toleratedError(1.0, s.StdDev()), "std dev is not one")
		} else {
			for i, v := range s.Values {
				assert.Equal(t, (*dfe)[c].Values[i], v, "categorical data has been changed")
			}
		}
	}
}
