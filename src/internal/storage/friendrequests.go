package storage

import (
	"context"
	"database/sql"
	"log"
)

type FriendRequest struct {
	Id             int    `json:"id"`
	SenderUserID   int    `json:"sender_user_id"`
	ReceiverUserID int    `json:"receiver_user_id"`
	Status         string `json:"status"`
	RequestDate    string `json:"request_date"`
	FriendsSince   string `json:"friends_since"`
	CreatedAt      string `json:"created_at"`
}

type FriendRequestStorage struct {
	db *sql.DB
}

// TODO Restriction to avoid duplicated requests
func (f *FriendRequestStorage) Create(ctx context.Context, fr *FriendRequest) error {
	query := `INSERT INTO friend_requests(sender_user_id,receiver_user_id,status,request_date) 
						VALUES($1,$2,$3,NOW()) RETURNING id`
	err := f.db.QueryRowContext(ctx, query, fr.SenderUserID, fr.ReceiverUserID, "pending").Scan(&fr.Id)
	if err != nil {
		return err
	}
	log.Printf("Friend request created, id: %v", fr.Id)
	return nil
}

func (f *FriendRequestStorage) GetByID(ctx context.Context, id string) (*FriendRequest, error) {
	var friends_since *string
	var fr FriendRequest
	query := `SELECT * FROM friend_requests WHERE ID=$1`
	row := f.db.QueryRowContext(ctx, query, id)
	err := row.Scan(&fr.Id, &fr.SenderUserID, &fr.ReceiverUserID, &fr.Status, &fr.RequestDate, &friends_since)
	if friends_since != nil {
		fr.FriendsSince = *friends_since
	}
	if err != nil {
		return nil, err
	}
	return &fr, nil
}

func (f *FriendRequestStorage) GetByUser(ctx context.Context, userId string) (*[]FriendRequest, error) {
	var frs []FriendRequest
	query := `SELECT * FROM friend_requests WHERE receiver_user_id=$1`
	rows, err := f.db.QueryContext(ctx, query, userId)
	if err != nil {
		return nil, err
	} else {
		fr := FriendRequest{}
		var friends_since *string
		for rows.Next() {
			err := rows.Scan(&fr.Id, &fr.SenderUserID, &fr.ReceiverUserID, &fr.Status, &fr.RequestDate, &friends_since)
			if err != nil {
				log.Fatal(err)
			}
			if friends_since != nil {
				fr.FriendsSince = *friends_since
			}
			frs = append(frs, fr)
		}
	}
	log.Println(frs)
	return &frs, nil
}

func (f *FriendRequestStorage) GetSentByUser(ctx context.Context, userId string) (*[]FriendRequest, error) {
	var frs []FriendRequest
	query := `SELECT * FROM friend_requests WHERE sender_user_id=$1`
	rows, err := f.db.QueryContext(ctx, query, userId)
	if err != nil {
		return nil, err
	} else {
		fr := FriendRequest{}
		var friends_since *string
		for rows.Next() {
			err := rows.Scan(&fr.Id, &fr.SenderUserID, &fr.ReceiverUserID, &fr.Status, &fr.RequestDate, &friends_since)
			if err != nil {
				log.Fatal(err)
			}
			if friends_since != nil {
				fr.FriendsSince = *friends_since
			}
			frs = append(frs, fr)
		}
	}
	log.Println(frs)
	return &frs, nil
}

func (f *FriendRequestStorage) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM friend_requests WHERE id=$1`
	result, err := f.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	row, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if row != 0 {
		log.Printf("Friend request id deleted succesfully, rows affected: %d", row)
	} else if row == 0 {
		log.Println("Invalid id")
	}
	return nil
}

func (f *FriendRequestStorage) Accept(ctx context.Context, id string) error {
	query := `UPDATE friend_requests set status='accepted' WHERE id=$1`
	result, err := f.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	row, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if row != 0 {
		log.Printf("Friend request id accepted succesfully, rows affected: %d", row)
	} else if row == 0 {
		log.Println("Invalid id")
	}
	return nil
}
