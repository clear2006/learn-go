package models

type Personne struct {
	Name string `json:"nom" db:"nom"`
	Age  int    `json:"age" db:"age"`
}

func AgeDiff(p1, p2 Personne) int {
	var agediff int = 0
	if p1.Age > p2.Age {
		agediff = p1.Age - p2.Age
	} else {
		agediff = p2.Age - p1.Age
	}
	return agediff
}
