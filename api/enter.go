package api

import "gin_vue_blog/gin_vue_blog_server/api/settings_api"

// Group 这是settings_api包下的汇总，这个包下定义的结构体都要在这里说明
type Group struct {
	SettingsAPI settings_api.SettingsAPI //settings_info下面的那个结构体
}

// GroupAPP new的作用是分配内存空间，并返回一个指向零值填充的结构体的指针
var GroupAPP = new(Group)
