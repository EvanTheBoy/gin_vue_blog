package main

import (
	"gin_vue_blog/gin_vue_blog_server/core"
	"gin_vue_blog/gin_vue_blog_server/global"
	"gin_vue_blog/gin_vue_blog_server/routers"
)

func main() {
	//读取配置文件
	core.InitCore()
	//初始化日志
	global.Log = core.InitLogger()
	//连接数据库
	global.DB = core.InitGorm()

	router := routers.InitRouter()
	addr := global.Config.System.Addr()
	global.Log.Infof("项目运行在:%s", addr)
	err := router.Run(addr)
	if err != nil {
		return
	}
}
