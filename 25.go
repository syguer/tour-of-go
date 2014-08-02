package main

import (
    "fmt"
    "math"
)

const Delta = 1e-10

func Sqrt(x float64) float64 {
    z := float64(1)
    prev := float64(0)
    for math.Abs(z - prev) > Delta {
        prev = z
        z = z - (z * z - x) / (2 * z)
    }
    return z
}

func main() {
    fmt.Println(Sqrt(2))
    fmt.Println(math.Sqrt(2))
}
