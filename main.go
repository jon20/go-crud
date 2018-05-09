package main

import (
	"encoding/json"
	"net/http"
	"log"
	"fmt"
)

type Res_Hello struct {
	Messsage string `json:"message"`
	}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", apiRequest)
	log.Fatal(http.ListenAndServe(":9000", mux))
}

func apiRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	
	res, _ := json.Marshal(Res_Hello{"Hello World!!"})
	fmt.Fprintf(w, string(res))
}
