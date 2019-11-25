package main

import (
	"encoding/json"
	"fmt"
	"sync"
)

var p = &sync.Pool{
	New: func() interface{} {
		return map[string]string{}
	},
}

func main() {
	a := p.Get().(map[string]string)
	fmt.Println(a)
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
	json.Unmarshal(data,&a)
	fmt.Println("a",a["a"])
}