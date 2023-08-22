package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// 上传文件
func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// 解析表单数据
		err := r.ParseMultipartForm(1000 << 20) // 限制最大文件大小为 1000MB
		if err != nil {
			http.Error(w, "无法解析表单数据", http.StatusInternalServerError)
			return
		}

		// 获取文件
		file, handler, err := r.FormFile("file")
		if err != nil {
			http.Error(w, "无法获取文件", http.StatusBadRequest)
			return
		}
		defer file.Close()

		// 创建保存文件的目标路径
		targetDir := "./uploads"
		os.MkdirAll(targetDir, os.ModePerm)

		// 创建保存文件的目标文件
		targetFile, err := os.Create(filepath.Join(targetDir, handler.Filename))
		if err != nil {
			http.Error(w, "无法创建目标文件", http.StatusInternalServerError)
			return
		}
		defer targetFile.Close()

		// 将上传的文件内容拷贝到目标文件
		_, err = io.Copy(targetFile, file)
		if err != nil {
			http.Error(w, "无法写入文件内容", http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "文件上传成功")
	} else {
		http.Error(w, "仅支持 POST 请求", http.StatusMethodNotAllowed)
	}

}

// 下载文件
func downloadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "bad request!", http.StatusBadRequest)
		return
	}
	filepath := r.FormValue("file")
	// 打开要下载的文件
	file, err := os.Open(filepath)

	if err != nil {
		http.Error(w, "无法打开文件", http.StatusInternalServerError)
		return
	}
	defer file.Close()
	nameArray := strings.Split(file.Name(), "\\")
	fileName := nameArray[len(nameArray)-1]
	//fmt.Println(fileName)
	// 设置响应头，指定文件名和内容类型
	w.Header().Set("Content-Disposition", "attachment; filename="+fileName)
	w.Header().Set("Content-Type", "application/octet-stream") // 通用的二进制流类型

	// 将文件内容写入响应体
	_, err = io.Copy(w, file)
	if err != nil {
		http.Error(w, "无法发送文件内容", http.StatusInternalServerError)
		return
	}
}
