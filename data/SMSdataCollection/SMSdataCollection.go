package main

import (
	"encoding/csv"
	"finalProject/service/entity"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func readAndParseFile() []entity.SMSData {
	fileName := "/Users/mac/go/src/finalProject/simulator/SMS.data"
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Ошибка при чтении файла:", err)
		return nil
	}

	lines := strings.Split(string(content), "\n")
	var result []entity.SMSData

	for _, line := range lines {
		sms, valid := validateAndParseLine(line)
		if valid {
			result = append(result, sms)
		}
	}

	return result
}

func checkCountryExistence(countryCode string) bool {
	countryFileContent, err := ioutil.ReadFile("/Users/mac/go/src/finalProject/allowed/country.txt")
	if err != nil {
		fmt.Println("Ошибка при чтении файла стран:", err)
		return false
	}

	countryList := strings.Split(string(countryFileContent), "\n")

	for _, country := range countryList {
		if country == countryCode {
			return true
		}
	}
	return false
}

func checkProviderValidity(provider string) bool {
	providerFileContent, err := ioutil.ReadFile("/Users/mac/go/src/finalProject/allowed/provider.txt")
	if err != nil {
		fmt.Println("Ошибка при чтении файла провайдеров:", err)
		return false
	}

	providerList := strings.Split(string(providerFileContent), "\n")

	for _, validProvider := range providerList {
		if validProvider == provider {
			return true
		}
	}
	return false
}

func validateAndParseLine(line string) (entity.SMSData, bool) {
	fields := strings.Split(line, ";")
	if len(fields) != 4 {
		return entity.SMSData{}, false
	}

	countryCode := fields[0]
	bandwidth := fields[1]
	responseTime := fields[2]
	provider := fields[3]

	countryExists := checkCountryExistence(countryCode)
	providerValid := checkProviderValidity(provider)

	return entity.SMSData{
		Country:      countryCode,
		Bandwidth:    bandwidth,
		ResponseTime: responseTime,
		Provider:     provider,
	}, countryExists && providerValid
}

func main() {
	smsData := readAndParseFile()
	reverse(smsData)
	fmt.Println(smsData)
}

func reverse(data []entity.SMSData) error {
	filename := "/Users/mac/go/src/finalProject/allowed/data_csv.csv"

	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("Ошибка при открытии файла: %v", err)
	}
	defer file.Close()

	rdr := csv.NewReader(file)
	lines, err := rdr.ReadAll()
	if err != nil {
		return fmt.Errorf("Ошибка чтения CSV: %v", err)
	}

	countryCodeToName := make(map[string]string)
	for _, line := range lines {
		countryCodeToName[line[1]] = line[0]
	}

	for i := range data {
		fullCountryName, ok := countryCodeToName[data[i].Country]
		if ok {
			data[i].Country = fullCountryName
		}
	}

	return nil
}
