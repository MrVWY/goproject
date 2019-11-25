package model

import (
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"sever-client/profect/common/message"
)

//在服务器启动后，就初始化一个userCurd实例，把它做成全局的变量
var (
	MyUserCurd *UserCurd
)

type UserCurd struct {
	pool *redis.Pool
}

//使用工厂模式。创建一个UserCurd实例
func NewUserCurd(pool *redis.Pool) (userCurd *UserCurd) {
	userCurd =&UserCurd{
		pool:pool,
	}
	return
}

//Complete login verification
func (this *UserCurd) getUser(conn redis.Conn,id int) (user *message.User , err error)  {
	//Query whether the user exists in redis by specifying ID
	result, err := redis.String(conn.Do("HGet", "users", fmt.Sprintf("%d", id)))

	fmt.Printf("result=%v\n", result)
	//如果对应的用户id,返回  ErrUserNotExist
	if err != nil {

		if err == redis.ErrNil {
			err = ERROR_USER_EXISTS
		}
		return
	}
	//声明一个 *User
	user = &message.User{}
	//将user 进行反序列化，并返回反序列化后的user实例
	err = json.Unmarshal([]byte(result), user)
	if err != nil {
		fmt.Printf("xxx~=%v\n", err)
		return
	}
	fmt.Printf("从数据库得到的user=%v", user)
	return
}

func (this *UserCurd) Login(userID int , userPwd string) (user *message.User , err error){
	conn := this.pool.Get()
	defer conn.Close()
	user , err = this.getUser(conn , userID)  //err = ERROR_USER_NOTEXISTS
	if err != nil{
		return
	}
	//校验密码
	fmt.Println(user.Pwd)
	if user.Pwd != userPwd {
		err = ERROR_USER_PWD
		return
	}
	return
}

//Complete Register verification
func (this *UserCurd) Register(user *message.User) (err error){
	//UserCurd 取出连接池中的一根链接
	//conn := this.pool.Get()
	//defer conn.Close()
	//_ , err = this.getUser(conn , user.Id)
	//if err != nil{
	//	err = ERROR_USER_EXISTS
	//	return
	//}
	////fmt.Println("Register")
	//data , err := json.Marshal(user)
	//if err != nil{
	//	return
	//}
	////输进数据库
	//_ ,err = conn.Do("HSet","users",fmt.Sprintf("%d",user.Id),string(data))
	//if err != nil{
	//	return
	//}
	//return
	//得到一个链接
	conn := this.pool.Get()
	//defer关闭链接
	defer conn.Close()

	_, err = this.getUser(conn, user.Id)
	if err == nil { //如果获取到一个用户，说明该用户存在了
		err = ERROR_USER_EXISTS //用户存在了...
		return
	}

	//序列化这个user
	data, err := json.Marshal(user)
	if err != nil {
		return
	}
	//存入数据库
	_, err = conn.Do("HSet", "users", fmt.Sprintf("%d", user.Id), string(data))
	if err != nil {
		return
	}
	return
}