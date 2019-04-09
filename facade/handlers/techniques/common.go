package techniques

import "github.com/kushao1267/facade/facade/handlers"

// Technique 必须实现Extract方法
type Technique interface {
	SetName(name string)
	GetName() string
	SetExtractor(extracted handlers.Extractor)
	Extract(html string) DirtyExtracted
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
func (t BaseTechnique) Extract(html string) DirtyExtracted {
	return DirtyExtracted{
		"titles":       []string{},
		"descriptions": []string{},
		"images":       []string{},
		"urls":         []string{},
	}
}

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
func (t HeadTagsTechnique) Extract(html string) DirtyExtracted {
	return DirtyExtracted{
		"titles":       []string{},
		"descriptions": []string{},
		"images":       []string{},
		"urls":         []string{},
	}
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
