package db

import (
	"time"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/kushao1267/facade/facade/config"
	"github.com/kushao1267/facade/facade/utils"
)

var LinkPreviewService *LinkPreview
var redisdb *redis.Client

// See use: https://github.com/go-redis/redis/blob/master/example_test.go
func init() {
	LinkPreviewService.Init()

	redisdb = NewRedis(config.AllConf.Redis["master"])
}

func NewRedis(c config.Redis) *redis.Client{
	db := redis.NewClient(&redis.Options{
		Addr:     c.Addr,
		Password: c.Password,
		DB:       c.DB, // use default DB
		PoolSize:     10,
		PoolTimeout:  30 * time.Second,
	})

	if err := db.Ping().Err(); err != nil {
		panic("初始化redis失败!")
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

func (l *LinkPreview) Init() {
	l.Url = "url"
	l.Title = "title"
	l.Description = "description"
	l.Image = "image"
	l.ImageStyle = "image_style"
	l.ImageHeight = "image_height"
	l.ImageWidth = "image_width"
}

func (l LinkPreview) GetKey(url string) string{
	hash := utils.GetMD5Hash(url)
	return fmt.Sprintf("link_preview_cache:%s", hash)
}

func (l LinkPreview) GetValues(url string, fields ...string) []string{
	key := l.GetKey(url)
	s := make([]string, len(fields))

	val, err :=redisdb.HMGet(key, fields...).Result()

	if err != nil {
		panic(err)
	}else if  err == redis.Nil { // key does not exists
		for i, _ := range fields {
			s[i] = ""
		}
	}else {
		for i, _ := range fields {
			s[i] = val[i].(string)
		}
	}
	return s
}

func (l LinkPreview) SetValues(url string, fields map[string]interface{}) {
	key := l.GetKey(url)

	if err := redisdb.HMSet(key, fields);err != nil {
		panic(err)
	}

	if err1 :=redisdb.Expire(key, config.AllConf.Expire).Err();err1 !=nil{
		panic(err1)
	}
}

func (l LinkPreview) Delete(url string) {
	key := l.GetKey(url)
	if err := redisdb.Del(key).Err();err!=nil{
		panic(err)
	}
}

