package techniques

import (
	"github.com/kushao1267/Facade/facade/utils"
)

// QQOMTechnique
type QQOMTechnique BaseTechnique

func (t QQOMTechnique) setName(name string) {
	t.Name = name
}

func (t QQOMTechnique) GetName() string {
	return t.Name
}

func (t QQOMTechnique) Extract(html string) DirtyExtracted {
	extracted := GetEmptyDirtyExtracted()
	t.setName("QQOMTechnique")

	/* 自定义提取信息 */
	// title
	titles := utils.MatchOneOf(html, `<meta name="tgm:shareTitle" content="(.+?)">`)
	if len(titles) > 1 {
		extracted[TitlesField] = append(extracted[TitlesField], titles[1:]...)
	}
	// image
	images := utils.MatchOneOf(html, `<img src="(.+?)"`)
	if len(images) > 1 {
		extracted[ImagesField] = append(extracted[ImagesField], images[1:]...)
	}
	// description
	descriptions := utils.MatchOneOf(html, `<meta name="tgm:shareDesc" content="(.+?)">`)
	if len(descriptions) > 1 {
		extracted[DescriptionsField] = append(extracted[DescriptionsField], descriptions[1:]...)
	}

	return extracted
}

