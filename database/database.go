package database

import (
	"fmt"
	"strconv"
	"time"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/postgres"

	"gokit-todos/lib"
	"gokit-todos/database/models"
)

var (
	ContextKey = lib.ContextKey("database")
)

func InitDb() *gorm.DB {
	db := openConnection()
	models.Automigrate(db)
	return db
}

func openConnection() *gorm.DB {
	// A connection attempt will be made `maxRetries` time.
	maxRetries, _ := strconv.Atoi(lib.GetEnv("DB_MAX_RETRIES", "3"))

	// The next connection attempt will be made after `waitTime` seconds.
	waitTime, _ := strconv.Atoi(lib.GetEnv("DB_WAIT_TIME", "5"))

	for i := 1; i <= maxRetries; i++ {
		fmt.Printf("Opening Connection; Attempt %d of %d...\n", i, maxRetries)
		dbConfig := "sslmode=disable host=db port=5432 dbname=todos user=go password=go"

		db, err := gorm.Open("postgres", dbConfig)
		if err != nil {
			// No need to sleep on the last attempt.
			if i != maxRetries {
				fmt.Printf("Cannot open connection (retrying in %ds): %v\n", waitTime, err)
				time.Sleep(time.Duration(waitTime) * time.Second)
			}
			continue
		}

		return db
	}

	panic(fmt.Errorf("database: gave up after %d retries", maxRetries))
}
