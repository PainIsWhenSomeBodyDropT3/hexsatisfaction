package pg

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/JesusG2000/hexsatisfaction/internal/config"
	_ "github.com/lib/pq"
)

// NewPg creates new  connection to pg database.
func NewPg(pgConfig config.PgConfig) (*sql.DB, error) {
	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%d", pgConfig.Host, pgConfig.User, pgConfig.Name, pgConfig.Password, pgConfig.Port)
	log.Println(dbURI)

	db, err := sql.Open(pgConfig.Dialect, dbURI)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	pgConfig.URI = dbURI

	return db, nil
}
