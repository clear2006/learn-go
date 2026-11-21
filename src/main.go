package main

import (
	"LearnGo/src/learn"
	"fmt"
)

func main() {
	var resultat1, resultat2 int = learn.Testfactoriel(3, 4)

	fmt.Printf("factoriel de 3factoriel: %d et factoriel de 4 : %d\n", resultat1, resultat2)
	var resRec int = learn.FactRec(3)
	fmt.Printf("factoriel de 3 en rec : %d", resRec)

	var testintp int = 45
	var teststrp string = "test"
	learn.TestPointeurs(&testintp, &teststrp)
	fmt.Printf("valeur changer pour 25 normalement :  %d \n", testintp)

}
