package main

import (
	"bytes"
	"io"
	"net/http"
)

// ssh 执行命令
func webssh(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "bad request!", http.StatusBadRequest)
		return
	}
	host := r.FormValue("host")
	port := r.FormValue("port")
	user := r.FormValue("user")
	psw := r.FormValue("psw")
	// 准备执行的远程命令
	cmd := r.FormValue("cmd")

	// 连接SSH服务器
	conn, err := sshConn(host, port, user, psw)
	if err != nil {
		http.Error(w, "Failed to dial: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	// 创建一个新的会话
	session, err := conn.NewSession()
	if err != nil {
		http.Error(w, "Failed to create session: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer session.Close()

	// 将命令输出重定向到标准输出
	//session.Stdout = os.Stdout
	//session.Stderr = os.Stderr

	// 创建一个缓冲区来捕获命令输出
	var stdoutBuf bytes.Buffer
	session.Stdout = &stdoutBuf

	// 执行远程命令
	err = session.Run(cmd)
	if err != nil {
		http.Error(w, "Failed to run: "+err.Error(), http.StatusInternalServerError)
		return
	}
	// 将缓冲区内容作为HTTP响应发送回客户端
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	io.Copy(w, &stdoutBuf)

}
