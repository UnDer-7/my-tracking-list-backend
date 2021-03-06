package service

import (
	"context"
	"my-tracking-list-backend/core/app_error"
	"my-tracking-list-backend/core/domain"
	"my-tracking-list-backend/core/ports/driven"
	"my-tracking-list-backend/core/ports/driver"
)

type userService struct {
	repository driven.UserRepository
}

func NewUserService(repository driven.UserRepository) driver.UserService {
	return &userService{
		repository: repository,
	}
}

func (service userService) SaveUser(ctx context.Context, user domain.User) (domain.User, error) {
	if !user.ID.IsZero() {
		return domain.User{}, app_error.ThrowBusinessError(
			"Um error ocorreu, tente novamente",
			"Nao pode criar usuario com Id",
		)
	}

	if existes, err := service.UserExists(ctx, user.Email); err != nil {
		return domain.User{}, err
	} else if existes {
		return domain.User{}, app_error.ThrowBusinessError(
			"Usuario ja cadastrado",
			"Usuario ja cadastrado",
		)
	}

	// todo: validar campos do user para n deixar inserir com vazios/nils
	userSave, err := service.repository.Persist(ctx, user)
	if err != nil {
		return domain.User{}, err
	}
	return userSave, nil
}

func (service userService) FindByEmail(ctx context.Context, email string) (domain.User, error) {
	// todo: validar email
	userFound, err := service.repository.GetByEmail(ctx, email)
	if err != nil {
		return domain.User{}, err
	}
	return userFound, nil
}

func (service userService) UserExists(ctx context.Context, email string) (bool, error) {
	return service.repository.ExistesByEmail(ctx, email)
}
