package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
)

type BillingData struct {
	CreateCustomer bool
	Purchase       bool
	Payout         bool
	Recurring      bool
	FraudControl   bool
	CheckoutPage   bool
}

func readBillingDataFile() (*BillingData, error) {
	fileName := "/Users/mac/go/src/finalProject/simulator/billing.data"
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	if len(content) < 1 {
		return nil, fmt.Errorf("Фаил %s пуст", fileName)
	}

	maskByte, err := strconv.ParseInt(string(content), 2, 64)
	if err != nil {
		return nil, err
	}

	billingData := &BillingData{
		CreateCustomer: maskByte&(1<<0) > 0,
		Purchase:       maskByte&(1<<1) > 0,
		Payout:         maskByte&(1<<4) > 0,
		Recurring:      maskByte&(1<<5) > 0,
		FraudControl:   maskByte&(1<<6) > 0,
		CheckoutPage:   maskByte&(1<<7) > 0,
	}

	return billingData, nil
}

func main() {
	billingData, err := readBillingDataFile()
	if err != nil {
		fmt.Println("Ошибка при чтении файла:", err)
		return
	}

	fmt.Println("CreateCustomer:", billingData.CreateCustomer)
	fmt.Println("Purchase:", billingData.Purchase)
	fmt.Println("Payout:", billingData.Payout)
	fmt.Println("Recurring:", billingData.Recurring)
	fmt.Println("FraudControl:", billingData.FraudControl)
	fmt.Println("CheckoutPage:", billingData.CheckoutPage)
}
