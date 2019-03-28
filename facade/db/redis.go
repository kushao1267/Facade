package db

import (
	"github.com/go-redis/redis"
	"github.com/kushao1267/facade/facade/config"
)

func init() {
	var LinkPreviewService LinkPreview
	if r, ok := NewRedis(config.AllConf.Redis["master"]); ok != nil {
		panic("初始化redis失败!")
	} else {
		LinkPreviewService.Redis = r
	}
}

func NewRedis(c config.Redis) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     c.Addr,
		Password: c.Password,
		DB:       c.DB, // use default DB
	})

	if err := rdb.Ping().Err(); err != nil {
		return nil, err
	}
	return rdb, nil
}

// LinkPreview: 链接预览缓存服务
type LinkPreview struct {
	Redis *redis.Client
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

func (l LinkPreview) GetKey() {

}
func (l LinkPreview) GetValues() {

}
func (l LinkPreview) SetValues() {

}
func (l LinkPreview) Delete() {

}
func (l LinkPreview) BatchDelete() {

}
