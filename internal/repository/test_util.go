package repository

import (
	"database/sql"

	"github.com/JesusG2000/hexsatisfaction/internal/config"
	"github.com/JesusG2000/hexsatisfaction/pkg/database/pg"
	"github.com/pkg/errors"
)

// Connect2Repositories connects to a pg database.
func Connect2Repositories() (*sql.DB, *Repositories, error) {

	cfg, err := config.Init()
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
