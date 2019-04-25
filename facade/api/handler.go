package api

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/kushao1267/Facade/facade/db"
	"github.com/kushao1267/Facade/facade/extractors"
	"github.com/kushao1267/Facade/facade/techniques"
	"github.com/kushao1267/Facade/facade/utils"
	"github.com/mgutz/ansi"
)

// Ping: test whether if the API server is running
func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

type ReturnData map[string]string

// LinkPreview: link preview API
func LinkPreview(c *gin.Context) {
	url := c.Request.FormValue("url")

	if !strings.HasPrefix(url, "https") && !strings.HasPrefix(url, "http") {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": FailCode,
			"msg":  "fail",
			"data": ReturnData{
				"request_url": url,
			},
		})
		return
	}

	var title, description, image string
	// 从缓存中取结果
	err, result := db.LinkPreviewService.GetValues(
		url,
		db.LinkPreviewService.Title,
		db.LinkPreviewService.Description,
		db.LinkPreviewService.Image)

	// 缓存存在
	if err == nil {
		title, description, image = result[0], result[1], result[2]
		log.Println("取到缓存")
	}else {
		// 缓存不存在
		// 1.根据域名判断需要使用的technique
		host, err := utils.GetHostName(url)
		tech, err1 := techniques.GetTechnique(host)
		if err == nil && err1 == nil {
			log.Println(ansi.Color("[使用technique]:", "green"), tech.GetName()) // 查到host对应的technique

			extractor := extractors.NewExtractor(
				false,
				tech,
			)
			extracted := extractor.Extract(utils.GetHtml(url), url)
			// test print
			title, description, image = utils.GetSafeFirst(extracted[techniques.TitlesField]),
				utils.GetSafeFirst(extracted[techniques.DescriptionsField]),
				utils.GetSafeFirst(extracted[techniques.ImagesField])
		} else {
			log.Println(ansi.Color("[未查到host对应的technique]:", "blue"), err1) // 未查到host对应的technique

			// 2.使用通用technique
			extractor := extractors.NewExtractor(
				false,
				techniques.HeadTagsTechnique{"HeadTagsTechnique"},
				techniques.HTML5SemanticTagsTechnique{"HTML5SemanticTagsTechnique"},
				techniques.SemanticTagsTechnique{"SemanticTagsTechnique"},
			)
			extracted := extractor.Extract(utils.GetHtml(url), url)
			// test print
			title, description, image = utils.GetSafeFirst(extracted[techniques.TitlesField]),
				utils.GetSafeFirst(extracted[techniques.DescriptionsField]),
				utils.GetSafeFirst(extracted[techniques.ImagesField])
		}

		// 缓存结果
		db.LinkPreviewService.SetValues(url,
			map[string]interface{}{
				"title":       title,
				"description": description,
				"image":       image,
			})
	}
	// 返回
	c.JSON(http.StatusOK, gin.H{
		"code": SuccessCode,
		"msg":  "success",
		"data": ReturnData{
			"title":       title,
			"description": description,
			"image":       image,
		},
	})
	return
}

