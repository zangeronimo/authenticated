package main

import (
	"fmt"
	"net/http"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/zangeronimo/authenticated/env"
	"github.com/zangeronimo/authenticated/src/controller/authentication"
	"github.com/zangeronimo/authenticated/src/db"
	gql "github.com/zangeronimo/authenticated/src/graphql"
)

func main() {
	// Start env package with global configurations
	env.New()

	// Start db package to start migration bases
	db.New()

	//GRAPHQL Server
	gql.New()

	http.HandleFunc("/auth", authentication.BasicAuth)

	fmt.Println("The server is running on port 4000")
	http.ListenAndServe(":4000", nil)

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
