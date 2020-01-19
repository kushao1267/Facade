package techniques

import (
	"testing"

	"github.com/kushao1267/Facade/facade/utils"
)

func TestToutiaoTechnique_Extract(t *testing.T) {
	var technique ToutiaoTechnique

	html, _ := utils.GetHTML("https://www.toutiao.com/a6673091667941130756/")

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
