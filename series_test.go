package gander

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func createTestSeries() *Series {
	s := NewSeries(
		"MySeries",
		[]float64{
			0, 2, 7, 1, 4, 1, 3, 7, 3, 4,
		},
	)
	return s
}

func createTestCategoricalSeries() *Series {
	s := NewCategoricalSeries(
		"MySeries",
		[]string{
			"a", "a", "b", "c", "b", "b", "d", "a", "c", "a",
		},
	)
	return s
}

func TestNewSeries(t *testing.T) {
	s := createTestSeries()

	assert.Equal(t, "MySeries", s.Name, "column name is not correct")
	assert.Equal(t, false, s.IsCategorical(), "column is categorical")
	assert.Equal(t, 10, len(s.Values), "wrong number of values")
}

func TestNewCategoricalSeries(t *testing.T) {
	s := createTestCategoricalSeries()

	assert.Equal(t, "MySeries", s.Name, "column name is not correct")
	assert.Equal(t, true, s.IsCategorical(), "column is not categorical")
	assert.Equal(t, 10, len(s.Values), "wrong number of values")
	assert.Equal(t, 4, len(s.CategoricalLabels), "wrong number of category labels")
	assert.Equal(t, 4, len(s.CategoricalValues), "wrong number of category values")
}

func TestSeriesSum(t *testing.T) {
	s := createTestSeries()
	assert.Equal(t, 32.0, s.Sum(), "sum of series is not correct")
}

func TestSeriesMean(t *testing.T) {
	s := createTestSeries()
	assert.Equal(t, 3.2, s.Mean(), "sum of series is not correct")
}
