package techniques

import (
	"github.com/kushao1267/Facade/facade/utils"
	"testing"
)

func TestQQOMTechnique_Extract(t *testing.T) {
	var technique QQOMTechnique

	_, html := utils.GetHtml("https://page.om.qq.com/page/OdcZGP3nLryW2tIbF8udTYrg0?source=omapp&appbar=omapp")

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
