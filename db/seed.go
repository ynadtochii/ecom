package db

import (
	"fmt"

	"github.com/ynadtochii/ecom/db/models"
)

func Seed() {
    users := []models.User{
        {Username: "John Doe" },
        {Username: "Jane Smith"},
    }

    for _, user := range users {
        DB.Create(&user)
    }
    fmt.Println("seeds have been planted")
}
