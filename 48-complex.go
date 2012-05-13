package main

import (
    "fmt"
    "math"
)


func Cbrt(x complex128) complex128 {
    z := x
        for i:=0; i<10000; i++ {
            z = z-(z*z*z-x)/(3*z*z)
        }
    return z
}

func main() {
    fmt.Println(Cbrt(2))
    fmt.Println(math.Cbrt(2))
}
