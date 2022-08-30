package amzx

import (
	"fmt"
	"os"
	"testing"
)

func TestSellerParse(t *testing.T) {
	testCases := []struct {
		File     string
		Excepted Seller
	}{
		{"shopcn1.html", Seller{
			Information: SellerInformation{},
			Feedback:    SellerFeedback{},
		}},
		{"shopcn2.html", Seller{
			Information: SellerInformation{},
			Feedback:    SellerFeedback{},
		}},
		{"shopus1.html", Seller{
			Information: SellerInformation{},
			Feedback:    SellerFeedback{},
		}},
	}

	for _, testCase := range testCases {
		s := Seller{}
		html, err := os.ReadFile("./testdata/" + testCase.File)
		if err != nil {
			t.Errorf("Read %s file error: %s", testCase.File, err.Error())
		}
		ss, err := s.Parse(string(html))
		if err != nil {
			t.Errorf("%s parse error: %s", testCase.File, err.Error())
		}
		fmt.Println(ss.Information.FullAddress())
	}
}
