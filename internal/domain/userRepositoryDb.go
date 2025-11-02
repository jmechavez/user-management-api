package domain

import (
	"database/sql"
	"log"
	"time"

	apperrors "github.com/jmechavez/user-management-api/internal/appErrors"
	_ "github.com/lib/pq" // PostgreSQL driver
)

type UserRepositoryDb struct {
	db *sql.DB
}

func (r UserRepositoryDb) FindByID(id int64) (*User, *apperrors.AppError) {
	query := `SELECT id_number, first_name, last_name, email FROM users WHERE id_number=$1`

	rows := r.db.QueryRow(query, id)
	var u User
	err := rows.Scan(&u.IdNumber, &u.FirstName, &u.LastName, &u.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, apperrors.NewNotFoundError("User not found")
		} else {
			log.Println("Error while scanning user:", err)
			return nil, apperrors.NewUnexpectedError("Unexpected database error")
		}
	}
	return &u, nil
}

func (r UserRepositoryDb) FindAllv2() ([]User, *apperrors.AppError) {
	query := `SELECT id_number, email FROM users`

	rows, err := r.db.Query(query)
	if err != nil {
		log.Println("Error executing query:", err)
		return nil, apperrors.NewNotFoundError("Unexpected database error")
	}
	defer rows.Close()

	users := make([]User, 0)
	for rows.Next() {
		var u User
		err := rows.Scan(&u.IdNumber, &u.Email)
		if err != nil {
			log.Println("Error while scanning user:", err)
			return nil, apperrors.NewUnexpectedError("Unexpected database error")
		}
		users = append(users, u)
	}

	if err = rows.Err(); err != nil {
		log.Println("Row iteration error:", err)
		return nil, apperrors.NewUnexpectedError("Unexpected database error")
	}

	return users, nil
}

func (r UserRepositoryDb) FindAll() ([]User, *apperrors.AppError) {
	query := `SELECT id_number, first_name, last_name, email FROM users`

	rows, err := r.db.Query(query)
	if err != nil {
		log.Println("Error executing query:", err)
		return nil, apperrors.NewNotFoundError("Unexpected database error")
	}
	defer rows.Close()

	users := make([]User, 0)
	for rows.Next() {
		var u User
		err := rows.Scan(&u.IdNumber, &u.FirstName, &u.LastName, &u.Email)
		if err != nil {
			log.Println("Error while scanning user:", err)
			return nil, apperrors.NewUnexpectedError("Unexpected database error")
		}
		users = append(users, u)
	}

	if err = rows.Err(); err != nil {
		return nil, apperrors.NewUnexpectedError("Unexpected database error")
	}

	return users, nil
}

func NewUserRepositoryDb() UserRepositoryDb {
	connStr := "user=admin password=admin123 dbname=mydb host=localhost port=5432 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return UserRepositoryDb{db}
}
