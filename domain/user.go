// Package domain provides ...
package domain

//User ....
type User struct {
	ID       NUUID
	Name     string
	Email    string
	Password string
	Phone    PhoneNumber
}

type PhoneNumber string

//UserRepository ....
type UserRepository interface {
	FindByEmail(e string) (User, error)
	Create(u User) error
	GetByID(ID NUUID) (User, error)
}
