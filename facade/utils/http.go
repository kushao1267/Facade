package utils

import (
	"encoding/json"
	"log"
	"net/url"
	"path"
	"time"

	"github.com/denisbrodbeck/striphtmltags"
	"github.com/levigross/grequests"
	"github.com/mgutz/ansi"
)

const (
	requestURLTimeout = 5
	userAgent         = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/11.1.2 Safari/605.1.15"
)

// GetComponent ...
func GetComponent(s string) (*url.URL, error) {
	u, err := url.Parse(s)
	return u, err
}

// GetHostName get host name
func GetHostName(s string) (string, error) {
	u, err := url.Parse(s)
	if err != nil {
		return "", err
	}
	return u.Host, nil
}

// URLJoin join url
func URLJoin(source, target string) string {
	return path.Join(source, target)
}

// GetJSON ...
func GetJSON(url string, v interface{}) {
	resp, err := grequests.Get(url, &grequests.RequestOptions{
		RequestTimeout: requestImgTimeout,
		UserAgent:      userAgent,
	})
	if err != nil {
		log.Println(ansi.Color("[GetJSON]:", "red"), err)
	}
	if err := json.Unmarshal(resp.Bytes(), &v); err != nil {
		log.Println(ansi.Color("[GetJSON]:", "red"), err)
	}
}

// GetHtml ...
func GetHtml(url string) (string, error) {
	resp, err := grequests.Get(url, &grequests.RequestOptions{
		RequestTimeout: requestURLTimeout * time.Second,

		UserAgent: userAgent,
	})
	if err != nil {
		log.Println(ansi.Color("[GetHtml]:", "red"), err)
		return "", err
	}

	return resp.String(), nil
}

// CleanHtmlTags 简单地处理html标签
func CleanHtmlTags(raw string) string {
	return striphtmltags.StripTags(raw)
}
