// Package usecases provides ...
package usecases

import "iohttps.com/live/realworld-medium-rewrite/domain"

type UserInteractor struct {
	UserRepo domain.UserRepository
}

func (itor UserInteractor) Register()
