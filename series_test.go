package gander

func createSampleSeries() *Series {
	s := Series{
		"Column 1",
		[]float64{
			0, 2, 7, 1, 4, 1, 3, 7, 3, 4,
		},
		nil,
		nil,
	}

	return &s
}
