package controllers

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/kushao1267/Facade/facade/extractors"
	"github.com/kushao1267/Facade/facade/forms"
	"github.com/kushao1267/Facade/facade/services"
	"github.com/kushao1267/Facade/facade/techniques"
	"github.com/kushao1267/Facade/facade/utils"
	"github.com/mgutz/ansi"
)

const (
	// FailCode 失败状态码
	FailCode = "0"
	// SuccessCode 成功状态码
	SuccessCode = "1" // 成功状态码
)

// NewRespData 用于构造返回结果
func NewRespData(code, msg string, data interface{}) gin.H {
	var resp = gin.H{
		"code": code,
		"msg":  msg,
	}
	if data != nil {
		resp["data"] = data
	}
	return resp
}

// LinkController ...
type LinkController struct{}
type returnData map[string]string

// Del delete link preview cache
func (ctrl LinkController) Del(c *gin.Context) {
	var linkForm forms.LinkForm
	if c.ShouldBind(&linkForm) != nil {
		c.JSON(http.StatusOK, NewRespData(FailCode, "Invalid form", linkForm))
		return
	}
	services.LinkPreviewService.Delete(linkForm.URL)

	c.JSON(http.StatusOK, NewRespData(SuccessCode, "success", nil))
	return
}

// Preview link preview API
func (ctrl LinkController) Preview(c *gin.Context) {
	var linkForm forms.LinkForm
	if c.ShouldBind(&linkForm) != nil {
		c.JSON(http.StatusOK, NewRespData(FailCode, "Invalid form", linkForm))
		return
	}
	url := linkForm.URL

	if !strings.HasPrefix(url, "https") && !strings.HasPrefix(url, "http") {
		c.JSON(http.StatusOK, NewRespData(FailCode, "fail", returnData{"request_url": url}))

		return
	}

	var title, description, image string
	// 从缓存中取结果
	result, err := services.LinkPreviewService.GetValues(url)

	// 缓存存在
	if err != nil { // 缓存不存在

		host, err := utils.GetHostName(url)
		tech, err1 := techniques.GetTechnique(host)
		var extractor extractors.Extractor
		if err == nil && err1 == nil {
			log.Println(ansi.Color("[使用technique]:", "green"), tech.GetName()) // 查到host对应的technique
			// 1.根据域名判断需要使用的technique
			extractor = extractors.NewExtractor(
				false,
				tech,
			)
		} else {
			log.Println(ansi.Color("[无对应technique]:", "blue"), err1) // 未查到host对应的technique
			// 2.使用通用technique
			extractor = extractors.NewExtractor(
				false,
				techniques.HeadTagsTechnique{Name: "HeadTagsTechnique"},
				techniques.HTML5SemanticTagsTechnique{Name: "HTML5SemanticTagsTechnique"},
				techniques.SemanticTagsTechnique{Name: "SemanticTagsTechnique"},
			)
		}
		html, err := utils.GetHTML(url)
		if err != nil {
			c.JSON(http.StatusOK, NewRespData(FailCode, "请求页面错误,"+err.Error(), nil))

			return
		}
		extracted := extractor.Extract(html, url)
		// test print
		title, description, image = utils.GetSafeFirst(extracted[techniques.TitlesField]),
			utils.GetSafeFirst(extracted[techniques.DescriptionsField]),
			utils.GetSafeFirst(extracted[techniques.ImagesField])

		// 缓存结果
		services.LinkPreviewService.SetValues(url,
			map[string]interface{}{
				"title":       title,
				"description": description,
				"image":       image,
			})
	} else {
		title, description, image = result[0], result[1], result[2]
		log.Println("取到缓存")
	}

	// 返回
	c.JSON(http.StatusOK, NewRespData(SuccessCode, "success", returnData{
		"title":       title,
		"description": description,
		"image":       image,
	}))
}
