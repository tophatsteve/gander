package gander

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDetectHeaderRowPresent(t *testing.T) {
	b := hasHeaderRow(createSampleDataWithHeaders())
	assert.Equal(t, true, b, "header row not detected")
}

func TestDetectHeaderRowAbsent(t *testing.T) {
	b := hasHeaderRow(createSampleDataWithoutHeaders())
	assert.Equal(t, false, b, "header row detected incorrectly")
}

func TestDetectHeaderRowAbsentWithMixedHeaders(t *testing.T) {
	b := hasHeaderRow(createSampleDataWithMixedHeaders())
	assert.Equal(t, false, b, "header row not detected")
}
