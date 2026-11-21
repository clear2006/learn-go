//une fonction avec une lettre minuscule en premiere lettre n'est pas accesible depuis un autre package
//tuto : https://go.developpez.com/tutoriels/go-par-l-exemple/#LX

package learn

import "fmt"

type person struct {
	name string
	age  int
}

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

func TestPointeurs(a *int, b *string) {
	fmt.Printf(" voila le pointeur sur a :%p  et la valeur de a : %d \n", a, *a)
	fmt.Printf(" voila le pointeur sur b :%p  et la valeur de b : %s \n", b, *b)
	*a = 25
}
