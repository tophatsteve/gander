package gander

import (
	"fmt"
	"math"
	"sort"
)

// A Series represents a column of data in a DataFrame.
type Series struct {
	Name              string
	Values            []float64
	categoricalLabels map[float64]string
	categoricalValues map[string]float64
}

// NewSeries creates a new Series with the specified name
// and values.
func NewSeries(name string, values []float64) *Series {
	s := Series{}
	s.Name = name
	s.Values = []float64{}

	for _, v := range values {
		s.Values = append(s.Values, v)
	}

	return &s
}

// NewCategoricalSeries create a new Series to contain categorical
// data. The data is passed in as a slice of strings. Internally
// the string values are converted to float64 and a map is maintained
// so that the original values can always be retrieved. No statistical
// operations can be carried out on a categorical series.
func NewCategoricalSeries(name string, values []string) *Series {
	categoryNumber := 0.0
	s := Series{}
	s.categoricalLabels = make(map[float64]string)
	s.categoricalValues = make(map[string]float64)
	s.Name = name

	s.Values = []float64{}

	for _, v := range values {
		if i, ok := s.categoricalValues[v]; ok == true {
			s.Values = append(s.Values, i)
		} else {
			s.Values = append(s.Values, categoryNumber)
			s.categoricalValues[v] = categoryNumber
			s.categoricalLabels[categoryNumber] = v
			categoryNumber += 1
		}
	}

	return &s
}

// Standardize scales the values in the Series
// to standard form.
func (s *Series) Standardize() {
	mu := s.Mean()
	sigma := s.StdDev()

	for i, v := range s.Values {
		s.Values[i] = (v - mu) / sigma
	}
}

// Sum adds together all the values in the Series.
func (s *Series) Sum() float64 {
	return sum(s.Values)
}

// Mean finds the mean of all the values in the Series.
func (s *Series) Mean() float64 {
	return s.Sum() / float64(len(s.Values))
}

// Median finds the median of all the values in the Series.
func (s *Series) Median() float64 {
	v := s.Sorted()

	if len(v)%2 == 0 {
		return (v[(len(v)/2)-1] + v[len(v)/2]) / 2
	}

	return v[(len(v) / 2)]
}

// Mode finds the mode of all the values in the Series. This returns
// a slice ofr float64 because a Series could have more than one mode.
func (s *Series) Mode() []float64 {
	m := []float64{}
	c := count(s.Values)

	var maxCount int

	for _, v := range c {
		if v > maxCount {
			maxCount = v
		}
	}

	for k := range c {
		if c[k] == maxCount {
			m = append(m, k)
		}
	}

	return m
}

// Variance finds the variance of the values in the Series.
func (s *Series) Variance() float64 {
	mu := s.Mean()
	sumOfSquares := sum(
		s.Apply(
			func(x float64) float64 {
				return math.Pow(x-mu, 2)
			}))

	return sumOfSquares / float64(len(s.Values))
}

// StdDev finds the standard deviation of the values in the Series.
func (s *Series) StdDev() float64 {
	return math.Sqrt(s.Variance())
}

func (s *Series) IsCategorical() bool {
	return s.categoricalLabels != nil
}

// Max returns the maximum value in the Series.
func (s *Series) Max() float64 {
	v := s.Sorted()
	return v[len(s.Values)-1]
}

// Min returns the minimum value in the Series.
func (s *Series) Min() float64 {
	v := s.Sorted()
	return v[0]
}

// Range returns the minimum and maximum values in the Series.
func (s *Series) Range() (float64, float64) {
	return s.Min(), s.Max()
}

// Apply applies a function to all values of the Series.
// It returns a new slice and does not affect the Series values.
func (s *Series) Apply(fn func(float64) float64) []float64 {
	r := []float64{}

	for _, v := range s.Values {
		r = append(r, fn(v))
	}

	return r
}

// Transform applies a function to all values of the Series,
// changing them in place.
func (s *Series) Transform(fn func(float64) float64) {
	for i, v := range s.Values {
		s.Values[i] = fn(v)
	}
}

// Sorted returns a slice of the sorted values in a Series.
// It does not change the values of the Series itself.
func (s *Series) Sorted() []float64 {
	r := make([]float64, len(s.Values))
	copy(r, s.Values)
	sort.Float64s(r)
	return r
}

// Hist returns a map of values to counts for categorical data.
// It returns an error is the Series does not contain categorical data.
func (s *Series) Hist() (map[string]int, error) {
	if s.IsCategorical() == false {
		return nil, fmt.Errorf("Series %s is not categorical", s.Name)
	}

	r := make(map[string]int)

	for _, v := range s.Values {
		c := s.categoricalLabels[v]
		if _, ok := r[c]; ok {
			r[c] += 1
		} else {
			r[c] = 1
		}
	}

	return r, nil
}

// Describe returns a summary of the statisical properties
// of all the Series.
func (s *Series) Describe() Summary {
	r := Summary{Name: s.Name}
	r.Mean = s.Mean()
	r.Median = s.Median()
	r.Mode = s.Mode()
	r.Min = s.Min()
	r.Max = s.Max()
	r.StdDev = s.StdDev()
	r.Variance = s.Variance()

	return r
}

func sum(r []float64) float64 {
	t := 0.0

	for _, v := range r {
		t += v
	}

	return t
}

func count(r []float64) map[float64]int {
	m := map[float64]int{}

	for _, v := range r {
		if _, ok := m[v]; ok {
			m[v] += 1
		} else {
			m[v] = 1
		}
	}

	return m
}

func (s *Series) dropRow(r int) {
	s.Values = append(s.Values[:r], s.Values[r+1:]...)
}
