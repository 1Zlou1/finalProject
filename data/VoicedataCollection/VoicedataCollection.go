package main

import (
	"finalProject/entity"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func readAndParseVoiceDataFile() []entity.VoiceCallData {
	fileName := "/Users/mac/go/src/finalProject/simulator/Voice.data"
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Ошибка при чтении файла:", err)
		return nil
	}

	lines := strings.Split(string(content), "\n")
	var result []entity.VoiceCallData

	for _, line := range lines {
		callData, valid := validateAndParseLine(line)
		if valid {
			result = append(result, callData)
		}
	}

	return result
}

func validateAndParseLine(line string) (entity.VoiceCallData, bool) {
	fields := strings.Split(line, ";")
	if len(fields) != 8 {
		return entity.VoiceCallData{}, false
	}

	country := fields[0]
	bandwidth := fields[1]
	responseTime := fields[2]
	provider := fields[3]
	connectionStability, err := strconv.ParseFloat(fields[4], 32)
	if err != nil {
		return entity.VoiceCallData{}, false
	}
	ttfb, err := strconv.Atoi(fields[5])
	if err != nil {
		return entity.VoiceCallData{}, false
	}
	voicePurity, err := strconv.Atoi(fields[6])
	if err != nil {
		return entity.VoiceCallData{}, false
	}
	medianOfCallsTime, err := strconv.Atoi(fields[7])
	if err != nil {
		return entity.VoiceCallData{}, false
	}

	if !checkCountryExistence(country) || !checkProviderValidity(provider) {
		return entity.VoiceCallData{}, false
	}

	return entity.VoiceCallData{
		Country:             country,
		Bandwidth:           bandwidth,
		ResponseTime:        responseTime,
		Provider:            provider,
		ConnectionStability: float32(connectionStability),
		TTFB:                ttfb,
		VoicePurity:         voicePurity,
		MedianOfCallsTime:   medianOfCallsTime,
	}, true
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

func main() {
	voiceCallData := readAndParseVoiceDataFile()
	fmt.Println(voiceCallData)
}
