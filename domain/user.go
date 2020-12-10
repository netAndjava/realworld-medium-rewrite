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

// PhoneNumber 类型
// Derived: PhoneNumber, TelPhoneNumber, MobilePhoneNumber
type PhoneNumber string

//UserRepository ....
type UserRepository interface {
	FindByEmail(e string) (User, error)
	Create(u User) error
	GetByID(ID NUUID) (User, error)
}
