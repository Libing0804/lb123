package process2

import "fmt"

var (
	userMgr *UserMgr
)
type UserMgr struct {
	onlineUsers map[int]*UserProcess
}



func init() {
	userMgr = &UserMgr{
		onlineUsers: make(map[int]*UserProcess, 1024),
	}

}

//增加
func (this *UserMgr) AddOnlineUser(up *UserProcess) {
	this.onlineUsers[up.UserId] = up
}

//删除
func (this *UserMgr) DelOnlineUser(userId int) {
	delete(this.onlineUsers, userId)
}

//返回当前用户
func (this *UserMgr) GetallOnlineUser() map[int]*UserProcess {
	return this.onlineUsers
}

//返回id对应的值
func (this *UserMgr) GetOnlinebyUserId(userId int) (err error, up *UserProcess) {
	up, ok := this.onlineUsers[userId]
	if !ok {
		fmt.Println("你要查找的用户当前不在线")
		return

	}
	return
}
