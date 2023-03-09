package global

import (
	"gin_vue_blog/gin_vue_blog_server/config"
	"gorm.io/gorm"
)

var (
	Config *config.Config
	DB     *gorm.DB
)
