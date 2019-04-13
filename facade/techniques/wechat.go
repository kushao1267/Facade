package techniques

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/kushao1267/facade/facade/utils"
	"strings"
)

// WeChatTechnique
type WeChatTechnique BaseTechnique

func (t WeChatTechnique) setName(name string) {
	t.Name = name
}

func (t WeChatTechnique) GetName() string {
	return t.Name
}

func (t WeChatTechnique) Extract(html string) DirtyExtracted {
	extracted := GetEmptyDirtyExtracted()
	t.setName("WeChatTechnique")

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return extracted
	}
	jsCode := ""
	doc.Find("script").Each(func(i int, selection *goquery.Selection) {
		if i == 12 {
			jsCode = selection.Text()
		}
	})

	/* 自定义提取信息 */
	// title
	titles := utils.MatchOneOf(jsCode, `msg_title = "(.+?)";`)
	if len(titles) > 1 {
		extracted[TitlesField] = append(extracted[TitlesField], titles[1:]...)
	}
	// image
	images := utils.MatchOneOf(jsCode, `msg_cdn_url = "(.+?)";`)
	if len(images) > 1 {
		extracted[ImagesField] = append(extracted[ImagesField], images[1:]...)
	}
	// description
	descriptions := utils.MatchOneOf(jsCode, `msg_desc = "(.+?)";`)
	if len(descriptions) > 1 {
		extracted[DescriptionsField] = append(extracted[DescriptionsField], descriptions[1:]...)
	}

	return extracted
}
