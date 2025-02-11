// package main

// import (
// 	"bufio"
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"os"
// 	"strings"
// )

// type Person struct {
// 	Nombre string `json:"nombre"`
// 	Direccion string `json:"direccion"`
// }

// func main() {

// 	reader := bufio.NewReader(os.Stdin)

// 	fmt.Print("Ingrese su nombre:")
//     nombre, err := reader.ReadString('\n')
// 	if err != nil {
// 		log.Fatal("Error al leer el nombre", err )
// 	}

// 	nombre = strings.TrimSpace(nombre)

// 	fmt.Print("Ingrese dirección:")
// 	direccion, err := reader.ReadString('\n')
// 	if err != nil {
// 		log.Fatal("Error al leer dirección", err )
// 	}

// 	direccion = strings.TrimSpace(direccion)

// 	persona := Person {
// 		Nombre: nombre,
// 		Direccion: direccion,
// 	}

// 	jsonData, err := json.Marshal(persona)
// 	if err != nil {
// 		log.Fatal("Error al convertir JSON", err)
// 	}

// 	fmt.Println("Objeto JSON:")
// 	fmt.Println(string(jsonData))

// }