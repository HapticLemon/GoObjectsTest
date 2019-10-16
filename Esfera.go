package main

import "fmt"

type Esfera struct{
	BaseObject
	Radio float64
}

func (e Esfera) Distancia() {
	fmt.Printf("Función de distancia de la esfera\n")
}