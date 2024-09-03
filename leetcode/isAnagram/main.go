package main

import (
	"fmt"
	"sort"
)

func main() {
    fmt.Println("isAnagram")
    fmt.Println(isAnagram("anagram", "nagaram"))
    fmt.Println(isAnagram("rat", "car"))

    fmt.Println("isAnagramWithHashMap")
    fmt.Println(isAnagramWithHashMap("anagram", "nagaram"))
    fmt.Println(isAnagramWithHashMap("rat", "car"))
}

func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    sliceS := []byte(s)
    sliceT := []byte(t)
    sort.Slice(sliceS, func(i, j int) bool { return sliceS[i] < sliceS[j] })
    sort.Slice(sliceT, func(i, j int) bool { return sliceT[i] < sliceT[j] })
    return string(sliceS) == string(sliceT)
}



func isAnagramWithHashMap(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    counts := make(map[byte]int)
    for i := 0; i < len(s); i++ {
        if _, ok := counts[s[i]]; ok {
            counts[s[i]]++
        } else {
            counts[s[i]] = 1
        }
    }
    for i := 0; i < len(t); i++ {
        if _, ok := counts[t[i]]; ok {
            if counts[t[i]] == 0 {
                return false
            }
            counts[t[i]]--
        } else {
            return false
        }
    }
    return true
}