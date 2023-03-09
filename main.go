package main

import (
	"fmt"
	"gin_vue_blog/gin_vue_blog_server/core"
	"gin_vue_blog/gin_vue_blog_server/global"
)

func main() {
	core.InitCore()
	fmt.Println(global.Config)
}
