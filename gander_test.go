package gander

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
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

func TestLoadCSVFromUrl(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "a,b,c,d,e")
		fmt.Fprintln(w, "1,3,7,9,2")
		fmt.Fprintln(w, "2,4,1,6,3")
	}))
	defer ts.Close()

	df, err := LoadCSVFromURL(ts.URL)
	assert.Equal(t, nil, err, "error is not nil")
	assert.Equal(t, 5, len(*df), "dataframe does not have the correct number of columns")
	assert.Equal(t, 2, df.Rows(), "dataframe does not have the correct number of rows")
}

func TestLoadCSVFromInvalidUrl(t *testing.T) {
	_, err := LoadCSVFromURL("http://doesnotexist.xyz.co/missing.csv")
	assert.Equal(t, true, strings.Contains(err.Error(), "no such host"), "error is not 'no such host'")
}

func TestLoadCSVFromPath(t *testing.T) {
	df, err := LoadCSVFromPath("./testdata/MOCK_DATA.csv")
	assert.Equal(t, nil, err, "error is not nil")
	assert.Equal(t, 6, len(*df), "dataframe does not have the correct number of columns")
	assert.Equal(t, 1000, df.Rows(), "dataframe does not have the correct number of rows")
}

func TestLoadCSVFromInvalidPath(t *testing.T) {
	_, err := LoadCSVFromPath("./testdata/MISSING.csv")
	assert.Equal(t, true, os.IsNotExist(err), "error is not 'does not exist'")
}
