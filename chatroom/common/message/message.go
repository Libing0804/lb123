package message

const (
	LoginMesType            = "LogiMes"
	LogiResMesType          = "LogiResMes"
	RegisterMesType         = "RegisterMes"
	RegisterResMesType      = "RgisterResMes"
	NotifyUserStatusMesType = "NotifyUserStatusMes"
	SmsMesType              = "SmsMes"
	SmsoneMesType			= "SmsoneMes"
)

//定义几 个户状态的常量
const (
	UserOnline = iota
	UserOffline
	serbusyline
)

type Message struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

//先定义消息
type LogiMes struct {
	UserId   int    `json:"userid"`   //用户id
	UserPwd  string `json:"userpwd"`  //用户密
	UserName string `json:"username"` //用户名
}
type LogiResMes struct {
	Code   int    `json:"code"`   //返回状态吗 00未注册 200登陆成功
	Error  string `json:"error"`  //返回错误信息
	UserId []int  `json:"userid"` //保存用户id的切片
}
type RegisterMes struct {
	User User `json:"user"`
}
type RegisterResMes struct {
	Code  int    `json:"code"` //400表示占用  20代表成功
	Error string `json:"error"`
}

//为了配合服务端推送用户状态的消息
type NotifyUserStatusMes struct {
	UserId int `json:"userId"` //用户id
	Status int `json:"status"` //用户状态
}

//增加smsMes  发送
type SmsMes struct {
	Content string `json:"content"`
	User           //匿名结构体
}
//增加smsMes  发送
type SmsoneMes struct {
	Content string `json:"content"`
	User           //匿名结构体
	UserId  int		`json:"userId"`
}