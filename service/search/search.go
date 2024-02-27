package main

import (
	"fmt"
)

type SMSData struct {
	Country      string
	Provider     string
	DeliveryTime int
}

type MMSData struct {
	Country      string
	Provider     string
	DeliveryTime int
}

type VoiceCallData struct {
	// ваша структура VoiceCallData
}

type EmailData struct {
	Country      string
	Provider     string
	DeliveryTime int
}

type BillingData struct {
	// ваша структура BillingData
}

type SupportData struct {
	// ваша структура SupportData
}

type IncidentData struct {
	Status string
	// другие поля для инцидентов
}

type ResultSetT struct {
	SMS     [][]SMSData
	MMS     [][]MMSData
	Voice   []VoiceCallData
	Email   map[string][]EmailData
	Billing BillingData
	Support struct {
		Load    int
		AvgWait int
	}
	Incidents []IncidentData
}

func sortByCountryAndProvider(sms []SMSData, orderByCountry bool) []SMSData {
	// подготовка и сортировка данных
	return sms
}

func getResultData() ResultSetT {
	// Получение SMS, MMS, VoiceCall, Email, Support, Billing и Incidents

	smsData := []SMSData{} // Получите ваши данные из источника

	// Получение данных по SMS и подготовить 2 отсортированных списка
	smsByProvider := sortByCountryAndProvider(smsData, false)
	smsByCountry := sortByCountryAndProvider(smsData, true)

	// Получение MMSData из источника
	mmsData := []MMSData{} // Получите ваши данные откуда-то

	// Получение данных по MMS и подготовить 2 отсортированных списка
	// Аналогично, как и со списками SMSData

	voiceData := []VoiceCallData{} // Получите ваши данные

	// Получение данных по VoiceCall

	emailData := []EmailData{} // Получите ваши данные

	// Получение данных по Email и преобразовать их в map[string][]EmailData

	supportData := SupportData{} // Получите ваши данные

	// Получение данных о системе Support

	incidentsData := []IncidentData{} // Получите ваши данные

	// Получение данных об истории инцидентов.

	result := ResultSetT{
		SMS:   [][]SMSData{smsByProvider, smsByCountry},
		MMS:   [][]MMSData{}, // аналогично с MMS
		Voice: voiceData,
		Email: map[string][]EmailData{}, // аналогично с Email
		Support: struct {
			Load    int
			AvgWait int
		}{},
		Incidents: incidentsData,
	}

	// Если есть данные о системе Support, заполните соответствующие поля.

	return result
}

func main() {
	resultData := getResultData()
	fmt.Println(resultData)
}
