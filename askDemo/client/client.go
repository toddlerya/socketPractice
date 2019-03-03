package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func connHandler(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(os.Stdin)
	buf := make([]byte, 1024)
	for {
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if input == "quit" {
			return
		}
		_, err := conn.Write([]byte(input))
		if err != nil {
			panic(err)
		}
		cnt, err := conn.Read(buf)
		if err != nil {
			fmt.Printf("无法读取数据, %s\n", err)
			continue
		}
		fmt.Print(string(buf[0:cnt]))
	}
}

func main() {
	port := "80"
	conn, err := net.Dial("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		fmt.Printf("无法连接到 %s 端口", port)
	}
	fmt.Printf("已经连接到 %s 端口，开始发消息吧！", port)
	connHandler(conn)
}
