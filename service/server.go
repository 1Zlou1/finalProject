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

	// не могу обработать ошибки об отсутствии значений

	result.Data = resultData

	resultJSON, err := json.Marshal(result)
	if err != nil {
		fmt.Fprint(w, "Error:", err)
		return
	}

	fmt.Fprint(w, string(resultJSON))
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
