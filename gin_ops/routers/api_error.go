package routers

import (
	"encoding/json"
	"fmt"
	"os"
)

var Error_code map[string]interface{}

func init() {
	//读取默认配置
	fmt.Println("尝试读取错误码文件")
	data, err := os.ReadFile("./def_config/error_codes.json")
	if err != nil {

		fmt.Println("读取错误码文件失败", err)
	}

	if err := json.Unmarshal(data, &Error_code); err != nil {
		fmt.Println("解析错误码文件失败", err)
	}

}
