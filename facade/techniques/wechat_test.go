package techniques

import "testing"
import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func TestWeChatTechnique_Extract(t *testing.T) {
	var technique WeChatTechnique

	res, err := http.Get("https://mp.weixin.qq.com/s/VRzeIxFO_sHTOHAyZRX7xw")
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		fmt.Println("status code error: " + string(res.StatusCode) + " " + res.Status)
	}
	bodyBytes, _ := ioutil.ReadAll(res.Body)

	extracted := technique.Extract(string(bodyBytes))
	allEmpty := true
	for _, value := range extracted {
		if len(value) > 0 {
			allEmpty = false
		}
	}
	if allEmpty {
		t.Fail()
	}
	fmt.Println(extracted)
}
