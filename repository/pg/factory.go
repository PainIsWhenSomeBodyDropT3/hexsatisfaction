package pg

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/pressly/goose"
	"github.com/spf13/viper"
)

type util struct {
	dialect  string
	host     string
	user     string
	dbname   string
	password string
	dbPort   string
}

// Factory represents the pg factory.
type Factory struct {
	util
	*sql.DB
	User
}

// User represents the users table.
type User struct {
	db *sql.DB
}

func (u *util) setup() {

	dialect := viper.GetString("db.dialect")
	host := viper.GetString("db.host")
	user := viper.GetString("db.username")
	dbname := viper.GetString("db.dbname")
	dbPort := viper.GetString("db.port")
	password := os.Getenv("DB_PASSWORD")

	u.dialect = dialect
	u.host = host
	u.user = user
	u.dbname = dbname
	u.password = password
	u.dbPort = dbPort

}

// NewFactory creates new pg factory.
func NewFactory() (*Factory, error) {
	var f Factory
	f.setup()
	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", f.host, f.user, f.dbname, f.password, f.dbPort)
	log.Println(dbURI)

	db, err := sql.Open(f.dialect, dbURI)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	err = goose.Up(db, "./migrations")
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
