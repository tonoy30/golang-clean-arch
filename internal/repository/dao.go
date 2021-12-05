package repository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Masterminds/squirrel"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

type DAO interface {
	NewCourseRepository() CourseRepository
}
type dao struct{}

var DB *sql.DB

func pgQb() squirrel.StatementBuilderType {
	return squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).RunWith(DB)
}
func NewSQLDao(db *sql.DB) DAO {
	DB = db
	return &dao{}
}

func NewSQLDB() (*sql.DB, error) {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln("cannot read from a config")
	}
	host := viper.GetString("database.host")
	port := viper.GetInt("database.port")
	user := viper.GetString("database.user")
	dbname := viper.GetString("database.dbname")
	password := viper.GetString("database.password")
	dbdriver := viper.GetString("database.dbdriver")
	// Starting a database
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable", host, port, user, dbname, password)
	DB, err = sql.Open(dbdriver, psqlInfo)
	if err != nil {
		return nil, err
	}
	return DB, nil
}
func (d dao) NewCourseRepository() CourseRepository {
	return &courseRepository{}
}
