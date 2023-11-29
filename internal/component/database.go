package component

import (
	"fmt"
	"log"

	"github.com/khairulharu/averincrud/internal/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetDatabaseConnection(cnf *config.Config) *gorm.DB {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		cnf.DB.User, cnf.DB.Pass, cnf.DB.Host, cnf.DB.Port, cnf.DB.Name)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("database connection failed: %v", err.Error())
	}

	fmt.Println("database connnected")

	return db
}
