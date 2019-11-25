package main

import (
	"fmt"
	"github.com/iGoogle-ink/gopay"
)

const privateKey  = "23 "
var client *gopay.AliPayClient

func init() {
	client = gopay.NewAliPayClient("2016101600703675",privateKey,false)
	client.SetAliPayRootCertSN()
	client.SetAppCertSN()
}

func bodymap(consignmentid1 string, number string, timeout string, cgprice string) gopay.BodyMap {
	bm := make(gopay.BodyMap)
	bm.Set("subject",consignmentid1)
	bm.Set("out_trade_no",number)
	bm.Set("quit_url","127.0.0.1:8888")
	bm.Set("timeout_express",timeout)
	bm.Set("total_amount",cgprice)
	bm.Set("product_code","QUICK_WAP_WAY")//销售产品码，商家和支付宝签约的产品码
	return bm
}

func pay(consignmentid1 string, number string, timeout string, cgprice string) (payUrl string) {
	fmt.Println("GoPay Version: ", gopay.Version)
	bm := bodymap(consignmentid1 , number , timeout , cgprice )
	payUrl, err := client.AliPayTradePagePay(bm)
	if err != nil{
		panic(err)
	}
	return payUrl
}
