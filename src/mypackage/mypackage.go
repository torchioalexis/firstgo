package mypackage

import "fmt"

//CarPublic con acceso publico
type CarPublic struct {
	Brand string
	Year  int
}

//PrintMessage con acceso publico
func PrintMessage(text string) {
	fmt.Println(text)
}
