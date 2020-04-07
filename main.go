package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/zangeronimo/authenticated/env"

	"github.com/dgrijalva/jwt-go"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

//SAMPLESECRET A secret frase
const SAMPLESECRET string = "Hello JWT"

type User struct {
	gorm.Model
	Code  string
	Price uint
	Email string `gorm:"type:varchar(100);unique_index"`
}

func main() {
	env.New()

	database := os.Getenv("SQLITE_BASE")
	db, err := gorm.Open("sqlite3", database)
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&User{})

	// Create
	//db.Create(&User{Code: "L1213", Price: 1000, Email: "luciano2@tudolinux.com.br"})

	// Read
	var user User
	//db.First(&user, 1)                   // find product with id 1
	db.First(&user, "id = ?", "6") // find product with code L1212

	// Update - update product´s price to 2000
	//db.Model(&user).Update("Price", 2000)

	// Delete - delete product
	//db.Delete(&user)

	fmt.Println(user.Email)

	http.HandleFunc("/auth", basicAuth)
	http.ListenAndServe(":4001", nil)
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
