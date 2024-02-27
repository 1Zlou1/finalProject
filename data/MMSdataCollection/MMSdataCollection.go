package main

import (
	"encoding/json"
	"finalProject/entity"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// Функция для отправки GET запроса и разбора полученного ответа в структуры
func sendGetRequest() []entity.MMSData {
	url := "http://127.0.0.1:8383"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Ошибка при отправке GET запроса:", err)
		return []entity.MMSData{}
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Ошибка! Получен некорректный код ответа:", resp.StatusCode)
		return []entity.MMSData{}
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ошибка при чтении ответа:", err)
		return []entity.MMSData{}
	}

	var mmsData []entity.MMSData
	err = json.Unmarshal(body, &mmsData)
	if err != nil {
		fmt.Println("Ошибка при разборе JSON:", err)
		return []entity.MMSData{}
	}

	// Удаление элементов, не соответствующих спискам стран и провайдеров
	countryList := make(map[string]struct{})
	providerList := make(map[string]struct{})

	countryFileContent, err := ioutil.ReadFile("/Users/mac/go/src/finalProject/allowed/Country.txt")
	if err != nil {
		fmt.Println("Ошибка при чтении файла стран:", err)
		return []entity.MMSData{}
	}
	for _, country := range strings.Split(string(countryFileContent), "\n") {
		countryList[country] = struct{}{}
	}

	providerFileContent, err := ioutil.ReadFile("/Users/mac/go/src/finalProject/allowed/provider.txt")
	if err != nil {
		fmt.Println("Ошибка при чтении файла провайдеров:", err)
		return []entity.MMSData{}
	}
	for _, provider := range strings.Split(string(providerFileContent), "\n") {
		providerList[provider] = struct{}{}
	}

	var validMMSData []entity.MMSData
	for _, data := range mmsData {
		if _, ok := countryList[data.Country]; ok {
			if _, ok := providerList[data.Provider]; ok {
				validMMSData = append(validMMSData, data)
			}
		}
	}

	return validMMSData
}

func main() {
	mmsData := sendGetRequest()
	fmt.Println(mmsData)
}
