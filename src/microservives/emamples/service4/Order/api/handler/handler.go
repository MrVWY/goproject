package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/micro/go-micro/client"
	"log"
	"net/http"
	O "microservives/emamples/service4/Order/osrv/proto"
	A "microservives/emamples/service4/athu/proto"
	"time"
)

var (
	serviceorder O.OService
	serviceathu A.AthuService
)

func Init(){
	serviceorder = O.NewOService("go.micro.web.srv.o",client.DefaultClient)
	serviceathu = A.NewAthuService("go.micro.web.srv.athu",client.DefaultClient)
}

func Creates(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("123")
	if r.Method != "POST" && len(r.Form.Get("username")) == 0 && len(r.Form.Get("password")) == 0  {
		log.Fatal("The request is fail")
		http.Error(w,"The request is illegal",404)
		return
	}
	r.ParseForm()
	rep := map[string]interface{}{
		"ref": time.Now().UnixNano(),
	}
	_, err := serviceathu.VerificationToken(context.TODO(), &A.Request{
		Token: r.Header.Get("token"),
	})
	if err != nil {
		log.Println("[Creates] Token is express , err：%s", err)
		rep["message"] = "Token is express,Please o-login"
		if err := json.NewEncoder(w).Encode(rep); err != nil{
			http.Error(w,err.Error(),500)
			return
		}
		return
	}

	res, err := serviceorder.Create(context.TODO(),&O.Request{
		Consignmentid:        r.Form.Get("Consignmentid"),
		Name:                 r.Form.Get("Name"),
		Price:                r.Form.Get("Price"),
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	fmt.Println(res)
	rep["message"] = res
	if err := json.NewEncoder(w).Encode(rep); err != nil{
		http.Error(w,err.Error(),500)
		return
	}
}

func GetAlls(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("123")
	if r.Method != "POST" {
		log.Fatal("The request is fail")
		http.Error(w,"The request is illegal",404)
		return
	}
	r.ParseForm()
	rep := map[string]interface{}{
		"ref": time.Now().UnixNano(),
	}
	_, err := serviceathu.VerificationToken(context.TODO(), &A.Request{
		Token: r.Header.Get("token"),
	})
	if err != nil {
		log.Println("[Creates] Token is express , err：%s", err)
		rep["message"] = "Token is express,Please o-login"
		if err := json.NewEncoder(w).Encode(rep); err != nil{
			http.Error(w,err.Error(),500)
			return
		}
		return
	}
	res, err := serviceorder.GetAll(context.TODO(),&O.Request{
		Name:                 r.Form.Get("Name"),
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	fmt.Println(res)
	rep["message"] = res
	if err := json.NewEncoder(w).Encode(rep); err != nil{
		http.Error(w,err.Error(),500)
		return
	}
}