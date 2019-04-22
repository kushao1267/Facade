package techniques

import (
	"github.com/kushao1267/Facade/facade/utils"
	"log"
	"testing"
)

func TestToutiaoTechnique_Extract(t *testing.T) {
	var technique ToutiaoTechnique

	html := utils.GetHtml("https://www.toutiao.com/a6673091667941130756/")

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
