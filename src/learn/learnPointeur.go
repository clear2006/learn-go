package learn

import "fmt"

func TestPointeurs(a *int, b *string) {
	fmt.Printf(" voila le pointeur sur a :%p  et la valeur de a : %d \n", a, *a)
	fmt.Printf(" voila le pointeur sur b :%p  et la valeur de b : %s \n", b, *b)
	*a = 25
}
