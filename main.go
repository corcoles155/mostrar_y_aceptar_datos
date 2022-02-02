package main

import (
	"fmt"
	"os"
	"bufio"
)

var numero1 int
var numero2 int
var resultado int
var descripcion string

func main()  {
	fmt.Println("Ingrese número 1: ")
	fmt.Scanln("%d", &numero1)

	fmt.Println("Ingrese número 2: ")
	fmt.Scanln("%d", &numero2)

	fmt.Println("Descripción: ")

	//Creamos un nuevo scanner donde nuestra entrada va a ser el teclado ((os.Stdin)
	scanner := bufio.NewcScanner(os.Stdin)
	//Si ha escaneado algo obtiene el texto
	if scanner.Scan(){
		//Recupero el valor y lo guarda en la variable leyenda
		descripcion = scanner.Text()
	}

	resultado = numero1 + numero2
	fmt.Println(descripcion, resultado)

}	
