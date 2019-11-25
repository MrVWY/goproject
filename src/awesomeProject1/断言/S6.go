package 断言
//断言实例1
import "fmt"

//声明定义一个接口
type Usb interface {
	//声明2个没有实现的方法
	Start()
	Stop()
}

type Phone struct {
	name string
}
//实现Usb接口的方法
func (p Phone) Start(){
	fmt.Println("手机开始工作..........")
}
func (p Phone) Stop()  {
	fmt.Println("手机停止工作..........")
}
func (p Phone) Call()  {
	fmt.Println("手机正在打电话..........")
}

type Camera struct {
	name string
}
//实现Usb接口方法
func (c Camera) Start(){
	fmt.Println("相机开始工作..........")
}
func (c Camera) Stop()  {
	fmt.Println("相机停止工作..........")
}

type Computer struct {

}
func (computer Computer) Working(usb Usb)  {
	usb.Start()
	if phone,ok := usb.(Phone); ok{
		phone.Call()
	}
	usb.Stop()
}

func main()  {
	var usbArr [3]Usb
	usbArr[0] = Phone{"ipone"}
	usbArr[1] = Phone{"vivo"}
	usbArr[2] = Camera{"ipad"}

	var computer Computer
	for _, v := range usbArr{
		computer.Working(v)
		fmt.Println()
	}
}