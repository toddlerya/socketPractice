package main

import (
	"log"
	"net"
	"time"
)

func main() {
	conn, err := net.DialTimeout("tcp", "104.236.176.96:80", 2*time.Millisecond)
	if err != nil {
		log.Printf("拨号失败哦：%s\n", err)
	}
	defer conn.Close()
	log.Println("拨号成功！")
}
