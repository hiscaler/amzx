package amzx

import "testing"

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
