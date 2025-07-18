package models

import (
	"fmt"
	"os"

	"github.com/mitchellh/mapstructure"
)

var Configs map[string]interface{}

var Configs_wed Configs_web_t
var Configs_user Configs_user_t
var Configs_file Configs_file_t

var Database_configs map[string]interface{}

var Allowed_avatar_ext = map[string]bool{
	".jpg":  true,
	".jpeg": true,
	".png":  true,
}

var Allowed_avatar_mime = map[string]bool{
	"image/jpeg": true,
	"image/png":  true,
}

func init() {

}

func All_config_init() {

	//初始化数据库
	Init_database()

	//读取web配置
	err := mapstructure.Decode(Configs["web"].(map[string]interface{}), &Configs_wed)
	if err != nil {
		panic(err)
	}

	//初始化user config
	err = mapstructure.Decode(Configs["user"].(map[string]interface{}), &Configs_user)
	if err != nil {
		panic(err)
	}

	//初始化file config
	err = mapstructure.Decode(Configs["file"].(map[string]interface{}), &Configs_file)
	if err != nil {
		panic(err)
	}
	//创建file的关键文件夹
	for _, value := range Configs_file.Pahts {
		err := os.MkdirAll(value, 0755)
		if err != nil {
			fmt.Printf("创建文件夹失败: %v\n", err)
			panic("创建文件夹失败")

		}
	}

}
