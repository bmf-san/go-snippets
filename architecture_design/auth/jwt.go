// Refered to https://github.com/EricLau1/go-api-login
package main

import (
	"encoding/json"
	"fmt"
	"html"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"

	"golang.org/x/crypto/bcrypt"
)

const (
	// Actualy, thease values comes from a form or something for getting user infomations.
	//These values are require validation.
	userName  = "bmf"
	userEmail = "foobar@example.com"
	userPass  = "password"
)

var secretKey = []byte("thisisexampleforauthjwt")

// ex.
// curl -X POST -H 'Content-Type:application/json' http://localhost:9999/login
func login(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Actualy, you need to get values from a form or somthing. ex. name, email, password.
		token, err := signIn(userEmail, userPass)
		if err != nil {
			toJSON(w, err.Error(), http.StatusUnauthorized)
			return
		}

		toJSON(w, token, http.StatusOK)
		return
	}

	toJSON(w, "Method not allowed", http.StatusMethodNotAllowed)
	return
}

func toJSON(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-type", "application/json; charset=UTF8")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Fatal(err)
	}
}

func signIn(userEmail string, userPass string) (string, error) {
	// Actualy, this values stored in a something storage so you need to get it from a something storage by using a something key.
	// ex. user := model.GetByEmail(email) → user.password
	// Here, hash a userPass for password verification(bcrypt.VerifyPassword).
	hashedUserPass, err := bcrypt.GenerateFromPassword([]byte(userPass), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedUserPass), []byte(userPass))
	if err != nil {
		return "", err
	}

	// If password verification is ok, creates and returns a jwt.
	jwt, err := generateJWT()
	if err != nil {
		return "", err
	}

	return jwt, nil
}

func generateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["user_email"] = userEmail
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()
	return token.SignedString(secretKey)
}

func jwtExtract(r *http.Request) (map[string]interface{}, error) {
	headerAuthorization := r.Header.Get("Authorization")
	bearerToken := strings.Split(headerAuthorization, " ")
	tokenString := html.EscapeString(bearerToken[1])
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return nil, err
	}
	return claims, nil
}

// ex.
// curl -H http://localhost:9999/public
func public(w http.ResponseWriter, r *http.Request) {
	toJSON(w, "public page", http.StatusOK)
	return
}

// ex.
// curl -H 'Content-Type:application/json' -H "Authorization:Bearer <JWT>" http://localhost:9999/private
func private(w http.ResponseWriter, r *http.Request) {
	jwtParams, err := jwtExtract(r)
	if err != nil {
		toJSON(w, err.Error(), http.StatusUnauthorized)
		return
	}
	email, ok := jwtParams["user_email"].(string)
	if !ok {
		toJSON(w, "payload invalid", http.StatusUnauthorized)
		return
	}
	toJSON(w, email, http.StatusOK)
	return
}

func middlewareAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authorizationHeader := r.Header.Get("Authorization")
		if authorizationHeader != "" {
			bearerToken := strings.Split(authorizationHeader, " ")
			if len(bearerToken) == 2 {
				token, err := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("Unauthorized")
					}
					return secretKey, nil
				})
				if err != nil {
					w.WriteHeader(http.StatusUnauthorized)
					w.Write([]byte(err.Error()))
					return
				}
				if token.Valid {
					next.ServeHTTP(w, r)
				}
			} else {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("Unauthorized"))
			}
		}
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/login", login)
	mux.HandleFunc("/public", public)
	mux.HandleFunc("/private", middlewareAuth(private))

	if err := http.ListenAndServe(":9999", mux); err != nil {
		fmt.Println(err)
	}
}
