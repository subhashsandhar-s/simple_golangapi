package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello World"))
}
func main() {
	router := mux.NewRouter()

	router.HandleFunc("/hello", helloWorld)
	http.ListenAndServe(":9001", router)
	fmt.Println("Server started running at port 9001")
}
