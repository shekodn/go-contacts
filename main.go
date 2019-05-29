package main

import (
    "github.com/gorilla/mux"
    "go-contacts/app"
    "go-contacts/controllers"
    "os"
    "fmt"
    "net/http"
)

func main() {
    router := mux.NewRouter()
    router.HandleFunc("/api/user/new", controllers.CreateAccount).Methods("POST")
    router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")
    router.HandleFunc("/api/contacts/new", controllers.CreateContact).Methods("POST")
    router.HandleFunc("/api/me/contacts", controllers.GetContactsFor).Methods("GET")

    router.Use(app.JwtAuthentication) //attach JWT auth middleware

    port := os.Getenv("PORT")

    if port == "" {
        port = "8000"
    }

    fmt.Println(port)
    // Launch the app, visit localhost:port/api
    err := http.ListenAndServe(":" + port, router)

    if err != nil {
        fmt.Println(err)
    }
}
