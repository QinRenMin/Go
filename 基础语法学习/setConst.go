package main
import "fmt"
import "unsafe"

//定义常量

const CONSTA int = 4

func main() {
	const(
		s = "serty"
		l = len(s)
		S = unsafe.Sizeof(s)
	)
	fmt.Printf("%d\n",CONSTA)
	println(s,l,S)
	
}
