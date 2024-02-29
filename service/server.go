package main

import (
	"encoding/json"
	"finalProject/service/data"
	"finalProject/service/entity"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func handleConnection(w http.ResponseWriter, r *http.Request) {
	result := &entity.ResultT{}

	resultData := data.GetAllResults()
	result.Data = resultData

	fmt.Println(result.Data)

	resultJSON, err := json.Marshal(result)
	if err != nil {
		http.Error(w, "Ошибка обработки данных", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(resultJSON)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api", handleConnection).Methods("GET")

	srv := &http.Server{
		Addr:    "127.0.0.1:8282",
		Handler: r,
	}

	fmt.Println("Сервер запущен на 127.0.0.1:8282/api")
	if err := srv.ListenAndServe(); err != nil {
		fmt.Printf("Ошибка запуска сервера: %v\n", err)
	}
}
