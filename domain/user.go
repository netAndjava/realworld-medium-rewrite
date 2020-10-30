// Package domain provides ...
package domain

//User
type User struct {
	ID       NUUID
	Name     string
	Email    string
	Password string
}
