package main

import (
	"errors"
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

	var gorraNegra = Gorra{
		marca:  "nike",
		color:  "negra",
		precio: 10.5,
		plana:  false,
	}

	fmt.Println("Hola Mundo")
	fmt.Println(otroTexto)
	fmt.Println(texto, suma)
	fmt.Println(gorraNegra)
	fmt.Println(gorraNegra.marca)
	fmt.Println(operacion(10, 5, "+"))
	fmt.Println(operacion(10, 5, "-"))
	fmt.Println(operacion(10, 5, "*"))
	fmt.Println(operacion(10, 5, "/"))
	fmt.Println(operacion(10, 5, "/+"))
	time.Sleep(time.Second * 1)
}

func operacion(n1 float32, n2 float32, op string) (float32, error) {
	switch op {
	case "+":
		return n1 + n2, nil
	case "-":
		return n1 - n2, nil
	case "*":
		return n1 * n2, nil
	case "/":
		return n1 / n2, nil
	}
	return 0, errors.New("No soportado")
}
