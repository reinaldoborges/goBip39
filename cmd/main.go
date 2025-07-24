package main

import (
    "bufio"
    "fmt"
    "goBip39/pkg"
    "os"
    "strings"
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

    fmt.Print("Enter a word to search: ")
    var input string
    fmt.Scanln(&input)
    input = strings.TrimSpace(input)

    file, err := os.Open("./data/english.txt")
    if err != nil {
        fmt.Printf("Error opening file: %v\n", err)
        return
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    found := false
    lineNumber := 1
    for scanner.Scan() {
        if strings.TrimSpace(scanner.Text()) == input {
            fmt.Printf("Word '%s' found at line %d\n", input, lineNumber)
            found = true
            break
        }
        lineNumber++
    }

    if err := scanner.Err(); err != nil {
        fmt.Printf("Error reading file: %v\n", err)
        return
    }

    if !found {
        fmt.Printf("Word '%s' not found in the list\n", input)
    }
}