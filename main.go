package main

import (
	"fmt"
	"log"
	"net/http"
)

//http://patorjk.com/software/taag/#p=display&f=3D-ASCII&t=app
var banner string = `
 ________  ________  ________   
|\   __  \|\   __  \|\   __  \  
\ \  \|\  \ \  \|\  \ \  \|\  \ 
 \ \   __  \ \   ____\ \   ____\
  \ \  \ \  \ \  \___|\ \  \___|
   \ \__\ \__\ \__\    \ \__\   
    \|__|\|__|\|__|     \|__|   
                                
`

func init() {
	fmt.Println(banner)
	InitConfig()
	fmt.Println("app initialized with port(s): ", Config.Port)
}
func main() {
	//全局过滤器
	http.Handle("/", WithTokenValidation(http.HandlerFunc(Handler)))
	/*
		//上传文件
		http.HandleFunc("/hs/fu", uploadHandler)
		//下载文件
		http.HandleFunc("/hs/fd", downloadHandler)
		//ssh 执行命令
		http.HandleFunc("/hs/se", webssh)
		//ssh 上传文件
		http.HandleFunc("/hs/su", sshUpload)*/
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", Config.Port), nil))
}
