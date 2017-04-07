package gander

type Series struct {
	Name   string
	Values []float64
}

type DataFrame []Series

func ReadCSV(path string) DataFrame {
	return DataFrame{}
}

func (d DataFrame) Columns() []string {
	return []string{}
}
