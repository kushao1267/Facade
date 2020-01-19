package techniques

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/kushao1267/Facade/facade/utils"
)

// WeiboArticleTechnique ...
type WeiboArticleTechnique BaseTechnique

func (t WeiboArticleTechnique) setName(name string) {
	t.Name = name
}

// GetName weibo article get name method
func (t WeiboArticleTechnique) GetName() string {
	return t.Name
}

// Extract weibo article extract method
func (t WeiboArticleTechnique) Extract(html string) DirtyExtracted {
	extracted := GetEmptyDirtyExtracted()
	t.setName("WeiboArticleTechnique")

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return extracted
	}
	jsCode := doc.Find("script").Text()
	/* 自定义提取信息 */
	// title
	titles := utils.MatchOneOf(jsCode, `"title": "(.+?)"`)
	if len(titles) > 1 {
		extracted[TitlesField] = append(extracted[TitlesField], titles[1:]...)
	}
	// image
	images := utils.MatchOneOf(jsCode, `"avatar_hd": "(.+?)"`)
	if len(images) > 1 {
		extracted[ImagesField] = append(extracted[ImagesField], images[1:]...)
	}
	// description
	descriptions := utils.MatchOneOf(jsCode, `"summary": "(.+?)"`)
	if len(descriptions) > 1 {
		extracted[DescriptionsField] = append(extracted[DescriptionsField], descriptions[1:]...)
	}

	return extracted
}
