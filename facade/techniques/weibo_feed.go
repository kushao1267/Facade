package techniques

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/kushao1267/Facade/facade/utils"
	"strings"
)

// WeiboTechnique
type WeiboTechnique BaseTechnique

func (t WeiboTechnique) setName(name string) {
	t.Name = name
}

func (t WeiboTechnique) GetName() string {
	return t.Name
}

func (t WeiboTechnique) Extract(html string) DirtyExtracted {
	extracted := GetEmptyDirtyExtracted()
	t.setName("WeiboTechnique")

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return extracted
	}
	jsCode := doc.Find("script").Text()

	/* 自定义提取信息 */
	// title
	titles := utils.MatchOneOf(jsCode, `"status_title": "(.+?)"`)
	if len(titles) > 1 {
		extracted[TitlesField] = append(extracted[TitlesField], titles[1:]...)
	}
	// image
	images := utils.MatchOneOf(jsCode, `"url": "(.+?)"`)
	if len(images) > 1 {
		extracted[ImagesField] = append(extracted[ImagesField], images[1:]...)
	}
	// description
	descriptions := utils.MatchOneOf(jsCode, `"text": "(.+?)"`)
	if len(descriptions) > 1 {
		extracted[DescriptionsField] = append(extracted[DescriptionsField], descriptions[1:]...)
	}

	return extracted
}
