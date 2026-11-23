package main

import (
	"LearnGo/src/learn"
	"fmt"
)

func main() {
	// =========================================================================
	//                             les Bases
	// =========================================================================
	var resultat1, resultat2 int = learn.Testfactoriel(3, 4)
	fmt.Println("factoriel de 3 : ", resultat1, "et factoriel de 4 :", resultat2)

	var resRec int = learn.FactRec(3)
	fmt.Printf("factoriel de 3 en rec : %d", resRec)

	// =========================================================================
	//                             les Pointeurs
	// =========================================================================
	var testIntPointeur int = 45
	var testStringPointeur = new(string)
	*testStringPointeur = "test"
	learn.TestPointeurs(&testIntPointeur, testStringPointeur)
	fmt.Printf("valeur changer pour 25 normalement :  %d \n", testIntPointeur)

	// =========================================================================
	//                             les Struct
	// =========================================================================
	var p1 learn.Person = learn.Person{Name: "ana", Age: 37}
	var p2 learn.Person = learn.Person{Name: "peter", Age: 22}
	var agediff int = learn.AgeDiff(p1, p2)
	fmt.Printf("la diff d'age entre peter et anna est  :  %d \n", agediff)

	// =========================================================================
	//                             les Tableaux
	// =========================================================================
	var p3 learn.Person = learn.Person{Name: "lola", Age: 18}
	var tabPers = learn.CreatTabPers(p1, p2, p3)
	fmt.Println("Tableau lors de la creation : ", tabPers)

	var p4 learn.Person = learn.Person{Name: "max", Age: 76}
	learn.ChangePers(&tabPers, p4, 0)
	fmt.Println("Tableau suuite au changement de ana en max : ", tabPers)
	// pour init un tab avec que des 0 :  var tab [5]in  ou   tab := [5]int{}

}
