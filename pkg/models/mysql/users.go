package mysql

import (
	"database/sql"
	"time"
)

type User struct {
	ID             int
	Name           string
	Email          string
	HashedPassword []byte
	Created        time.Time
}

type UserModel struct {
	DB *sql.DB
}

// We'll use the Insert method to add a new record to the "users" table.
func (dbm *DBModel) Insert(name, email, password string) error {
	return nil
}

// We'll use the Authenticate method to verify whether a user exists with
// the provided email address and password. This will return the relevant
// user ID if they do.
func (dbm *DBModel) Authenticate(email, password string) (int, error) {
	return 0, nil
}

// We'll use the Exists method to check if a user exists with a specific ID.
func (dbm *DBModel) Exists(id int) (bool, error) {
	return false, nil
}
