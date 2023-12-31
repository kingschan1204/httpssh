package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Yaml struct {
	App appConfig `yaml:"app"`
}

type appConfig struct {
	Port     string `yaml:"port"`     // bind port
	Security string `yaml:"security"` // http token
	Default  *def   `yaml:def`        // default sever seting
}

type def struct {
	Host string `yaml:"host"` // host
	Port string `yaml:"port"` // port
	User string `yaml:"user"` // user
	Psw  string `yaml:"psw"`  // psw
}

////////////////////////////////////////////////////////

var Config appConfig

func InitConfig() {
	//获取当前目录
	//fmt.Println(os.Getwd())
	filename := "./app.yaml"
	y := new(Yaml)
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("read gateway.yaml file error %v\n", err)
	}
	err = yaml.Unmarshal(yamlFile, y)
	if err != nil {
		log.Fatalf("yaml 解码失败: %v\n", err)
	}
	Config = y.App

}
