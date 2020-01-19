package services

import (
	"errors"
	"log"
	"time"

	"github.com/kushao1267/Facade/facade/config"
	"github.com/kushao1267/Facade/facade/db"
	"github.com/kushao1267/Facade/facade/utils"
)

// LinkPreview store field
// url         string
// title       string
// description string
// image       string
// ImageStyle  string
// ImageHeight string
// ImageWidth  string
type LinkPreview struct{}

// LinkPreviewService 链接预览缓存服务
var LinkPreviewService = new(LinkPreview)

// GetFields 获取preview的所有字段
func (l LinkPreview) GetFields() []string {
	return []string{"title", "description", "image"}
}

// GetKey link preview cache like "link_preview_cache:${url_hash}"
func (l LinkPreview) GetKey(url string) string {
	hash := utils.GetMD5Hash(url)
	return "link_preview_cache:" + hash
}

// GetValues get link preview cache
func (l LinkPreview) GetValues(url string) ([]string, error) {
	key := l.GetKey(url)
	fields := l.GetFields()
	s := make([]string, len(fields))
	var empty []string

	val, err := db.RedisDB.HMGet(key, fields...).Result()

	if err != nil {
		return empty, err
	}
	for i := range fields {
		if val[i] != nil {
			s[i] = val[i].(string)
		} else {
			return empty, errors.New("not find")
		}
	}

	return s, nil
}

// SetValues set link preview cache
func (l LinkPreview) SetValues(url string, fields map[string]interface{}) {
	key := l.GetKey(url)

	db.RedisDB.HMSet(key, fields)

	if err1 := db.RedisDB.Expire(key, config.AllConf.Redis.Expire*time.Second).Err(); err1 != nil {
		log.Println(err1)
	}
}

// Delete delete link preview cache
func (l LinkPreview) Delete(url string) {
	key := l.GetKey(url)
	if err := db.RedisDB.Del(key).Err(); err != nil {
		panic(err)
	}
}
