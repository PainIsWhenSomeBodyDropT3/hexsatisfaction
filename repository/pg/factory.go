package pg

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

type dbParam struct {
	dialect  string
	host     string
	user     string
	name     string
	password string
	port     string
}

// Repository represents the pg repository.
type Repository struct {
	dbParam
	*sql.DB
	User
}

// User represents the users table.
type User struct {
	db *sql.DB
}

// NewUserRepository creates new User repository.
func (f *Repository) NewUserRepository() *User {
	return &User{f.DB}
}

// NewPgRepository creates new pg repository.
func NewPgRepository() (*Repository, error) {
	var f Repository
	err := initConfig()
	if err != nil {
		log.Fatalf("errors initializing configs: %s", err.Error())
	}

	f.setup()

	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", f.host, f.user, f.name, f.password, f.port)
	log.Println(dbURI)

	db, err := sql.Open(f.dialect, dbURI)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	f.DB = db
	return &f, nil
}

func (u *dbParam) setup() {

	dialect := viper.GetString("db.dialect")
	host := viper.GetString("db.host")
	user := viper.GetString("db.username")
	name := viper.GetString("db.name")
	port := viper.GetString("db.port")
	password := os.Getenv("DB_PASSWORD")

	u.dialect = dialect
	u.host = host
	u.user = user
	u.name = name
	u.password = password
	u.port = port

}

func initConfig() error {

	env := ".env"
	path, err := os.Getwd()
	if err != nil {
		return err
	}
	path = strings.SplitAfter(path, "hexsatisfaction")[0]
	if err := godotenv.Load(path + "/" + env); err != nil {
		log.Fatal("Error loading .env file")
	}

	viper.AddConfigPath(path)
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
