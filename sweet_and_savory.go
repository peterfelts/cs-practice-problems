package main

import (
    "sort"
)

func SweetAndSavory(dishes []int, target int) []int {

    bestPair := []int{0,0}
    bestDelta := int(^uint(0) >> 1)

    
    // Sort the arry, then split it in to two
    // different arrays (sweet and savory)
    sort.Ints(dishes)

    // Find the point where we positive 
    // integers start
    startIndex := 0
    for i:=0; i<len(dishes); i++ {
        if dishes[i] > 0 {
            startIndex = i
            break
        }
    }
    
    sweetDishes := dishes[:startIndex]
    savoryDishes := dishes[startIndex:]

    sweetIndex := len(sweetDishes)-1 // negative index
    savoryIndex := 0 // positive index
    
    for {
        if sweetIndex <0 || savoryIndex > len(savoryDishes)-1 {
            break
        }

        total := sweetDishes[sweetIndex] + savoryDishes[savoryIndex]

        if total <= target {
            delta := target - total

            // if the current delta is <= bestDelta, update bestDelta
            if delta < bestDelta {
                bestPair = []int{sweetDishes[sweetIndex], savoryDishes[savoryIndex]}
                bestDelta = delta
            }
            savoryIndex++

        }else {
            sweetIndex--
        }
    }

    return bestPair
}
