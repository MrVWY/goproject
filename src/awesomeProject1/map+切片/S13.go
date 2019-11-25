package map_切片
//map 切片使用
import "fmt"

type student struct {
	Name string
	Age int
}

func main(){
	//声明map切片
	var monsters []map[string]string
	//make
	monsters = make([]map[string]string,2)
	if monsters[0] == nil {
		monsters[0] = map[string]string{
			"name" : "A",
			"Age" : "18",
		}
	}
	if monsters[1] == nil {
		monsters[0] = map[string]string{
			"name" : "B",
			"Age" : "28",
		}
	}
	var newsmonster map[string]string = map[string]string{
		"name" : "C",
		"Age" : "23",
	}
	monsters = append(monsters,newsmonster)
	fmt.Println(monsters)

	students := make(map[string]student,10)
	student1 := student{
		Name: "A",
		Age:  12,
	}
	student2 := student{
		Name: "B",
		Age:  13,
	}
	students = map[string]student{
		"no1":student1,
		"no2":student2,
	}
	fmt.Println(students)
}