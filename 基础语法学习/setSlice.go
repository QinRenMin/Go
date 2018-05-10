package main
import "fmt"

//定义切片

func main() {
    //定义切片

    slice := []int{1,2,3,4,5,6}
    printSlice(slice)
    //打印完整切片
    fmt.Println("slice = ",slice)
    fmt.Printf("\n")
    //打印从开头到4（不包含4）
    fmt.Println("slice[:4] = ",slice[:4])
    fmt.Printf("\n")

    numbers1 := make([]int,0,2)
    printSlice(numbers1)

    numbers2 := make([]int ,2,3)
    printSlice(numbers2)

    var k []int
    if(k == nil) {
        fmt.Printf("这是个空切片\n:")
    }
}

func printSlice(s []int ) {
    fmt.Printf("len = %d  cap = %d slice = %v\n",len(s),cap(s),s)
}
