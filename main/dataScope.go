package main

type User struct { //heap
	ID   int
	Name string
}

func createUser(id int, name string) *User { //scape analysis ----determinds that bether variable store in stack or promoted in heap
	user := User{
		ID:   id,
		Name: name,
	}
	return &user //heap
}

func main() {
	user := createUser(1, "John Doe")                     //heap
	println("User ID:", user.ID, "User Name:", user.Name) //heap
}
