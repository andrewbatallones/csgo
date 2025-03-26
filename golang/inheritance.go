package golang

import "fmt"

// Inheritance
// Go does not support inheritance out of the box so how would golang implement inheritance.

type User struct {
	Name string
}

// One way is the have the parent struct live as an attribute on the child struct.
// An issue with this is that if a function needs the parent class, it will not have access to the child class.
// I guess that would make sense with other languages, if it accepts the parent class, the function wouldn't really
// have access to the child class.
type Manager struct {
	Position string
	User     User
}

func PrintName(u *User) {
	fmt.Printf("%s\n", u.Name)
}

// To call the above func, it would pass in the field
func TestInheritence() {
	manager := Manager{
		Position: "marketing",
		User: User{
			Name: "Foo Bar",
		},
	}

	PrintName(&manager.User)
	fmt.Printf("%s\n", manager.Position)
}
