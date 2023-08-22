package main

import (
	"fmt"
	"github.com/pkg/sftp"
	"io"
	"net/http"
	"os"
)

// 向运程ssh服务器上传本地文件
func sshUpload(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "bad request!", http.StatusBadRequest)
		return
	}
	host := r.FormValue("host")
	port := r.FormValue("port")
	user := r.FormValue("user")
	psw := r.FormValue("psw")
	//本地服务器文件
	localFilePath := r.FormValue("localFilePath")
	//远程服务器文件
	remoteFilePath := r.FormValue("remoteFilePath")

	// 连接SSH服务器
	conn, err := sshConn(host, port, user, psw)
	if err != nil {
		http.Error(w, "Failed to dial: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	// 创建SFTP客户端
	client, err := sftp.NewClient(conn)
	if err != nil {
		http.Error(w, "Failed to create SFTP client:"+err.Error(), http.StatusInternalServerError)
		return
	}
	defer client.Close()

	// 打开本地文件
	localFile, err := os.Open(localFilePath)
	if err != nil {
		http.Error(w, "Failed to open local file: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer localFile.Close()

	// 创建远程文件
	remoteFile, err := client.Create(remoteFilePath)
	if err != nil {
		http.Error(w, "Failed to create remote file:"+err.Error(), http.StatusInternalServerError)
		return
	}
	defer remoteFile.Close()

	// 将本地文件内容复制到远程文件
	_, err = io.Copy(remoteFile, localFile)
	if err != nil {
		http.Error(w, "Failed to copy file content: "+err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("File transferred successfully.")
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}
