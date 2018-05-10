package main
import "fmt"

//定义指针数组

func main(){
    var a [3] int
    var p [3] *int

    a = [3] int{1,2,3}
    //给每个指针数组赋值地址值
    for i:=0; i < 3; i++{
        p[i] = &a[i]
    }

    //通过指针数组输出数据

    for i:=0; i < 3; i++{
        fmt.Printf("%d  ",*p[i])
    }
}
