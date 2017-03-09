package main
import "fmt"

func main() {
    var numInt, numRot int32
    fmt.Scanf("%d %d", &numInt, &numRot)
    A := make([]int32, numInt)
    for i:=int32(0); i<numInt; i++ {
        fmt.Scanf("%d", &A[i])
    }
    B := make([]int32, numInt)
    // rotate
    for i:=int32(0); i<numRot; i++ {
        B = append(A[1:], A[0])
        A = B
    }
    // cetak
    for i:=int32(0); i<numInt; i++ {
        fmt.Printf("%d ", A[i])
    }
    
}
