package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	h := map[string]string{}
	m := map[string]string{
		"level":   "debug",
		"message": "File not found",
		"a" :   "true",
	}
	fmt.Println(m["level"])
	data, err := json.Marshal(m)
	if  err == nil {
		fmt.Printf("%s\n", data)
	}
	json.Unmarshal(data,&h)
	fmt.Println(h)
	if h["a"] == "true"{
		fmt.Println("1231321233")
	}
}