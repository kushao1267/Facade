package services

import (
	"errors"
	"github.com/kushao1267/Facade/facade/config"
	"github.com/kushao1267/Facade/facade/db"
	"github.com/kushao1267/Facade/facade/utils"
	"log"
	"time"
)

// LinkPreview store field
//Url         string
//Title       string
//Description string
//Image       string
//ImageStyle  string
//ImageHeight string
//ImageWidth  string
type LinkPreview struct{}

// LinkPreviewService: 链接预览缓存服务
var LinkPreviewService = new(LinkPreview)

func (l LinkPreview) GetFiels() []string {
	return []string{"title", "description", "image"}
}

// GetKey link preview cache like "link_preview_cache:${url_hash}"
func (l LinkPreview) GetKey(url string) string {
	hash := utils.GetMD5Hash(url)
	return "link_preview_cache:" + hash
}

// GetValues get link preview cache
func (l LinkPreview) GetValues(url string) (error, []string) {
	key := l.GetKey(url)
	fields := l.GetFiels()
	s := make([]string, len(fields))
	var empty []string

	val, err := db.RedisDB.HMGet(key, fields...).Result()

	if err != nil {
		return err, empty
	}
	for i := range fields {
		if val[i] != nil {
			s[i] = val[i].(string)
		} else {
			return errors.New("not find"), empty
		}
	}

	return nil, s
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
