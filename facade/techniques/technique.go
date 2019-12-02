package techniques

import (
	"errors"
)

var (
	EmptyString = ""
	EmptyList   = []string{}
	// extracted field
	TitlesField       = "titles"
	DescriptionsField = "descriptions"
	ImagesField       = "images"
	UrlsField         = "urls"
	VideosField       = "videos"
	FeedsField        = "feeds"
)

func GetEmptyDirtyExtracted() DirtyExtracted {
	return DirtyExtracted{
		TitlesField:       EmptyList,
		DescriptionsField: EmptyList,
		ImagesField:       EmptyList,
		UrlsField:         EmptyList,
		VideosField:       EmptyList,
		FeedsField:        EmptyList,
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
