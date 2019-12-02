package techniques

import (
	"errors"
)

var (
	emptyList = []string{}
	// EmptyString empty string
	EmptyString = ""
	// TitlesField titles field
	TitlesField = "titles"
	// DescriptionsField ...
	DescriptionsField = "descriptions"
	// ImagesField ...
	ImagesField = "images"
	// UrlsField ...
	UrlsField = "urls"
	// VideosField ...
	VideosField = "videos"
	// FeedsField ...
	FeedsField = "feeds"
)

// GetEmptyDirtyExtracted get empty extracted
func GetEmptyDirtyExtracted() DirtyExtracted {
	return DirtyExtracted{
		TitlesField:       emptyList,
		DescriptionsField: emptyList,
		ImagesField:       emptyList,
		UrlsField:         emptyList,
		VideosField:       emptyList,
		FeedsField:        emptyList,
	}
}

// Technique 必须实现的方法
type Technique interface {
	setName(name string)
	GetName() string
	Extract(html string) DirtyExtracted
}

// DirtyExtracted :未经过clean的提取结果
type DirtyExtracted map[string][]string

// GetTechnique 从hostname获取相应的technique
func GetTechnique(host string) (Technique, error) {
	// 从hostname获取相应的technique
	techMap := map[string]Technique{
		"mp.weixin.qq.com": WeChatTechnique{"WeChatTechnique"},
		"www.toutiao.com":  ToutiaoTechnique{"ToutiaoTechnique"},
		"m.zjbyte.com":     ToutiaoTechnique{"ToutiaoTechnique"},
		"page.om.qq.com":   QQOMTechnique{"QQOMTechnique"},
		"post.mp.qq.com":   QQMPTechnique{"QQMPTechnique"},
		"m.weibo.cn":       WeiboTechnique{"WeiboTechnique"},
		"media.weibo.cn":   WeiboArticleTechnique{"WeiboArticleTechnique"},
	}

	if val, ok := techMap[host]; ok {
		return val, nil
	}
	return BaseTechnique{}, errors.New("Not Implement technique for:" + host)
}
