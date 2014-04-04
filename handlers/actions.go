package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/wurkhappy/WH-Payments/models"
	"net/http"
)

func UpdateAction(params map[string]interface{}, body []byte, userID string) ([]byte, error, int) {
	id := params["id"].(string)
	var payment *models.Payment
	payment, err := models.FindPaymentByID(id)
	if err != nil {
		return nil, fmt.Errorf("%s", "Error finding payment"), http.StatusBadRequest
	}

	var action *models.Action
	json.Unmarshal(body, &action)

	payment.LastAction = action

	payment.Update()

	a, _ := json.Marshal(action)
	return a, nil, http.StatusOK
}
