package process

import (
	"encoding/binary"
	"encoding/json"
	"fmt"

	"net"
	"os"

	"github.com/tcp/chatroom/client/utils"
	"github.com/tcp/chatroom/common/message"
)

type UserProcess struct {
}

//注册
func (this *UserProcess) Register(userId int, userPwd string, userName string) (err error) {
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("conn dial failed err:", err)
		return
	}
	defer conn.Close()
	var mes message.Message
	mes.Type = message.RegisterMesType
	//创建一个registerMes
	var registerMes message.RegisterMes
	registerMes.User.UserId = userId
	registerMes.User.UserPwd = userPwd
	registerMes.User.UserName = userName
	//将registerMes序列化
	data, err := json.Marshal(registerMes)
	if err != nil {
		fmt.Println("json.Marshal is failed err:", err)
		return
	}
	//把data赋给mer.data字段
	mes.Data = string(data)
	//将mes进行序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal is failed err:", err)
		return
	}
	tf := &utils.Transfer{
		Conn: conn,
	}
	err = tf.WritePkg(data)
	if err != nil {

		fmt.Println(" register write failed err:", err)
		return
	}
	mes, err = tf.ReadPkg()

	if err != nil {
		fmt.Println("readpkg failed err", err)
		return
	}
	var registerResMes message.RegisterResMes
	err = json.Unmarshal([]byte(mes.Data), &registerResMes) //别忘了取地址
	if registerResMes.Code == 200 {
		fmt.Println("注册成功 重新登录")
		os.Exit(1)
	} else {
		fmt.Println(registerResMes.Error)
		os.Exit(1)

	}
	return
}
func (this *UserProcess) Login(userId int, userPwd string) (err error) {
	//fmt.Printf("userId=%d,userPwd=%s\n", userId,userPwd)
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("conn dial failed err:", err)
		return
	}
	defer conn.Close()
	//准备通过conn发送消息
	var mes message.Message
	mes.Type = message.LoginMesType
	//创建一个LogiMes
	var logiMes message.LogiMes
	logiMes.UserId = userId
	logiMes.UserPwd = userPwd
	//将logiMes序列化
	data, err := json.Marshal(logiMes)
	if err != nil {
		fmt.Println("json.Marshal is failed err:", err)
		return
	}
	//把data赋给mer.data字段
	mes.Data = string(data)
	//将mes进行序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal is failed err:", err)
		return

	}
	//data 就是我们要发的消息
	//发送信息的长度 但是conn只发送消息的切片 所以需要转换
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)
	n, err := conn.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Println("send is failed err:", err)
		return
	}
	//fmt.Println("客户端发送数据长度成")
	fmt.Println(string(data))
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("send is failed err:", err)
	}
	//这里换需要服务器返回的消息
	//time.Sleep(time.Secod * 20)
	tf := &utils.Transfer{
		Conn: conn,
	}
	mes, err = tf.ReadPkg()

	if err != nil {
		fmt.Println("readpkg failed err", err)
		return
	}
	var loginResMes message.LogiResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes) //别忘了取地址
	if loginResMes.Code == 200 {
		fmt.Println("登陆成功")
		//初始化CurUser
		curUser.Conn=conn
		curUser.UserId=userId
		curUser.UserStatus=message.UserOnline
		//打印在线列表
		fmt.Println("当前在线列表")
		for _, value := range loginResMes.UserId {
			fmt.Printf("用户id=%d\n", value)
			//完成初始化
			user:=&message.User{
				UserId: value,
				UserStatus: message.UserOnline,
			}
			onlineUsers[value]=user
		}
		//起一个协程与客户端保持联系  接收并显示在客户的终端
		go ServerProcessMes(conn)
		//显示登成功的菜单
		for {
			ShowMenu()

		}
	} else {
		fmt.Println(loginResMes.Error)
	}
	return
}
