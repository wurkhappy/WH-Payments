package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/wurkhappy/WH-Payments/models"
	"net/http"
)

func CreatePaymentsByVersionID(params map[string]interface{}, body []byte) ([]byte, error, int) {
	versionID := params["id"].(string)
	var payments []*models.Payment

	err := json.Unmarshal(body, &payments)
	if err != nil {
		return nil, fmt.Errorf("%s", "Wrong value types"), http.StatusBadRequest
	}

	for _, payment := range payments {
		payment.VersionID = versionID
		//TODO: this should really be a transaction
		//because if one save goes bad and others have already been saved then it could
		//lead to weird zombie payments
		err = payment.Upsert()
		if err != nil {
			return nil, fmt.Errorf("%s %s", "Error saving: ", err.Error()), http.StatusBadRequest
		}
	}

	a, _ := json.Marshal(payments)

	events := Events{&Event{"created.payment", a}}
	go events.Publish()

	return a, nil, http.StatusOK

}

func GetPaymentsByVersionID(params map[string]interface{}, body []byte) ([]byte, error, int) {
	id := params["id"].(string)
	payments, err := models.FindPaymentsByVersionID(id)
	if err != nil {
		return nil, fmt.Errorf("%s", "Error finding agreement"), http.StatusBadRequest
	}

	p, _ := json.Marshal(payments)
	return p, nil, http.StatusOK
}

func UpdatePayment(params map[string]interface{}, body []byte) ([]byte, error, int) {
	id := params["id"].(string)

	var updatedPayment *models.Payment
	json.Unmarshal(body, &updatedPayment)

	payment, err := models.FindPaymentByID(id)
	if err != nil {
		return nil, fmt.Errorf("%s", "Error finding agreement"), http.StatusBadRequest
	}

	payment.PaymentItems = updatedPayment.PaymentItems

	err = payment.Update()
	if err != nil {
		return nil, fmt.Errorf("%s %s", "Error saving: ", err.Error()), http.StatusBadRequest
	}

	jsonString, _ := json.Marshal(payment)
	return jsonString, nil, http.StatusOK

}
