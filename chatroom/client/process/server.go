package process

import (
	"encoding/json"
	"fmt"
	"net"
	"os"

	"github.com/tcp/chatroom/client/utils"
	"github.com/tcp/chatroom/common/message"
)

//显示登陆成功后的界面
func ShowMenu() {
	fmt.Println("用户登录界面：恭喜XXX登录成功")
	fmt.Println("1.显示用户在线列表")
	fmt.Println("2.发送消息")
	fmt.Println("3.私聊")
	fmt.Println("4.退出系统")
	fmt.Println("请选择1-4")
	var key int
	var content string
	var userId int
	//我们总会使用到SMSProcess实例定义一个实例
	smsProcess:=&SmsProcess{}
	fmt.Scanf("%d\n", &key)
	switch key {
	case 1:
		//fmt.Println("1.显示用户在线列表")
		outputOnlineUser()
	case 2:
		//fmt.Println("2.发送消息")
		fmt.Println("请输入你想对大家说的话：")
		fmt.Scanf("%s\n",&content)
		smsProcess.SendGroupMes(content)

	case 3:
		//fmt.Println("3.私聊")
		fmt.Println("输入对方的id")
		fmt.Scanf("%d\n",&userId)
		fmt.Println("请输入你想对ta说的话：")
		fmt.Scanf("%s\n",&content)
		smsProcess.SendoneMes(content,userId)
	case 4:
		fmt.Println("4.退出系统")
		os.Exit(1)
	default:
		fmt.Println("输入错误")
	}
	//和服务器端保持通信
}
func ServerProcessMes(conn net.Conn) {
	tf := &utils.Transfer{
		Conn: conn,
	}
	for {
		mes, err := tf.ReadPkg()
		if err != nil {

			fmt.Println("read failed err:", err)
			return
		}
		//如果读取到消息
		switch mes.Type {
		case message.NotifyUserStatusMesType:
			//取出NotifyUserStatusMes信息
			var notifyUserStatusMes message.NotifyUserStatusMes
			json.Unmarshal([]byte(mes.Data), &notifyUserStatusMes)
			//保存的客户端的map中
			updateUserStatus(&notifyUserStatusMes)
		case message.SmsMesType://有人群发消息
			outputGroupMes(&mes)
		case message.SmsoneMesType:
			outputoneMes(&mes)
		default:
			fmt.Println("服务端的返回未知消息类型")

		}
		fmt.Println(mes)
	}
}
