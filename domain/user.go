// Package domain provides ...
package domain

//User ....
type User struct {
	ID       NUUID
	Name     string
	Email    string
	Password string
}

//UserRepository ....
type UserRepository interface {
	/* TODO: add methods */
}
