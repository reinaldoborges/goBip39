package main

import (
    "bufio"
    "embed"
    "fmt"
    "os"
    "strings"
)

//go:embed data/english.txt
var englishTxt embed.FS

func main() {
    fmt.Println("Welcome to the BIP39 word search tool!")
    fmt.Println("You can search for a word or a hexadecimal line number in the BIP39 English word list.")
    fmt.Println("To search for a word, type the word (or the first 4 characters of it).")
    fmt.Println("To search for a line number, type the line number in hexadecimal format (e.g., 0x001).\n")

    reader := bufio.NewReader(os.Stdin)
    for {
        fmt.Print("Enter a word or hex line number to search (empty to exit): ")
        input, _ := reader.ReadString('\n')
        input = strings.TrimSpace(input)
        if input == "" {
            fmt.Println("Exiting.")
            break
        }

        file, err := englishTxt.Open("data/english.txt")
        if err != nil {
            fmt.Printf("Error opening embedded file: %v\n", err)
            continue
        }

        // Check if input is a hexadecimal number (starts with "0x" or "0X")
        if strings.HasPrefix(input, "0x") || strings.HasPrefix(input, "0X") {
            var lineNum int
            hexStr := strings.TrimPrefix(strings.ToLower(input), "0x")
            _, err := fmt.Sscanf(hexStr, "%x", &lineNum)
            if err != nil {
                fmt.Printf("Invalid hexadecimal input: %v\n", err)
                file.Close()
                continue
            }
            scanner := bufio.NewScanner(file)
            currentLine := 0
            found := false
            for scanner.Scan() {
                if currentLine == lineNum {
                    word := strings.TrimSpace(scanner.Text())
                    fmt.Printf("Line 0x%03X: %s\n", lineNum, word)
                    found = true
                    break
                }
                currentLine++
            }
            if !found {
                fmt.Printf("Line 0x%03X not found in the list\n", lineNum)
            }
            file.Close()
            continue
        }

        // Otherwise, search for the word (compare only first 4 chars if longer)
        scanner := bufio.NewScanner(file)
        found := false
        lineNumber := 1
        for scanner.Scan() {
            line := strings.TrimSpace(scanner.Text())
            compareInput := input
            compareLine := line
            if len(input) > 4 {
                compareInput = input[:4]
            }
            if len(line) > 4 {
                compareLine = line[:4]
            }
            if compareLine == compareInput {
                fmt.Printf("Word '%s' found at line 0x%03X\n", input, lineNumber-1)
                found = true
                break
            }
            lineNumber++
        }

        if err := scanner.Err(); err != nil {
            fmt.Printf("Error reading file: %v\n", err)
            file.Close()
            continue
        }

        if !found {
            fmt.Printf("Word '%s' not found in the list\n", input)
        }
    }
}