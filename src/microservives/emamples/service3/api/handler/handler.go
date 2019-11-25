package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/micro/go-micro/client"
	"log"
	"net/http"
	"micro/emamples/service3/Usrv/proto"
	O "micro/emamples/service3/athu/proto"
	"time"
)

var (
	serviceuser proto.UserService
	serviceathu O.AthuService
)

func Init(){
	serviceuser = proto.NewUserService("go.micro.web.svr.user",client.DefaultClient)
	serviceathu = O.NewAthuService("go.micro.web.srv.athu",client.DefaultClient)
}

func Login(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("123")
	if r.Method != "POST" && len(r.Form.Get("username")) == 0 && len(r.Form.Get("password")) == 0 {
		log.Fatal("The request is illegal")
		http.Error(w,"The request is illegal",404)
		return
	}
	r.ParseForm()
	fmt.Println(r.Form.Get("username"),r.Form.Get("password"))
	res, err := serviceuser.Login(context.TODO(), &proto.Request{
		Name: r.Form.Get("username"),
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	rep := map[string]interface{}{
		"ref": time.Now().UnixNano(),
	}
	fmt.Println(res)
	if res.User.Pwd == r.Form.Get("password") {
		rep["success"] = true

		// 干掉密码返回
		res.User.Pwd = ""
		rep["data"] = res.User
		log.Println("[Login] 密码校验完成，生成token...")

		// 生成token
		rsp2, err := serviceathu.MakeToken(context.TODO(), &O.Request{
			Username: res.User.Name,
		})
		if err != nil {
			log.Println("[Login] 创建token失败，err：%s", err)
			http.Error(w, err.Error(), 500)
			return
		}

		log.Println("[Login] token %s", rsp2.Token)
		rep["token"] = rsp2.Token

		// 同时将token写到cookies中
		w.Header().Add("set-cookie", "application/json; charset=utf-8")
		// 过期30分钟
		expire := time.Now().Add(30 * time.Minute)
		cookie := http.Cookie{Name: "remember-me-token", Value: rsp2.Token, Path: "/", Expires: expire, MaxAge: 90000}
		http.SetCookie(w, &cookie)

	} else {
		rep["success"] = false
		rep["error"] = "fail"
	}

	w.Header().Add("Content-Type", "application/json; charset=utf-8")

	// 返回JSON结构
	if err := json.NewEncoder(w).Encode(rep); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func Register(response http.ResponseWriter, r *http.Request)  {
	if r.Method != "POST" && len(r.Form.Get("username")) == 0 && len(r.Form.Get("password")) == 0 {
		log.Fatal("The request is illegal")
		http.Error(response,"The request is illegal",404)
		return
	}
	r.ParseForm()
	res, err := serviceuser.Register(context.TODO(),&proto.Request{
		Name:                 r.Form.Get("userName"),
		Password:             r.Form.Get("password"),
	})
	if err != nil {
		http.Error(response, err.Error(), 500)
		return
	}
	res.User.Pwd = ""
	rep := map[string]interface{}{
		"ref": time.Now().UnixNano(),
		"status" : "Successfully",
		"data": res.User,
	}

	response.Header().Add("Content-Type","application/json; charset=utf-8")

	if err := json.NewEncoder(response).Encode(rep) ; err != nil{
		http.Error(response,err.Error(),5000)
		return
	}

}