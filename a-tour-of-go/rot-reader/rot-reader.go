package main

import (
    "fmt"
    "io"
    "os"
    "strings"
)

type rot13Reader struct {
    r io.Reader
}

func (r rot13Reader) Read(b []byte) (int, error) {
    n, err := r.r.Read(b)
    if err == nil {
        for i := 0; i < n; i++ {
            if b[i] > 'A' && b[i] < 'Z' {
                b[i] = 'A' + ((b[i] - 'A' + 13) % 26)
            } else if b[i] > 'a' && b[i] < 'z' {
                b[i] = 'a' + ((b[i] - 'a' + 13) % 26)
            }
        }
    }
    return n, err
}

func main() {
    s := strings.NewReader("Lbh penpxrq gur pbqr!")
    r := rot13Reader{s}
    // goland:noinspection GoUnhandledErrorResult
    _, e := io.Copy(os.Stdout, &r)
    if e != nil {
        fmt.Println("Could not write to stdout ")
    }
}
