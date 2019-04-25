package db

import (
	"log"
	"time"

	"github.com/go-redis/redis"
	"github.com/kushao1267/Facade/facade/config"
	"github.com/kushao1267/Facade/facade/utils"
	"github.com/mgutz/ansi"
	"errors"
)

var LinkPreviewService *LinkPreview
var redisdb *redis.Client

// See use: https://github.com/go-redis/redis/blob/master/example_test.go
func init() {
	LinkPreviewService = NewLinkPreview()

	redisdb = NewRedis(config.AllConf.Redis)
}

func NewRedis(c config.Redis) *redis.Client {
	log.Println(c)
	db := redis.NewClient(&redis.Options{
		Addr:        c.Addr,
		Password:    c.Password,
		DB:          c.DB, // use default DB
		PoolSize:    10,
		PoolTimeout: 30 * time.Second,
	})

	if err := db.Ping().Err(); err != nil {
		log.Println(ansi.Color("[初始化redis失败]:", "red"), err)
	}
	return db
}

// LinkPreview: 链接预览缓存服务
type LinkPreview struct {
	// store field
	Url         string
	Title       string
	Description string
	Image       string
	ImageStyle  string
	ImageHeight string
	ImageWidth  string
}

func NewLinkPreview() *LinkPreview {
	return &LinkPreview{
		"url",
		"title",
		"description",
		"image",
		"image_style",
		"image_height",
		"image_width",
	}
}

func (l LinkPreview) GetKey(url string) string {
	hash := utils.GetMD5Hash(url)
	return "link_preview_cache:" + hash
}

func (l LinkPreview) GetValues(url string, fields ...string) (error, []string) {
	key := l.GetKey(url)
	s := make([]string, len(fields))
	var empty []string

	val, err := redisdb.HMGet(key, fields...).Result()

	if err != nil {
		return err, empty
	}
	for i, _ := range fields {
		if val[i] != nil {
			s[i] = val[i].(string)
		} else {
			return errors.New("not find"), empty
		}
	}

	return nil, s
}

func (l LinkPreview) SetValues(url string, fields map[string]interface{}) {
	key := l.GetKey(url)

	if err := redisdb.HMSet(key, fields); err != nil {
		log.Println(err)
	}

	if err1 := redisdb.Expire(key, config.AllConf.Redis.Expire * time.Second).Err(); err1 != nil {
		log.Println(err1)
	}
}

func (l LinkPreview) Delete(url string) {
	key := l.GetKey(url)
	if err := redisdb.Del(key).Err(); err != nil {
		panic(err)
	}
}
