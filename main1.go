package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
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

	var peliculas [3]string

	peliculas[0] = "1p"
	peliculas[1] = "2p"
	peliculas[2] = "3p"

	otraPeli := [3]string{
		"asda",
		"werwer",
		"iuyoiuyo",
	}

	var comentarios [3][2]string

	comentarios[0][0] = "0asd"
	comentarios[0][1] = "1asd"
	comentarios[1][0] = "2asd"
	comentarios[1][1] = "3asd"
	comentarios[2][0] = "4asd"
	comentarios[2][1] = "5asd"

	var pelis []string
	pelis = append(pelis, "ASD")
	pelis2 := append(pelis, "123")

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
	pantalon("rojo", "corto", "saraza")
	fmt.Println(peliculas)
	fmt.Println(peliculas[1])
	fmt.Println(otraPeli)
	fmt.Println(otraPeli[1])
	fmt.Println(comentarios)
	fmt.Println(pelis)
	fmt.Println(pelis2)
	fmt.Println((os.Args))
	valor, _ := strconv.Atoi(os.Args[2])
	fmt.Println(valor + 1)

}

func pantalon(caracteristicas ...string) {
	for _, c := range caracteristicas {
		fmt.Println(c)
	}
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
