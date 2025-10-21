package domain

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq" // PostgreSQL driver
)

type UserRepositoryDb struct {
	db *sql.DB
}

func (r UserRepositoryDb) FindAll() ([]User, error) {
	query := `SELECT id_number, first_name, last_name, email FROM users`

	rows, err := r.db.Query(query)
	if err != nil {
		log.Println("Error executing query:", err)
		return nil, err
	}
	defer rows.Close()

	users := make([]User, 0)
	for rows.Next() {
		var u User
		err := rows.Scan(&u.IdNumber, &u.FirstName, &u.LastName, &u.Email)
		if err != nil {
			log.Println("Error while scanning user:", err)
			return nil, err
		}
		users = append(users, u)
	}

	if err = rows.Err(); err != nil {
		return nil, err
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
