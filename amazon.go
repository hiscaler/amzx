package amzx

import (
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

// ValidAsin 是否为有效的 asin
func ValidAsin(asin string) bool {
	if asin == "" || len(asin) != 10 || !strings.HasPrefix(strings.ToUpper(asin), "B0") {
		return false
	}
	return true
}
