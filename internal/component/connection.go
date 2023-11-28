package component

import (
	"fmt"
	"log"

	"github.com/khairulharu/marketplace/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabase(config *config.Config) *gorm.DB {
	dsn := fmt.Sprintf("host=%v "+"port=%v "+"user=%v "+"password=%v "+"dbname=%v "+"sslmode=%v ",
		config.DB.Host, config.DB.Port, config.DB.User, config.DB.Pass, config.DB.Name, config.DB.SSL)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("connect database failed: %s", err.Error())
	}
	fmt.Println("connect to databse")
	return db
}
