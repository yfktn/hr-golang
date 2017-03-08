package main
import "fmt"

func main() {
    var a [5]int64
    var min, max, sum int64
    /*for i:=0; i<5; i++ {
        fmt.Scan(&a[i])
    }*/
    fmt.Scanf("%d %d %d %d %d", &a[0], &a[1], &a[2], &a[3], &a[4])
    init := true
    for no:=0; no<5; no++ {
        sum = 0
        for i,v := range a {
            if i != no {
                sum += v
            }
        }
        if sum >= max || init {
            max = sum
        }
        if sum <= min || init {
            min = sum
        }
        init = false
    }
    fmt.Printf("%d %d", min, max)
}
