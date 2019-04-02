package handler

import (
	"github.com/kushao1267/facade/facade/config"
	"github.com/kushao1267/facade/facade/utils"
)

const (
	title_field         = "title"
	image_field         = "image"
)

type ReturnData struct {
	// meta field
	Error       error
	Exception 	string
	Url         string
	Hostname    string
	ShouldCache bool
	// return field
	Title       string
	Description string
	Images      []string

	Bestimage  *utils.Image
}

func NewReturnData(hostname, exception, title, description string, err error,images []string) *ReturnData{
	r := &ReturnData{}
	r.Error = err

	if exception != ""{
		r.ShouldCache = false
	}else {
		r.ShouldCache = true
	}

	if hasDefault(hostname){
		r.Exception = exception
	}

	if len(description) > 125{
		r.Description = description[:125] + "..."
	}

	if title != ""{
		r.Title = title
	}else {
		r.Title = getDefault(hostname, title_field)
	}

	if isForceDefault(hostname) || len(r.Images) == 0{
		r.Images = append(r.Images, getDefault(hostname, image_field))
	}else {
		r.Images = images
	}

	r.Bestimage = utils.GetBestImage(r.Images)

	return r
}

func hasDefault(hostname string) bool{
	if _,ok := config.AllConf.ReturnMap[hostname]; ok {
		return true
	}
	return false
}

func isForceDefault(hostname string) bool{
	if val, ok := config.AllConf.ReturnMap[hostname]; ok {
		return val.ForceDefault
	}
	return false
}

func getDefault(hostname ,field string) string{
	if val, ok := config.AllConf.ReturnMap[hostname]; ok {
		if field == title_field{
			return val.Title
		}else if field == image_field{
			return val.Image
		}
	}

	return ""
}

// Handle 封装共用的requests方法以及显示声明handle,extract等接口
type Handle interface {
	Requests() //请求方法
	Extract()  // 提取方法
	Handle()   // 处理方法，是handle的出口
}
