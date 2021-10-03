package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var messages = []string{"Hello", "World!"}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func allMessages(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "%s", messages[])
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(messages)
}

func singleMessage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	i, _ := strconv.Atoi(key)
	fmt.Fprintf(w, messages[i])
}

func newMessage(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	rawdata := string(reqBody)
	message := rawdata[1 : len([]rune(rawdata))-1]
	messages = append(messages, message)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(messages)
}

func changeMessage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	i, _ := strconv.Atoi(key)
	reqBody, _ := ioutil.ReadAll(r.Body)
	rawdata := string(reqBody)
	message := rawdata[1 : len([]rune(rawdata))-1]
	messages[i] = message
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(messages)
}

func deleteMessage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	i, _ := strconv.Atoi(key)
	messages = append(messages[:i], messages[i+1:]...)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(messages)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", mainHandler)
	router.HandleFunc("/api/messages", allMessages)
	router.HandleFunc("/api/message", newMessage).Methods("POST")
	router.HandleFunc("/api/message/{id}", changeMessage).Methods("PUT")
	router.HandleFunc("/api/message/{id}", deleteMessage).Methods("DELETE")
	router.HandleFunc("/api/message/{id}", singleMessage)
	http.ListenAndServe(":5000", router)
}
