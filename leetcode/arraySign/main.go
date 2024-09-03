package main

import "fmt"

func main() {
    fmt.Println(arraySign([]int{-1,-2,-3,-4,3,2,1}))
    fmt.Println(arraySign([]int{1,5,0,2,-3}))
    fmt.Println(arraySign([]int{-1,1,-1,1,-1}))
}

func arraySign(nums []int) int {
    isNegative := 1
    for _, n := range nums {
        if n == 0 {
            return 0
        }
        if n < 0 {
            isNegative = -1 * isNegative
        }
    }
    return isNegative
}