package main

import (
	"bytes"
	"golang.org/x/crypto/ssh"
	"log"
	"os"
)

var logConn = log.New(os.Stdout, "sshConn:", log.LstdFlags)

// 创建ssh连接
func sshConn(host, port, user, psw string) (*ssh.Client, error) {
	// SSH配置信息
	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			// You can use password or key authentication
			ssh.Password(psw),
			// or
			// ssh.PublicKeys(readPrivateKey("path_to_private_key_file")),
		},
		// You might need to adjust this based on your use case
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	logConn.Println("create conn :", host, port)
	// 连接SSH服务器
	conn, err := ssh.Dial("tcp", host+":"+port, config)
	if err != nil {
		logConn.Println("Failed to dial: ", err.Error())
		return nil, err
	}
	//defer conn.Close()
	return conn, nil
}

// 执行命令
func executeCmd(host, port, user, psw, cmd string) (*bytes.Buffer, error) {
	// 连接SSH服务器
	conn, err := sshConn(host, port, user, psw)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	// 创建一个新的会话
	session, err := conn.NewSession()
	if err != nil {
		return nil, err
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
		return nil, err
	}
	return &stdoutBuf, nil
}
