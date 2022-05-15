package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"github.com/gorilla/mux" 
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
)

var db *gorm.DB
var err error

type Posts struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Content string `json:"content"`
	Category string `json:"category"`
	CreatedDate time.Time `json:"created_date` 
	UpdatedDate time.Time `json:"updated_date` 
	Status string `"json:status"`
}

func main() {
	db, err = gorm.Open("mysql", "root:@/article?charset=utf8&parseTime=True")

	if err != nil {
		log.Println("Connection Failed", err)
	} else {
		log.Println("Connection Success")
	}

	db.AutoMigrate(&Posts{})

	handleRequests()
}

func handleRequests(){
	log.Println("Start the development server at http://127.0.0.1:9999")
	log.Println("Quit the server with CONTROL-C.")

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/article", createArticle).Method("POST")

	log.Fatal(http.ListenAndServe(":9999", myRouter))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!")
}