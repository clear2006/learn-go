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

func ChangeTabToSlice(tab [3]Personne) []Personne {
	slice := make([]Personne, len(tab)) //cree un slice avec deja 3 emplacement avec les valeur par defaut des champ de personne(pas utiliser si append)
	copy(slice, tab[:])                 //transforme le tab en slice
	return slice
}

func AddToSlice(tab [3]Personne, slice []Personne) []Personne {
	slice = append(slice, tab[:]...) //transforme le tableau puis y ajoute tout ces elements
	return slice
}

//numbers := []int{10, 20, 30, 40, 50}
//slice1 := numbers[1:3]  Éléments de l’indice 1 à l’indice 3 (3 exclu)
//slice2 := numbers[:2]   Du début jusqu’à l’indice 2 (2 exclu)
//slice3 := numbers[3:]   De l’indice 3 jusqu’à la fin
//s1 = append(s1, s2...)
// Équivaut à écrire : append(s1, 3, 4, 5)
