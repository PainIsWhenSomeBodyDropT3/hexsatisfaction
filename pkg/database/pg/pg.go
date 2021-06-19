package pg

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/JesusG2000/hexsatisfaction/internal/config"
	// pg driver
	_ "github.com/lib/pq"
)

// NewPg creates new connection to pg database.
func NewPg(pgConfig config.PgConfig) (*sql.DB, error) {
	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%d", pgConfig.Host, pgConfig.User, pgConfig.DatabaseName, pgConfig.Password, pgConfig.Port)
	log.Println(dbURI)

	db, err := sql.Open(pgConfig.DatabaseDialect, dbURI)
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
