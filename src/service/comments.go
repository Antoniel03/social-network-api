package service

import (
	"context"
	"log"

	"github.com/Antoniel03/social-network-api/internal/storage"
)

func (u *Service) PublishComment(comment *storage.Comment, ctx context.Context) error {
	log.Println("Comment service reached")
	err := u.Repository.Comments.Create(ctx, comment)
	if err != nil {
		return err
	}
	return nil
}

func (u *Service) GetComment(id string, ctx context.Context) (*storage.Comment, error) {
	log.Println("Comment service reached")
	comment, err := u.Repository.Comments.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return comment, nil
}

func (u *Service) GetCommentsByPost(id string, ctx context.Context) (*[]storage.Comment, error) {
	log.Println("Comment service reached")
	posts, err := u.Repository.Comments.GetByPost(ctx, id)
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (u *Service) DeleteComment(id string, ctx context.Context) error {
	log.Println("Comment service reached")
	err := u.Repository.Comments.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
