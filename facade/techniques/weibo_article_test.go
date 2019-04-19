package techniques

import (
	"github.com/kushao1267/facade/facade/utils"
	"log"
	"testing"
)

func TestWeiboArticleTechnique_Extract(t *testing.T) {
	var technique WeiboArticleTechnique

	html := utils.GetHtml("https://media.weibo.cn/article?id=2309404362621859024154&jumpfrom=weibocom")

	extracted := technique.Extract(html)
	allEmpty := true
	for _, value := range extracted {
		if len(value) > 0 {
			allEmpty = false
		}
	}
	if allEmpty {
		t.Fail()
	}
	log.Println(extracted)
}
