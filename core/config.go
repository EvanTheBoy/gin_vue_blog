package core

import (
	"fmt"
	"gin_vue_blog/gin_vue_blog_server/config"
	"gin_vue_blog/gin_vue_blog_server/global"
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

// InitCore 函数名首字母大写表示可以被其他包访问
func InitCore() {
	const ConfigFile = "settings.yaml"
	c := &config.Config{}
	yamlConfig, err := os.ReadFile(ConfigFile)
	if err != nil {
		panic(fmt.Errorf("get yamlConf error: %s", err))
	}
	err = yaml.Unmarshal(yamlConfig, c)
	if err != nil {
		log.Fatalf("config Init Unmarshal: %v", err)
	}
	log.Println("config yamlFile load Init success.")
	global.Config = c
}
