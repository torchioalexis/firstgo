package main

import (
	"fmt"
	"strings"

	"github.com/torchioalexis/firstgo/src/mypackage"
)

func isPalindromo(text string) {
	text = strings.ToLower(text)
	var textReverse string
	for i := len(text) - 1; i >= 0; i-- {
		textReverse += string(text[i])
	}
	if text == textReverse {
		fmt.Println("Es un palíndromo")
	} else {
		fmt.Println("No es un palíndromo")
	}
}

type pc struct {
	ram   int
	brand string
	disk  int
}

func (myPC pc) String() string {
	return fmt.Sprintf("Tengo %d GB de RAM, %d GB de disco y soy marca %s", myPC.ram, myPC.disk, myPC.brand)
}

func main() {
	// CALCULAR ÁREA RECTÁNGULO
	baseRectangulo := 10
	alturaRectangulo := 5
	areaRectangulo := baseRectangulo * alturaRectangulo
	fmt.Println("El área del rectángulo es:", areaRectangulo)

	// CALCULAR ÁREA TRAPECIO
	baseTrapecio1 := 10
	baseTrapecio2 := 20
	alturaTrapecio := 8
	areaTrapecio := (baseTrapecio1 + baseTrapecio2) * alturaTrapecio / 2
	fmt.Println("El área del trapecio es:", areaTrapecio)

	// CALCULAR ÁREA CIRCULO
	pi := 3.14
	var radioCirculo float64 = 5
	areaCirculo := pi * radioCirculo * radioCirculo
	fmt.Println("El área del círculo es:", areaCirculo)

	//FOR INVERTER

	reverseCounter := 21
	for reverseCounter >= 0 {
		fmt.Println(reverseCounter)
		reverseCounter--
	}

	//CONDICIONALES

	toTest := 19

	if toTest%2 == 0 {
		fmt.Println("El número", toTest, "es par")
	} else {
		fmt.Println("El número", toTest, "es impar")
	}

	user := ""
	pass := ""

	if user == "" && pass == "" {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}

	//RECORRIDO DE SLICES CON RANGE PARA COMPROBAR PALÍNDROMOS

	isPalindromo("CAFE = efac")

	//PRUEBAS DE STRUCTS PRIVADAS Y PUBLICAS
	var myCar mypackage.CarPublic
	myCar.Brand = "Peugeot"
	myCar.Year = 2021
	fmt.Println(myCar)

	//PRUEBAS DE FUNCIONES PRIVADAS Y PUBLICAS
	mypackage.PrintMessage("¡¡¡VAMOS CANAYA!!!")

	//PERSONALIZACIÓN DE OUTPUT EN CONSULA PARA STRUCTS
	myPC := pc{ram: 16, brand: "msi", disk: 100}

	fmt.Println(myPC)
}
