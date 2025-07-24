package main

import (
    "fmt"
    "goBip39/pkg"
)

func main() {
    fmt.Println("Welcome to my Golang project!")
    // Example usage of utility functions
    str := "123"
    num, err := pkg.StringToInt(str)
    if err != nil {
        fmt.Printf("Error converting string to int: %v\n", err)
        return
    }
    fmt.Printf("Converted string '%s' to integer: %d\n", str, num)

    intStr := pkg.IntToString(num)
    fmt.Printf("Converted integer %d back to string: '%s'\n", num, intStr)
}