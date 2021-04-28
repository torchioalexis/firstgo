package main

import "fmt"

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
	const controlUser = ""

	if user == "" && pass == "" {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}
}
