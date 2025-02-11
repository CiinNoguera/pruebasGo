// package main

// import (
// 	"fmt"
// 	"sort"
// )

// func main() {

// 	var numbers []int
// 	var num int

// 	fmt.Println("Introduce números enteros. O presiona '0' para salir:")

// 	for {
// 		fmt.Print("Introduce un número: ")
// 		_, err := fmt.Scanln(&num)

// 		if err != nil {
// 			fmt.Println("Entrada inválida. Intentelo nuevamente")
// 			continue
// 		}

// 		if num == 0 {
// 			break
// 		}

// 		numbers = append(numbers, num)

// 		sort.Ints(numbers)

// 		fmt.Println("Slice ordenada: ", numbers)
// 	}

// 	fmt.Println("Programa finalizado")


// }