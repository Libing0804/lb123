package model

//定义一个用户的结构体
type User struct {
	//为了序列化反序列化成功我么必须使用户的信息json字符串的key 与字段的tag相同
	UserId   int    `json:"userid"`
	UserName string `json:"username"`
	UserPwd  string `json:"userpwd"`
}
