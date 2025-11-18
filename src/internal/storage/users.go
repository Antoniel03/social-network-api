package storage

import (
	"context"
	"database/sql"
	"errors"
	"log"
)

type User struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	BirthDate string `json:"birth_date"`
	CreatedAt string `json:"created_at"`
}

type UserStorage struct {
	db *sql.DB
}

func (u *UserStorage) Create(ctx context.Context, user *User) error {
	err := validateCreateUser(user)
	if err != nil {
		return err
	}
	query := `INSERT INTO users(name,lastname,email,password,birth_date,created_at) 
						VALUES($1,$2,$3,$4,CURRENT_DATE,NOW()) RETURNING id`
	err = u.db.QueryRowContext(ctx, query, user.Name, user.Lastname, user.Email, user.Password).Scan(&user.Id)
	if err != nil {
		return err
	}
	log.Printf("User created, id: %v", user.Id)
	return nil
}
func (u *UserStorage) GetByID(ctx context.Context, id string) (*User, error) {
	var user User
	query := `SELECT * FROM users WHERE ID=$1`
	row := u.db.QueryRowContext(ctx, query, id)
	err := row.Scan(&user.Id, &user.Name, &user.Lastname, &user.Email, &user.Password, &user.BirthDate, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserStorage) GetByEmail(ctx context.Context, email string) (*User, error) {
	var user User
	query := `SELECT * FROM users WHERE email=$1`
	row := u.db.QueryRowContext(ctx, query, email)
	err := row.Scan(&user.Id, &user.Name, &user.Lastname, &user.Email, &user.Password, &user.BirthDate, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserStorage) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM users WHERE id=$1`
	result, err := u.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	row, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if row != 0 {
		log.Printf("User id deleted succesfully, rows affected: %d", row)
	} else if row == 0 {
		log.Println("Invalid id")
	}
	return nil
}

func (u *UserStorage) UpdateName(ctx context.Context, id string, name string) error {
	query := `UPDATE users SET name=$1 WHERE id=$2`
	result, err := u.db.ExecContext(ctx, query, name, id)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	log.Printf("Name updated, rows affected: %d", rows)
	return nil
}

func (u *UserStorage) UpdateLastname(ctx context.Context, id string, lastname string) error {
	query := `UPDATE users SET lastname=$1 WHERE id=$2`
	result, err := u.db.ExecContext(ctx, query, lastname, id)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	log.Printf("Lastname updated, rows affected: %d", rows)
	return nil
}

func (u *UserStorage) UpdateEmail(ctx context.Context, id string, email string) error {
	query := `UPDATE users SET email=$1 WHERE id=$2`
	result, err := u.db.ExecContext(ctx, query, email, id)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	log.Printf("Email updated, rows affected: %d", rows)
	return nil
}

func (u *UserStorage) UpdatePassword(ctx context.Context, id string, password string) error {
	query := `UPDATE users SET password=$1 WHERE id=$2`
	result, err := u.db.ExecContext(ctx, query, password, id)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	log.Printf("Password updated, rows affected: %d", rows)
	return nil
}

func validateCreateUser(user *User) error {
	err := errors.New("Invalid parameter for create function")

	if user.Name == "" {
		return err
	} else if user.Lastname == "" {
		return err
	} else if user.Email == "" {
		return err
	} else if user.Password == "" {
		return err
	}
	return nil
}
