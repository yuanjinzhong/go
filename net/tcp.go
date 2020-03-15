package main

import (
	"fmt"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:3569")
	if err != nil {
		fmt.Println("发生错误:", err)
		return
	}
	for true {
		conon, err := listen.Accept()
		if err != nil {
			fmt.Println("监听发生错误:", err)
			continue
		}
		go Read(conon)
	}
}

func Read(conn net.Conn) {
	data := make([]byte, 1000)
	for {
		n, err := conn.Read(data)
		if err != nil {
			fmt.Println("从链接读取数据发生异常:", err)
			break
		}
		fmt.Println(string(data[0:n]))
	}
}
