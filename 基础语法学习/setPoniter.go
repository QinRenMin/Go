package main
import "fmt"

//定义指针变量，理解存储的内容

func main(){
    //一级指针
    var a int
    var p *int

    a = 20
    p = &a //p指向a

    fmt.Printf("a的地址是:%x\n",&a)
    fmt.Printf("a的地址是:%x\n",p)
    fmt.Printf("a的值是:%d\n",*p)

    //指向指针的指针
    var pp **int
    pp = &p //指向p的地址
    fmt.Printf("p的地址是:%x\n",pp)
    fmt.Printf("a的值是：%d\n",**pp)

    //定义空指针
    var ptr *int
    fmt.Printf("%T\n",ptr)
    if(ptr == nil){
        fmt.Printf("ptr是空指针\n")

    }
}
