package learn

type Person struct {
	Name string
	Age  int
}

func AgeDiff(p1, p2 Person) int {
	var agediff int = 0
	if p1.Age > p2.Age {
		agediff = p1.Age - p2.Age
	} else {
		agediff = p2.Age - p1.Age
	}
	return agediff
}
