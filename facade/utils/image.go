package utils

import (
	"image"
	"time"

	"github.com/levigross/grequests"
)

const (
	requestImgTimeout = 4 * time.Second
	// imageHeightRule 图片高度限制
	imageHeightRule = 170
	// imageWidthRule 图片宽度限制
	imageWidthRule = 340
)

// Image 图片类，常用属性
type Image struct {
	url    string
	height int
	width  int
	style  int
}

// ImageEntity 图片实体
type ImageEntity struct {
	url      string
	height   int
	width    int
	ocuppies int
}

// Height image height
func (i ImageEntity) Height() int {
	return i.height
}

// Width image width
func (i ImageEntity) Width() int {
	return i.width
}

// Ocuppies image ocuppies
func (i ImageEntity) Ocuppies() int {
	return i.ocuppies
}

func (i *ImageEntity) getImage() {
	resp, err := grequests.Get(i.url, &grequests.RequestOptions{
		RequestTimeout: requestImgTimeout,
	})

	imgConfig, _, err1 := image.DecodeConfig(resp)

	if err != nil || err1 != nil { // 请求失败 or 解析图片失败
		i.width, i.height, i.ocuppies = 0, 0, 0
		return
	}

	i.width, i.height = imgConfig.Width, imgConfig.Height
	i.ocuppies = len(resp.Bytes()) / 1024 // 获取占用空间大小，单位KB
}

func (i ImageEntity) isSatisfySize() bool {
	if i.height > imageHeightRule && i.width > imageWidthRule {
		return true
	}
	return false
}

func (i ImageEntity) isSatisfyOccupies() {}

// GetBestImage 获取最优图片
func GetBestImage(urls []string) *Image {
	return &Image{}
}
