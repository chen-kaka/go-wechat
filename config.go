package main

import (
	"github.com/jimlawless/cfg"
	"fmt"
	"log"
)

func main() {
	mymap := make(map[string]string)
	err := cfg.Load("config.conf", mymap)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", mymap["token"])
}

func GetCofigMap() map[string]string {
	configMap := make(map[string]string)
	err := cfg.Load("/Users/kakachan/go/conf/go-wechat/config.conf", configMap)
	if err != nil {
		log.Fatal(err)
	}
	return configMap
}