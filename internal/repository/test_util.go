package repository

import (
	"database/sql"
	"log"

	"github.com/JesusG2000/hexsatisfaction/internal/config"
	"github.com/JesusG2000/hexsatisfaction/pkg/database/pg"
)

const configPath = "config/main"

func Connect2Repositories() (*sql.DB, *Repositories, error) {

	cfg, err := config.Init(configPath)
	if err != nil {
		log.Fatal("Init config error", err)
	}

	db, err := pg.NewPg(cfg.Pg)
	if err != nil {
		return nil, nil, err
	}

	repos := NewRepositories(db)

	return db, repos, nil
}
