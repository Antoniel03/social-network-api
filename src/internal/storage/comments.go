package storage

import (
	"context"
	"database/sql"
	"errors"
	"log"
)

type Comment struct {
	Id          int    `json:"id"`
	UserId      int    `json:"user_id"`
	PostId      int    `json:"post_id"`
	TextContent string `json:"text_content"`
	CreatedAt   string `json:"created_at"`
}

type CommentStorage struct {
	db *sql.DB
}

func (c *CommentStorage) Create(ctx context.Context, comment *Comment) error {
	if comment.TextContent == "" {
		return errors.New("Empty text content.")
	}
	query := `INSERT INTO comments(user_id,post_id,text_content,created_at) VALUES($1,$2,$3,NOW()) RETURNING id`
	err := c.db.QueryRowContext(ctx, query, comment.UserId, comment.PostId, comment.TextContent).Scan(&comment.Id)
	if err != nil {
		return err
	}
	log.Printf("Comment created, id: %v", comment.Id)
	return nil
}

func (c *CommentStorage) GetByID(ctx context.Context, id string) (*Comment, error) {
	var comment Comment
	query := `SELECT * FROM comments WHERE ID=$1`
	row := c.db.QueryRowContext(ctx, query, id)
	err := row.Scan(&comment.Id, &comment.UserId, &comment.PostId, &comment.TextContent, &comment.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &comment, nil
}

func (c *CommentStorage) GetByPost(ctx context.Context, postId string) (*[]Comment, error) {
	var comments []Comment
	query := `SELECT * FROM comments WHERE post_id=$1`
	rows, err := c.db.QueryContext(ctx, query, postId)
	if err != nil {
		return nil, err
	} else {
		comment := Comment{}
		for rows.Next() {
			err := rows.Scan(&comment.Id, &comment.UserId, &comment.PostId, &comment.TextContent, &comment.CreatedAt)
			if err != nil {
				log.Fatal(err)
			}
			comments = append(comments, comment)
		}
	}
	return &comments, nil
}

func (c *CommentStorage) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM comments WHERE id=$1`
	result, err := c.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	row, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if row != 0 {
		log.Printf("Comment id deleted succesfully, rows affected: %d", row)
	} else if row == 0 {
		log.Println("Invalid id")
	}
	return nil
}
