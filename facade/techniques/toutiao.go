package techniques

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/kushao1267/Facade/facade/utils"
	"strings"
)

// ToutiaoTechnique
type ToutiaoTechnique BaseTechnique

func (t ToutiaoTechnique) setName(name string) {
	t.Name = name
}

func (t ToutiaoTechnique) GetName() string {
	return t.Name
}

func (t ToutiaoTechnique) Extract(html string) DirtyExtracted {
	extracted := GetEmptyDirtyExtracted()
	t.setName("ToutiaoTechnique")

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return extracted
	}
	jsCode := ""
	doc.Find("script").Each(func(i int, selection *goquery.Selection) {
		if i == 6 {
			jsCode = selection.Text()
		}
	})

	/* 自定义提取信息 */
	// title
	titles := utils.MatchOneOf(jsCode, `title: '(.+?)',`)
	if len(titles) > 1 {
		extracted[TitlesField] = append(extracted[TitlesField], titles[1:]...)
	}
	// image
	images := utils.MatchOneOf(jsCode, `coverImg: '(.+?)'`)
	if len(images) > 1 {
		extracted[ImagesField] = append(extracted[ImagesField], images[1:]...)
	}
	// description
	descriptions := utils.MatchOneOf(jsCode, `content: '(.+?)',`)
	if len(descriptions) > 1 {
		extracted[DescriptionsField] = append(extracted[DescriptionsField], descriptions[1:]...)
	}

	return extracted
}
