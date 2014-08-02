package main

import (
    "fmt"
    "math/cmplx"
)

const Delta = 1e-10

func Cbrt(x complex128) complex128 {
    z,prev := complex128(1),complex128(0)
    for cmplx.Abs(z - prev) > Delta {
        prev = z
        z = z - (cmplx.Pow(z, 3) - x) / (3 * cmplx.Pow(z, 2))
    }
    return z
}

func main() {
    fmt.Println(Cbrt(2))
    fmt.Println(cmplx.Pow(2,1.0/3.0))
}
