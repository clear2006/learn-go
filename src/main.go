package main

import (
	"LearnGo/src/learn"
	"fmt"
)

func main() {
	var resultat1, resultat2 int = learn.Testfactoriel(3, 4)

	fmt.Println("factoriel de 3 : ", resultat1, "et factoriel de 4 :", resultat2)
	var resRec int = learn.FactRec(3)
	fmt.Printf("factoriel de 3 en rec : %d", resRec)

	var testintp int = 45
	var teststrp string = "test"
	learn.TestPointeurs(&testintp, &teststrp)
	fmt.Printf("valeur changer pour 25 normalement :  %d \n", testintp)

	var p1 learn.Person = learn.Person{Name: "ana", Age: 37}
	var p2 learn.Person = learn.Person{Name: "peter", Age: 22}
	var agediff int = learn.AgeDiff(p1, p2)
	fmt.Printf("la diff d'age entre peter et anna est  :  %d \n", agediff)
}
