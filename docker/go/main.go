package main

import (
	"encoding/json"
	"log"
	"net/http"
	//mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/gorilla/mux"
	. "./dao"
	. "./models"
	. "./config"
)

var config = Config{}
var dao = UserDAO{}

func SendHelloWorld(w http.ResponseWriter, r *http.Request) {
	respondWithJson(w, http.StatusOK,  map[string]string{"meeesage": "Hello World"})
	}


func AllUsersEndPoint(w http.ResponseWriter, r *http.Request) {
	users, err := dao.FindAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJson(w, http.StatusOK, users)
}

func FindUserEndpoint(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	user, err := dao.FindById(param["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid user id")
		return
	}
	respondWithJson(w, http.StatusOK, user)
}

func CreateUserEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	user.Created_at = bson.Now()
	user.Updated_at = bson.Now()
	if err := dao.Insert(user); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, user)
}

func UpdateUserEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var user User
	param := mux.Vars(r)

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := dao.Update(user, param["id"]); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

func DeleteUserEndPoint(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	if err := dao.Delete(param["id"]); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(code)
	w.Write(response)
}


func init() {
	config.Read()
	dao.Server = config.Server
	dao.Database = config.Database
	dao.Connect()
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", SendHelloWorld).Methods("GET")
	r.HandleFunc("/users", AllUsersEndPoint).Methods("GET")
	r.HandleFunc("/users/{id}", FindUserEndpoint).Methods("GET")
	r.HandleFunc("/users", CreateUserEndPoint).Methods("POST")
	r.HandleFunc("/users/{id}", UpdateUserEndPoint).Methods("PUT")
	r.HandleFunc("/users/{id}", DeleteUserEndPoint).Methods("DELETE")
	if err := http.ListenAndServe(":9000", r); err != nil {
		log.Fatal(err)
	}
}
