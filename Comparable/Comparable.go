package comparable

type Person struct {
	firstName  string
	lastName   string
	age        int
	middleName []string
}

func NewPerson(fn, ls string, a int, mn []string) Person {
	return Person{
		firstName:  fn,
		lastName:   ls,
		age:        a,
		middleName: mn,
	}
}

func (p *Person) Equals(p1 Person) bool {
	if len(p.middleName) == len(p1.middleName) {
		return true

	}
	return false
}
