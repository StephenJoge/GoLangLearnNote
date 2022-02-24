package main

import "net/http"

//测试文件服务路径访问
func main() {
	http.Handle("/", http.FileServer(http.Dir("..")))
	http.ListenAndServe(":8080", nil)
}
