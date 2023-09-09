package main

import (
	"hello/routes"
	"net/http"
)

func main() {
	// 定义路由
	r := routes.SetupRouter()

	// 监听端口
	http.ListenAndServe(":8080", r)
}
