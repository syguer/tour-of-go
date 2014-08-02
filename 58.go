package main

import (
    "fmt"
    "math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string{
     return fmt.Sprintf("cannot Sqrt negative number: %d\n", int64(e))
}

const Delta = 1e-10

func Sqrt(f float64) (float64, error) {
    if f < 0 {
        e := ErrNegativeSqrt(f)
        return f, e
    }
    z, prev := 2.0, 0.0
   
    for math.Abs(prev - z) > Delta {
        prev = z
        z = z - (math.Pow(z, 2) - f) / (2 * z)
    }
    return z, nil
}

func main() {
    fmt.Println(Sqrt(2))
    fmt.Println(Sqrt(-2))
}
