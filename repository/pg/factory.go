package pg

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/pressly/goose"
	"github.com/spf13/viper"
)

type dbParam struct {
	dialect  string
	host     string
	user     string
	dbname   string
	password string
	dbPort   string
}

// Factory represents the pg factory.
type Factory struct {
	dbParam
	*sql.DB
	User
}

// User represents the users table.
type User struct {
	db *sql.DB
}

// NewUserRepository creates new User repository.
func (f *Factory) NewUserRepository() *User {
	return &User{f.DB}
}

// NewFactory creates new pg factory.
func NewFactory() (*Factory, error) {
	var f Factory
	migrations := "migrations"
	path, err := initConfig()
	if err != nil {
		log.Fatalf("errors initializing configs: %s", err.Error())
	}

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

	err = goose.Up(db, path+"/"+migrations)
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

func initConfig() (string, error) {

	env := ".env"
	path, err := os.Getwd()
	if err != nil {
		return path, err
	}
	path = strings.SplitAfter(path, "hexsatisfaction")[0]
	fmt.Println(path)
	if err := godotenv.Load(path + "/" + env); err != nil {
		log.Fatal("Error loading .env file")
	}

	viper.AddConfigPath(path)
	viper.SetConfigName("config")

	return path, viper.ReadInConfig()
}
