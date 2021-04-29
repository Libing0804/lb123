package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net"

	"github.com/tcp/chatroom/common/message"
	"github.com/tcp/chatroom/server/process2"
	"github.com/tcp/chatroom/server/utils"
)

type Processor struct {
	Conn net.Conn
}

//编写一个serverprocessMes函数
//功能：根据客户端发来信息的种类的不同决定调用那个函数来处理
func (this *Processor) serverProcessMes(mes *message.Message) (err error) {
	//fmt.Println(mes)
	switch mes.Type {
	case message.LoginMesType:
		//处理登录
		up := &process2.UserProcess{
			Conn: this.Conn,
		}
		err = up.ServrProcessLogin(mes)
	case message.RegisterMesType:
		up := &process2.UserProcess{
			Conn: this.Conn,
		}
		up.ServrProcessRegister(mes)
	case message.SmsMesType:
		//创建一个SMSProcess实例完成群发消息
		smsProcess:=&process2.SmsProcess{}
		smsProcess.SendGroupMes(mes)
	case message.SmsoneMesType:
		var smsoneMes message.SmsoneMes
		json.Unmarshal([]byte(mes.Data),&smsoneMes)
		smsProcess:=&process2.SmsProcess{}
		smsProcess.SendoneMes(mes,smsoneMes.UserId)
	default:
		fmt.Println("消息类型不存在")
	}
	return
}
func (this *Processor) process2() (err error) {
	for {
		//读取数据封装成一个函数
		tf := &utils.Transfer{
			Conn: this.Conn,
		}
		mes, err := tf.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出  服务端也要退出")
				return err
			} else {
				fmt.Println("readpkg is failed err:", err)
			}
		}
		err = this.serverProcessMes(mes)
		if err != nil {
			fmt.Println(err)
			return err
		}
	}
}
