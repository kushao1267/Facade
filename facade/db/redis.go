package db

import (
	"github.com/go-redis/redis"
	"github.com/kushao1267/Facade/facade/config"
	"github.com/kushao1267/Facade/facade/utils"
	"github.com/mgutz/ansi"
	"log"
	"time"
)

var LinkPreviewService *LinkPreview
var redisdb *redis.Client

// See use: https://github.com/go-redis/redis/blob/master/example_test.go
func init() {
	LinkPreviewService = NewLinkPreview()

	redisdb = NewRedis(config.AllConf.Redis["master"])
}

func NewRedis(c config.Redis) *redis.Client {
	db := redis.NewClient(&redis.Options{
		Addr:        c.Addr,
		Password:    c.Password,
		DB:          c.DB, // use default DB
		PoolSize:    10,
		PoolTimeout: 30 * time.Second,
	})

	if err := db.Ping().Err(); err != nil {
		log.Println(ansi.Color("[初始化redis失败]:","red"), err)
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

func (l LinkPreview) GetValues(url string, fields ...string) []string {
	key := l.GetKey(url)
	s := make([]string, len(fields))

	val, err := redisdb.HMGet(key, fields...).Result()

	if err != nil {
		log.Println(ansi.Color("[link_preview_cache:GetValues]:","red"), err)
		return []string{}
	} else if err == redis.Nil { // key does not exists
		for i, _ := range fields {
			s[i] = ""
		}
	} else {
		for i, _ := range fields {
			s[i] = val[i].(string)
		}
	}
	return s
}

func (l LinkPreview) SetValues(url string, fields map[string]interface{}) {
	key := l.GetKey(url)

	if err := redisdb.HMSet(key, fields); err != nil {
		panic(err)
	}

	if err1 := redisdb.Expire(key, config.AllConf.Expire).Err(); err1 != nil {
		panic(err1)
	}
}

func (l LinkPreview) Delete(url string) {
	key := l.GetKey(url)
	if err := redisdb.Del(key).Err(); err != nil {
		panic(err)
	}
}
