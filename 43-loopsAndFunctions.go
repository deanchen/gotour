package main

import (
    "fmt"
    "math"
)

func Sqrt(x float64) float64 {
    z := float64(1)
    for i:=0; i<1000; i++ {
        z = z - (z*z-x)/2*z
    }
    return z
}

func main() {
    fmt.Println("My function", Sqrt(2))
    fmt.Println("Real function", math.Sqrt(2))
}
