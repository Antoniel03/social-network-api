package service

import (
	"context"
	"github.com/Antoniel03/social-network-api/internal/storage"
	"log"
)

func (u *Service) SendFriendRequest(fr *storage.FriendRequest, ctx context.Context) error {
	log.Println("Friend request service reached")
	err := u.Repository.FriendRequests.Create(ctx, fr)
	if err != nil {
		return err
	}
	return nil
}

func (u *Service) GetFriendRequest(id string, ctx context.Context) (*storage.FriendRequest, error) {
	log.Println("Friend requests service reached")
	post, err := u.Repository.FriendRequests.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (u *Service) GetUserFriendRequests(id string, ctx context.Context) (*[]storage.FriendRequest, error) {
	log.Println("Friend requests service reached")
	frs, err := u.Repository.FriendRequests.GetByUser(ctx, id)
	if err != nil {
		return nil, err
	}
	return frs, nil
}

func (u *Service) GetSentFriendRequests(id string, ctx context.Context) (*[]storage.FriendRequest, error) {
	log.Println("Friend requests service reached")
	frs, err := u.Repository.FriendRequests.GetSentByUser(ctx, id)
	if err != nil {
		return nil, err
	}
	return frs, nil
}

func (u *Service) DeleteFriendRequest(id string, ctx context.Context) error {
	log.Println("Friend request service reached")
	err := u.Repository.FriendRequests.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (u *Service) AcceptFriendRequest(id string, ctx context.Context) error {
	log.Println("Friend request service reached")
	err := u.Repository.FriendRequests.Accept(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
