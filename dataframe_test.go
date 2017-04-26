package gander

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

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
