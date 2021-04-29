package process2

import (
	"encoding/json"
	"fmt"
	"net"

	"github.com/tcp/chatroom/common/message"
	"github.com/tcp/chatroom/server/utils"
)

type SmsProcess struct {
}

//转发消息
func (this *SmsProcess) SendGroupMes(mes *message.Message) {
	//遍历服务端的 onlineUser map
	//取出mes中的内容
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("json.unmar failed err:", err)
		return
	}
	data, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("json.mar failed err:", err)
		return
	}

	for id, up := range userMgr.onlineUsers {
		//过滤掉自己
		if id == smsMes.UserId {
			continue
		}
		this.SendMEsToEachOnlineUser(data, up.Conn)
	}

}

//转发私聊消息
func (this *SmsProcess) SendoneMes(mes *message.Message, userId int) {
	//遍历服务端的 onlineUser map
	//取出mes中的内容
	var smsoneMes message.SmsoneMes
	err := json.Unmarshal([]byte(mes.Data), &smsoneMes)
	if err != nil {
		fmt.Println("json.unmar failed err:", err)
		return
	}
	data, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("json.mar failed err:", err)
		return
	}

	for id, up := range userMgr.onlineUsers {
		//过滤掉自己
		if id == userId {
			this.SendMEsToEachOnlineUser(data, up.Conn)

		}
	}
}
func (this *SmsProcess) SendMEsToEachOnlineUser(data []byte, conn net.Conn) {
	tf := &utils.Transfer{
		Conn: conn,
	}
	err := tf.WritePkg(data)
	if err != nil {
		fmt.Println("消息转发失败:", err)
		return
	}

}
