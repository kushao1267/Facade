package techniques

import (
	"github.com/kushao1267/Facade/facade/utils"
)

// QQOMTechnique ...
type QQMPTechnique BaseTechnique

func (t QQMPTechnique) setName(name string) {
	t.Name = name
}

func (t QQMPTechnique) GetName() string {
	return t.Name
}

func (t QQMPTechnique) Extract(html string) DirtyExtracted {
	extracted := GetEmptyDirtyExtracted()
	t.setName("QQMPTechnique")
	/* 自定义提取信息 */
	// title
	titles := utils.MatchOneOf(html, `data-article-title="(.+?)"`)
	if len(titles) > 1 {
		extracted[TitlesField] = append(extracted[TitlesField], titles[1:]...)
	}
	// image
	images := utils.MatchOneOf(html, `data-cover-url="(.+?)"`)
	if len(images) > 1 {
		extracted[ImagesField] = append(extracted[ImagesField], images[1:]...)
	}
	// description
	descriptions := utils.MatchOneOf(html, `data-article-desc="(.+?)"`)
	if len(descriptions) > 1 {
		extracted[DescriptionsField] = append(extracted[DescriptionsField], descriptions[1:]...)
	}

	return extracted
}
