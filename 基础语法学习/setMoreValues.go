package main
import "fmt"
//声明多个变量
func main() {
	//当多变量类型相同时
	var a,b,c int
	a,b,c = 1,2,3
	
	var q,w= true,2
	
	x,y:= 1.3,5
	
	//声明全局变量
	var (
		z int 
		h bool
	)
	fmt.Println(a,b,c,q,w,x,y,z,h)
}
