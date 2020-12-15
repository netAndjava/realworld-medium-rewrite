// Package interfaces provides ...
package interfaces

import (
	"iohttps.com/live/realworld-medium-rewrite/domain"
	"iohttps.com/live/realworld-medium-rewrite/infrastructure/database"
)

//UserRepo .....
type UserRepo database.DbRepo

func NewUserRepo(helper database.DbHandler) domain.UserRepository {
	return &UserRepo{Handler: helper}
}

func (repo *UserRepo) Create(u domain.User) error {
	_, err := repo.Handler.Execute(`insert into t_user (id,name,password,email,phone) values(?,?,?,?,?)`, u.ID, u.Name, u.Password, u.Email, u.Phone)
	return err
}

func (repo *UserRepo) FindByPhone(phone domain.PhoneNumber) (domain.User, error) {
	var user domain.User
	user.Phone = phone
	row := repo.Handler.QueryRow(`select id from t_user where phone=?`, string(phone))
	err := row.Scan(&user.ID)
	// TODO: 从err中判断是否是不存在数据错误 <15-12-20, nqq> //
	return user, err
}

func (repo *UserRepo) GetByEmail(e domain.Email) (domain.User, error) {
	row := repo.Handler.QueryRow(`select id,password where email=?`, string(e))
	var user domain.User
	user.Email = e
	err := row.Scan(&user.ID, &user.Password)
	// TODO: 从err中判断是否是不存在数据错误 <15-12-20, nqq> //
	return user, err
}
