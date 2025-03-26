package examples

import "fmt"

type User struct {
	name string
}

func (u *User) PrintName() {
	fmt.Printf("%s\n", u.name)
}

// I would have the subclass here and have the parent class be an attribute
type Manager struct {
	position string
	user     User
}
