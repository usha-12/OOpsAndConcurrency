package Struct1

type Person struct {
	firstName string
	lastName  string
	age       int
}
//constractor - here in golang go constactor concept we can simply create object
//create a simple function with return type of person 
//constarctor also set the value for person 
//we can use one either constractor or set method  
func NewPerson(fn, ls string, a int) Person {
	return Person{
		firstName: fn,
		lastName:  ls,
		age:       a,
	}
}

func (p *Person) SetFristName(f string) {//setter function set the value
	p.firstName = f
}

func (p *Person) GetFirstName() string {
	return p.firstName
}