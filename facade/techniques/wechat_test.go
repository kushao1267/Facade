package techniques

import (
	"github.com/kushao1267/Facade/facade/utils"
	"log"
	"testing"
)

func TestWeChatTechnique_Extract(t *testing.T) {
	var technique WeChatTechnique

	html := utils.GetHtml("https://mp.weixin.qq.com/s?__biz=MzA5MzY4NTQwMA==&mid=2651009603&idx=1&" +
		"sn=8bfdbbdfbd39657a42bb11de8247748b&chksm=8bad81b4bcda08a2f871c3e98bc19b17e28770cbcec8b4faf5db5975" +
			"e6d8e7718f3e2fadd37e&scene=38&ascene=0&devicetype=android-26&version=27000364&nettype=WIFI&abtes" +
				"t_cookie=BAABAAoACwASABMABQAjlx4AVpkeAMmZHgDamR4A3JkeAAAA&lang=zh_CN&pass_ticket=Zie1T1S73St7" +
					"j%2FX5Y%2B52AD86py31gJUONapTiG%2Fvt%2BNITYW5wxzba8PFWPIUf3sf&wx_header=1")

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
