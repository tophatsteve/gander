package gander

import (
	"fmt"
	"math"
	"testing"

	"errors"
	"github.com/stretchr/testify/assert"
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

func createOddTestSeries() *Series {
	s := NewSeries(
		"MySeries",
		[]float64{
			0, 2, 7, 1, 5, 5, 3, 7, 4,
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

func toleratedError(e, a float64) bool {
	if math.Abs(e-a) < 0.0000000001 {
		return true
	}
	return false
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

func TestSeriesMedianOddValues(t *testing.T) {
	s := createOddTestSeries() // 0, 1, 2, 3, 4, 5, 5, 7, 7
	assert.Equal(t, 4.0, s.Median(), "median is not correct")
}

func TestSeriesMedianEvenValues(t *testing.T) {
	s := createTestSeries()
	assert.Equal(t, 3.0, s.Median(), "median is not correct")
}

func TestSeriesMode(t *testing.T) {
	s := createTestSeries()
	m := s.Mode()
	assert.Equal(t, 4, len(m), "number of mode items is not correct")
}

func TestApplyFunctionChangesResult(t *testing.T) {
	s := createTestSeries()
	r := s.Apply(func(x float64) float64 {
		return 0
	})

	for _, v := range r {
		assert.Equal(t, 0.0, v, "function not applied to series")
	}
}

func TestTransformFunctionChangesValues(t *testing.T) {
	s := createTestSeries()
	s.Transform(func(x float64) float64 {
		return 0
	})

	for _, v := range s.Values {
		assert.Equal(t, 0.0, v, "function not applied to series")
	}
}

func TestApplyFunctionDoesNotChangeValues(t *testing.T) {
	s := createTestSeries()
	e := createTestSeries()
	s.Apply(func(x float64) float64 {
		return 0
	})

	for i, v := range s.Values {
		assert.Equal(t, e.Values[i], v, "function was applied to series")
	}
}

func TestSortedSortsResults(t *testing.T) {
	s := createTestSeries()
	r := s.Sorted()
	e := []float64{0, 1, 1, 2, 3, 3, 4, 4, 7, 7}

	for i, v := range e {
		assert.Equal(t, v, r[i], "values not sorted correctly")
	}
}

func TestSortedDoesNotChangeValues(t *testing.T) {
	s := createTestSeries()
	s.Sorted()
	e := []float64{0, 2, 7, 1, 4, 1, 3, 7, 3, 4}

	for i, v := range e {
		assert.Equal(t, v, s.Values[i], "values sorted inplace")
	}
}

func TestMaxReturnsMaximum(t *testing.T) {
	s := createTestSeries()
	v := s.Max()
	assert.Equal(t, 7.0, v, "max is not correct")
}

func TestMinReturnsMinimum(t *testing.T) {
	s := createTestSeries()
	v := s.Min()
	assert.Equal(t, 0.0, v, "min is not correct")
}

func TestVariance(t *testing.T) {
	s := createTestSeries()
	v := s.Variance()
	assert.Equal(t, 5.16, v, "variance is not correct")
}

func TestStdDev(t *testing.T) {
	s := createTestSeries()
	v := s.StdDev()
	assert.Equal(t, "2.27156", fmt.Sprintf("%.5f", v), "std dev is not correct")
}

func TestStandardizedValues(t *testing.T) {
	s := createTestSeries()
	s.Standardize()
	assert.Equal(t, true, toleratedError(0.0, s.Mean()), "mean is not zero(ish)")
	assert.Equal(t, 1.0, s.StdDev(), "std dev is not one")
}

func TestHistOfCategoricalData(t *testing.T) {
	s := createTestCategoricalSeries()
	c, err := s.Hist()
	assert.Equal(t, nil, err, "error is not nil")
	assert.Equal(t, 4, c["a"], "category a count is not correct")
	assert.Equal(t, 3, c["b"], "category b count is not correct")
	assert.Equal(t, 2, c["c"], "category c count is not correct")
	assert.Equal(t, 1, c["d"], "category d count is not correct")
}

func TestHistOfNonCategoricalData(t *testing.T) {
	s := createTestSeries()
	_, err := s.Hist()
	assert.Equal(t, errors.New("Series MySeries is not categorical"), err, "did not return correct error")
}
