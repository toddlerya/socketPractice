package main

import (
	"fmt"
	"net"
	"strings"
)

func connHandler(conn net.Conn) {
	buf := make([]byte, 4096)
	for {
		cnt, err := conn.Read(buf)
		if cnt == 0 || err != nil {
			err := conn.Close()
			if err != nil {
				panic(err)
			}
			break
		}
		inStr := strings.TrimSpace(string(buf[0:cnt]))
		inputs := strings.Split(inStr, " ")
		switch inputs[0] {
		case "ping":
			_, err = conn.Write([]byte("pong\n"))
			if err != nil {
				panic(err)
			}
		case "echo":
			echoStr := strings.Join(inputs[1:], " ") + "\n"
			_, err = conn.Write([]byte(echoStr))
			if err != nil {
				panic(err)
			}
		case "quit":
			err = conn.Close()
			if err != nil {
				panic(err)
			}
		default:
			prompt := fmt.Sprintf("不支持的参数哦，仅支持\n" +
				"[1] ping\n" +
				"[2] echo some text\n" +
				"[3] quit\n")
			_, err = conn.Write([]byte(prompt))
			if err != nil {
				panic(err)
			}
		}
	}
	fmt.Printf("已断开来自 %s 的连接\n", conn.RemoteAddr())
}

func main() {
	port := "80"
	server, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		fmt.Printf("无法在%s端口启动服务\n", port)
		return
	}
	fmt.Printf("服务已经在%s端口启动\n", port)
	for {
		conn, err := server.Accept()
		fmt.Printf("已接受来自 %s 的连接\n", conn.RemoteAddr())
		if err != nil {
			fmt.Printf("连接失败：%s\n", err)
			return
		}
		go connHandler(conn)
	}
}
