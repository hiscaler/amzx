package amzx

import (
	"errors"
	"github.com/PuerkitoBio/goquery"
	"strconv"
	"strings"
)

// 解析卖家资料

type Seller struct {
	Information SellerInformation `json:"information"`
	Feedback    SellerFeedback    `json:"feedback"`
}

type SellerInformation struct {
	Name     string `json:"name"`
	Country  string `json:"country"`
	Province string `json:"province"`
	City     string `json:"city"`
	Area     string `json:"area"`
	Address  string `json:"address"`
	Postcode string `json:"postcode"`
}

type Feedback struct {
	Positive float64 `json:"positive"` // 好评百分比
	Neutral  float64 `json:"neutral"`  // 中评百分比
	Negative float64 `json:"negative"` // 差评百分比
	Count    int     `json:"count"`    // 数量
}

type SellerFeedback struct {
	ThirtyDays   Feedback `json:"thirty_days"`   // 30 天
	NinetyDays   Feedback `json:"ninety_days"`   // 90 天
	TwelveMonths Feedback `json:"twelve_months"` // 12 个月
	Lifetime     Feedback `json:"lifetime"`      // 所有
}

func (s SellerInformation) FullAddress() string {
	values := make([]string, 0)
	if s.Province != "" {
		values = append(values, s.Province)
	}
	if s.City != "" {
		values = append(values, s.City)
	}
	if s.Area != "" {
		values = append(values, s.Area)
	}
	if s.Address != "" {
		values = append(values, s.Address)
	}
	sep := ""
	switch strings.ToUpper(s.Country) {
	case CountryCN:
		sep = ""
	default:
		sep = " "
	}
	return strings.Join(values, sep)
}

func SellerParse(html string) (s Seller, err error) {
	if html == "" {
		return s, errors.New("html is empty")
	}

	var doc *goquery.Document
	doc, err = goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return
	}

	// Information
	information := SellerInformation{}
	addresses := make([]string, 0)
	businessAddressIndex := -1
	doc.Find("#page-section-detail-seller-info .a-row").Each(func(i int, selection *goquery.Selection) {
		if i != 0 {
			value := strings.TrimSpace(selection.Text())
			switch i {
			case 1:
				value = strings.NewReplacer("公司名称:", "", "Business Name:", "").Replace(value)
				information.Name = strings.TrimSpace(value)
			default:
				if businessAddressIndex == -1 && (value == "公司地址:" || value == "Business Address:") {
					businessAddressIndex = i + 1
				} else if businessAddressIndex != -1 {
					addresses = append(addresses, value)
				}
			}
		}
	})

	// Set information values
	if len(addresses) > 0 {
		country := addresses[len(addresses)-1]
		information.Country = country
		switch strings.ToUpper(country) {
		case CountryCN:
			for i, value := range addresses {
				switch i {
				case 0:
					information.Address = value
				case 1:
					information.City = value
				case 2:
					information.Area = value
				case 3:
					information.Province = value
				case 4:
					information.Postcode = value
				case 5:
					information.Country = value
				}
			}
		case CountryUS:
			for i, value := range addresses {
				switch i {
				case 0:
					information.Address = value
				case 1:
					information.City = value
				case 2:
					information.Province = value
				case 3:
					information.Postcode = value
				case 4:
					information.Country = value
				}
			}
		default:
			information.Address = addresses[0]
		}
	}
	s.Information = information

	// Feedback
	feedback := SellerFeedback{}
	replacer := strings.NewReplacer("%", "", "-", "", ",", "")
	doc.Find("#feedback-summary-table tr").Each(func(i int, selection *goquery.Selection) {
		if i != 0 {
			var v1, v2, v3, v4 float64
			selection.Find("td").Each(func(i int, selection *goquery.Selection) {
				value := replacer.Replace(strings.TrimSpace(selection.Text()))
				if value != "" {
					v, _ := strconv.ParseFloat(value, 64)
					switch i {
					case 1:
						v1 = v
					case 2:
						v2 = v
					case 3:
						v3 = v
					case 4:
						v4 = v
					}
				}
			})
			switch i {
			case 1:
				feedback.ThirtyDays.Positive = v1
				feedback.NinetyDays.Positive = v2
				feedback.TwelveMonths.Positive = v3
				feedback.Lifetime.Positive = v4
			case 2:
				feedback.ThirtyDays.Neutral = v1
				feedback.NinetyDays.Neutral = v2
				feedback.TwelveMonths.Neutral = v3
				feedback.Lifetime.Neutral = v4
			case 3:
				feedback.ThirtyDays.Negative = v1
				feedback.NinetyDays.Negative = v2
				feedback.TwelveMonths.Negative = v3
				feedback.Lifetime.Negative = v4
			case 4:
				feedback.ThirtyDays.Count = int(v1)
				feedback.NinetyDays.Count = int(v2)
				feedback.TwelveMonths.Count = int(v3)
				feedback.Lifetime.Count = int(v4)
			}
		}
	})
	s.Feedback = feedback

	return
}
