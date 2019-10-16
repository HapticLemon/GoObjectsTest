package main

func main() {
	var objetoEsfera Objeto
	var objetoCaja Objeto

	// Están declarados como objetos genéricos pero cada
	// uno se extiende de una forma distinta.
	//
	objetoEsfera = Esfera{
		BaseObject{0,0},
		2,
	}
	objetoCaja = Caja{
		BaseObject{1,0},
		1,
		1,
		1,
	}

	// Monto una slice genérica de tipo Objeto
	//
	var anything []Objeto
	anything = append( anything , objetoEsfera )
	anything = append( anything , objetoCaja )

	// Recorremos el slice y llamamos a Distancia sin que nos importe el tipo
	// de objeto que estemos usando.
	//
	for _, elemento := range anything{
		elemento.Distancia()
	}

}