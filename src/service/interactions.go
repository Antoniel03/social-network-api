package service

import (
	"context"
	"github.com/Antoniel03/social-network-api/internal/storage"
	"log"
)

func (u *Service) CreateInteraction(interaction *storage.Interaction, ctx context.Context) error {
	log.Println("Interaction service reached")
	err := u.Repository.Interactions.Create(ctx, interaction)
	if err != nil {
		return err
	}
	return nil
}

func (u *Service) GetInteractionData(id string, ctx context.Context) (*storage.Interaction, error) {
	log.Println("Interaction service reached")
	interaction, err := u.Repository.Interactions.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return interaction, nil
}

func (u *Service) GetInteractionDataByUser(id string, ctx context.Context) (*[]storage.Interaction, error) {
	log.Println("Interaction service reached")
	interactions, err := u.Repository.Interactions.GetByUser(ctx, id)
	if err != nil {
		return nil, err
	}
	return interactions, nil
}
