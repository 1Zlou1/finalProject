package main

import (
	"encoding/json"
	"finalProject/entity"
	"fmt"
	"io/ioutil"
	"net/http"
)

func getSupportData() []entity.SupportData {
	url := "http://127.0.0.1:8484"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Ошибка при отправке GET запроса:", err)
		return []entity.SupportData{}
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Ошибка! Получен некорректный код ответа:", resp.StatusCode)
		return []entity.SupportData{}
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ошибка при чтении ответа:", err)
		return []entity.SupportData{}
	}

	var supportData []entity.SupportData
	if err := json.Unmarshal(body, &supportData); err != nil {
		fmt.Println("Ошибка при разборе JSON:", err)
		return []entity.SupportData{}
	}

	return supportData
}

func main() {
	supportData := getSupportData()
	fmt.Println(supportData)
}
