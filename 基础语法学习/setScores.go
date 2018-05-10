package main
import "fmt"

func main(){
    var scoreRadeMap map[string] int
    scoreRadeMap = make(map[string] int)

    scoreRadeMap["best"] = 100
    scoreRadeMap["good"] = 90
    scoreRadeMap["bad"] = 59

    //输出MAP信息
    for score := range scoreRadeMap {
        fmt.Println(score,scoreRadeMap[score])
    }
    //判断是否有信息存在map的相关信息
    score,ok := scoreRadeMap["bad"]
    fmt.Println(score)
    fmt.Println(ok)
    if(ok){
        fmt.Println("存在",score)
    } else{
        fmt.Println("不存在")
    }
}
