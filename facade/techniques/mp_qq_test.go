package techniques

import (
	"github.com/kushao1267/Facade/facade/utils"
	"log"
	"testing"
)

func TestQQMPTechnique_Extract(t *testing.T) {
	var technique QQMPTechnique

	html := utils.GetHtml("http://post.mp.qq.com/kan/video/2500617653-0585c9" +
		"6135a584ah-u0853798wbx.html?_wv=2281701505&sig=ecd52b1b427e9e93b6da83c199c" +
			"e2709&time=1553339450&iid=Mzc4Njg4NjQ3Nw==&sourcefrom=6")

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
