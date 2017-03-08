package main
import (
    "fmt"
    "strconv"
)

func main() {
    ins := "00:00:00PM"
    fmt.Scanf("%s", &ins)    
    isPM := ins[len(ins)-2:] == "PM"
    ins = ins[0:len(ins)-2]
    h,_ := strconv.Atoi(ins[0:2])
    if isPM {
        if h != 12 {
            h += 12
        }
        ins = fmt.Sprint(h) + ins[2:len(ins)]
    } else if !isPM && h == 12 {
        ins = "00" + ins[2:len(ins)]
    } 
    fmt.Printf("%s", ins)
}
