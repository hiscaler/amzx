package amzx

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestSellerParse(t *testing.T) {
	testCases := []struct {
		File     string
		Excepted Seller
	}{
		{"shopcn1.html", Seller{
			Information: SellerInformation{
				Name:     "qingdaojingyixinjiajuyouxiangongsi",
				Province: "shandongsheng",
				City:     "licangqu",
				Area:     "qingdaoshi",
				Address:  "longshuilu318hao1lou251shi",
				Postcode: "276000",
				Country:  "CN",
			},
			Feedback: SellerFeedback{
				ThirtyDays:   Feedback{75, 0, 25, 8},
				NinetyDays:   Feedback{80, 0, 20, 15},
				TwelveMonths: Feedback{90, 0, 10, 31},
				Lifetime:     Feedback{87, 3, 11, 38},
			},
		}},
		{"shopcn2.html", Seller{
			Information: SellerInformation{
				Name:     "chen min",
				Province: "广东省",
				City:     "深圳市",
				Area:     "龙岗区",
				Address:  "平湖街道华南城2号交易广场5B-020号",
				Postcode: "518000",
				Country:  "CN",
			},
			Feedback: SellerFeedback{
				ThirtyDays:   Feedback{0, 0, 0, 0},
				NinetyDays:   Feedback{0, 0, 0, 0},
				TwelveMonths: Feedback{86, 0, 14, 7},
				Lifetime:     Feedback{89, 1, 10, 291},
			},
		}},
		{"shopus1.html", Seller{
			Information: SellerInformation{
				Name:     "GK GRAND LLC",
				Province: "Florida",
				City:     "ST AUGUSTINE",
				Area:     "",
				Address:  "202 Marshall Circle",
				Postcode: "32086",
				Country:  "US",
			},
			Feedback: SellerFeedback{
				ThirtyDays:   Feedback{100, 0, 0, 2},
				NinetyDays:   Feedback{100, 0, 0, 14},
				TwelveMonths: Feedback{97, 1, 1, 76},
				Lifetime:     Feedback{99, 0, 1, 0},
			},
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
		assert.Equal(t, testCase.Excepted, *ss, testCase.File)
	}
}
