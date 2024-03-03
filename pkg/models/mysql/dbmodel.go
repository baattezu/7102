package mysql

import "database/sql"

// DBModel represents the model for working with the MySQL database.
type DBModel struct {
	DB *sql.DB
}

// InitDB opens a connection to the MySQL database.
func (dbm *DBModel) InitDB(dataSourceName string) error {
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return err
	}
	dbm.DB = db
	return nil
}
