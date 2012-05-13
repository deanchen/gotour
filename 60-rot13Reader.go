package main

import (
    "io"
    "os"
    "strings"
)

type rot13Reader struct {
    r io.Reader
}

func (r *rot13Reader) Read(p []byte) (n int, err error) {
    n, err = r.r.Read(p)
    for i := range p[:n] {
        switch {
        case 'A' <= p[i] && p[i]<= 'M':
            p[i] = (p[i] - 'A') + 'N'
        case 'N' <= p[i] && p[i] <= 'Z':
            p[i]= (p[i] - 'N') + 'A'
        case 'a' <= p[i]&& p[i] <= 'm':
            p[i]= (p[i] - 'a') + 'n'
        case 'n' <= p[i]&& p[i] <= 'z':
            p[i]= (p[i] - 'n') + 'a'
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
