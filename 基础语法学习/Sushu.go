package main
import "fmt"
//输出1-100之间的素数,使用GOTO语句
func main(){
	var C,c int //变量声明
	C = 1
	A: for C < 100{
		C++
		for c = 2; c < C; c++{
			if C%c == 0{
				goto A //发现因子不是素数
				}
			}
			fmt.Println(C,"是素数")
	}
}
