package migrations

import (
	"github.com/ynadtochii/ecom/db"
	"github.com/ynadtochii/ecom/db/models"
)

func Migrate() {
	db.DB.AutoMigrate(
		&models.User{},
	)
}
