package storage

import (
	"context"
	"database/sql"
	"errors"
	"log"
)

// TODO add name string attr string attr
type Media struct {
	Id         int    `json:"id"`
	UserId     int    `json:"user_id"`
	PostId     int    `json:"post_id"`
	Location   string `json:"location"`
	FileSize   int    `json:"file_size"`
	UploadedAt string `json:"uploaded_at"`
}

type MediaStorage struct {
	db *sql.DB
}

func (m *MediaStorage) Create(ctx context.Context, media *Media) error {
	query := `INSERT INTO media(user_id,post_id,file_size,location,uploaded_at) VALUES($1,$2,$3,$4,NOW()) RETURNING id`
	err := m.db.QueryRowContext(ctx, query, media.UserId, media.PostId, media.FileSize, media.Location).Scan(&media.Id)
	if err != nil {
		return err
	}
	log.Printf("Media uploaded, id: %v", media.Id)
	return nil
}

func (m *MediaStorage) GetByID(ctx context.Context, id string) (*Media, error) {
	var media Media
	query := `SELECT * FROM media WHERE ID=$1`
	row := m.db.QueryRowContext(ctx, query, id)
	err := row.Scan(&media.Id, &media.UserId, &media.PostId, &media.Location, &media.FileSize, &media.UploadedAt)
	if err != nil {
		return nil, err
	}
	return &media, nil
}

func (m *MediaStorage) GetByUser(ctx context.Context, userId string) (*[]Media, error) {
	var mediaSet []Media
	query := `SELECT * FROM media WHERE user_id=$1`
	rows, err := m.db.QueryContext(ctx, query, userId)
	if err != nil {
		return nil, err
	} else {
		media := Media{}
		for rows.Next() {
			err := rows.Scan(&media.Id, &media.UserId, &media.PostId, &media.Location, &media.FileSize, &media.UploadedAt)
			if err != nil {
				log.Fatal(err)
			}
			mediaSet = append(mediaSet, media)
		}
	}
	return &mediaSet, nil
}

func (m *MediaStorage) GetByPost(ctx context.Context, postId string) (*[]Media, error) {
	var mediaSet []Media
	query := `SELECT * FROM media WHERE post_id=$1`
	rows, err := m.db.QueryContext(ctx, query, postId)
	if err != nil {
		return nil, err
	} else {
		media := Media{}
		for rows.Next() {
			err := rows.Scan(&media.Id, &media.UserId, &media.PostId, &media.Location, &media.FileSize, &media.UploadedAt)
			if err != nil {
				log.Fatal(err)
			}
			mediaSet = append(mediaSet, media)
		}
	}
	return &mediaSet, nil
}

func (m *MediaStorage) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM media WHERE id=$1`
	result, err := m.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	row, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if row != 0 {
		log.Printf("Media id deleted succesfully, rows affected: %d", row)
	} else if row == 0 {
		return errors.New("Invalid id")
	}
	return nil
}
