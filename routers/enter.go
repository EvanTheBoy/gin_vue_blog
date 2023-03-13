package routers

import (
	"gin_vue_blog/gin_vue_blog_server/api"
	"gin_vue_blog/gin_vue_blog_server/global"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	router := gin.Default()
	settingsAPI := api.GroupAPP.SettingsAPI
	router.GET("", settingsAPI.SettingsAPIView)
	return router
}
