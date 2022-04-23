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
	if len(queries) > 0 {
		sb.WriteRune('?')
		for _, v := range queries {
			v = strings.TrimSpace(v)
			i := len(v)
			if i == 0 || v == "&" || v == "?" {
				continue
			}
			firstChar := v[0:1]
			if firstChar == "?" || firstChar == "&" {
				v = v[1:]
				i--
			}
			lastChar := v[i-1:]
			if lastChar == "?" || lastChar == "&" {
				v = v[0 : i-1]
			}
			sb.WriteString(v)
			sb.WriteRune('&')
		}
		url = sb.String()
		url = url[0 : len(url)-1]
	} else {
		url = sb.String()
	}
	return
}
