package data

import (
	Billing "finalProject/data/Billding"
	"finalProject/data/Email"
	"finalProject/data/Incident"
	"finalProject/data/MMS"
	"finalProject/data/SMS"
	"finalProject/data/Support"
	"finalProject/data/Voice"
	"finalProject/service/entity"
	"finalProject/service/sort"
)

func getSmsData() [][]entity.SMSData {
	return sort.SortAndSetSMS(SMS.SMSRun())
}
func getMMSData() [][]entity.MMSData {
	return sort.SortAndSetMMS(MMS.MMSRun())
}
func getVoiceData() []entity.VoiceCallData {
	return Voice.RunVoice()
}
func getSupportData() []int {
	return sort.CalculateSupportStatus(Support.RunSupport())
}
func getIncidentData() []entity.IncidentData {
	s := Incident.RunIncident()
	sort.SortIncidents(s)
	return s
}
func getEmailData() map[string][][]entity.EmailData {
	return sort.SortAndSetEmail(Email.RunEmail())
}
func getBillingData() *entity.BillingData {
	return Billing.RunBilling()
}

func GetAllResults() entity.ResultSetT {
	result := entity.ResultSetT{
		SMS:       getSmsData(),
		MMS:       getMMSData(),
		VoiceCall: getVoiceData(),
		Support:   getSupportData(),
		Incidents: getIncidentData(),
		Email:     getEmailData(),
		Billing:   *getBillingData(),
	}
	return result
}
