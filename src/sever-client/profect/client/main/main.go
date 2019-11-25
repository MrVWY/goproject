package main

import (
	"sever-client/profect/client/process"
	"fmt"
	"os"
)

var userid int
var userPwd string
var userName string

func main(){
	//接受用户的选择
	var key int
	//判断是否还继续显示菜单


	for true {
		fmt.Println("------------欢迎登陆----------------------")
		fmt.Println("\t\t 1、登录")
		fmt.Println("\t\t 2、注册")
		fmt.Println("\t\t 3、退出")

		fmt.Scanf("%d\n",&key)
		switch key {
			case 1:
				fmt.Println("login")
				//login
				fmt.Println("Your ID")
				fmt.Scanf("%d\n",&userid)
				fmt.Println("Your Password")
				fmt.Scanf("%s\n",&userPwd)
				up := &process.UserProcess{}
				up.Alogin(userid,userPwd)
			case 2:
				fmt.Println("create new account")
				fmt.Println("Your ID")
				fmt.Scanf("%d\n",&userid)
				fmt.Println("Your Password")
				fmt.Scanf("%s\n",&userPwd)
				fmt.Println("Your name")
				fmt.Scanf("%s\n",&userName)
				up := &process.UserProcess{}
				up.Register(userid,userPwd,userName)
			case 3:
				fmt.Println("exit")
				os.Exit(0)
			default:
				fmt.Println("please input again")
		}
	}
}