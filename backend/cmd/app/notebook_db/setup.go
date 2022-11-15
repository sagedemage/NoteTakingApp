package notebook_db

import (
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"github.com/joho/godotenv"
	"os"
	"gorm.io/gorm"
)

const DB_USERNAME = "root"
const DB_NAME = "notebook"
const DB_HOST = "127.0.0.1"
const DB_PORT = "3306"

func InitDB() *gorm.DB {
	/* Setup the Database for MySQL */
	err := godotenv.Load(".env")

	if err != nil {
		panic(err)
	}

	var db_password string = os.Getenv("DB_PASSWORD")

	// Open database
	dsn := DB_USERNAME + ":" + db_password + "@tcp(" + DB_HOST + ")/" + DB_NAME + "?" + 
	"charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	
	if err != nil {
		panic(err)
	}

	// Generate table structure
	db.AutoMigrate(&User{}) // user table
	db.AutoMigrate(&Note{}) // notes table

	return db
}

func UseSQLite(database_path string) *gorm.DB {
	/* Setup the Database */

	// Open database
	db, err := gorm.Open(sqlite.Open(database_path), &gorm.Config{})
	
	if err != nil {
		panic(err)
	}

	// Generate table structure
	db.AutoMigrate(&User{}) // user table
	db.AutoMigrate(&Note{}) // notes table

	return db
}
