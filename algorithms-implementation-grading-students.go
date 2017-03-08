package main
import "fmt"

func main() {
    ns := 0
    curr := 0
    fmt.Scan(&ns)
    for i:=1; i<=ns; i++ {
        fmt.Scan(&curr)
        pad := curr % 10
        puluhan := curr - pad
        if curr < 38 {
            fmt.Println(curr)
        } else if pad > 5 && puluhan + 10 - curr < 3 {
            fmt.Println(puluhan+10)
        } else if pad < 5 && puluhan + 5 - curr < 3 {
            fmt.Println(puluhan+5)
        } else {
            fmt.Println(curr)
        }
    }
}
