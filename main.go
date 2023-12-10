package main

import (
	"bufio"
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
	var originalString string

	if len(os.Args) > 1 {
		originalString = os.Args[1]
	} else {
		scanner := bufio.NewScanner(os.Stdin)

		if scanner.Scan() {
			originalString = scanner.Text()
		} else {
			fmt.Println("Error reading from stdin.")
			return
		}
	}

	if originalString == "" {
		fmt.Println("No input provided. Usage: sadword \"Your Original String\"")
		return
	}

	replacedString := go1337_(originalString)
	fmt.Println(replacedString)
}
