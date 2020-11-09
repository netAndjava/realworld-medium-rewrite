// Package domain provides ...
package domain

//User ....
type User struct {
	ID         NUUID
	Name       string
	Email      string
	Password   string
	IfVerified bool //邮箱是否校验过
}

//UserRepository ....
type UserRepository interface {
	/* TODO: add methods */
}
