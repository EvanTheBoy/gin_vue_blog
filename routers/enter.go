package routers

import (
	"gin_vue_blog/gin_vue_blog_server/global"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	router := gin.Default()
	router.GET("", func(c *gin.Context) {
		c.String(200, "浩文, 你一定可以申请到很好的top级别的美国大学的, 要相信自己呀!")
	})
	return router
}
