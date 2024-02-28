package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func handleConnection(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "ok")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handleConnection).Methods("GET")

	srv := &http.Server{
		Addr:    "localhost:8282",
		Handler: r,
	}

	fmt.Println("Сервер запущен на localhost:8282")
	if err := srv.ListenAndServe(); err != nil {
		fmt.Printf("Ошибка запуска сервера: %v\n", err)
	}
}
