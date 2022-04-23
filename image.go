package amzx

import (
	"path"
	"strings"
)

// OriginalImage 获取原图地址
// 该函数不会判断图片地址有效性，只会简单的将亚马逊正常的缩略图地址变成原图地址，比如：
// 缩略图：https://images-na.ssl-images-amazon.com/images/I/41rjwBnXx5L._SY300_SX300_QL70_FMwebp_.jpg
// 原图：https://images-na.ssl-images-amazon.com/images/I/41rjwBnXx5L.jpg
func OriginalImage(url string) string {
	url = strings.TrimSpace(url)
	n := len(url)
	if n != 0 && url[n-2:n-1] != "/" {
		name := path.Base(url)
		firstIndex := strings.Index(name, ".")
		lastIndex := strings.LastIndex(name, ".")
		if firstIndex == lastIndex {
			// "." 不存在或者只有一个
			return url
		}
		newName := name[0:firstIndex] + name[lastIndex:]
		return strings.ReplaceAll(url, name, newName)
	}
	return url
}
