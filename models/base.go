package models

import (
  _ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/jinzhu/gorm"
	"os"
	"github.com/joho/godotenv"
	"fmt"
)

var db *gorm.DB

//Load .env file
func init() {
    e := godotenv.Load()
    if e != nil {
        fmt.Print(e)
    }

    username := os.Getenv("POSTGRES_USER")
    password := os.Getenv("POSTGRES_PASSWORD")
    dbName := os.Getenv("POSTGRES_DB")
    dbHost := os.Getenv("db_host")

    //Build connection string
    dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password)
    fmt.Println(dbUri, "\n")

    conn, err := gorm.Open("postgres", dbUri)

    if err != nil {
        fmt.Print("Error:", err, "\n")
    }

    db = conn
    db.Debug().AutoMigrate(&Account{}, &Contact{}) // DB migration # TODO
}

func GetDB () *gorm.DB {
    return db
}
