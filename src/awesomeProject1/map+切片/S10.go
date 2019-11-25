package map_切片
//切片 是一个可以动态变化的数组  切片的定义和遍历 内置函数append copy
import "fmt"

func main() {
	var intarr [5]int = [...]int{1,22,55,332,452}
	//声明定义一个切片
	slice := intarr[1:3]
	fmt.Println("intarr = ",intarr)
	fmt.Println("slice 的元素是 = ",slice)
	fmt.Println("slice 的元素个数 = ",len(slice))
	fmt.Println("slice 的容量 = ",cap(slice))
	for i:=0 ; i<len(slice) ;i++{
		fmt.Printf("slice[%v] = %v",i,slice[i])
		fmt.Println()
	}
	fmt.Println()
	fmt.Println("----------------------------")

	var slice1 []float64 = make([]float64,5,10)
	slice1[1] = 10
	slice1[3] = 20
	fmt.Println(slice1)
	fmt.Println("slice1 的元素个数 = ",len(slice1))
	fmt.Println("slice1 的容量 = ",cap(slice1))
	for i,v := range slice1{
		fmt.Printf("i=%v v=%v \n",i,v)
		fmt.Println()
	}
	fmt.Println("----------------------------")

	var slice2 []string = []string{"T","S","A"}
	fmt.Println("slice2 的元素是 = ",slice2)
	fmt.Println("slice2 的元素个数 = ",len(slice2))
	fmt.Println("slice2 的容量 = ",cap(slice2))
}