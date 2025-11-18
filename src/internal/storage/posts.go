package storage

import (
	"context"
	"database/sql"
	"log"
)

type Post struct {
	Id          int    `json:"id"`
	UserId      int    `json:"user_id"`
	CreatedAt   string `json:"created_at"`
	TextContent string `json:"text_content"`
	Likes       int    `json:"likes"`
	Dislikes    int    `json:"dislikes"`
}

type PostStorage struct {
	db *sql.DB
}

func (p *PostStorage) Create(ctx context.Context, post *Post) error {
	query := `INSERT INTO posts(user_id,text_content,likes,dislikes,created_at) 
						VALUES($1,$2,$3,$4,NOW()) RETURNING id`
	err := p.db.QueryRowContext(ctx, query, post.UserId, post.TextContent, 0, 0).Scan(&post.Id)
	if err != nil {
		return err
	}
	log.Printf("Post created, id: %v", post.Id)
	return nil
}

func (p *PostStorage) GetByUser(ctx context.Context, userId string) (*[]Post, error) {
	var posts []Post
	query := `SELECT * FROM posts WHERE user_id=$1`
	rows, err := p.db.QueryContext(ctx, query, userId)
	if err != nil {
		return nil, err
	} else {
		post := Post{}
		for rows.Next() {
			err := rows.Scan(&post.Id, &post.UserId, &post.CreatedAt, &post.TextContent, &post.Likes, &post.Dislikes)
			if err != nil {
				log.Fatal(err)
			}
			posts = append(posts, post)
		}
	}
	log.Println(posts)
	return &posts, nil
}

func (p *PostStorage) GetByID(ctx context.Context, id string) (*Post, error) {
	query := `SELECT * FROM posts WHERE id=$1`
	row := p.db.QueryRowContext(ctx, query, id)
	post := Post{}
	err := row.Scan(&post.Id, &post.UserId, &post.CreatedAt, &post.TextContent, &post.Likes, &post.Dislikes)
	if err != nil {
		return nil, err
	}
	log.Println(post)
	return &post, nil
}

func (p *PostStorage) LikePost(ctx context.Context, id string) error {
	query := `UPDATE posts SET likes=likes+1 WHERE id=$1`
	result, err := p.db.ExecContext(ctx, query, id)
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

func (p *PostStorage) DislikePost(ctx context.Context, id string) error {
	query := `UPDATE posts SET dislikes=dislikes+1 WHERE id=$1`
	result, err := p.db.ExecContext(ctx, query, id)
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

func (p *PostStorage) RemoveLike(ctx context.Context, id string) error {
	query := `UPDATE posts SET likes=likes-1 WHERE id=$1`
	result, err := p.db.ExecContext(ctx, query, id)
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

func (p *PostStorage) RemoveDislike(ctx context.Context, id string) error {
	query := `UPDATE posts SET dislikes=dislikes-1 WHERE id=$1`
	result, err := p.db.ExecContext(ctx, query, id)
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

func (p *PostStorage) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM posts WHERE id=$1`
	result, err := p.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	row, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if row != 0 {
		log.Printf("Post %s deleted succesfully, rows affected: %d", id, row)
	} else if row == 0 {
		log.Println("Invalid id")
	}
	return nil
}
