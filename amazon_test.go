package amzx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOriginalImagePath(t *testing.T) {
	images := map[string]string{
		"https://images-na.ssl-images-amazon.com/images/I/41rjwBnXx5L._SY300_SX300_QL70_FMwebp_.jpg":     "https://images-na.ssl-images-amazon.com/images/I/41rjwBnXx5L.jpg",
		"https://images-na.ssl-images-amazon.com/images/I/51tHj5lVYTL.__AC_SX300_SY300_QL70_FMwebp_.jpg": "https://images-na.ssl-images-amazon.com/images/I/51tHj5lVYTL.jpg",
		"https://images-na.ssl-images-amazon.com/images/a.jpg":                                           "https://images-na.ssl-images-amazon.com/images/a.jpg",
		"https://images-na.ssl-images-amazon.com/images/":                                                "https://images-na.ssl-images-amazon.com/images/",
		"https://images-na.ssl-images-amazon.com/1.jpg":                                                  "https://images-na.ssl-images-amazon.com/1.jpg",
		"https://images-na.ssl-images-amazon.com/images/I/61S-VTI6XmL.__AC_SX300_SY300_QL70_FMwebp_.jpg": "https://images-na.ssl-images-amazon.com/images/I/61S-VTI6XmL.jpg",
		"https://m.media-amazon.com/images/I/61czNd06BPS._AC_UY218_.jpg":                                 "https://m.media-amazon.com/images/I/61czNd06BPS.jpg",
	}
	for k, v := range images {
		img := OriginalImage(k)
		if img != v {
			t.Errorf("%s => %s parse failed.", k, img)
		}
	}
}

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
