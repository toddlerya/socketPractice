package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	server, err := net.Listen("tcp", ":80")
	if err != nil {
		fmt.Printf("服务端启动监听失败：%s\n", err)
		return
	}
	defer server.Close()
	log.Println("服务端启动监听成功！")
	var count int
	for {
		time.Sleep(10 * time.Second)
		if _, err := server.Accept(); err != nil {
			log.Println("接受连接请求失败！", err)
			break
		}
		count++
		log.Printf("%d: 接受一个新的连接请求！\n", count)
	}

}
