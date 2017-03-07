package main
import "fmt"

func main() {
    cnt := 0
    sum := 0
    c := 0
    fmt.Scan(&cnt)
    for i:=0;i<cnt;i++ {
        fmt.Scan(&c)
        sum = sum + c
    }
    fmt.Print(sum)
}
