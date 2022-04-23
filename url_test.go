package amzx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestURL(t *testing.T) {
	testCases := []struct {
		tag      string
		asin     string
		country  string
		queries  []string
		expected string
		hasError bool
	}{
		{"t1", "B0", CountryUS, []string{}, "", true},
		{"t2.1", "B092M62439", CountryUS, []string{}, "https://www.amazon.com/dp/B092M62439", false},
		{"t2.2", "B092M62439", CountryUS, nil, "https://www.amazon.com/dp/B092M62439", false},
		{"t2.3", "B092M62439", CountryUS, []string{"a=b"}, "https://www.amazon.com/dp/B092M62439?a=b", false},
		{"t2.3", "B092M62439", CountryUS, []string{"?a=b"}, "https://www.amazon.com/dp/B092M62439?a=b", false},
	}
	for _, testCase := range testCases {
		url, err := URL(testCase.asin, testCase.country, testCase.queries...)
		assert.Equal(t, testCase.expected, url, testCase.tag)
		assert.Equal(t, testCase.hasError, err != nil, testCase.tag)
		if !testCase.hasError && err != nil {
			t.Logf("Error: %s", err.Error())
		}
	}
}
