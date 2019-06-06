package utils

import (
	"encoding/json"
	"github.com/denisbrodbeck/striphtmltags"
	"github.com/levigross/grequests"
	"github.com/mgutz/ansi"
	"log"
	"net/url"
	"path"
	"time"
)

const (
	requestUrlTimeout = 6
	userAgent         = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/11.1.2 Safari/605.1.15"
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

// GetJson
func GetJson(url string, v interface{}) {
	resp, err := grequests.Get(url, &grequests.RequestOptions{
		RequestTimeout: requestImgTimeout,
		UserAgent:      userAgent,
	})
	if err != nil {
		log.Println(ansi.Color("[GetJson]:", "red"), err)
	}
	if err := json.Unmarshal(resp.Bytes(), &v); err != nil {
		log.Println(ansi.Color("[GetJson]:", "red"), err)
	}
}

// GetHtml
func GetHtml(url string) string {
	resp, err := grequests.Get(url, &grequests.RequestOptions{
		RequestTimeout: requestUrlTimeout * time.Second,
		UserAgent:      userAgent,
	})
	if err != nil {
		log.Println(ansi.Color("[GetHtml]:", "red"), err)
	}

	return resp.String()
}

// CleanHtmlTags 简单地处理html标签
func CleanHtmlTags(raw string) string {
	return striphtmltags.StripTags(raw)
}
