package main

import "code.google.com/p/go-tour/pic"

func Pic(dx, dy int) [][]uint8 {
    matrix := make([][]uint8, dy)
    for y := range matrix {
        matrix[y] = make([]uint8, dx)
        for x := range matrix[y] {
            matrix[y][x] = uint8((x+y)/2)
        }
    }
    return matrix
}

func main() {
    pic.Show(Pic)
}
