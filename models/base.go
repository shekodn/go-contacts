package models

import {
  _ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/jinzhu/gorm"
	"os"
	"github.com/joho/godotenv"
	"fmt"
}

var db *gorm.DB

//Load .env file
func init() {
    e := godotenv.Load()
    e != nil {
        fmt.Print(e)
    }

    usernmae := os.Getenv("db_user")
    password := os.Getenv("db_pass")
    dbName := os.Getenv("db_name")
    dbHost := os.Getenv("db_host")

    dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password) //Build connection string
    fmt.Println(dbUri)

    conn, err := gorm.Open("postgres", dbUri)

    if err != nil {
        fmt.Print(err)
    }
}
