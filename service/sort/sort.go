package sort

import (
	"finalProject/service/entity"
	"math"
	"sort"
)

func SortAndSetSMS(input []entity.SMSData) [][]entity.SMSData {
	var result [][]entity.SMSData

	var providers []entity.SMSData
	var countries []entity.SMSData
	for _, data := range input {

		providers = append(providers, data)

		countries = append(countries, data)
	}

	// Сортируем оба созданных среза
	sort.Slice(providers, func(i, j int) bool {
		return providers[i].Country < providers[j].Country
	})

	sort.Slice(countries, func(i, j int) bool {
		return countries[i].Provider < countries[j].Provider
	})

	result = append(result, providers)
	result = append(result, countries)

	return result
}

func SortAndSetMMS(input []entity.MMSData) [][]entity.MMSData {
	var result [][]entity.MMSData

	var providers []entity.MMSData
	var countries []entity.MMSData
	for _, data := range input {

		providers = append(providers, data)

		countries = append(countries, data)
	}

	sort.Slice(providers, func(i, j int) bool {
		return providers[i].Country < providers[j].Country
	})

	sort.Slice(countries, func(i, j int) bool {
		return countries[i].Provider < countries[j].Provider
	})

	result = append(result, providers)
	result = append(result, countries)

	return result
}

func SortAndSetEmail(input []entity.EmailData) map[string][][]entity.EmailData {
	result := make(map[string][][]entity.EmailData)

	providersByCountry := make(map[string][]entity.EmailData)
	for _, data := range input {
		providersByCountry[data.Country] = append(providersByCountry[data.Country], data)
	}

	for countryCode, providers := range providersByCountry {
		sort.Slice(providers, func(i, j int) bool {
			return providers[i].DeliveryTime < providers[j].DeliveryTime
		})

		var fastestProviders [][]entity.EmailData
		var slowestProviders [][]entity.EmailData

		if len(providers) >= 3 {
			fastestProviders = append(fastestProviders, providers[:3])
			slowestProviders = append(slowestProviders, providers[len(providers)-3:])
		} else {

			fastestProviders = append(fastestProviders, providers)
		}

		result[countryCode] = append(result[countryCode], fastestProviders...)
		result[countryCode] = append(result[countryCode], slowestProviders...)
	}

	return result
}

func CalculateSupportStatus(input []entity.SupportData) []int {
	totalTickets := 0
	for _, d := range input {
		totalTickets += d.ActiveTickets
	}

	averageMinutesPerTicket := 60 / 18
	potentialWaitTime := int(math.Round(float64(totalTickets) * float64(averageMinutesPerTicket)))

	var load int
	if totalTickets <= 9 {
		load = 1
	} else if totalTickets <= 16 {
		load = 2
	} else {
		load = 3
	}
	EmailData := []int{load, potentialWaitTime}

	return EmailData
}

func SortIncidents(data []entity.IncidentData) {
	sort.SliceStable(data, func(i, j int) bool {
		if data[i].Status == "active" {
			return true
		}
		return false
	})
}
