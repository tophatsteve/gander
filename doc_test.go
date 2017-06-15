package gander

import (
	"fmt"
	"log"
)

func ExampleLoadCSVFromPath() {
	df, err := LoadCSVFromPath("testdata/MOCK_DATA.csv")
	if err != nil {
		log.Panic(err)
	}
	fmt.Printf("%v\n", df.Columns())
	// Output: 6
}

func ExampleLoadCSVFromURL() {
	df, err := LoadCSVFromURL("http://download.tensorflow.org/data/iris_training.csv")
	if err != nil {
		log.Panic(err)
	}
	fmt.Printf("%v\n", df.Columns())
	// Output: 5
}

func ExampleNewDataFrame() {
	df, _ := NewDataFrame(
		[][]string{
			{"a", "b", "c", "d", "e"},
			{"1", "2", "3", "4", "5"},
			{"3", "5", "2", "2", "4"},
			{"7", "6", "1", "3", "3"},
			{"4", "2", "4", "7", "6"},
		})
	fmt.Printf("%v\n", df.Rows())
	// Output: 4
}

func ExampleDataFrame_DropColumns() {
	df, _ := NewDataFrame(
		[][]string{
			{"a", "b", "c", "d", "e"},
			{"1", "2", "3", "4", "5"},
			{"3", "5", "2", "2", "4"},
			{"7", "6", "1", "3", "3"},
			{"4", "2", "4", "7", "6"},
		})
	df.DropColumns(0, 2)
	fmt.Printf("%v\n", df.Columns())
	// Output: 3
}

func ExampleDataFrame_DropColumnsByName() {
	df, _ := NewDataFrame(
		[][]string{
			{"a", "b", "c", "d", "e"},
			{"1", "2", "3", "4", "5"},
			{"3", "5", "2", "2", "4"},
			{"7", "6", "1", "3", "3"},
			{"4", "2", "4", "7", "6"},
		})
	df.DropColumnsByName("b", "d", "e")
	fmt.Printf("%v\n", df.Columns())
	// Output: 2
}

func ExampleDataFrame_DropColumnsWhere() {
	df, _ := NewDataFrame(
		[][]string{
			{"a", "b", "c", "d", "e"},
			{"1", "2", "3", "4", "5"},
			{"3", "5", "2", "2", "4"},
			{"7", "6", "1", "3", "3"},
			{"4", "2", "4", "7", "6"},
		})
	df.DropColumnsWhere(func(s *Series) bool {
		return s.Name == "c"
	})
	fmt.Printf("%v\n", df.Columns())
	// Output: 4
}
