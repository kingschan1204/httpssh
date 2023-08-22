package main

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"log"
	"os"
)

var logConn = log.New(os.Stdout, "sshConn:", log.Llongfile|log.LstdFlags)

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
		fmt.Println("Failed to dial: ", err.Error())
		return nil, err
	}
	//defer conn.Close()
	return conn, nil
}
