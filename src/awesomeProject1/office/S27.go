package main
//做对这三道题，表明你基本已经如火纯青般地掌握了for-range的用法
import "fmt"

type T struct {
	n int
}

func main() {
	//数组
	ts := [2]T{}
	for i, t := range ts {
		switch i {
		case 0:
			t.n = 3
			ts[1].n = 9
		case 1:
			fmt.Print(t.n, " ")
		}
	}
	fmt.Print(ts)
    //数组指针
	a := [2]T{}
	for i, t := range &a {
		switch i {
		case 0:
			t.n = 3
			a[1].n = 9
		case 1:
			fmt.Print(t.n, " ")
		}
	}
	fmt.Print(a)
	//切片
	s := [2]T{}
	for i := range s[:] {
		t := &s[i]
		switch i {
		case 0:
			t.n = 3
			s[1].n = 9
		case 1:
			fmt.Print(t.n, " ")
		}
	}
	fmt.Print(s)

	 x := 1
	fmt.Println(x)
	 {
			 fmt.Println(x)
			i,x := 2,2
			fmt.Println(i,x)
	 }
	 fmt.Println(x)  // print ?


}
