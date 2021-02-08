package config

import (
	"fmt"
	"io/ioutil"

	jsoniter "github.com/json-iterator/go"
)

type sysconfig struct {
	Port       string `json:"Port"`
	DBUserName string `json:"DBUserName"`
	DBPassword string `json:"DBPassword"`
	DBIp       string `json:"DBIp"`
	DBPort     string `json:"DBPort"`
	DBName     string `json:"DBName"`
}

var Sysconfig = &sysconfig{}

func init() {
	b, err := ioutil.ReadFile("./config.json")
	if err != nil {
		fmt.Printf("error:%s", err)
		panic(err)
	} else {
		err = jsoniter.Unmarshal(b, Sysconfig)
		if err != nil {
			fmt.Printf("error:%s", err)
			panic(err)
		}
	}
}
