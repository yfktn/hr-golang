package main
import "fmt"

func main() {
    N := 0
    fmt.Scan(&N)
    var a, b, buff = 0, N-1, 0
    var cnta, cntb = 0, 0
    for line:=0; line<N; line++ {
        for col:=0; col<N; col++ {
            fmt.Scan(&buff)
            if col == a {
                cnta = cnta + buff
            }
            if col == b {
                cntb = cntb + buff
            }
        }
        a = a + 1
        b = b - 1
    }
    diff := cnta - cntb
    if(diff < 0) {
        diff = diff * -1
    }
    fmt.Println(diff)
}
