package main

import "fmt"

func main() {
    input := []int{1,2,3}
    fmt.Printf("input: %v\n", input)
    fmt.Printf("output: %v\n", plusOne(input))
    input = []int{9}
    fmt.Printf("input: %v\n", input)
    fmt.Printf("output: %v\n", plusOne(input))
    input = []int{9,9,9}
    fmt.Printf("input: %v\n", input)
    fmt.Printf("output: %v\n", plusOne(input))
}


func plusOne(digits []int) []int {
    for i := len(digits) - 1 ; i >= 0 ; i-- {
        if digits[i] < 9{
            digits[i] = digits[i] + 1
            break
        }
        
        digits[i] = 0
        if i == 0 {
            digits = append([]int{1},digits...)
        }
    }
    return digits
}