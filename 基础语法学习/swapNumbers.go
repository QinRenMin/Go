package main
import "fmt"

//使用指针作为函数的形参，完成两个值的交换

func main(){
    var a,b int
    a = 10
    b = 20

    fmt.Printf("交换前a,b的值：%d  %d \n",a,b)

    swap(&a,&b)
    fmt.Printf("交换后a,b的值：%d  %d \n",a,b)
}

func swap(x ,y *int){
    var temp  int

    temp = *x
    *x = *y
    *y = temp

}
