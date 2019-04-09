package handlers

import "fmt"
import (
	"github.com/kushao1267/facade/facade/handlers/techniques"
	"github.com/kushao1267/facade/facade/utils"
	"strings"
)

const MarkTechnique = false

// Extracted :Contains data extracted from a page.
type Extracted map[string][]string

func (e Extracted) Represent() []string {
	maxShown := 40
	var detailStr []string

	for name, values := range e {
		count := len(values)
		if count > 0 {
			value := values[0]
			if count-1 > 0 {
				detailStr = append(
					detailStr,
					fmt.Sprintf("(%s: '%s', %s more)", name, value[:maxShown], count-1),
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

func (e Extracted) Title() string {
	if val, ok := e["title"]; ok {
		if len(val) > 0 {
			return val[0]
		}
	}
	return ""
}

func (e Extracted) Image() string {
	if val, ok := e["title"]; ok {
		if len(val) > 0 {
			return val[0]
		}
	}
	return ""
}

func (e Extracted) Video() string {
	if val, ok := e["title"]; ok {
		if len(val) > 0 {
			return val[0]
		}
	}
	return ""
}

func (e Extracted) Description() string {
	if val, ok := e["title"]; ok {
		if len(val) > 0 {
			return val[0]
		}
	}
	return ""
}

func (e Extracted) Url() string {
	if val, ok := e["title"]; ok {
		if len(val) > 0 {
			return val[0]
		}
	}
	return ""
}

func (e Extracted) Feed() string {
	if val, ok := e["title"]; ok {
		if len(val) > 0 {
			return val[0]
		}
	}
	return ""
}

// Extractor :Extracts title, image and description from an HTML document.
type Extractor struct {
	urlTypes    []string
	textTypes   []string
	strictTypes bool
	techniques  []techniques.Technique
}

func NewDictExtractor(techniques []techniques.Technique, strictTypes bool) Extractor {
	d := Extractor{}

	d.urlTypes = []string{"images", "urls", "feeds", "videos"}
	d.textTypes = []string{"titles", "descriptions"}
	d.strictTypes = strictTypes

	if len(techniques) > 0 {
		d.techniques = techniques
	}
	return d
}

// Run a given technique against the HTML.
// Technique is a string including the full module path
// and class name for the technique, for example::
// extraction.techniques.FacebookOpengraphTags
// HTML is a string representing an HTML document.
func (d Extractor) runTechnique(technique techniques.Technique, html string) techniques.DirtyExtracted {
	technique.SetExtractor(d)
	return technique.Extract(html)
}

// cleanUpText Cleanup text values like titles or descriptions.
func (d Extractor) cleanUpText(value, mark string) string {
	text := strings.TrimSpace(value)
	if mark != "" {
		text = mark + " " + text
	}
	return text
}

// Transform relative URLs into absolute URLs if possible.
// If the value_url is already absolute, or we don't know the
// source_url, then return the existing value. If the value_url is
// relative, and we know the source_url, then try to rewrite it.
func (d Extractor) cleanUpUrl(valueUrl, sourceUrl, mark string) string {
	netloc, _ := utils.GetHostName(valueUrl)

	var url string
	if netloc != "" || sourceUrl == "" {
		url = valueUrl
	} else {
		url = utils.UrlJoin(sourceUrl, valueUrl)
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
// 1. removes multiple whitespaces
// 2. rewrite relative URLs as absolute URLs if source_url is specified
// 3. filter out duplicate values
// 4. marks the technique that produced the result
// 5. returns only specified text_types and url_types depending on self.strict_types
func (d Extractor) cleanUp(results techniques.DirtyExtracted, technique techniques.Technique, sourceUrl string) Extracted {
	var cleanedResults Extracted

	var mark string
	if MarkTechnique {
		mark = "#" + technique.GetName() // 接口无法定义字段，只能通过method的方式来set和get
	} else {
		mark = ""
	}

	for dataType, dataValues := range results {
		var values []string
		if utils.StringInSlice(dataType, d.textTypes) {
			for _, dataValue := range dataValues {
				if dataValue != "" {
					values = append(values, d.cleanUpText(dataValue, mark))
				}
			}
		} else if utils.StringInSlice(dataType, d.urlTypes) {
			for _, dataValue := range dataValues {
				values = append(values, d.cleanUpUrl(dataValue, sourceUrl, mark))
			}
		} else if d.strictTypes {
			continue
		}

		var uniqueValues []string
		for _, dataValue := range dataValues {
			if utils.StringInSlice(dataValue, uniqueValues) {
				uniqueValues = append(uniqueValues, dataValue)
			}
		}

		cleanedResults[dataType] = uniqueValues
	}
	return cleanedResults
}

// Extracts contents from an HTML document.
// >>> from extraction import Extractor
// >>> import requests
// >>> html = requests.get("http://lethain.com/").text
// >>> extracted = Extractor().extract(html)
// >>> print extracted
// `source_url` is optional, but allows for a certain level of
// cleanup to be performed, such as converting relative URLs
// into absolute URLs and such.
func (d Extractor) Extract(html, sourceUrl string) Extracted {
	var extracted Extracted

	for _, technique := range d.techniques {

		techniqueExtracted := d.runTechnique(technique, html)
		techniqueCleaned := d.cleanUp(techniqueExtracted, technique, sourceUrl)

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
