package utils

import (
	"encoding/json"
	"github.com/kushao1267/facade/facade/logger"
	"github.com/levigross/grequests"
	"net/url"
	"path"
	"regexp"
	"time"
)

const requestUrlTimeout = 6

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
	})
	if err!=nil{
		logger.JsonLogger.Error("GetJson: ", err)
	}
	if err := json.Unmarshal(resp.Bytes(), &v); err != nil {
		logger.JsonLogger.Error("GetJson: ", err)
	}
}

// GetHtml
func GetHtml(url string) string {
	resp, err := grequests.Get(url, &grequests.RequestOptions{
		RequestTimeout: requestUrlTimeout * time.Second,
		Headers:  map[string]string{
			"user-agent":
				"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKi" +
				"t/605.1.15 (KHTML, like Gecko) Version/11.1.2 Safari/605.1.15"},
	})
	if err!=nil{
		logger.JsonLogger.Error("GetHtml", err)
	}

	return resp.String()
}

// CleanHtmlTags 简单地处理html标签
func CleanHtmlTags(raw string) string {
	cleaner := regexp.MustCompile("<.*?>")
	return cleaner.ReplaceAllString(raw, "")
}
