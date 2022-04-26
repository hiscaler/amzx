package amzx

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"github.com/disintegration/imaging"
	"github.com/otiai10/gosseract/v2"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"math/big"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

// https://www.amazon.com/errors/validateCaptcha

var (
	re  = regexp.MustCompile("[0-9]+")
	re2 = regexp.MustCompile("[A-Z]")
)

func Captcha(imagePath string) (text string, err error) {
	var imageExt string
	isUrl := false
	u, err := url.Parse(imagePath)
	isUrl = err == nil && u.Scheme != "" && u.Host != ""
	if isUrl {
		var response *http.Response
		response, err = http.Get(imagePath)
		if err != nil {
			return
		}
		var b []byte
		b, err = ioutil.ReadAll(response.Body)
		if err != nil {
			return
		}
		defer func() {
			err = response.Body.Close()
		}()

		switch http.DetectContentType(b) {
		case "image/jpeg":
			imageExt = "jpg"
		default:
			imageExt = filepath.Ext(imagePath)
		}

		fnUniqueFileName := func() string {
			str := "1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
			s := ""
			b := bytes.NewBufferString(str)
			bigInt := big.NewInt(int64(b.Len()))
			for i := 0; i < 12; i++ {
				randomInt, _ := rand.Int(rand.Reader, bigInt)
				s += string(str[randomInt.Int64()])
			}
			return s
		}
		imagePath = fmt.Sprintf("%s/%s.%s", os.TempDir(), fnUniqueFileName(), imageExt)
		err = os.WriteFile(imagePath, b, 0666)
		if err != nil {
			return
		}
		defer os.Remove(imagePath)
	} else {
		// Is local file
		imageExt = filepath.Ext(imagePath)
		if imageExt != "" {
			imageExt = imageExt[1:]
		}
	}

	var imageFile *os.File
	imageFile, err = os.Open(imagePath)
	if err != nil {
		return
	}
	defer imageFile.Close()

	var img image.Image
	if imageExt != "" {
		imageExt = strings.ToLower(imageExt)
	}
	switch imageExt {
	case "jpg", "jpeg":
		img, _ = jpeg.Decode(imageFile)
	case "png":
		img, _ = png.Decode(imageFile)
	default:
		return "", fmt.Errorf("unknown image format")
	}

	// use image histogram and threshold to separate 6 letter
	x1 := img.Bounds().Min.X
	x2 := img.Bounds().Max.X
	y1 := img.Bounds().Min.Y
	y2 := img.Bounds().Max.Y
	width := x2 - x1
	height := y2 - y1

	columnMean := make([]int, 0)
	for i := 0; i < width; i++ {
		total := 0
		for j := 0; j < height; j++ {
			colorVal := img.At(i, j)
			colorValStr := fmt.Sprintf("%v", colorVal)
			cVal, _ := strconv.Atoi(re.FindAllString(colorValStr, 1)[0])
			total += cVal
		}
		mean := total / height
		columnMean = append(columnMean, mean)
	}

	retryCount := 0
	threshold := 241
	separatorIndex := make([]int, 0)
	for len(separatorIndex) != 7 {
		retryCount += 1
		if threshold >= 255 || threshold <= 0 || retryCount >= 15 {
			return "", fmt.Errorf("can not crop correct number of letters")
		} else if len(separatorIndex) > 7 {
			threshold += 1
		} else if len(separatorIndex) < 7 {
			threshold -= 1
		}

		colMeanIndex := make([]int, 0)
		for i, colMean := range columnMean {
			if colMean < threshold {
				colMeanIndex = append(colMeanIndex, i)
			}
		}
		l := len(colMeanIndex)
		colMeanMinIndex := colMeanIndex[0]
		colMeanMaxIndex := colMeanIndex[l-1]

		separatorIndex = []int{colMeanMinIndex}
		for i, val := range colMeanIndex[:l-1] {
			if val != colMeanIndex[i+1]-1 {
				separatorIndex = append(separatorIndex, val)
			}
		}
		separatorIndex = append(separatorIndex, colMeanMaxIndex)
	}

	// rotate each character
	client := gosseract.NewClient()
	defer client.Close()
	for i, val := range separatorIndex[:6] {
		rotateAngle := 0
		if i%2 == 0 {
			rotateAngle = 15
		} else {
			rotateAngle = -15
		}

		// crop and rotate image
		croppedImage := imaging.Crop(img, image.Rect(val+1, 0, separatorIndex[i+1], height))
		rotatedImage := imaging.Rotate(croppedImage, float64(rotateAngle), color.RGBA{
			R: 255,
			G: 255,
			B: 255,
			A: 255,
		})

		buf := new(bytes.Buffer)
		switch imageExt {
		case "jpg", "jpeg":
			e := jpeg.Encode(buf, rotatedImage, nil)
			if e != nil {
				err = fmt.Errorf("jpeg.Encode() error: %s", e.Error())
			}
		case "png":
			e := png.Encode(buf, rotatedImage)
			if e != nil {
				err = fmt.Errorf("png.Encode() error: %s", e.Error())
			}
		}

		e := client.SetImageFromBytes(buf.Bytes())
		if e != nil {
			err = fmt.Errorf("client.SetImageFromBytes() error: %s", e.Error())
		}
		char, e := client.Text()
		if e != nil {
			err = fmt.Errorf("client.Text() error: %s", e.Error())
		}

		if char == "" {
			return "", fmt.Errorf("got empty letter QQ")
		} else if len(char) > 1 {
			charArr := strings.Split(char, "")
			for _, ca := range charArr {
				if re2.Match([]byte(ca)) {
					text += ca
					break
				}

			}
		} else if !re2.Match([]byte(char)) {
			return "", fmt.Errorf("no english letter %v", char)
		} else {
			text += char
		}

	}
	return
}
