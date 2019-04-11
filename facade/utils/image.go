package utils

import (
	"github.com/levigross/grequests"
	"image"
	"time"
)

const (
	requestImgTimeout = 4 * time.Second
	ImageHeightRule   = 170
	ImageWidthRule    = 340
)

// Image 图片类，常用属性
type Image struct {
	url    string
	height int
	width  int
	style  int
}

type ImageEntity struct {
	Url      string
	height   int
	width    int
	ocuppies int
}

// modify this three field value is not allowed
func (i ImageEntity) Height() int {
	return i.height
}

func (i ImageEntity) Width() int {
	return i.width
}

func (i ImageEntity) Ocuppies() int {
	return i.ocuppies
}

func (i *ImageEntity) getImage() {
	resp, err := grequests.Get(i.Url, &grequests.RequestOptions{
		RequestTimeout: requestImgTimeout,
	})

	imgConfig, _, err1 := image.DecodeConfig(resp)

	if err != nil || err1 != nil { //请求失败 or 解析图片失败
		i.width, i.height, i.ocuppies = 0, 0, 0
		return
	}

	i.width, i.height = imgConfig.Width, imgConfig.Height
	i.ocuppies = len(resp.Bytes()) / 1024 // 获取占用空间大小，单位KB
}

func (i ImageEntity) isSatisfySize() bool {
	if i.height > ImageHeightRule && i.width > ImageWidthRule {
		return true
	}
	return false
}

func (i ImageEntity) isSatisfyOccupies() {}

func GetBestImage(urls []string) *Image {
	return &Image{}
}
