package gander

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func createSampleDataWithHeaders() [][]string {
	return [][]string{
		{"a", "b", "c", "d", "e"},
		{"1", "2", "3", "4", "5"},
		{"3", "5", "2", "2", "4"},
		{"7", "6", "1", "3", "3"},
		{"4", "2", "4", "7", "6"},
	}
}

func createSampleDataWithoutHeaders() [][]string {
	return [][]string{
		{"1", "2", "3", "4", "5"},
		{"3", "5", "2", "2", "4"},
		{"7", "6", "1", "3", "3"},
		{"4", "2", "4", "7", "6"},
	}
}

func createSampleDataWithCategoricalData() [][]string {
	return [][]string{
		{"a", "b", "c", "d", "e"},
		{"1", "2", "3", "a", "5"},
		{"3", "5", "2", "b", "4"},
		{"7", "6", "1", "b", "3"},
		{"4", "2", "4", "a", "6"},
	}
}

func createSampleDataWithMixedHeaders() [][]string {
	return [][]string{
		{"1", "2", "c", "3", "e"},
		{"1", "2", "3", "4", "5"},
		{"3", "5", "2", "2", "4"},
		{"7", "6", "1", "3", "3"},
		{"4", "2", "4", "7", "6"},
	}
}

func createLargerSampleData() [][]string {
	return [][]string{
		{"a", "b", "c", "d", "e"},
		{"1", "2", "3", "4", "5"},
		{"3", "5", "2", "2", "4"},
		{"7", "6", "1", "3", "3"},
		{"4", "2", "4", "7", "6"},
		{"1", "2", "3", "4", "5"},
		{"3", "5", "2", "2", "4"},
		{"7", "6", "1", "3", "3"},
		{"4", "2", "4", "7", "6"},
		{"1", "2", "3", "4", "5"},
		{"3", "5", "2", "2", "4"},
		{"7", "6", "1", "3", "3"},
		{"4", "2", "4", "7", "6"},
		{"1", "2", "3", "4", "5"},
		{"3", "5", "2", "2", "4"},
		{"7", "6", "1", "3", "3"},
		{"4", "2", "4", "7", "6"},
	}
}

func TestLoadCSVFromFile(t *testing.T) {
	assert.Equal(t, true, false, "TestLoadCSVFromFile not implemented")
}

func TestLoadCSVFromUrl(t *testing.T) {
	assert.Equal(t, true, false, "TestLoadCSVFromUrl not implemented")
}
