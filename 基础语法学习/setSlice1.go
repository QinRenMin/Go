package main
import "fmt"

//使用copy和append函数对切片进行操作
func main() {
    var nums []int
    printSlice(nums)
    //添加一个元素
    nums = append(nums,0)
    printSlice(nums)
    //同时添加多个元素
    nums = append(nums,5,6,7)
    printSlice(nums)
    //创建切片，并设置其容量值是之前的两倍
    nums1 := make([]int,len(nums),2*(cap(nums)))
    //将nums复制到nums1
    copy(nums1,nums)
    printSlice(nums1)
}
func printSlice(x []int){
    fmt.Printf("len = %d map =  %d  slice = %v\n",len(x),cap(x),x)
}
