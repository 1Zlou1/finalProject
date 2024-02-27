package main

import (
	"encoding/json"
	"finalProject/service/entity"
	"fmt"
	"io/ioutil"
	"net/http"
)

func getIncidentData() []entity.IncidentData {
	url := "http://127.0.0.1:8585"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Ошибка при отправке GET запроса:", err)
		return []entity.IncidentData{}
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Ошибка! Получен некорректный код ответа:", resp.StatusCode)
		return []entity.IncidentData{}
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ошибка при чтении ответа:", err)
		return []entity.IncidentData{}
	}

	var incidentData []entity.IncidentData
	if err := json.Unmarshal(body, &incidentData); err != nil {
		fmt.Println("Ошибка при разборе JSON:", err)
		return []entity.IncidentData{}
	}

	return incidentData
}

func main() {
	incidentData := getIncidentData()
	fmt.Println(incidentData)
}
