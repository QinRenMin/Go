package main
import (
    "fmt"
    "math"
)
//函数的使用方法

func main(){
    /*声明函数变量*/
    getSquear :=  func(x float64  ) float64 {
        return math.Sqrt(x)
    }

    fmt.Println(getSquear(3))
    //闭包调用
    add_func := add(1,2)
    fmt.Println(add_func())
    fmt.Println(add_func())
}

func add(x,y int) func()(int,int) {
    i:=0
    return func()(int,int) {
        i++
        return i,x+y
    }
}

