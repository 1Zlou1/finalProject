package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type IncidentData struct {
	Topic  string `json:"topic"`
	Status string `json:"status"` // возможные статусы active и closed
}

func getIncidentData() []IncidentData {
	url := "http://127.0.0.1:8585"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Ошибка при отправке GET запроса:", err)
		return []IncidentData{}
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Ошибка! Получен некорректный код ответа:", resp.StatusCode)
		return []IncidentData{}
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ошибка при чтении ответа:", err)
		return []IncidentData{}
	}

	var incidentData []IncidentData
	if err := json.Unmarshal(body, &incidentData); err != nil {
		fmt.Println("Ошибка при разборе JSON:", err)
		return []IncidentData{}
	}

	return incidentData
}

func main() {
	incidentData := getIncidentData()
	fmt.Println(incidentData)
}
