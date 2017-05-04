package gander

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

func (s *Series) Sum() float64 {
	t := 0.0

	for _, v := range s.Values {
		t += v
	}

	return t
}

func (s *Series) Mean() float64 {
	return s.Sum() / float64(len(s.Values))
}

func (s *Series) Median() float64 {
	return 0
}

func (s *Series) Mode() float64 {
	return 0
}

func (s *Series) Variance() float64 {
	return 0
}

func (s *Series) StdDev() float64 {
	return 0
}

func (s *Series) IsCategorical() bool {
	return s.CategoricalLabels != nil
}
