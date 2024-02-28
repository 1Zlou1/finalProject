package main

import (
	"finalProject/data/MMS"
	"finalProject/data/SMS"
	"finalProject/service/entity"
)

func main() {

	// Получение данных
	smsData := SMS.SMSRun()
	mmsData := MMS.MMSRun()
	emailData := getEmailDataFromSource()
	// ...

	// Задание 2: Получение данных по sms и подготовка 2 отсортированных списка
	smsByProvider := sortByProvider(smsData)
	smsByCountry := sortByCountry(smsData)

	// Задание 3: Получение данных по mms и подготовка 2 отсортированных списка
	mmsByProvider := sortByProvider(mmsData)
	mmsByCountry := sortByCountry(mmsData)

	// Задание 5: Получение данных по электронной почте и формирование map со списками провайдеров
	emailByCountry := prepareEmailByCountryData(emailData)

	// Задание 7: Получение данных о системе Support и расчет общего состояния нагрузки
	supportLoad, avgWaitTime := calculateSupportData(supportData)

	// Получение данных об истории инцидентов и сортировка
	sortedIncidents := sortIncidents(incidentsData)

	// Формирование ResultSetT
	resultSet := ResultSetT{
		main2.SMS: [][]SMSData{smsByProvider, smsByCountry},
		MMS:       [][]MMSData{mmsByProvider, mmsByCountry},
		Email:     emailByCountry,
		Support: struct {
			Load    int
			AvgWait int
		}{Load: supportLoad, AvgWait: avgWaitTime},
		Incidents: sortedIncidents,
	}
	// ...
}

func getResultData() entity.ResultSetT {
	entity.ResultSetT{}
}
