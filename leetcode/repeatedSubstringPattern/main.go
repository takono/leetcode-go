package main

import "fmt"

func main() {
   fmt.Println(repeatedSubstringPattern("abab"))
    fmt.Println(repeatedSubstringPattern("aba"))
    fmt.Println(repeatedSubstringPattern("abcabcabcabc"))
}


func repeatedSubstringPattern(s string) bool {
    for i := 1; i <= len(s)/2; i++ {
        if len(s)%i == 0 {
            for j := 0; j < len(s); j += i {
                if s[:i] != s[j : j+i]{
                    break
                }
                if j == len(s)-i {
                    return true
                }
            }
        }
    }
    return false
}