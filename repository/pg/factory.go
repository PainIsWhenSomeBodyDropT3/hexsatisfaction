package pg

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type util struct {
	dialect  string
	host     string
	user     string
	name     string
	password string
	dbPort   string
}

// Factory represents the pg factory .
type Factory struct {
	util
	*sql.DB
	User
}

// User represents the User table.
type User struct {
	db *sql.DB
}

func (u *util) setup() {

	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found\n")
	}

	dialect, ok := os.LookupEnv("DIALECT")
	if !ok {
		dialect = "postgres"
	}

	host, ok := os.LookupEnv("HOST")
	if !ok {
		host = "localhost"
	}

	user, ok := os.LookupEnv("DB_USER")
	if !ok {
		user = "postgres"
	}

	name, ok := os.LookupEnv("NAME")
	if !ok {
		name = "postgres"
	}

	password, ok := os.LookupEnv("PASSWORD")
	if !ok {
		password = "18051965q"
	}

	dbPort, ok := os.LookupEnv("DB_PORT")
	if !ok {
		dbPort = "5432"
	}

	u.dialect = dialect
	u.host = host
	u.user = user
	u.name = name
	u.password = password
	u.dbPort = dbPort

}

// NewFactory creates new pg factory
func NewFactory() (*Factory, error) {
	var f Factory
	f.setup()
	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", f.host, f.user, f.name, f.password, f.dbPort)
	log.Println(dbURI)

	db, err := sql.Open(f.dialect, dbURI)
	if err != nil {
		return nil, err
	}
	f.DB = db
	return &f, nil
}

// NewUserRepository creates new User repository.
func (f *Factory) NewUserRepository() *User {
	return &User{f.DB}
}
