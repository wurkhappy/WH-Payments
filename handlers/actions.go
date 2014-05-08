package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/wurkhappy/WH-Payments/models"
	"net/http"
	"time"
)

func UpdateAction(params map[string]interface{}, body []byte) ([]byte, error, int) {
	id := params["id"].(string)
	var payment *models.Payment
	payment, err := models.FindPaymentByID(id)
	if err != nil {
		return nil, fmt.Errorf("%s", "Error finding payment"), http.StatusBadRequest
	}

	var action *models.Action
	json.Unmarshal(body, &action)
	action.Date = time.Now()
	action.UserID = params["userID"].(string)

	payment.LastAction = action

	if action.Name == models.ActionAccepted {
		fmt.Println("paid")
		payment.SetAsPaid()
	}

	payment.Update()
	fmt.Println(payment)

	go createAndSendEvents(body, payment)

	a, _ := json.Marshal(action)

	return a, nil, http.StatusOK
}

func createAndSendEvents(body []byte, payment *models.Payment) {
	var m map[string]interface{}
	json.Unmarshal(body, &m)
	data := map[string]interface{}{
		"paymentID":      payment.ID,
		"versionID":      payment.VersionID,
		"creditSourceID": m["creditSourceID"],
		"debitSourceID":  m["debitSourceID"],
		"userID":         payment.LastAction.UserID,
		"amount":         payment.AmountDue,
		"date":           payment.LastAction.Date,
		"paymentItems":   payment.PaymentItems,
	}
	j, _ := json.Marshal(data)
	events := Events{&Event{"payment." + payment.LastAction.Name, j}}
	events.Publish()
}
