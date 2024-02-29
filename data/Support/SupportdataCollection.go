package Support

import (
	"encoding/json"
	"finalProject/service/entity"
	"fmt"
	"io/ioutil"
	"net/http"
)

func getSupportData() []entity.SupportData {
	url := "http://127.0.0.1:8383/support"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Ошибка при отправке GET запроса:", err)
		return []entity.SupportData{}
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Ошибка! Получен некорректный код ответа:", resp.StatusCode)
		return []entity.SupportData{}
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ошибка при чтении ответа:", err)
		return []entity.SupportData{}
	}

	var supportData []entity.SupportData
	if err := json.Unmarshal(body, &supportData); err != nil {
		fmt.Println("Ошибка при разборе JSON:", err)
		return []entity.SupportData{}
	}

	return supportData
}

func RunSupport() []entity.SupportData {
	supportData := getSupportData()
	return supportData
}
