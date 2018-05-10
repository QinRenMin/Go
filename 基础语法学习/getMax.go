package main
import "fmt"

//设置获得最大函数值
func max(num1,num2 int ) int {
    var result int

    if (num1 < num2){
        result = num2
    } else {
        result = num1
    }
    return result
}

func main() {
    var a,b int
    a = 10
    b = 20
    fmt.Printf("虽大值:%d",max(a,b))
}
