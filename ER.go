package main

import (
	"fmt"
	"regexp"
)

func main() {
	//Test word
	w := "Aaa2a.e"

	//Parameters
	minLength := 7
	minSpecialChar := 1
	minBigLetters := 2
	minLittleLetters := 3
	minDigits := 1

	//Conditions
	con1 := fmt.Sprintf("%s%d%s", ".{", minLength, ",}")                 //.{7,}
	con2 := fmt.Sprintf("%s%d%s", "[^A-Za-z0-9]{", minSpecialChar, ",}") //[^A-Za-z0-9]{1,}
	con3 := fmt.Sprintf("%s%d%s", "[A-Z]{", minBigLetters, ",}")         //[A-Z]{2,}
	con4 := fmt.Sprintf("%s%d%s", "[a-z]{", minLittleLetters, ",}")      //[a-z]{3,}
	con5 := fmt.Sprintf("%s%d%s", "[0-9]{", minDigits, ",}")             //[0-9]{1,}

	//Regular Expression
	if !regexp.MustCompile(con1).MatchString(w) {
		fmt.Println("La longitud debe ser mayor a ", minLength, ".")
	} else if !regexp.MustCompile(con2).MatchString(w) {
		fmt.Println("Debe tener al menos ", minSpecialChar, " caracter(es) especial(es).")
	} else if !regexp.MustCompile(con3).MatchString(w) {
		fmt.Println("Debe tener al menos ", minBigLetters, " letra(s) mayuscula(s).")
	} else if !regexp.MustCompile(con4).MatchString(w) {
		fmt.Println("Debe tener al menos ", minLittleLetters, " letra(s) minuscula(s).")
	} else if !regexp.MustCompile(con5).MatchString(w) {
		fmt.Println("Debe tener al menos ", minDigits, " numero(s).")
	} else {
		fmt.Println("Constrasena aceptada.")
	}

}
