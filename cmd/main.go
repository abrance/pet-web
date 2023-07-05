package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// 静态文件目录
	fs := http.FileServer(http.Dir("build"))

	// 所有请求都会被重定向到 index.html
	//http.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
	//	http.ServeFile(w, r, "build/index.html")
	//})

	// 将静态文件目录和路由处理器绑定到服务器
	http.Handle("/", http.StripPrefix("/", fs))

	fmt.Println("Server is running on port 9000")
	log.Fatal(http.ListenAndServe(":9000", nil))
}
