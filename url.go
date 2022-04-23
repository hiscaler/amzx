package amzx

import (
	"fmt"
	"strings"
)

// Site 根据国家获取站点地址
func Site(country string) (url string, err error) {
	switch country {
	case CountryCA:
		url = "https://www.amazon.ca"
	case CountryGB:
		url = "https://www.amazon.co.uk"
	case CountryIT:
		url = "https://www.amazon.it"
	case CountryDE:
		url = "https://www.amazon.de"
	case CountryFR:
		url = "https://www.amazon.fr"
	case CountryES:
		url = "https://www.amazon.es"
	case CountryUS:
		url = "https://www.amazon.com"
	default:
		err = fmt.Errorf("invalid country: %s", country)
	}
	return
}

// URL 根据 asin、国家、附加参数生成 asin 完整地址
func URL(asin, country string, queries ...string) (url string, err error) {
	if !ValidAsin(asin) {
		err = fmt.Errorf("invalid asin: %s", asin)
		return
	}

	url, err = Site(country)
	if err != nil {
		return
	}

	sb := strings.Builder{}
	sb.WriteString(url)
	sb.WriteString("/dp/")
	sb.WriteString(asin)
	n := len(queries)
	if n > 0 {
		sb.WriteRune('?')
		for i, v := range queries {
			v = strings.TrimSpace(v)
			if v != "" {
				switch v[0:1] {
				case "?", "&":
					v = v[1:]
				}
				sb.WriteString(v)
			}
			if n != i+1 {
				sb.WriteRune('&')
			}
		}
	}
	url = sb.String()
	return
}
