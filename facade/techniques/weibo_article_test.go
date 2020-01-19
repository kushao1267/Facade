package techniques

import (
	"testing"

	"github.com/kushao1267/Facade/facade/utils"
)

func TestWeiboArticleTechnique_Extract(t *testing.T) {
	var technique WeiboArticleTechnique

	html, _ := utils.GetHTML("https://media.weibo.cn/article?id=2309404362621859024154&jumpfrom=weibocom")

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
	t.Log(extracted)
}
