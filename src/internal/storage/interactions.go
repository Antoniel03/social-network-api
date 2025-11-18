package storage

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"strconv"
)

type Interaction struct {
	Id              int    `json:"id"`
	UserId          int    `json:"user_id"`
	PostId          int    `json:"post_id"`
	InteractedAt    string `json:"interacted_at"`
	InteractionType string `json:"interaction_type"`
}

type InteractionStorage struct {
	db *sql.DB
}

// Create It gets the interaction info and saves it in the database
func (i *InteractionStorage) Create(ctx context.Context, interaction *Interaction) error {
	query := `INSERT INTO interactions(user_id,post_id,interacted_at,interaction_type) 
						VALUES($1,$2,NOW(), $3) RETURNING id`
	err := i.db.QueryRowContext(ctx, query, interaction.UserId, interaction.PostId, interaction.InteractionType).Scan(&interaction.Id)
	if err != nil {
		return err
	}

	postId := strconv.Itoa(interaction.PostId)
	if interaction.InteractionType == "like" {
		err = i.increaseLikeCounter(ctx, postId)
		if err != nil {
			return err
		}
	} else if interaction.InteractionType == "dislike" {
		err = i.increaseDislikeCounter(ctx, postId)
		if err != nil {
			return err
		}
	}
	log.Printf("Interaction created, id: %v", interaction.Id)
	return nil
}

func (i *InteractionStorage) Delete(ctx context.Context, id string, postId string, interactionType string) error {
	query := `DELETE FROM interactions USING posts WHERE interactions.post_id=posts.id  AND interactions.id=$1 AND posts.id=$2`
	result, err := i.db.ExecContext(ctx, query, id, postId)
	if err != nil {
		return err
	}
	row, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if row == 0 {
		return errors.New("Invalid id")
	}

	if row != 0 {
		if interactionType == "like" {
			err = i.decreaseLikeCounter(ctx, postId)
			if err != nil {
				return err
			}
		} else if interactionType == "dislike" {
			err = i.decreaseDislikeCounter(ctx, postId)
			if err != nil {
				return err
			}
		}
		log.Printf("Interaction id deleted succesfully, rows affected: %d", row)
	}
	return nil
}

func (i *InteractionStorage) GetByID(ctx context.Context, id string) (*Interaction, error) {
	var interaction Interaction
	query := `SELECT * FROM interactions WHERE ID=$1`
	row := i.db.QueryRowContext(ctx, query, id)
	err := row.Scan(&interaction.Id, &interaction.UserId, &interaction.PostId, &interaction.InteractedAt, &interaction.InteractionType)
	if err != nil {
		return nil, err
	}
	return &interaction, nil
}

func (i *InteractionStorage) GetByUser(ctx context.Context, userId string) (*[]Interaction, error) {
	var interactions []Interaction
	query := `SELECT * FROM interactions WHERE user_id=$1`
	rows, err := i.db.QueryContext(ctx, query, userId)
	if err != nil {
		return nil, err
	} else {
		interaction := Interaction{}
		for rows.Next() {
			err := rows.Scan(&interaction.Id, &interaction.UserId, &interaction.PostId, &interaction.InteractedAt, &interaction.InteractionType)
			if err != nil {
				log.Fatal(err)
			}
			interactions = append(interactions, interaction)
		}
	}
	log.Println(interactions)
	return &interactions, nil
}

func (i *InteractionStorage) GetByPostAndUser(ctx context.Context, userId string, postId string) (*Interaction, error) {
	var interaction Interaction
	query := `SELECT * FROM interactions WHERE user_id=$1 AND post_id=$2`
	row := i.db.QueryRowContext(ctx, query, userId, postId)
	err := row.Scan(&interaction.Id, &interaction.UserId, &interaction.PostId, &interaction.InteractedAt, &interaction.InteractionType)
	if err != nil {
		return nil, err
	}
	return &interaction, nil
}

func (i *InteractionStorage) increaseDislikeCounter(ctx context.Context, id string) error {
	query := `UPDATE posts SET dislikes=dislikes+1 WHERE id=$1`
	result, err := i.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	log.Printf("Updated dislikes, rows affected: %d", rows)
	return nil
}

func (i *InteractionStorage) increaseLikeCounter(ctx context.Context, id string) error {
	query := `UPDATE posts SET likes=likes+1 WHERE id=$1`
	result, err := i.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	log.Printf("Updated likes, rows affected: %d", rows)
	return nil
}

func (i *InteractionStorage) decreaseLikeCounter(ctx context.Context, id string) error {
	query := `UPDATE posts SET likes=likes-1 WHERE id=$1`
	result, err := i.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	} else if rows == 0 {
		return errors.New("Invalid id")
	}

	log.Printf("Updated likes, rows affected: %d", rows)
	return nil
}

func (i *InteractionStorage) decreaseDislikeCounter(ctx context.Context, id string) error {
	query := `UPDATE posts SET dislikes=dislikes-1 WHERE id=$1`
	result, err := i.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	} else if rows == 0 {
		return errors.New("Invalid id")
	}

	log.Printf("Updated dislikes, rows affected: %d", rows)
	return nil
}
