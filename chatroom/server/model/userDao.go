package model

import (
	"encoding/json"
	"fmt"

	"github.com/garyburd/redigo/redis"
	"github.com/tcp/chatroom/common/message"


)

//我们希望服务器启动后就立马初始化一个userDao实例
//把他做成全局的变量
var (
	MyUserDao *UserDao
)

type UserDao struct {
	pool *redis.Pool
}

//使用工厂模式创建 userDao
func NewUserDao(pool *redis.Pool) (userDao *UserDao) {
	userDao = &UserDao{
		pool: pool,
	}
	return
}

//方法

//根据用户id返回用户实例或者error
func (this *UserDao) GetUserById(conn redis.Conn, id int) (user *User, err error) {
	//通过给定的id去redis中查询
	res, err := redis.String(conn.Do("hget", "users", id))
	if err != nil {
		if err == redis.ErrNil {
			fmt.Println("没找到对应的id")
			err = ERROR_USER_NOTEXIXTS
		}
		return
	}
	user = &User{}
	//这里需要把res反序列化
	err = json.Unmarshal([]byte(res), user)
	if err != nil {
		fmt.Println("json.unmarshal is failed :err", err)
		return
	}
	return

}

//完成用户登录验证
func (this *UserDao) Login(userId int, userPwd string) (user *User, err error) {
	//从userdao中取连接
	conn := this.pool.Get()
	defer conn.Close()
	user, err = this.GetUserById(conn, userId)
	if err != nil {
		return
	}
	//到这里用户已经获取到  现在开始获取密码
	if user.UserPwd != userPwd {
		err = ERROR_USER_PWD
		return
	}
	return
}
func (this *UserDao) Register(user *message.User) (err error) {
	//从userdao中取连接
	conn := this.pool.Get()
	defer conn.Close()
	_, err = this.GetUserById(conn, user.UserId)
	if err == nil {
		err = ERROR_USER_EXIXTS
		return
	}
	data, err := json.Marshal(user)
	if err != nil {
		return
	}
	_, err = conn.Do("hset", "users", user.UserId, string(data))
	if err != nil {
		fmt.Println("保存注册错误")
		return
	}
	return
}
