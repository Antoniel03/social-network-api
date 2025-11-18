package service

import (
	"context"
	"errors"
	"github.com/Antoniel03/social-network-api/internal/storage"
	"log"
)

func (u *Service) RegisterUser(user *storage.User, ctx context.Context) error {
	log.Println("User service reached")
	err := u.Repository.Users.Create(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func (u *Service) Login(email string, password string, ctx context.Context) (*storage.User, error) {
	user, err := u.Repository.Users.GetByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	if password != user.Password {
		return nil, errors.New("Invalid password")
	}
	return user, nil
}

func (u *Service) GetUser(id string, ctx context.Context) (*storage.User, error) {
	log.Println("User service reached")
	user, err := u.Repository.Users.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *Service) DeleteUser(id string, ctx context.Context) error {
	log.Println("User service reached")
	err := u.Repository.Users.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (u *Service) UpdateUser(attr string, data string, id string, ctx context.Context) error {
	log.Println("User service reached")
	var err error
	switch attr {
	case "name":
		err = u.Repository.Users.UpdateName(ctx, id, data)
	case "lastname":
		err = u.Repository.Users.UpdateLastname(ctx, id, data)
	case "email":
		err = u.Repository.Users.UpdateEmail(ctx, id, data)
	case "password":
		err = u.Repository.Users.UpdatePassword(ctx, id, data)
	}
	if err != nil {
		return err
	}
	return nil
}
