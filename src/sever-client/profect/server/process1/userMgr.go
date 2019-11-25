package process1

import (
	"fmt"
)

//
var (
	userMgr *UserMgr
)

type UserMgr struct {
	onlineUsers map[int]*UserloginProcess  //  int-*UserloginProcess()
}

//完成对userMgr初始化工作
func init() {
	userMgr = &UserMgr{
		onlineUsers: make(map[int]*UserloginProcess,1024),
	}
}

//完成对onlineUsers添加
func (this *UserMgr) AddonlineUsers(u *UserloginProcess) {
	this.onlineUsers[u.UserId] = u

}

//完成对onlineUsers删除
func (this *UserMgr) DeleteonlineUsers(userId int)  {
	delete(this.onlineUsers,userId)
}

//返回当前所有在前的用户
func (this *UserMgr) GetAlleonlineUsers() map[int] *UserloginProcess {
	return this.onlineUsers
}

//根据id返回对应的值
func (this *UserMgr) ReturneonlineUsers(userId int) (u *UserloginProcess ,err error) {
	//如何从map取出一个值，带检测方式
	u , ok := this.onlineUsers[userId]
	if !ok  {
		//说明查询的用户不在线
		err =fmt.Errorf("用户%d 不存在",userId)
		return
	}
	return
}