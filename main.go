package main

import (
	"log"

	"gin-study/internal"
)

// main 程序入口
//
//	@title						gin学习项目
//	@version					1.0
//	@description				使用gin框架学习web开发时会使用到的功能.
//	@host						localhost:8000
//	@BasePath					/v1
//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						token
//	@description				使用token传递用户认证信息
func main() {
	if err := internal.Run("./config/user.yaml"); err != nil {
		log.Fatal(err)
	}
}
