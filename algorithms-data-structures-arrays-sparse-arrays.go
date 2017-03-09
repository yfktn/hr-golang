package main
import "fmt"

func main() {
    n1st, n2nd := 0, 0
    fmt.Scanf("%d", &n1st)
    N := make([]string, n1st)
    for i:=0;i<n1st;i++ {
        fmt.Scanf("%s", &N[i])
    }
    fmt.Scanf("%d", &n2nd)
    var Q string
    var nN int
    for i:=0;i<n2nd;i++ {
        nN = 0
        fmt.Scanf("%s", &Q)
        for j:=0;j<n1st;j++ {
            if Q == N[j] {
                nN += 1
            }
        }
        fmt.Println(nN)
    }
}
