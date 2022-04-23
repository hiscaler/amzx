package amzx

import (
	"fmt"
	"net/url"
	"strings"
)

const (
	CountryCN = "CN" // 中国
	CountryUS = "US" // 美国
	CountryCA = "CA" // 加拿大
	CountryDE = "DE" // 德国
	CountryGB = "GB" // 英国
	CountryFR = "FR" // 法国
	CountryES = "ES" // 西班牙
	CountryIT = "IT" // 意大利
	CountryJP = "JP" // 日本
	CountryMX = "MX" // 墨西哥
	CountryAU = "AU" // 澳大利亚
	CountryIN = "IN" // 印度
	CountryAE = "AE" // 阿联酋
	CountryTR = "TR" // 土耳其
	CountrySG = "SG" // 新加坡
	CountryNL = "NL" // 荷兰
	CountryBR = "BR" // 巴西
	CountrySA = "SA" // 沙特阿拉伯
	CountrySE = "SE" // 瑞典
	CountryPL = "PL" // 波兰
)

func ValidAsin(key string) bool {
	if key == "" || len(key) != 10 || !strings.HasPrefix(strings.ToUpper(key), "B0") {
		return false
	}
	return true
}

// OriginalImage 获取原图地址
func OriginalImage(imagePath string) string {
	imagePath = strings.TrimSpace(imagePath)
	if imagePath != "" {
		u, err := url.Parse(imagePath)
		if err == nil && u.Path != "" && strings.Index(u.Path, "/") != -1 {
			paths := strings.Split(u.Path, "/")
			filename := paths[len(paths)-1]
			if strings.Index(filename, ".") != -1 {
				pair := strings.Split(filename, ".")
				if len(pair) >= 2 {
					imagePath = strings.ReplaceAll(imagePath, filename, fmt.Sprintf("%s.%s", pair[0], pair[len(pair)-1]))
				}
			}
		}
	}
	return imagePath
}

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
