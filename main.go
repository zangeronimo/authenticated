package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/zangeronimo/authenticated/env"
	"github.com/zangeronimo/authenticated/models"

	"github.com/dgrijalva/jwt-go"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//SAMPLESECRET A secret frase
const SAMPLESECRET string = "Hello JWT"

func main() {
	env.New()

	dbBase := os.Getenv("DB_BASE")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	db, err := gorm.Open("mysql", fmt.Sprintf("%v:%v@/%v?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbBase))
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&models.Company{})
	db.AutoMigrate(&models.Product{})
	db.AutoMigrate(&models.User{})

	// Create
	db.Create(&models.User{Name: "Luciano Zangeronimo", Email: "zangeronimo@tudolinux.com.br", Products: []models.Product{{Title: "Authenticated"}}})

	// Read
	var user models.User
	db.First(&user, 1) // find user with id 1
	//db.First(&user, "Name = ?", "Luciano") // find user with name is Luciano

	// Update - update product´s price to 2000
	//db.Model(&user).Update("Products", []models.Product{{Title: "Authenticated"}})

	// Delete - delete product
	//db.Delete(&user)

	fmt.Println(user.Products)

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
