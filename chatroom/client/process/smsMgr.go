package process

import (
	"encoding/json"
	"fmt"

	"github.com/tcp/chatroom/common/message"
)

func outputGroupMes(mes *message.Message) { //这里来的消息是smsmestype类型
	//显示消息
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("json.Unmarshal failed err:", err)
		return
	}
	info := fmt.Sprintf("用户:[%d]\t对大家说:%s\n", smsMes.UserId, smsMes.Content)
	fmt.Println(info)
}
func outputoneMes(mes *message.Message) { //这里来的消息是smsonemestype类型
	//显示消息
	var smsoneMes message.SmsoneMes
	err := json.Unmarshal([]byte(mes.Data), &smsoneMes)
	if err != nil {
		fmt.Println("json.Unmarshal failed err:", err)
		return
	}
	info := fmt.Sprintf("用户:[%d]\t对ta说:%s\n", smsoneMes.UserId, smsoneMes.Content)
	fmt.Println(info)
}

