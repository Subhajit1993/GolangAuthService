package database

import (
	"Authentication/internal/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"sync"
)

var (
	PgDB   *gorm.DB
	oncePG sync.Once
)

func InitPgDatabase() error {
	var err error
	oncePG.Do(func() {
		model := []interface{}{&entities.Users{}, &entities.Passwordless{}, &entities.RefreshTokens{}}
		dsn := "host=" + os.Getenv("SELF_POSTGRES_HOST") + " user=" + os.Getenv("SELF_POSTGRES_USER") + " password=" + os.Getenv("SELF_POSTGRES_PASSWORD") + " dbname=" + os.Getenv("SELF_POSTGRES_DB") + " port=" + os.Getenv("SELF_POSTGRES_PORT") + " sslmode=" + os.Getenv("SELF_SSL")
		PgDB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			return
		}
		// AutoMigrate is a variadic function that accepts a list of models and creates tables for them
		// Reference: https://stackoverflow.com/questions/46654132/go-gorm-db-automigrate
		PgDB.AutoMigrate(model...)
		log.Default().Println("AuthService Database Connected")
	})
	return err
}
