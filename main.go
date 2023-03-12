package main

import (
	"gin_vue_blog/gin_vue_blog_server/core"
	"gin_vue_blog/gin_vue_blog_server/global"
)

func main() {
	//读取配置文件
	core.InitCore()
	//初始化日志
	global.Log = core.InitLogger()
	global.Log.Warnln("啊哈哈")
	global.Log.Error("嘿嘿嘿")
	global.Log.Infof("啊呀呀呀")
	//连接数据库
	global.DB = core.InitGorm()
}
