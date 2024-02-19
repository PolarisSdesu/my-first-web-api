package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", corsWrapper(handler))
	http.ListenAndServe(":8080", nil)
}

// handler 是原始的请求处理函数
func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // Content-Type 头部
	fmt.Fprintf(w, `{"message": "Hello, World!"}`)
}

// corsWrapper 创建一个包装器函数来处理 CORS 头部信息
func corsWrapper(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")             // 允许任何源
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS") // 允许 GET 和 OPTIONS 方法（这里仅做示例，实际应根据需求配置）

		if r.Method == "OPTIONS" {
			// 如果是预检（OPTIONS）请求，则结束处理并返回
			w.WriteHeader(http.StatusNoContent)
			return
		}

		// 调用原始的请求处理函数
		h(w, r)
	}
}
