package main

import (
	"log"
	"net"
	"time"
)

func establishConn(i int) net.Conn {
	conn, err := net.Dial("tcp", ":80")
	if err != nil {
		log.Printf("%d: 拨号失败哦：%s", i, err)
		return nil
	}
	log.Printf("%d: 拨号成功！", i)
	return conn
}

func main() {
	var connArray []net.Conn
	for i := 1; i < 1000; i++ {
		conn := establishConn(i)
		if conn != nil {
			connArray = append(connArray, conn)
		}
	}
	time.Sleep(10000 * time.Second)
}
