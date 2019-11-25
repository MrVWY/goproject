package main

import (
	//"awesomeProject/factory/model"
	"awesomeProject1/factory/model"
	"fmt"
)

func main()  {
	//创建要给Teacher
	var stu = model.Teacher{
		Name:"Tom",
		Age:18,
	}

	fmt.Println(stu)

	var stu1 = model.Newstudent("JACK",88)
	fmt.Println(*stu1)

	fmt.Println("name",stu.Name,"score=",stu1.GetScore())
}
