package settings_api

import "github.com/gin-gonic/gin"

// SettingsAPIView 这个方法是绑定给enter.go中的SettingsAPI结构体的
func (SettingsAPI) SettingsAPIView(c *gin.Context) {
	//以下代码中，gin.H{}是Go语言中常用的一个返回数据格式，也是map[string]interface{}的简写
	//gin.H{}可以被用来构造JSON, XML, yaml等格式的数据
	c.JSON(200, gin.H{"message": "浩文，你一定可以的，你一定可以申请到美国TOP级别的大学的！"})
}
