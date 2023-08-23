package main

import (
	"fmt"
	"net/http"
)

var urlMap map[string]http.HandlerFunc

func init() {
	//路由映射
	urlMap = map[string]http.HandlerFunc{
		"/hs/fu":     uploadHandler,
		"/hs/fd":     downloadHandler,
		"/hs/se":     webssh,
		"/hs/su":     sshUpload,
		"/hs/def/se": webSshDef,
	}
}

// 路由
func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("url:", r.URL.Path, " RequestURI:", r.RequestURI)
	doHandle, exists := urlMap[r.URL.Path]
	if exists {
		doHandle(w, r)
	} else {
		http.Error(w, r.URL.Path+" 404 not found !", http.StatusNotFound)
	}

}

// 全局过滤器
func WithTokenValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 从请求头中获取Token
		token := r.Header.Get("security")
		// 验证Token
		if token != Config.Security {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		// 继续处理下一个中间件或处理程序
		next.ServeHTTP(w, r)
	})
}
