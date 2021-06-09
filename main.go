package main

import (
	"github.com/weilaim/blog-api/conf"
	"github.com/weilaim/blog-api/server"
)

func main(){
	//读取配置
	conf.Init()

	//装载路由
	r := server.NewRouter()
	r.Run(":3000")
}