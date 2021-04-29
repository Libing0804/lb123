package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "10.52.242.177:8080")
	if err != nil {
		fmt.Println("connect failed err:", err)
		return
	}
	for {
		reader := bufio.NewReader(os.Stdin)
		str, err := reader.ReadString('\n')
		str = strings.Trim(str, " \r\n")
		if str == "exit" {
			break
		}
		if err != nil {
			fmt.Println("reader failed err:", err)

		}
		n, err := conn.Write([]byte(str))
		if err != nil {
			fmt.Println("send faild err:", err)
		}
		fmt.Printf("此次发送了%d个字节/n", n)
	}
}
