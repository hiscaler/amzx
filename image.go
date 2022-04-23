package amzx

import (
	"fmt"
	"net/url"
	"strings"
)

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
