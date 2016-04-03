package main

import (
	"encoding/json"
	_ "fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	logEntry := string(r.Method) + ": " + string(r.RequestURI)
	log.Println(logEntry)
	response := []byte("Hello World!")
	w.Header().Set("Content-Type", "application/html; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func JSONhandler(w http.ResponseWriter, r *http.Request) {
	logEntry := string(r.Method) + ": " + string(r.RequestURI)
	log.Println(logEntry)
	responseObject := make(map[string]string)
	responseObject["message"] = "Hello World!"
	jsonResponse, _ := json.Marshal(responseObject)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func main() {
	log.Println("starting up and ready for action")
	http.HandleFunc("/", handler)
	http.HandleFunc("/json", JSONhandler)
	http.ListenAndServe(":3030", nil)
}
