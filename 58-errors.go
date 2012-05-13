package main

import (
    "fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
    return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(f float64) (float64, error) {
    z := float64(1)
    if z < 0 {return 0, ErrNegativeSqrt(z)}
    for i:=0; i<10; i++ {
        z = z - (z*z-f)/2*z
    }
    return z, nil
}

func main() {
    fmt.Println(Sqrt(-2))
    fmt.Println(Sqrt(2))
}
