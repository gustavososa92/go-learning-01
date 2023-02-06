package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	fmt.Println("**** Lector ****")
	nuevoTexto := "\n" + os.Args[1]

	file, err := os.OpenFile("file.txt", os.O_APPEND, 0777)
	showError(err)

	escribir, err := file.WriteString(nuevoTexto)
	showError(err)
	fmt.Println(escribir)

	file.Close()

	texto, err := ioutil.ReadFile("file.txt")
	showError(err)

	fmt.Println(string(texto))
}

func showError(err error) {
	if err != nil {
		panic(err)
	}
}
