package common

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/kushao1267/facade/facade/techniques"
	"strings"
)

// Extract info from standard HTML metatags like title
// This is usually a last-resort, low quality, but reliable parsing mechanism.
// HeadTagsTechnique
type HeadTagsTechnique struct {
	Name string
}

func (t HeadTagsTechnique) setName(name string) {
	t.Name = name
}

func (t HeadTagsTechnique) GetName() string {
	return t.Name
}

func (t HeadTagsTechnique) getMetaNameMap() *map[string]string {
	return &map[string]string{
		techniques.DescriptionsField: techniques.DescriptionsField,
		"author":                     "authors",
	}
}

// Extract :Extract data from a string representing an HTML document.
func (t HeadTagsTechnique) Extract(html string) techniques.DirtyExtracted {
	extracted := techniques.GetEmptyDirtyExtracted()
	t.setName("HeadTagsTechnique")

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return extracted
	}
	// Find the title
	if title := doc.Find("title").First().Text(); title != techniques.EmptyString {
		extracted[techniques.TitlesField] = append(extracted[techniques.TitlesField], title)
	}
	// extract data from meta tags
	metaNameMap := *t.getMetaNameMap()
	doc.Find("meta").Each(func(i int, selection *goquery.Selection) {
		name, e1 := selection.Attr("name")
		content, e2 := selection.Attr("content")
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
				if _, ok3 := extracted[techniques.UrlsField]; ok3 {
					extracted[techniques.UrlsField] = append(extracted[techniques.UrlsField], href)
				}
			} else if strings.Contains(rel, "alternate") && ok1 && ok2 && _type == "application/rss+xml" {
				if _, ok3 := extracted[techniques.FeedsField]; ok3 {
					extracted[techniques.FeedsField] = append(extracted[techniques.FeedsField], href)
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

func (t HTML5SemanticTagsTechnique) setName(name string) {
	t.Name = name
}

func (t HTML5SemanticTagsTechnique) GetName() string {
	return t.Name
}

// The HTML5 `article` tag, and also the `video` tag give us some useful
// hints for extracting page information for the sites which happen to
// utilize these tags.
func (t HTML5SemanticTagsTechnique) Extract(html string) techniques.DirtyExtracted {
	extracted := techniques.GetEmptyDirtyExtracted()
	t.setName("HTML5SemanticTagsTechnique")

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return extracted
	}

	doc.Find("article").Each(func(i int, selection *goquery.Selection) {
		if title := selection.Find("h1").Text(); title != techniques.EmptyString {
			extracted[techniques.TitlesField] = append(extracted[techniques.TitlesField], title)
		}
		if desc := selection.Find("p").Text(); desc != techniques.EmptyString {
			extracted[techniques.DescriptionsField] = append(extracted[techniques.DescriptionsField], desc)
		}
	})

	doc.Find("video").Each(func(i int, selection *goquery.Selection) {
		selection.Find("source").Each(func(i int, selection *goquery.Selection) {
			if src, ok := selection.Attr("src"); ok {
				extracted[techniques.VideosField] = append(extracted[techniques.VideosField], src)
			}
		})
	})
	return extracted
}

// This technique relies on the basic tags themselves--for example,
// all IMG tags include images, most H1 and H2 tags include titles,
// and P tags often include text usable as descriptions.
//
// This is a true last resort technique.
// SemanticTagsTechnique
type SemanticTagsTechnique struct {
	Name string
}

func (t SemanticTagsTechnique) setName(name string) {
	t.Name = name
}

func (t SemanticTagsTechnique) GetName() string {
	return t.Name
}

// tuple形式的结构可以用struct构造
type extractString struct {
	tag       string
	dest      string
	maxStores int
}

type extractAttr struct {
	tag       string
	dest      string
	attr      string
	maxStores int
}

// list to support ordering of semantics, e.g. h1
// is higher quality than h2 and so on
// format is {"name of tag", "destination list", store_first_n}
func (t SemanticTagsTechnique) getExtractString() *[]extractString {
	return &[]extractString{
		{
			"h1", techniques.TitlesField, 3,
		},
		{
			"h2", techniques.TitlesField, 3,
		},
		{
			"h3", techniques.TitlesField, 1,
		},
		{
			"p", techniques.DescriptionsField, 5,
		},
	}
}

// list to support ordering of semantics, e.g. h1
// is higher quality than h2 and so on
// format is {"name of tag", "destination list", store_first_n}
func (t SemanticTagsTechnique) getExtractAttribute() *[]extractAttr {
	return &[]extractAttr{
		{"img", techniques.ImagesField, "src", 3},
	}
}

// Extract :Extract data from a string representing an HTML document.
func (t SemanticTagsTechnique) Extract(html string) techniques.DirtyExtracted {
	extracted := techniques.GetEmptyDirtyExtracted()
	t.setName("SemanticTagsTechnique")
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return extracted
	}

	extractStr := *t.getExtractString()
	extractAttr := *t.getExtractAttribute()
	for _, val := range extractStr {
		stores := 0
		doc.Find(val.tag).Each(func(i int, selection *goquery.Selection) {
			if stores < val.maxStores {
				extracted[val.dest] = append(extracted[val.dest], selection.Text())
			}
			stores += 1
		})
	}

	for _, val := range extractAttr {
		stores := 0
		doc.Find(val.tag).Each(func(i int, selection *goquery.Selection) {
			if stores < val.maxStores {
				if attr, ok := selection.Attr(val.attr); ok {
					extracted[val.dest] = append(extracted[val.dest], attr)
				}
			}
			stores += 1
		})
	}
	return extracted
}
