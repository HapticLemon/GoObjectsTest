package main

import "fmt"

type Caja struct{
	BaseObject
	Alto int
	Ancho int
	Fondo int
}

func (c Caja) Distancia() {
	fmt.Printf("Función de distancia de la caja\n")
}