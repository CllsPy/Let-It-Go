package main

import (
	"fmt"
)

func median(arr []int){
	sum := 0
	for _, value := range arr{
		sum += value
	}
	fmt.Println("Median:", sum/len(arr))
}

func main() {
	arr := []int{10, 2}
	median(arr)

}
