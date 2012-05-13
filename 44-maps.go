package main

import (
    "strings"
    "code.google.com/p/go-tour/wc"
)

func WordCount(s string) map[string]int {
    count := make(map[string]int)
    for _, value := range strings.Fields(s) {
        count[value]++
    }
    return count
}

func main() {
    wc.Test(WordCount)
}
