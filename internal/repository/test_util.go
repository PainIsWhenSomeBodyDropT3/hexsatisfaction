package repository

import (
	"database/sql"

	"github.com/JesusG2000/hexsatisfaction/internal/config"
	"github.com/JesusG2000/hexsatisfaction/pkg/database/pg"
	"github.com/pkg/errors"
)

const configPath = "config/main"

// Connect2Repositories connects to a pg database.
func Connect2Repositories() (*sql.DB, *Repositories, error) {

	cfg, err := config.Init(configPath)
	if err != nil {
		return nil, nil, errors.Wrap(err, "couldn't init config")
	}

	db, err := pg.NewPg(cfg.Pg)
	if err != nil {
		return nil, nil, errors.Wrap(err, "couldn't create pg model")
	}

	repos := NewRepositories(db)

	return db, repos, nil
}
