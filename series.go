package gander

import (
	"math"
	"sort"
)

type Series struct {
	Name              string
	Values            []float64
	CategoricalLabels map[float64]string
	CategoricalValues map[string]float64
}

func NewSeries(name string, values []float64) *Series {
	s := Series{}
	s.Name = name
	s.Values = []float64{}

	for _, v := range values {
		s.Values = append(s.Values, v)
	}

	return &s
}

func NewCategoricalSeries(name string, values []string) *Series {
	categoryNumber := 0.0
	s := Series{}
	s.CategoricalLabels = make(map[float64]string)
	s.CategoricalValues = make(map[string]float64)
	s.Name = name

	s.Values = []float64{}

	for _, v := range values {
		if i, ok := s.CategoricalValues[v]; ok == true {
			s.Values = append(s.Values, i)
		} else {
			s.Values = append(s.Values, categoryNumber)
			s.CategoricalValues[v] = categoryNumber
			s.CategoricalLabels[categoryNumber] = v
			categoryNumber += 1
		}
	}

	return &s
}

// Standardize scales the values in the Series
// to standard form.
func (s *Series) Standardize() {
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

func (s *Series) Mode() float64 {
	return 0
}

func (s *Series) Variance() float64 {
	sumOfSquares := sum(
		s.Apply(
			func(x float64) float64 {
				return math.Pow(x, 2)
			}))

	return sumOfSquares / float64(len(s.Values))
}

func (s *Series) StdDev() float64 {
	return math.Sqrt(s.Variance())
}

func (s *Series) IsCategorical() bool {
	return s.CategoricalLabels != nil
}

func (s *Series) Max() float64 {
	v := s.Sorted()
	return v[len(s.Values)-1]
}

func (s *Series) Min() float64 {
	v := s.Sorted()
	return v[0]
}

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

func sum(r []float64) float64 {
	t := 0.0

	for _, v := range r {
		t += v
	}

	return t
}
