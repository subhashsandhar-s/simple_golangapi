package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	podName := os.Getenv("HOSTNAME")
	log.Printf("Request received in Pod: %s, Path: %s, Method: %s\n", podName, r.URL.Path, r.Method)

	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello World"))
}
func main() {
	router := mux.NewRouter()

	router.HandleFunc("/hello", helloWorld)
	log.Println("Server started running at 9001")
	http.ListenAndServe(":9001", router)
}
