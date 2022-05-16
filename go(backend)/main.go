package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	CreatedDate time.Time `json:"created_date"` 
	UpdatedDate time.Time `json:"updated_date"` 
	Status string `json:"status"`
}

type Result struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
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
	
	myRouter.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)

		res := Result{Code: 404, Message: "Method not found"}
		response, _ := json.Marshal(res)
		w.Write(response)
	})

	myRouter.MethodNotAllowedHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)

		res := Result{Code: 403, Message: "Method not allowed"}
		response, _ := json.Marshal(res)
		w.Write(response)
	})

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/article/", createArticle).Methods("POST")
	myRouter.HandleFunc("/article/", getArticles).Methods("GET")
	myRouter.HandleFunc("/article/{id}", getArticle).Methods("GET")
	myRouter.HandleFunc("/article/{id}", updateArticle).Methods("PUT")
	myRouter.HandleFunc("/article/{id}", deleteArticle).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":9999", myRouter))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!")
}

func createArticle(w http.ResponseWriter, r *http.Request){
	payloads, _ := ioutil.ReadAll(r.Body)

	var posts Posts
	json.Unmarshal(payloads, &posts)

	db.Create(&posts)

	res := Result{Code: 200, Data: posts, Message: "Success create article"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func getArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: get article")

	posts := []Posts{}
	db.Find(&posts)

	res := Result{Code: 200, Data: posts, Message: "Success get all article"}
	results, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(results)
}

func getArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	postsID := vars["id"]

	var posts Posts

	db.First(&posts, postsID)

	res := Result{Code: 200, Data: posts, Message: "Success get article"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func updateArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	postsID := vars["id"]

	payloads, _ := ioutil.ReadAll(r.Body)

	var postsUpdates Posts
	json.Unmarshal(payloads, &postsUpdates)

	var posts Posts
	db.First(&posts, postsID)
	db.Model(&posts).Updates(postsUpdates)

	res := Result{Code: 200, Data: posts, Message: "Success update article"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

