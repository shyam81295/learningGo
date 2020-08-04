package main

import (
    // format package
    "fmt"
    "runtime"
)

func main() {
    // runtim.GOOS gives runtime program's OS
    fmt.Println("Hello from", runtime.GOOS)
}
