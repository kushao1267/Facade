package techniques

import (
	"github.com/PuerkitoBio/goquery"
	"strings"
)

var emptyDirtyExtracted = DirtyExtracted{
	"titles":       []string{},
	"descriptions": []string{},
	"images":       []string{},
	"urls":         []string{},
}

// Technique 必须实现的方法
type Technique interface {
	SetName(name string)
	GetName() string
	Extract(html string) DirtyExtracted
}

// DirtyExtracted :未经过clean的提取结果
type DirtyExtracted map[string][]string

// Technique
type BaseTechnique struct {
	Name string
}

func (t BaseTechnique) SetName(name string) {
	t.Name = name
}

func (t BaseTechnique) GetName() string {
	return t.Name
}

// Extract :Extract data from a string representing an HTML document.
func (t BaseTechnique) Extract(html string) DirtyExtracted {
	return emptyDirtyExtracted
}

// Extract info from standard HTML metatags like title, for example:
//
// <head>
// <meta http-equiv="content-type" content="text/html; charset=UTF-8" />
// <meta name="author" content="Will Larson" />
// <meta name="description" content="Will Larson&#39;s blog about programming and other things." />
// <meta name="keywords" content="Blog Will Larson Programming Life" />
// <link rel="alternate" type="application/rss+xml" title="Page Feed" href="/feeds/" />
// <link rel="canonical" href="http://lethain.com/digg-v4-architecture-process/">
// <title>Digg v4&#39;s Architecture and Development Processes - Irrational Exuberance</title>
// </head>
//
// This is usually a last-resort, low quality, but reliable parsing mechanism.
// HeadTagsTechnique
type HeadTagsTechnique struct {
	Name string
}

func (t HeadTagsTechnique) SetName(name string) {
	t.Name = name
}

func (t HeadTagsTechnique) GetName() string {
	return t.Name
}

func (t HeadTagsTechnique) GetMetaNameMap() *map[string]string {
	return &map[string]string{
		"description": "descriptions",
		"author":      "authors",
	}
}

// Extract :Extract data from a string representing an HTML document.
func (t HeadTagsTechnique) Extract(html string) DirtyExtracted {
	extracted := emptyDirtyExtracted
	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return extracted
	}
	// Find the titlew
	if title := doc.Find("title").First().Text(); title != "" {
		extracted["title"] = append(extracted["title"], title)
	}
	// extract data from meta tags
	metaNameMap := *t.GetMetaNameMap()
	doc.Find("meta").Each(func(i int, selection *goquery.Selection) {
		name, e1 := selection.Attr("name")
		content, e2 := selection.Attr("")
		if e1 && e2 {
			if nameDest, ok := metaNameMap[name]; ok {
				if _, ok1 := extracted[nameDest]; ok1 {
					extracted[nameDest] = append(extracted[nameDest], content)
				}
			}
		}
	})

	// extract data from link tags
	doc.Find("link").Each(func(i int, selection *goquery.Selection) {
		if rel, ok := selection.Attr("rel"); ok {
			href, ok1 := selection.Attr("href")
			_type, ok2 := selection.Attr("type")
			if strings.Contains(rel, "canonical") && ok1 {
				if _, ok3 := extracted["urls"]; ok3 {
					extracted["urls"] = append(extracted["urls"], href)
				}
			} else if strings.Contains(rel, "alternate") && ok1 && ok2 && _type == "application/rss+xml" {
				if _, ok3 := extracted["feeds"]; ok3 {
					extracted["feeds"] = append(extracted["feeds"], href)
				}
			}
		}

	})

	return extracted
}

// HTML5SemanticTagsTechnique
type HTML5SemanticTagsTechnique struct {
	Name string
}

func (t HTML5SemanticTagsTechnique) SetName(name string) {
	t.Name = name
}

func (t HTML5SemanticTagsTechnique) GetName() string {
	return t.Name
}

// Extract :Extract data from a string representing an HTML document.
func (t HTML5SemanticTagsTechnique) Extract(html string) DirtyExtracted {
	return DirtyExtracted{
		"titles":       []string{},
		"descriptions": []string{},
		"images":       []string{},
		"urls":         []string{},
	}
}

// SemanticTagsTechnique
type SemanticTagsTechnique struct {
	Name string
}

func (t SemanticTagsTechnique) SetName(name string) {
	t.Name = name
}

func (t SemanticTagsTechnique) GetName() string {
	return t.Name
}

// Extract :Extract data from a string representing an HTML document.
func (t SemanticTagsTechnique) Extract(html string) DirtyExtracted {
	return DirtyExtracted{
		"titles":       []string{},
		"descriptions": []string{},
		"images":       []string{},
		"urls":         []string{},
	}
}
