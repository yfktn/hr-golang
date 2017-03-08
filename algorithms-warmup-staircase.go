package main
import "fmt"

func main() {
    n := 0
    fmt.Scanf("%d", &n)
    offset := n-1
    for i:=0; i<n; i++ {
        for j:=0; j<n; j++ {
            if j>=offset {
                fmt.Print("#")
            } else {
                fmt.Print(" ")
            }
        }
        fmt.Print("\n")
        offset = offset - 1
    }
    fmt.Println("")
}
