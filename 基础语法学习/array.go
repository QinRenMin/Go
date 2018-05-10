package main
import "fmt"
//定义一位数组和多维数组，并对其进行输出，注意语法格式
func main(){
    var a[5] int
    a = [5] int{1,2,3,4,5}
    for i:=1; i < 5; i++ {
        fmt.Printf("%d\n",a[i])
    }
   b:= [2][3] int{
       {1,2,3},
       {4,5,6}}
    fmt.Printf("%d\n",b[1][2])
}
