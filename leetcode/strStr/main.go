package main

import "fmt"

func main() {
    fmt.Println(strStr("sadbutsad", "sad"))
    fmt.Println(strStr("xadbutsad", "sad"))
    fmt.Println(strStr("leet", "leeto"))
}


// Sliding window
func strStr(haystack string, needle string) int {
    startIndex :=  -1
    for i := 0; i < len(haystack) - len(needle) + 1; i++ {
        for j := 0; j < len(needle); j++ {
            if haystack[i+j] != needle[j] {
                break
            }
            if j == len(needle) - 1 {
                startIndex = i
                break
            }
        }
        if startIndex != -1 {
            break
        }
    }
    return startIndex
}