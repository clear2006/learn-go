package learn

import (
	. "LearnGo/src/models"
	"fmt"
)

func AfficherNoms(personne []Personne) {
	for _, p := range personne {
		fmt.Println(p.Name)
	}
}

func ChangeTabToSlice(tab []Personne) []Personne {
	slice := make([]Personne, len(tab))
	copy(slice, tab)
	return slice
}
