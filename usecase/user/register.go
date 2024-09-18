package user

import (
	"context"
)

type RegisterUserIn struct {
	Email    string
	Password string
}

type RegisterUserOut struct {
	ID string
}

func (u *UseCase) RegisterUser(ctx context.Context, in RegisterUserIn) (RegisterUserOut, error) {
	return RegisterUserOut{}, nil
}
