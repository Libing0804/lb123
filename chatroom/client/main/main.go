package main

import (
	"fmt"
	"os"

	"github.com/tcp/chatroom/client/process"
)

var userId int
var userPwd string
var userName string

func main() {
	//接受用户的选择
	var key int
	//判断是否继续显示菜单
	//var loop = true
	for {
		fmt.Println("************欢迎登录多人聊天系统***********")
		fmt.Println("****************1.登录系统*************")
		fmt.Println("****************2.注册用户*************")
		fmt.Println("****************3.退出系统*************")
		fmt.Println("***************请选择1.2.3*************")

		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			fmt.Println("denglu")
			fmt.Println("输入用户id")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("输入用户密码")
			fmt.Scanf("%s\n", &userPwd)
			up := &process.UserProcess{}
			up.Login(userId, userPwd)
			//loop = false
		case 2:
			fmt.Println("注册用户")
			fmt.Println("请输入id")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入密码")
			fmt.Scanf("%s\n", &userPwd)
			fmt.Println("请输入昵称")
			fmt.Scanf("%s\n", &userName)
			up := &process.UserProcess{}

			up.Register(userId, userPwd, userName)

			//loop = false

		case 3:
			fmt.Println("退出系统")
			os.Exit(1)
		default:
			fmt.Println("重来")
			//loop = false

		}
	}
	// if key == 1 {
	// 	fmt.Println("输入用户id")
	// 	fmt.Scanf("%d\n", &userId)
	// 	fmt.Println("输入用户密码")
	// 	fmt.Scanf("%s\n", &userPwd)
	// 	login(userId, userPwd)

	// } else if key == 2 {
	// 	fmt.Println("等一等")

	// }
}
