package main

import (
    "io"
    "os"
    "strings"
)

type rot13Reader struct {
    r io.Reader
}

func (r *rot13Reader) Read (p []byte) (n int, err error) {
    n,err = r.r.Read(p)
    for i := range p {
        if ('A' <= p[i] && p[i] <= 'M') || ('a' <= p[i] && p[i] <= 'm') {
            p[i] += 13
        } else if ('N' <= p[i] && p[i] <= 'Z') || ('n' <= p[i] && p[i] <= 'z') {
            p[i] -= 13
        }
    }
    return
}

func main() {
    s := strings.NewReader(
        "Lbh penpxrq gur pbqr!")
    r := rot13Reader{s}
    io.Copy(os.Stdout, &r)
}
