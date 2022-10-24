package data

import (
	"context"
	"database/sql"
	"log"
	"time"
)

const dbTimeOut = time.Second * 3

var db *sql.DB

func New(dbPool *sql.DB) Models {
	db = dbPool

	return Models{
		User: User{},
	}
}

type Models struct {
	User User
}

type User struct {
	ID        int       `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt`
	UpdatedAt time.Time `json:"updatedAt`
}

func (u *User) GetAll() ([]*User, error) {

	ctx, cancel := context.WithTimeout(context.Background(), dbTimeOut)
	defer cancel()

	query := `select id, email, password, createdAt, updatedAt`

	rows, err := db.QueryContext(ctx, query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []*User

	for rows.Next() {
		var user User
		err := rows.Scan(
			&user.ID,
			&user.Email,
			&user.Password,
			&user.CreatedAt,
			&user.UpdatedAt,
		)

		if err != nil {
			log.Println("Error scanning", err)
			return nil, err
		}

		users = append(users, &user)
	}

	return users, nil

}

func (u *User) InsertOne(user User) (int, error) {

	ctx, cancel := context.WithTimeout(context.Background(), dbTimeOut)
	defer cancel()

	insertCommand := `insert into users ( email, password, createdAt, updatedAt) values ($1,$2,$3,$4) returning id`

	var newId int

	err := db.QueryRowContext(ctx, insertCommand,
		user.Email,
		user.Password, time.Now(), time.Now()).Scan(&newId)

	if err != nil {
		return 0,err
	}

	return newId ,nil
}
