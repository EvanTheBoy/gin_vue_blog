package main

import (
	"gin_vue_blog/gin_vue_blog_server/core"
	"gin_vue_blog/gin_vue_blog_server/global"
)

func main() {
	//读取配置文件
	core.InitCore()
	//连接数据库
	global.DB = core.InitGorm()

}
