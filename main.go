package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/zangeronimo/authenticated/db"
	"github.com/zangeronimo/authenticated/env"
	gql "github.com/zangeronimo/authenticated/graphql"

	"github.com/dgrijalva/jwt-go"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//SAMPLESECRET A secret frase
const SAMPLESECRET string = "Hello JWT"

func main() {
	// Start env package with global configurations
	env.New()

	// Start db package to start migration bases
	db.New()

	//GRAPHQL Server
	gql.New()

	fmt.Println("The server is running on port 4001")
	http.ListenAndServe(":4001", nil)

	//GORM EXAMPLE
	// Create
	//dbase.Create(&db.User{Name: "Luciano Zangeronimo", Email: "zangeronimo@tudolinux.com.br", Products: []db.Product{{Title: "Authenticated"}}})

	// Read
	//var user db.User
	//dbase.First(&user, 1) // find user with id 1
	//dbase.First(&user, "Name = ?", "Luciano") // find user with name is Luciano

	// Update - update product´s price to 2000
	//dbase.Model(&user).Update("Products", []db.Product{{Title: "Authenticated"}})

	// Delete - delete product
	//dbase.Delete(&user)

	//fmt.Println(user.Name)

	//http.HandleFunc("/auth", basicAuth)
	//http.ListenAndServe(":4001", nil)
}

func basicAuth(w http.ResponseWriter, r *http.Request) {
	// Check basic login for basic authentication, first stap to generate a JWT
	auth := strings.SplitN(r.Header.Get("Authorization"), " ", 2)

	if len(auth) != 2 || auth[0] != "Basic" {
		http.Error(w, "authorization failed", http.StatusUnauthorized)
		return
	}

	payload, _ := base64.StdEncoding.DecodeString(auth[1])
	pair := strings.SplitN(string(payload), ":", 2)

	if len(pair) != 2 || !validate(pair[0], pair[1]) {
		http.Error(w, "authorization failed", http.StatusUnauthorized)
		return
	}

	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foo": "bar",
		"nbf": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(SAMPLESECRET))

	fmt.Println(tokenString, err)
	fmt.Fprintf(w, "Hello, %s - Err %s!", tokenString, err)
}

func validate(username, password string) bool {

	basicUsername := os.Getenv("BASIC_USERNAME")
	basicPassword := os.Getenv("BASIC_PASSWORD")

	if username == basicUsername && password == basicPassword {
		return true
	}
	return false
}
