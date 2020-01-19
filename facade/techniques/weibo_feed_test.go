package techniques

import (
	"testing"

	"github.com/kushao1267/Facade/facade/utils"
)

func TestWeiboTechnique_Extract(t *testing.T) {
	var technique WeiboTechnique

	html, _ := utils.GetHTML("https://m.weibo.cn/5187664653/4354456894352205")

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
