package message

//确定消息类型type
const  (
	LoginMesType  = "LoginMes"
	LoginResMesType = "LoginResMes"
	RegisterMesType = "RegisterMes"
	RegisterResMesType = "RegisterResMes"
	NotifyUserStatusMesType = "NotifyUserStatusMes"
	SmMesType = "SmMes"
)

//用户在线状态
const (
	UserOnline = iota
	UserOffline
	UserBusyStatus
)

//要发送的消息主体
type Message struct {
	Type string `json:"type"`
	Date string	 `json:"date"`
}

//定义2个消息 即Date
type LoginMes struct {
	UserID int `json:"user_id"`
	UserPwd string `json:"user_pwd"`
	Username string `json:"username"`
}

//type User1 struct {
//	UserIds[]int               //保存用户ID的切片
//	UserName []string
//}

type LoginResMes struct {
	Code int  `json:"code"`
	UserIds[]int               //保存用户ID的切片
	//UserName []string
	//User1
	Error string `json:"error"`
}

type RegisterMes struct {
	User User  `json:"user"`
}

type RegisterResMes struct {
	Code int  `json:"code"`// 500表示未注册  200 表示登录成功
	Error string `json:"error"`
}

//配合服务器端推送的用户状态消息
type NotifyUserStatusMes struct {
	UserId int `json:"user_id"`
	Satus int `json:"satus"`
}

type SmMes struct {
	Content string `json:"content"`
	User
}