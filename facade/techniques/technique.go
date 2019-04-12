package techniques

var (
	EmptyString       = ""
	TitlesField       = "titles"
	DescriptionsField = "descriptions"
	ImagesField       = "images"
	UrlsField         = "urls"
	VideosField       = "videos"
	FeedsField        = "feeds"
)

func GetEmptyDirtyExtracted() DirtyExtracted {
	return DirtyExtracted{
		TitlesField:       []string{},
		DescriptionsField: []string{},
		ImagesField:       []string{},
		UrlsField:         []string{},
		VideosField:       []string{},
		FeedsField:        []string{},
	}
}

// Technique 必须实现的方法
type Technique interface {
	SetName(name string)
	GetName() string
	Extract(html string) DirtyExtracted
}

// DirtyExtracted :未经过clean的提取结果
type DirtyExtracted map[string][]string
