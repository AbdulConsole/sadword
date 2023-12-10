package main

import (
    "fmt"
    "os"
    "strings"
)

func go1337_(s string) string {
    replacements := map[rune]rune{
        'G': 'g',
        'a': '4',
        'A': '4',
        'e': '3',
        'E': '3',
        'i': '1',
        'I': '1',
        'T': '7',
        't': '7',
        'o': '0',
        'O': '0',
    }

    return strings.Map(func(r rune) rune {
        if replacement, ok := replacements[r]; ok {
            return replacement
        }
        return r
    }, s)
}

func main() {
    if len(os.Args) != 2 {
        fmt.Println("Usage: sadword \"Your Original String\"")
        return
    }

    originalString := os.Args[1]
    replacedString := go1337_(originalString)
    fmt.Println(replacedString)
}
