package Interface
//interface
import (
	"fmt"
	"math/rand"
	"sort"
)
type hero struct {
	Age int
}
type HeroSlice []hero


func (hs HeroSlice) Len() int  {
	return len(hs)
}
func (hs HeroSlice) Less(i,j int) bool  {
	return hs[i].Age  < hs[j].Age
}
func (hs HeroSlice) Swap(i,j int)  {
	hs[i],hs[j] = hs[j],hs[i]
}

func main()  {
	var intSlice = []int{-1,10,2,3,32,25,25}
	sort.Ints(intSlice)
	fmt.Println(intSlice)

	fmt.Println("-----------------------------------------")

	var heros  HeroSlice
	for i := 0; i<10 ; i++{
		hero := hero{
			Age : rand.Intn(100),
		}
		heros = append(heros,hero)
	}

	//before
	fmt.Println("----------------before-------------------------")
	for _,v := range heros{
		fmt.Println(v)
	}
	//after
	fmt.Println("------------------after-----------------------")
	sort.Sort(heros)
	for _,v := range heros{
		fmt.Println(v)
	}
}