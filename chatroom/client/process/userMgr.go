package process

import (
	"fmt"
	"github.com/tcp/chatroom/common/message"
	"github.com/tcp/chatroom/client/model"

)

var onlineUsers map[int]*message.User = make(map[int]*message.User, 10)

//全局的变量用来维护conn
var curUser model.CurUser //我们在用户登陆成功后进行初始化
//显示当前在线用户
func outputOnlineUser() {
	//遍历
	fmt.Println("当前在线用户列表：")
	for id, user := range onlineUsers {
		fmt.Printf("用户id：%d,状态：%d\n", id, user.UserStatus)
	}
}

//编写一个方法处理返回的信息
func updateUserStatus(notifyUserStatusMes *message.NotifyUserStatusMes) {
	user, ok := onlineUsers[notifyUserStatusMes.UserId]
	if !ok { //原来没有
		user = &message.User{
			UserId: notifyUserStatusMes.UserId,
		}
	}
	user.UserStatus = notifyUserStatusMes.Status
	onlineUsers[notifyUserStatusMes.UserId] = user
	outputOnlineUser()
}
