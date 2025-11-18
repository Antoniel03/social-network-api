package storage

import (
	"context"
	"database/sql"
)

type Storage struct {
	Users interface {
		Create(ctx context.Context, user *User) error
		GetByID(ctx context.Context, id string) (*User, error)
		GetByEmail(ctx context.Context, email string) (*User, error)
		Delete(ctx context.Context, id string) error
		UpdatePassword(ctx context.Context, id string, password string) error
		UpdateEmail(ctx context.Context, id string, email string) error
		UpdateName(ctx context.Context, id string, name string) error
		UpdateLastname(ctx context.Context, id string, lastname string) error
	}

	Posts interface {
		Create(ctx context.Context, post *Post) error
		GetByUser(ctx context.Context, userId string) (*[]Post, error)
		GetByID(ctx context.Context, id string) (*Post, error)
		LikePost(ctx context.Context, id string) error
		DislikePost(ctx context.Context, id string) error
		RemoveLike(ctx context.Context, id string) error
		RemoveDislike(ctx context.Context, id string) error
		Delete(ctx context.Context, id string) error
	}

	Interactions interface {
		Create(ctx context.Context, interaction *Interaction) error
		GetByID(ctx context.Context, id string) (*Interaction, error)
		GetByPostAndUser(ctx context.Context, userId string, postId string) (*Interaction, error)
		GetByUser(ctx context.Context, userId string) (*[]Interaction, error)
		Delete(ctx context.Context, id string, postId string, interactionType string) error
	}

	Media interface {
		Create(ctx context.Context, media *Media) error
		GetByID(ctx context.Context, id string) (*Media, error)
		GetByUser(ctx context.Context, userId string) (*[]Media, error)
		GetByPost(ctx context.Context, postId string) (*[]Media, error)
		Delete(ctx context.Context, id string) error
	}

	Comments interface {
		Create(ctx context.Context, comment *Comment) error
		GetByID(ctx context.Context, id string) (*Comment, error)
		GetByPost(ctx context.Context, postId string) (*[]Comment, error)
		Delete(ctx context.Context, id string) error
	}
	FriendRequests interface {
		Create(ctx context.Context, fr *FriendRequest) error
		GetByID(ctx context.Context, id string) (*FriendRequest, error)
		GetByUser(ctx context.Context, userId string) (*[]FriendRequest, error)
		GetSentByUser(ctx context.Context, userId string) (*[]FriendRequest, error)
		Accept(ctx context.Context, id string) error
		Delete(ctx context.Context, id string) error
	}
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Users:          &UserStorage{db},
		Posts:          &PostStorage{db},
		Comments:       &CommentStorage{db},
		Media:          &MediaStorage{db},
		FriendRequests: &FriendRequestStorage{db},
		Interactions:   &InteractionStorage{db},
	}
}
