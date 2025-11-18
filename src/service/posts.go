package service

import (
	"context"
	"github.com/Antoniel03/social-network-api/internal/storage"
	"log"
	"strconv"
)

func (u *Service) PublishPost(post *storage.Post, ctx context.Context) error {
	log.Println("Post service reached")
	err := u.Repository.Posts.Create(ctx, post)
	if err != nil {
		return err
	}
	return nil
}

func (u *Service) GetPost(id string, ctx context.Context) (*storage.Post, error) {
	log.Println("Post service reached")
	post, err := u.Repository.Posts.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (u *Service) GetPostsDataByUser(id string, ctx context.Context) (*[]storage.Post, error) {
	log.Println("Post service reached")
	posts, err := u.Repository.Posts.GetByUser(ctx, id)
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (u *Service) DeletePost(id string, ctx context.Context) error {
	log.Println("Post service reached")
	err := u.Repository.Posts.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (u *Service) AddLike(userId string, postId string, ctx context.Context) error {
	log.Println("Post service reached")

	intPostId, err := strconv.Atoi(postId)
	if err != nil {
		return err
	}

	intUserId, err := strconv.Atoi(postId)
	if err != nil {
		return err
	}

	interaction, err := u.Repository.Interactions.GetByPostAndUser(ctx, userId, postId)
	if err != nil || interaction == nil {

		interaction := storage.Interaction{UserId: intUserId, PostId: intPostId, InteractionType: "like"}
		err = u.Repository.Interactions.Create(ctx, &interaction)
		if err != nil {
			return err
		}

		log.Println("post liked")
	} else {

		id := strconv.Itoa(interaction.Id)
		err = u.Repository.Interactions.Delete(ctx, id, postId, "like")
		if err != nil {
			return err
		}

		log.Println("post like removed")
	}

	return nil
}

func (u *Service) AddDislike(userId string, postId string, ctx context.Context) error {
	log.Printf("Post service reached, with %s %s", userId, postId)

	intPostId, err := strconv.Atoi(postId)
	if err != nil {
		return err
	}

	intUserId, err := strconv.Atoi(postId)
	if err != nil {
		return err
	}

	interaction, err := u.Repository.Interactions.GetByPostAndUser(ctx, userId, postId)
	if err != nil || interaction == nil {

		interaction := storage.Interaction{UserId: intUserId, PostId: intPostId, InteractionType: "dislike"}
		err = u.Repository.Interactions.Create(ctx, &interaction)
		if err != nil {
			return err
		}

		log.Println("post disliked")
	} else {

		id := strconv.Itoa(interaction.Id)
		err = u.Repository.Interactions.Delete(ctx, id, postId, "dislike")
		if err != nil {
			return err
		}

		log.Println("post dislike removed")
	}

	return nil
}
