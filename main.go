package main

import (
	"fmt"
	"gin_vue_blog/gin_vue_blog_server/core"
	"gin_vue_blog/gin_vue_blog_server/global"
)

func main() {
	core.InitCore()
	global.DB = core.InitGorm()
	fmt.Println(global.DB)
}
