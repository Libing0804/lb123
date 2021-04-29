package main

import (
	"fmt"
	"net"
	"time"

	"github.com/tcp/chatroom/server/model"
)

// //编写一个serverProcessLogin函数处理登录
// func serverProcessLogin(conn net.Conn, mes *message.Message) (err error) {
// 	//1. 先从mes中取出mes.data在反序列化
// 	var loginMes message.LogiMes
// 	err = json.Unmarshal([]byte(mes.Data), &loginMes) //地址传入
// 	if err != nil {
// 		fmt.Println("json.unmarshal is failed err:", err)
// 		return
// 	}
// 	var resmes message.Message
// 	resmes.Type = message.RegisterMesType
// 	//在生命一个LoginresMes
// 	var loginResMes message.LogiResMes
// 	//如果用户的账号是100 密码是123456就是对的
// 	if loginMes.UserId == 100 && loginMes.UserPwd == "123456" {
// 		//合法
// 		loginResMes.Code = 200

// 	} else {
// 		//不合法
// 		loginResMes.Code = 500 //表示该用户不存在
// 		loginResMes.Error = "表示该用户不存在"

// 	}
// 	//将loginResMes序列化
// 	data, err := json.Marshal(loginResMes)
// 	if err != nil {
// 		fmt.Println("loginResMes json marshal failed err:", err)
// 		return
// 	}
// 	resmes.Data = string(data)
// 	//对resmes序列化
// 	data, err = json.Marshal(resmes)
// 	if err != nil {
// 		fmt.Println("resmes json.marshal failed err:", err)
// 		return
// 	}
// 	err = writePkg(conn, data)

// 	return
// }

// //编写一个serverprocessMes函数
// //功能：根据客户端发来信息的种类的不同决定调用那个函数来处理
// func serverProcessMes(conn net.Conn, mes *message.Message) (err error) {
// 	switch mes.Type {
// 	case message.LoginMesType:
// 		//处理登录
// 		err = serverProcessLogin(conn, mes)
// 	case message.RegisterMesType:
// 		//处理注册的
// 	default:
// 		fmt.Println("消息类型不存在")
// 	}
// 	return
// }
// func readPkg(conn net.Conn) (mes message.Message, err error) {
// 	buf := make([]byte, 8094)

// 	_, err = conn.Read(buf[:4])
// 	if err != nil {
// 		//err = errors.New("read head error")
// 		return
// 	}
// 	//fmt.Println("读到的数据是buf=", buf[:n])
// 	var pkgLen uint32
// 	pkgLen = binary.BigEndian.Uint32(buf[:4])
// 	n, err := conn.Read(buf[:pkgLen])
// 	if n != int(pkgLen) || err != nil {
// 		//err = errors.New("read pkg body error")
// 		return
// 	}
// 	//把pkglen反序列化成message.Message
// 	err = json.Unmarshal(buf[:pkgLen], &mes)
// 	if err != nil {
// 		fmt.Println("json unmarshal failed err:", err)
// 		return
// 	}
// 	return
// }
// func writePkg(conn net.Conn, data []byte) (err error) {
// 	//先发送对方一个长度给对方
// 	var pkgLen uint32
// 	pkgLen = uint32(len(data))
// 	var buf [4]byte
// 	binary.BigEndian.PutUint32(buf[0:4], pkgLen)
// 	n, err := conn.Write(buf[0:4])
// 	if n != 4 || err != nil {
// 		fmt.Println("send is failed err:", err)
// 		return
// 	}
// 	//发送数据本身
// 	n, err = conn.Write(data)
// 	if n != int(pkgLen) || err != nil {
// 		fmt.Println("send is failed err:", err)
// 		return
// 	}
// 	return
// }
func process(conn net.Conn) {
	//读取客户端的信息

	defer conn.Close()
	//创建一个总控
	processor := &Processor{
		Conn: conn,
	}
	err := processor.process2()
	if err != nil {
		fmt.Println("客户端和服务端协成出错误")
		return
	}
}

//编写一个函数完成userDao的初始化
func initUserDao() {
	model.MyUserDao = model.NewUserDao(pool) //传进来的额pool就是redis.go里的全局pool
}
func main() {
	//初始化连接池
	initPool("localhost:6379", 16, 0, 300*time.Second)
	//初始化userDao
	initUserDao()
	fmt.Println("服务器在8889监听")
	listen, err := net.Listen("tcp", "0.0.0.0:8889")
	if err != nil {
		fmt.Println("net.listen is failed:,err", err)
		return
	}
	defer listen.Close()
	for {
		fmt.Println("等待客户端来连接")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.accpte is failed errr:", err)
		}
		//一旦连接成功启动一个协成
		go process(conn)
	}
}
