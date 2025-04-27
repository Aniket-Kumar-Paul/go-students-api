package sqlite

import (
	"database/sql"

	"github.com/Aniket-Kumar-Paul/go-students-api/internal/config"
	_ "github.com/mattn/go-sqlite3" // _ => indirect usage
)

type Sqlite struct {
	Db *sql.DB
}

func New(cfg *config.Config) (*Sqlite, error) {
	db, err := sql.Open("sqlite3", cfg.StoragePath)
	if err!= nil {
		return nil, err
	}

	_, errr := db.Exec(`CREATE TABLE IF NOT EXISTS students (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		email TEXT,
		age INTEGER
	)`)
	if errr!=nil {
		return nil, errr
	}

	return &Sqlite{
		Db: db,
	}, nil
}