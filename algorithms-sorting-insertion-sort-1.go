package main
import "fmt"

func cetak(a []int, size int) {
    for i:=0; i<size;i++ {
        fmt.Printf("%d ", a[i])
    }
    fmt.Print("\n")
}

func main() {
    ns := 0
    fmt.Scan(&ns)
    ai := make([]int, ns)
    for i:=0; i<ns; i++ {
        fmt.Scan(&ai[i])
    }
    e := ai[ns-1]
    i:=0
    for i=ns-2; i>=0; i-- {
        if ai[i] >= e {
            ai[i+1] = ai[i]
            cetak(ai, ns)
        } else {
            ai[i+1] = e
            cetak(ai, ns)
            break
        }
    }
    if i < 0 {
        ai[0] = e
        cetak(ai, ns)
    }
}
