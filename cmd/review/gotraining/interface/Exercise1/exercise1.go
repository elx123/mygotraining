package main
//https://github.com/ardanlabs/gotraining/tree/master/topics/go/language/interfaces
import "fmt"

type speaker interface{
	speak()
}

type english struct{
}

func (e english) speak(){
	fmt.Println("Hello World")
}

type chinese struct{
}

func (c chinese) speak(){
	fmt.Println("你好世界")
}

func main(){
	var spinterface speaker
	en := english{}
	ch := chinese{}
	spinterface = &en
	spinterface.speak()
	spinterface = ch
	spinterface.speak()
}
