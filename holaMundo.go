package main

import (
	"fmt"
	"time"
)

func main() {
	var suma int = 8 + 9
	var texto string = "El resultado es:"
	otroTexto := "LALALA"

	fmt.Println("Hola Mundo")
	fmt.Println(otroTexto)
	time.Sleep(time.Second * 1)
	fmt.Println(texto, suma)
}
