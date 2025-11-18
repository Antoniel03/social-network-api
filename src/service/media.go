package service

import (
	"context"
	"github.com/Antoniel03/social-network-api/internal/storage"
	"log"
)

func (u *Service) SaveMedia(media *storage.Media, ctx context.Context) error {
	log.Println("Media service reached")
	err := u.Repository.Media.Create(ctx, media)
	if err != nil {
		return err
	}
	return nil
}
