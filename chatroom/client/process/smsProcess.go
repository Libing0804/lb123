package process

import (
	"encoding/json"
	"fmt"

	"github.com/tcp/chatroom/common/message"
	"github.com/tcp/chatroom/server/utils"
)


type SmsProcess struct {
}

func (this *SmsProcess) SendGroupMes(content string) (err error) {
	//创建一个message
	var mes message.Message
	mes.Type=message.SmsMesType
	//创建一个SMSMes实例
	var smsMes message.SmsMes
	smsMes.Content=content
	smsMes.UserId=curUser.UserId
	smsMes.UserStatus=curUser.UserStatus
	//将SMSMes序列化
	data,err:=json.Marshal(smsMes)
	if err!=nil{
		fmt.Println("sendGroupMes json.Mar failed err:",err)
		return
	}
	//对mes序列化
	mes.Data=string(data)
	data,err=json.Marshal(mes)
	if err!=nil{
		fmt.Println("sendGroupMes json.Mar failed err:",err)
		return
	}
//将mes发送
	tf:=&utils.Transfer{
		Conn: curUser.Conn,
	}
	err=tf.WritePkg(data)
	if err!=nil{
		fmt.Println("sendGroup send failed err:",err)
		return
	}
		return
}
func (this *SmsProcess) SendoneMes(content string,userId int) (err error) {
	//创建一个message
	var mes message.Message
	mes.Type=message.SmsoneMesType
	//创建一个SMSMes实例
	var smsoneMes message.SmsoneMes
	smsoneMes.Content=content
	smsoneMes.UserId=userId
	smsoneMes.UserStatus=curUser.UserStatus
	//将SMSMes序列化
	data,err:=json.Marshal(smsoneMes)
	if err!=nil{
		fmt.Println("sendoneMes json.Mar failed err:",err)
		return
	}
	//对mes序列化
	mes.Data=string(data)
	data,err=json.Marshal(mes)
	if err!=nil{
		fmt.Println("sendoneMes json.Mar failed err:",err)
		return
	}
	//将mes发送
	tf:=&utils.Transfer{
		Conn: curUser.Conn,
	}
	err=tf.WritePkg(data)
	if err!=nil{
		fmt.Println("sendoneMes send failed err:",err)
		return
	}
		return
}