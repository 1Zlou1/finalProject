package Email

import (
	"finalProject/service/entity"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func readAndParseEmailDataFile() []entity.EmailData {
	fileName := "/Users/mac/go/src/finalProject/simulator/email.data"
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Ошибка при чтении файла:", err)
		return nil
	}

	lines := strings.Split(string(content), "\n")
	var result []entity.EmailData

	for _, line := range lines {
		fields := strings.Split(line, ";")
		if len(fields) == 3 {
			country := fields[0]
			provider := fields[1]
			deliveryTime, err := strconv.Atoi(fields[2])
			if err != nil {
				continue
			}

			if !checkCountryExistence(country) || !checkProviderValidity(provider) {
				continue
			}

			email := entity.EmailData{
				Country:      country,
				Provider:     provider,
				DeliveryTime: deliveryTime,
			}
			result = append(result, email)
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

	countryList := make(map[string]struct{})
	for _, country := range strings.Split(string(countryFileContent), "\n") {
		countryList[country] = struct{}{}
	}

	_, exists := countryList[countryCode]
	return exists
}

func checkProviderValidity(provider string) bool {
	providerFileContent, err := ioutil.ReadFile("/Users/mac/go/src/finalProject/allowed/provider.txt")
	if err != nil {
		fmt.Println("Ошибка при чтении файла провайдеров:", err)
		return false
	}

	providerList := make(map[string]struct{})
	for _, provider := range strings.Split(string(providerFileContent), "\n") {
		providerList[provider] = struct{}{}
	}

	_, valid := providerList[provider]
	return valid
}

func RunEmail() []entity.EmailData {
	emailData := readAndParseEmailDataFile()
	return emailData
}
