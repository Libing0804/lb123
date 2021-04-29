package process2

import (
	"encoding/json"
	"fmt"
	"net"

	"github.com/tcp/chatroom/common/message"
	"github.com/tcp/chatroom/server/model"
	"github.com/tcp/chatroom/server/utils"
)

type UserProcess struct {
	Conn   net.Conn
	UserId int
}

//编写通知所有在线用户的方法

func (this *UserProcess) NotifyOthersOnlineUser(userId int) {
	//便利onlineUsers
	for id, up := range userMgr.onlineUsers {
		if id == userId {
			continue
		}
		//开始通知
		up.NotifyMeOnlineUser(userId)
	}
}
func (this *UserProcess) NotifyMeOnlineUser(userId int) {
	//组装NotifyUserStatusMes消息
	var mes message.Message
	mes.Type = message.NotifyUserStatusMesType
	var notifyUserStatusMes message.NotifyUserStatusMes
	notifyUserStatusMes.UserId = userId
	notifyUserStatusMes.Status = message.UserOnline
	data, err := json.Marshal(notifyUserStatusMes)
	if err != nil {
		fmt.Println("json.mar err:", err)
		return
	}
	mes.Data = string(data)
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.mar err:", err)
		return
	}
	//发送
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("write failed err:", err)
		return
	}
}
func (this *UserProcess) ServrProcessRegister(mes *message.Message) (err error) {
	var registerMes message.RegisterMes
	err = json.Unmarshal([]byte(mes.Data), &registerMes) //地址传入
	if err != nil {
		fmt.Println("json.unmarshal is failed err:", err)
		return
	}
	var resmes message.Message
	resmes.Type = message.RegisterMesType
	var registerResMes message.RegisterResMes
	err = model.MyUserDao.Register(&registerMes.User)
	if err != nil {
		if err == model.ERROR_USER_EXIXTS {
			registerResMes.Code = 505
			registerResMes.Error = model.ERROR_USER_EXIXTS.Error()
		} else {
			registerResMes.Code = 505
			registerResMes.Error = "buzhidao错误"
		}
	} else {
		registerResMes.Code = 200

	}
	data, err := json.Marshal(registerResMes)
	if err != nil {
		fmt.Println("loginResMes json marshal failed err:", err)
		return
	}
	resmes.Data = string(data)
	//对resms序列化
	data, err = json.Marshal(resmes)
	if err != nil {
		fmt.Println("resms json.marshal failed err:", err)
		return
	}
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	return
}

//编写一个serverProcessLogin函数处登录
func (this *UserProcess) ServrProcessLogin(mes *message.Message) (err error) {
	//1. 先从mes中取出mes.data在反序列化
	var loginMes message.LogiMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes) //地址传入
	if err != nil {
		fmt.Println("json.unmarshal is failed err:", err)
		return
	}
	var resmes message.Message
	resmes.Type = message.LogiResMesType
	//在生命一个LoginresMes
	var loginResMes message.LogiResMes
	// //果用户的账号是100 密码是123456就对的
	// if loginMes.UserId = 100 && loginMes.UserPwd == "123456 {
	// 	/合法
	// 	logiResMes.Code = 00

	// } els {
	// 	//合法
	// 	loginResMes.Code = 500 //表示该用户存在
	// 	loginResMes.Error = "表示该用户不在"

	//  }
	//数据库验证
	user, err := model.MyUserDao.Login(loginMes.UserId, loginMes.UserPwd)
	if err != nil {
		if err == model.ERROR_USER_NOTEXIXTS {
			loginResMes.Code = 500
			loginResMes.Error = err.Error()
		} else if err == model.ERROR_USER_PWD {
			loginResMes.Code = 403
			loginResMes.Error = err.Error()

		} else {
			loginResMes.Code = 505
			loginResMes.Error = "服务器内部错误"

		}

	} else {
		loginResMes.Code = 200
		//登陆成功放进map列表之中
		this.UserId = loginMes.UserId
		userMgr.AddOnlineUser(this)
		//通知其他在线用户我上线了
		this.NotifyOthersOnlineUser(loginMes.UserId)
		//便利
		for id,_:= range userMgr.onlineUsers {
			loginResMes.UserId = append(loginResMes.UserId, id)
		}
		fmt.Println(user, "登录成功")

	}
	//将logiResMes序化
	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("loginResMes json marshal failed err:", err)
		return
	}
	resmes.Data = string(data)
	//对resms序列化
	data, err = json.Marshal(resmes)
	if err != nil {
		fmt.Println("resms json.marshal failed err:", err)
		return
	}
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)

	return
}
