package map_切片
// map基本使用 增删改查
import "fmt"

func main() {
	fmt.Println("----------1---------------")
	//map的声明和注意事项
	var a map[string]string
	//在使用map前，需要先make，作用是给map分配数据空间
	a = make(map[string]string)
	a["no1"] = "AAA"
	a["no2"] = "BBB"
	a["no3"] = "CCC"
	//a["no1"] = "AAA"
	a["no4"] = "AAA"
	fmt.Println(a)

	fmt.Println("----------2---------------")

	b := make(map[string]string)
	b["no1"] = "AAAA"
	b["no2"] = "BBBB"
	b["no3"] = "CCCC"
	fmt.Println(b)

	fmt.Println("----------3---------------")

	var c map[string]string = map[string]string{
		"no1" : "AAAAA",
		"no2" : "BBBBB",
	}
	fmt.Println(c)

	fmt.Println("----------4---------------")

	var d map[string]string = map[string]string{
		"no1" : "AAAAAA",
		"no2" : "BBBBBB",
		"no3" : "CCCCCC",
	}
	fmt.Println(d)

	fmt.Println("-----3个学生信息包括学生名字和年龄---------------------")

	studentMap := make(map[string]map[string]string)
	studentMap["1"] = map[string]string{
		"name" : "A",
		"sex" : "18",
	}

	studentMap["2"] = map[string]string{
		"name" : "B",
		"sex" : "18",
	}
	fmt.Println(studentMap)
	//遍历
	for k1,v1 := range studentMap{
		fmt.Println("k1=",k1)
		for k2,v2 := range v1{
			fmt.Printf("\t k2=%v v2=%v",k2,v2)
		}
		fmt.Println()
	}
}