package main

import "fmt"

const (
    _ = iota // 0
    KB = 1 << (iota * 10)
    MB = 1 << (iota * 10)
    GB = 1 << (iota * 10)
    TB = 1 << (iota * 10)
)

func main() {
    fmt.Println("binary\t\tdecimal")
    fmt.Printf("%b\t", KB)
    fmt.Printf("%d\n", KB)

    fmt.Printf("%b\t", MB)
    fmt.Printf("%d\n", MB)

    fmt.Printf("%b\t", TB)
    fmt.Printf("%d\n", TB)
}
