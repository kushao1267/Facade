package utils

// Image 图片类，常用属性
type Image struct {
	url    string
	height string
	width  string
	style  string
}

func GetBestImage(urls []string) *Image{
	return &Image{}
}
