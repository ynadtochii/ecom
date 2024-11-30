
package db

import (
    "fmt"
    "log"
    "os"
    "time"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
    dsn := fmt.Sprintf(
        "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
        os.Getenv("DB_HOST"),
        os.Getenv("DB_USER"),
        os.Getenv("DB_PASSWORD"),
        os.Getenv("DB_NAME"),
        os.Getenv("DB_PORT"),
    )

    var err error
    maxRetries := 5
    for i := 0; i < maxRetries; i++ {
        DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
        if err == nil {
            log.Println("Successfully connected to the database")
            return
        }
        log.Printf("Failed to connect to the database (attempt %d/%d): %v", i+1, maxRetries, err)
        time.Sleep(5 * time.Second)
    }
    log.Fatalf("Failed to connect to the database after %d attempts", maxRetries)
}

func Close() {
	if DB != nil {
		sqlDB, err := DB.DB()
		if err != nil {
			log.Printf("Error getting database instance: %v", err)
			return
		}
		err = sqlDB.Close()
		if err != nil {
			log.Printf("Error closing database connection: %v", err)
		}
	}
}




