// Package domain provides ...
package domain

import "errors"

//User ....
type User struct {
	ID       NUUID
	Name     string
	Email    Email
	Password string
	Phone    PhoneNumber
}

func (u User) Check() error {
	if len(u.Name) == 0 {
		return errors.New("请输入用户名")
	}
	if len(u.Password) == 0 {
		return errors.New("请输入密码")
	}
	if err := u.Email.Check(); err != nil {
		return err
	}
	return u.Phone.Check()
}

func (e Email) Check() error {
	if len(e) == 0 {
		return errors.New("请输入邮箱")
	}
	return nil
}

// PhoneNumber 类型
// Derived: PhoneNumber, TelPhoneNumber, MobilePhoneNumber
type PhoneNumber string

//Email
type Email string

//UserRepository ....
type UserRepository interface {
	FindByPhone(phone PhoneNumber) (User, error)
	Create(u User) error
	GetByID(ID NUUID) (User, error)
	GetByEmail(e Email) (User, error)
}
