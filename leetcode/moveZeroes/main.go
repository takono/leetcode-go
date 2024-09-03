package main

import "fmt"

func main() {
    lists := [][]int{{0,1,0,3,12}, {0}, {0,0,1}}
    for _, list := range lists {
        moveZeroes(list)
    }
}


func moveZeroes(nums []int) {
    fmt.Printf("input: %v \n", nums)
    zeroCount := 0
    for i := 0; i < len(nums) - zeroCount; {
        if nums[i] == 0 {
            nums = append(nums[:i], nums[i+1:]...)
            nums = append(nums, 0)
            zeroCount++
        } else{
            i++
        }
    }
    fmt.Printf("output: %v \n", nums)
}

func moveZeroesTwoArray(nums []int) {
    fmt.Printf("input: %v \n", nums)
    zeroCount := 0
    for i := 0; i < len(nums) - zeroCount; {
        if nums[i] == 0 {
            nums = append(nums[:i], nums[i+1:]...)
            nums = append(nums, 0)
            zeroCount++
        } else{
            i++
        }
    }
    fmt.Printf("output: %v \n", nums)
}