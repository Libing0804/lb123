package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()
	for {
		var buf = make([]byte, 1024)
		fmt.Printf("服务器在等待%s发送信息\n", conn.RemoteAddr().String())
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read failed err:", err)
			return
		}
		fmt.Println(string(buf[:n]))
	}

}
func main() {
	fmt.Println("开始监听")
	listen, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		fmt.Println("listen is failed:err", err)
		return
	}
	defer listen.Close()
	for {
		fmt.Println("等待连接")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept failed err:", err)

		} else {
			fmt.Println("accept is succed  cilent ip:", conn.RemoteAddr().String())
		}
		go process(conn)
	}
}
