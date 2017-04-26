package gander

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateDataFrameFromCSV(t *testing.T) {
	df, err := NewDataFrame(createSampleDataWithHeaders())
	assert.Equal(t, nil, err, "error is not nil")
	assert.Equal(t, 5, len(*df), "dataframe does not have the correct number of columns")
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
