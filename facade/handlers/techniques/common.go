package techniques

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/kushao1267/facade/facade/handlers"
	"strings"
)

var emptyDirtyExtracted = &DirtyExtracted{
	"titles":       []string{},
	"descriptions": []string{},
	"images":       []string{},
	"urls":         []string{},
}

// Technique 必须实现的方法
type Technique interface {
	SetName(name string)
	GetName() string
	SetExtractor(extracted handlers.Extractor)
	Extract(html string) *DirtyExtracted
}

// DirtyExtracted :未经过clean的提取结果
type DirtyExtracted map[string][]string

// Technique
type BaseTechnique struct {
	Name      string
	Extractor handlers.Extractor
}

func (t BaseTechnique) SetName(name string) {
	t.Name = name
}

func (t BaseTechnique) GetName() string {
	return t.Name
}

func (t BaseTechnique) SetExtractor(extractor handlers.Extractor) {
	t.Extractor = extractor
}

// Extract :Extract data from a string representing an HTML document.
func (t BaseTechnique) Extract(html string) *DirtyExtracted {
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
	Name      string
	Extractor handlers.Extractor
}

func (t HeadTagsTechnique) SetName(name string) {
	t.Name = name
}

func (t HeadTagsTechnique) GetName() string {
	return t.Name
}

func (t HeadTagsTechnique) SetExtractor(extractor handlers.Extractor) {
	t.Extractor = extractor
}

// Extract :Extract data from a string representing an HTML document.
func (t HeadTagsTechnique) Extract(html string) *DirtyExtracted {
	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return emptyDirtyExtracted
	}
	// Find the review items
	title := doc.Find("title").First().Text()
	&DirtyExtracted{"titles": []string{title}}
	return emptyDirtyExtracted
}

// HTML5SemanticTagsTechnique
type HTML5SemanticTagsTechnique struct {
	Name      string
	Extractor handlers.Extractor
}

func (t HTML5SemanticTagsTechnique) SetName(name string) {
	t.Name = name
}

func (t HTML5SemanticTagsTechnique) GetName() string {
	return t.Name
}

func (t HTML5SemanticTagsTechnique) SetExtractor(extractor handlers.Extractor) {
	t.Extractor = extractor
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
	Name      string
	Extractor handlers.Extractor
}

func (t SemanticTagsTechnique) SetName(name string) {
	t.Name = name
}

func (t SemanticTagsTechnique) GetName() string {
	return t.Name
}

func (t SemanticTagsTechnique) SetExtractor(extractor handlers.Extractor) {
	t.Extractor = extractor
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
