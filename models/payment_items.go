package models

type PaymentItem struct {
	TaskID    string  `json:"taskID"`
	SubTaskID string  `json:"subtaskID"`
	Hours     float64 `json:"hours"`
	AmountDue float64 `json:"amountDue"`
	Rate      float64 `json:"rate"`
	Title     string  `json:"title"`
}

type PaymentItems []*PaymentItem
