package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kushao1267/facade/facade/db"
	"github.com/kushao1267/facade/facade/extrator"
	"github.com/kushao1267/facade/facade/logger"
	"github.com/kushao1267/facade/facade/techniques"
	"github.com/kushao1267/facade/facade/utils"
	"net/http"
	"strings"
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

	if strings.HasPrefix(url, "https") || strings.HasPrefix(url, "http") {
		// 从缓存中取结果
		result := db.LinkPreviewService.GetValues(
			url,
			db.LinkPreviewService.Title,
			db.LinkPreviewService.Description,
			db.LinkPreviewService.Image)

		if len(result) > 0{ // 缓存存在
			c.JSON(http.StatusOK, gin.H{
				"code": SuccessCode,
				"msg":  "success",
				"data": ReturnData{
					"title": result[0],
					"description":result[1],
					"image": result[2],
				},
			})
			return
		}
		// 抓取
		// 1.根据域名判断需要使用的technique
		host, err:=utils.GetHostName(url)
		tech, err1 := techniques.GetTechnique(host)
		if err == nil && err1 == nil {
			extractor:= extrator.NewExtractor(
				false,
				tech,
				techniques.HeadTagsTechnique{"HeadTagsTechnique"},
				techniques.HTML5SemanticTagsTechnique{"HTML5SemanticTagsTechniques"},
				techniques.SemanticTagsTechnique{"SemanticTagsTechnique"},
			)
			extracted := extractor.Extract(utils.GetHtml(url), url)
			// test print
			c.JSON(http.StatusOK, gin.H{
				"code": SuccessCode,
				"msg":  "success",
				"data": ReturnData{
					"title": extracted[techniques.TitlesField][0],
					"description":extracted[techniques.DescriptionsField][0],
					"image": extracted[techniques.ImagesField][0],
				},
			})
			return
		}
		logger.JsonLogger.Info(err1) // 未查到host对应的technique

		// 2.使用通用technique
		extractor := extrator.NewExtractor(
			false,
			techniques.HeadTagsTechnique{"HeadTagsTechnique"},
			techniques.HTML5SemanticTagsTechnique{"HTML5SemanticTagsTechnique"},
			techniques.SemanticTagsTechnique{"SemanticTagsTechnique"},
		)
		extracted := extractor.Extract(utils.GetHtml(url), url)
		// test print
		c.JSON(http.StatusOK, gin.H{
			"code": SuccessCode,
			"msg":  "success",
			"data": ReturnData{
				"title": utils.GetSafeFirst(extracted[techniques.TitlesField]),
				"description":utils.GetSafeFirst(extracted[techniques.DescriptionsField]),
				"image": utils.GetSafeFirst(extracted[techniques.ImagesField]),
			},
		})
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{
		"code": FailCode,
		"msg":  "fail",
		"data": url,
	})
}


