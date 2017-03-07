package main
import "fmt"

func check(a,b int, bob * int, alice * int) {
    if a<b {
        *bob = *bob + 1
    } else if a>b {
        *alice = *alice + 1
    }
}

func main() {
    var a0, a1, a2, b0, b1, b2 int
    var bob, alice = 0, 0
    fmt.Scanf("%d %d %d", &a0, &a1, &a2)
    fmt.Scanf("%d %d %d", &b0, &b1, &b2)
    check(a0, b0, &bob, &alice)
    check(a1, b1, &bob, &alice)
    check(a2, b2, &bob, &alice)
    fmt.Printf("%d %d", alice, bob)
}
