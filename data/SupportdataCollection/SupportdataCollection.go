package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type SupportData struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Product string `json:"product"`
}

func getSupportData() []SupportData {
	url := "http://127.0.0.1:8484"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Ошибка при отправке GET запроса:", err)
		return []SupportData{}
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Ошибка! Получен некорректный код ответа:", resp.StatusCode)
		return []SupportData{}
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ошибка при чтении ответа:", err)
		return []SupportData{}
	}

	var supportData []SupportData
	if err := json.Unmarshal(body, &supportData); err != nil {
		fmt.Println("Ошибка при разборе JSON:", err)
		return []SupportData{}
	}

	return supportData
}

func main() {
	supportData := getSupportData()
	fmt.Println(supportData)
}
