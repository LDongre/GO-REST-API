package main

import (
	"log"
	"net/http"

	"./controller"
	"./model"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	mux := controller.Register()
	db := model.Connect()
	defer db.Close() // finally of java
	log.Fatal(http.ListenAndServe("localhost:3000", mux))
}
