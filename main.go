package main

import (
	"errors"
	"fmt"
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
	fmt.Println(gorras(45, "$"))
}

func gorras(pedido float32, moneda string) (string, string, float32) {
	precio := func() float32 {
		return pedido * 7
	}
	return "EL precio total de gorras es:", moneda, precio()
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
