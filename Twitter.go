package main

import (
	"fmt"
	"regexp"
)

func main() {
	//Test word
	usuario := "@aa_"

	//Regular Expression
	r := regexp.MustCompile("^@[A-Za-z0-9-_]+$")

	if r.MatchString(usuario) {
		fmt.Println("Usurario aceptado.")
	} else {
		fmt.Println("Usurario denegado.")
	}

}
