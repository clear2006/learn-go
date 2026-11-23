package learn

func CreatTabPers(p1, p2, p3 Person) []Person {
	return []Person{p1, p2, p3}
}

func ChangePers(tab *[]Person, pers Person, ind int) {
	(*tab)[ind] = pers
}
