package main

import (
	"fmt"
	"time"
)

type Gorra struct {
	marca  string
	color  string
	precio float32
	plana  bool
}

func main() {
	var suma int = 8 + 9
	var texto string = "El resultado es:"
	otroTexto := "LALALA"

	var gorraNegra = Gorra  {
		marca:"nike",
		color: "negra",
		precio: 10.5,
		plana: false,
	}

	fmt.Println("Hola Mundo")
	fmt.Println(otroTexto)
	fmt.Println(texto, suma)
	fmt.Println(gorraNegra)
	fmt.Println(gorraNegra.marca)
	time.Sleep(time.Second * 1)
}
