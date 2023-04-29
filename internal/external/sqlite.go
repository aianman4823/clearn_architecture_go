package external

import (
	"fmt"
	"os"

	"github.com/aianman4823/clearn_architecture_go/internal/config"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func SetupDB() {
	dbname := config.DBName
	db, err := gorm.Open(sqlite.Open(dbname), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	DB = db
}
