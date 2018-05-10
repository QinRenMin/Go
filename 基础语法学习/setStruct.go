package main
import "fmt"

//定义结构体，存放不同类型的变量
type Person struct {
    name string
    age int
    hobby string
}
func main(){
    var p  Person
    p.name = "jane"
    p.age = 19
    p.hobby = "study"
    fmt.Printf("name: %s\n",p.name)
    fmt.Printf("age: %d\n",p.age)
    fmt.Printf("hobby: %s\n",p.hobby)
    PrintPerson(p)
}
//通过调用函数输出
func PrintPerson(p1 Person) {
    fmt.Printf("通过调用函数输出\n")
    fmt.Printf("name: %s \n",p1.name)
}
