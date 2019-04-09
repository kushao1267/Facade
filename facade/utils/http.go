package utils

import (
	"net/url"
	"path"
	"regexp"
)

func GetComponent(s string) (*url.URL, error) {
	u, err := url.Parse(s)
	return u, err
}

func GetHostName(s string) (string, error) {
	u, err := url.Parse(s)
	if err != nil {
		return "", err
	}
	return u.Host, nil
}

func UrlJoin(source, target string) string {
	return path.Join(source, target)
}

// CleanHtmlTags 简单地处理html标签
func CleanHtmlTags(raw string) string {
	cleaner := regexp.MustCompile("<.*?>")
	return cleaner.ReplaceAllString(raw, "")
}
