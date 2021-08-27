package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/state", stateHandler)
	http.HandleFunc("/lock", lockHandler)
	http.HandleFunc("/unlock", unlockHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func stateHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("state")
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	bodyString := string(bodyBytes)
	log.Println(bodyString)
}

func lockHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("lock")
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	bodyString := string(bodyBytes)
	log.Println(bodyString)
}

func unlockHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("unlock")
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	bodyString := string(bodyBytes)
	log.Println(bodyString)
}
