//une fonction avec une lettre minuscule en premiere lettre n'est pas accesible depuis un autre package
//tuto : https://go.developpez.com/tutoriels/go-par-l-exemple/#LX

package learn

import "fmt"

func Testfactoriel(a, b int) (int, int) {
	fmt.Printf(" on veut faire factoriel de %d et factoriel de : %d \n", a, b)
	var res1, res2 int = 1, 1
	for i := 1; i <= a; i++ {
		res1 *= i
	}
	for i := 1; i <= b; i++ {
		res2 *= i
	}
	return res1, res2
}

func FactRec(a int) int {
	if a == 0 {
		return 1
	}
	return a * FactRec(a-1)
}

//types de for
/*
classique
for i := 0; i < 5; i++ {
	fmt.Println(i)
}

conditionelle sert de while
n := 1
for n < 5 {
	n *= 2
}

for range
permetent de parcourir un tableau, slice, map, string
for index, valeur := range fruits {
	fmt.Println(index, valeur)
}

// Utilisez _ pour ignorer l'index ou la valeur
for _, valeur := range fruits {
	fmt.Println(valeur)
}

*/
