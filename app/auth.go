package app

import (
  "net/http"
  u "go-contacts/utils"
  "strings"
  "go-contacts/models"
  jwt "github.com/dgrijalva/jwt-go"
  "os"
  "context"
  "fmt"
)


var JwtAuthentication = func(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // List of endpoints that do not requiere auth
        notAuth := []string{"/api/user/new", "/api/user/login"}
        //current request path
        requestPath := r.URL.Path
        for _, value := range notAuth {
            if value == requestPath {
                next.ServeHTTP(w, r)
                return
            }
        }

        response := make(map[string]interface{})
        // Grabs the token from the header
        tokenHeader := r.Header.Get("Authorization")

        // Token is missing, so returns with error code 403 (Unauthorized)
        if tokenHeader == "" {
            response = u.Message(false, "Missing auth token")
            w.WriteHeader(http.StatusForbidden)
            w.Header().Add("Content-type", "application/json")
            u.Respond(w, response)
            return
        }
        //The token comes in format `Bearer {token-body}`, we check if the retrieved token matched this requirement
        splitted := strings.Split(tokenHeader, " ")
        if len(splitted) != 2 {
            response = u.Message(false, "Invalid/Malformed token")
            w.WriteHeader(http.StatusForbidden)
            w.Header().Add("Content-Type", "application/json")
            u.Respond(w, response)
            return
        }
        // Grab the token part
        tokenPart := splitted[1]
        tk := &models.Token{}



        token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {
            return []byte(os.Getenv("token_password")), nil
        })

        // Malformed token, returns with http code 403
        if err != nil {
            response  = u.Message(false, "Malformed auth token")
            w.WriteHeader(http.StatusForbidden)
            w.Header().Add("Content-Type", "application/json")
            u.Respond(w, response)
        }

        //  Token is Invalid, maybe not signed on this server
        if !token.Valid {
            response = u.Message(false, "Token is not Valid")
            w.WriteHeader(http.StatusForbidden)
            w.Header().Add("Content-Type", "application/json")
            u.Respond(w, response)
            return
        }

        // Everything went well
        fmt.Sprintf("user %", tk.UserId)
        ctx := context.WithValue(r.Context(), "user", tk.UserId)
        r = r.WithContext(ctx)
        next.ServeHTTP(w, r)
    })
}
