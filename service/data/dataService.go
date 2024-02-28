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
)

func getSmsData() []entity.SMSData {
	return SMS.SMSRun()
}
func getMMSData() []entity.MMSData {
	return MMS.MMSRun()
}
func getVoiceData() []entity.VoiceCallData {
	return Voice.RunVoice()
}
func getSupportData() []entity.SupportData {
	return Support.RunSupport()
}
func getIncidentData() []entity.IncidentData {
	return Incident.RunIncident()
}
func getEmailData() []entity.EmailData {
	return Email.RunEmail()
}
func getBillingData() *entity.BillingData {
	return Billing.RunBilling()
}

func GetAllResults() *entity.ResultSetT {

	smsData := getSmsData()
	mmsData := getMMSData()

}
