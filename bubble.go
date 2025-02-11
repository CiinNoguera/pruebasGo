package main

import (
	"fmt"
)


func Swap(arr []int, i int) {
	arr[i], arr[i+1] = arr[i+1], arr[i]
}


func BubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				Swap(arr, j)
			}
		}
	}
}

func main() {

	numeros1 := []int{5, 2, 9, 1, 6, 8, 4, 3, 7, 10}
	fmt.Println("Antes de ordenar:", numeros1)
	BubbleSort(numeros1)
	fmt.Println("Después de ordenar:", numeros1)

	numeros2 := []int{-3, 10, -1, 8, 7, -5, 2, 4, 0, 6}
	fmt.Println("\nAntes de ordenar:", numeros2)
	BubbleSort(numeros2)
	fmt.Println("Después de ordenar:", numeros2)
}
