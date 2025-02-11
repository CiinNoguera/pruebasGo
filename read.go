// 
package main

import (
	"fmt"
	"os"
)

type Name struct {
	Fname string
	Lname string
}

func main() {

	fmt.Print("Enter the name of the text file: ")
	var filename string
	fmt.Scanln(&filename)

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var names []Name

	buffer := make([]byte, 40)

	for {
		n, err := file.Read(buffer)
		if n == 0 || err != nil {
			break
		}

		fname := string(buffer[:20])
		lname := string(buffer[20:])

		name := Name{
			Fname: fname,
			Lname: lname,
		}

		names = append(names, name)
	}

	fmt.Println("\nNames from file:")
	for _, name := range names {
		fmt.Println(name.Fname, name.Lname)
	}
}
