package main
import "fmt"

func main() {
    var numSeq int64
    numQueries := 0
    fmt.Scanf("%d %d", &numSeq, &numQueries)
    //S := [numSeq][1] int64
    S := make([][]int64, numSeq)
    var x, y, lastAns int64
    var t, Soff, size int
    lastAns = 0
    // scan queries
    for i:=0; i<numQueries; i++ {
        fmt.Scanf("%d %d %d", &t, &x, &y)
        if t == 1 {
            Soff = int((x ^ lastAns) % numSeq)
            S[Soff] = append(S[Soff], y)
        } else if t == 2 {
            Soff = int((x ^ lastAns) % numSeq)
            size = len(S[Soff])
            SoffY := y % int64(size)
            //fmt.Println(size, " ", Soff, " ", SoffY, S[Soff])
            lastAns = S[Soff][SoffY]
            fmt.Printf("%d\n", lastAns)
        }
    }
    
}
