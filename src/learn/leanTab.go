package learn

import . "LearnGo/src/models"

func CreatTabPers(p1, p2, p3 Personne) [3]Personne {
	return [...]Personne{p1, p2, p3}
}

func ChangePers(tab *[3]Personne, pers Personne, ind int) {
	(*tab)[ind] = pers
}
