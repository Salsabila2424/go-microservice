package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type TokenResponse struct {
	Token string `json:"token"`
}

var jwtKey = []byte("myapps-secret-key")

func generateToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})
	return token.SignedString(jwtKey)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if user.Username == "admin" && user.Password == "password" {
		token, err := generateToken(user.Username)
		if err != nil {
			http.Error(w, "Error generating token", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(TokenResponse{Token: token})
		return
	}
	http.Error(w, "Invalid credentials", http.StatusUnauthorized)
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Welcome to Authentication Service")
	})
	http.HandleFunc("/login", loginHandler)
	fmt.Println("Authentication service is running on port 8081...")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
