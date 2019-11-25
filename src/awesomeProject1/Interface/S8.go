package Interface
//interface
import "fmt"

//定义二个接口
type AInterface interface {
	Say()
}
type BInterface interface {
	Hello()
}

//定义结构体
type Stu struct {
	Name string
}
func (stu Stu) Say()  {
	fmt.Println("stu Say")
}

type integer int
func (i integer) Say(){
	fmt.Println("integer Say i = ",i)
}

type Monster struct {

}
func (m Monster) Hello(){
	fmt.Println("Monster Hello")
}

func main()  {
	var stu Stu
	var a AInterface = stu
	a.Say()

	fmt.Println("----------------------------------------------")

	var math integer  = 111
	var b AInterface = math
	b.Say()

	fmt.Println("----------------------------------------------")

	var monster Monster
	var c BInterface = monster
	c.Hello()

	fmt.Println("----------------------------------------------")

	var t interface{} = stu
	var num float64 = 1.5
	t = num
	fmt.Println(t)
}