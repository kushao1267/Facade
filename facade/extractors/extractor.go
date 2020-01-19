package extractors

import "fmt"
import (
	"html"
	"strings"

	"github.com/kushao1267/Facade/facade/techniques"
	"github.com/kushao1267/Facade/facade/utils"
)

const markTechnique = false

// Extracted :Contains data extracted from a page.
type Extracted map[string][]string

var emptyData = ""

// represent print all Extracted content
func (e Extracted) represent() []string {
	maxShown := 40
	var detailStr []string

	for name, values := range e {
		count := len(values)
		if count > 0 {
			value := values[0]
			if count-1 > 0 {
				detailStr = append(
					detailStr,
					fmt.Sprintf("(%s: '%s', %s more)", name, value[:maxShown], string(count-1)),
				)
			} else {
				detailStr = append(
					detailStr,
					fmt.Sprintf("(%s: '%s')", name, value[:maxShown]),
				)
			}
		}
	}

	return detailStr
}

func (e Extracted) title() string {
	if val, ok := e["title"]; ok {
		if len(val) > 0 {
			return val[0]
		}
	}
	return emptyData
}

func (e Extracted) image() string {
	if val, ok := e["image"]; ok {
		if len(val) > 0 {
			return val[0]
		}
	}
	return emptyData
}

func (e Extracted) video() string {
	if val, ok := e["video"]; ok {
		if len(val) > 0 {
			return val[0]
		}
	}
	return emptyData
}

func (e Extracted) description() string {
	if val, ok := e["description"]; ok {
		if len(val) > 0 {
			return val[0]
		}
	}
	return emptyData
}

func (e Extracted) url() string {
	if val, ok := e["url"]; ok {
		if len(val) > 0 {
			return val[0]
		}
	}
	return emptyData
}

func (e Extracted) feed() string {
	if val, ok := e["feed"]; ok {
		if len(val) > 0 {
			return val[0]
		}
	}
	return emptyData
}

// Extractor Extracts title, image and description from an HTML document.
type Extractor struct {
	urlTypes    []string
	textTypes   []string
	strictTypes bool
	techniques  []techniques.Technique
}

// NewExtractor new an Extractor
func NewExtractor(strictTypes bool, techniques ...techniques.Technique) Extractor {
	d := Extractor{
		[]string{"images", "urls", "feeds", "videos"},
		[]string{"titles", "descriptions"},
		strictTypes,
		techniques,
	}

	return d
}

// Run a given technique against the HTML.
// Technique is a string including the full module path
// and class name for the technique,
// HTML is a string representing an HTML document.
func (d Extractor) runTechnique(technique techniques.Technique, html string) techniques.DirtyExtracted {
	return technique.Extract(html)
}

// cleanUpText Cleanup text values like titles or descriptions.
func (d Extractor) cleanUpText(value, mark string) string {
	// 去掉空格
	text := strings.TrimSpace(value)

	// 去除标签, 其中Unescape text是将"&lt;&gt;"这样的字符变为标签"<>"以便于去除.
	text = utils.CleanHTMLTags(html.UnescapeString(text))

	// 去掉非utf-8字符
	text = utils.ValidUTF8(text)

	// 长度限制
	runeText := []rune(text)
	if len(runeText) > 125 {
		runeText = append(runeText[:125], []rune("...")...)
	}
	text = string(runeText)

	if mark != "" {
		text = mark + " " + text
	}
	return text
}

// Transform relative URLs into absolute URLs if possible.
func (d Extractor) cleanUpURL(valueURL, sourceURL, mark string) string {
	netloc, _ := utils.GetHostName(valueURL)

	var url string
	if netloc != "" || sourceURL == "" {
		url = valueURL
	} else {
		url = utils.URLJoin(sourceURL, valueURL)
	}

	if strings.HasPrefix(url, "//") {
		url = "http:" + url
	}

	if mark != "" {
		url += mark
	}
	return url
}

// Allows standardizing extracted contents, at this time:
func (d Extractor) cleanUp(results techniques.DirtyExtracted, technique techniques.Technique, sourceURL string) Extracted {
	cleanedResults := Extracted{}

	var mark string
	if markTechnique {
		mark = "#" + technique.GetName() // 接口无法定义字段，只能通过method的方式来set和get
	} else {
		mark = ""
	}

	for dataType, dataValues := range results {

		if utils.StringInSlice(dataType, d.textTypes) {
			for _, dataValue := range dataValues {
				if dataValue != "" {
					cleanedText := d.cleanUpText(dataValue, mark)
					cleanedResults[dataType] = append(cleanedResults[dataType], cleanedText)
				}
			}
		} else if utils.StringInSlice(dataType, d.urlTypes) {
			for _, dataValue := range dataValues {
				cleanedURL := d.cleanUpURL(dataValue, sourceURL, mark)
				cleanedResults[dataType] = append(cleanedResults[dataType], cleanedURL)
			}
		} else if d.strictTypes {
			continue
		}
		// 去重
		cleanedResults[dataType] = utils.RemoveDuplicateString(cleanedResults[dataType])
	}
	return cleanedResults
}

// Extract contents from an HTML document.
func (d Extractor) Extract(html, sourceURL string) Extracted {
	var extracted = Extracted{}

	for _, technique := range d.techniques {

		techniqueExtracted := d.runTechnique(technique, html)
		techniqueCleaned := d.cleanUp(techniqueExtracted, technique, sourceURL)

		for dataType, dataValues := range techniqueCleaned {
			var uniqueDataValues []string

			if len(dataValues) > 0 {
				if _, ok := extracted[dataType]; !ok {
					extracted[dataType] = []string{}
				}
				// don't include duplicate values
				for _, x := range dataValues {
					if !utils.StringInSlice(x, extracted[dataType]) {
						uniqueDataValues = append(uniqueDataValues, x)
					}
				}

				extracted[dataType] = append(extracted[dataType], uniqueDataValues...)
			}
		}
	}

	return extracted
}
