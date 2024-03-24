package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var nombre, direccion string
	fmt.Print("Introduce un nombre: ")
	fmt.Scanln(&nombre)
	fmt.Print("Introduce una dirección: ")
	fmt.Scanln(&direccion)

	datos := make(map[string]string)
	datos["nombre"] = nombre
	datos["direccion"] = direccion

	jsonData, err := json.Marshal(datos)
	if err != nil {
		fmt.Println("Error al convertir a JSON:", err)
		return
	}

	fmt.Println(string(jsonData))
}
