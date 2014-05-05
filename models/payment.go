package models

import (
	"database/sql"
	"encoding/json"
	_ "github.com/bmizerany/pq"
	"github.com/nu7hatch/gouuid"
	"github.com/wurkhappy/WH-Payments/DB"
	"log"
	"time"
)

type Payment struct {
	ID           string       `json:"id"`
	VersionID    string       `json:"versionID"`
	Title        string       `json:"title"`
	DateExpected time.Time    `json:"dateExpected"`
	PaymentItems PaymentItems `json:"paymentItems"`
	LastAction   *Action      `json:"lastAction"`
	IsDeposit    bool         `json:"isDeposit"`
	AmountDue    float64      `json:"amountDue"`
	AmountPaid   float64      `json:"amountPaid"`
	Number       int64        `json:"number"`
}

//for unmarshaling purposes
type payment struct {
	ID           string       `json:"id"`
	VersionID    string       `json:"versionID"`
	Title        string       `json:"title"`
	DateExpected time.Time    `json:"dateExpected"`
	PaymentItems PaymentItems `json:"paymentItems"`
	LastAction   *Action      `json:"lastAction"`
	IsDeposit    bool         `json:"isDeposit"`
	AmountDue    float64      `json:"amountDue"`
	AmountPaid   float64      `json:"amountPaid"`
	Number       int64        `json:"number"`
}

func NewPayment() *Payment {
	id, _ := uuid.NewV4()
	return &Payment{
		ID: id.String(),
	}
}

func (p *Payment) Save() (err error) {
	jsonByte, _ := json.Marshal(p)
	_, err = DB.SavePayment.Exec(p.ID, string(jsonByte))
	if err != nil {
		log.Print(err)
		return err
	}
	return nil
}

func (p *Payment) Upsert() (err error) {
	jsonByte, _ := json.Marshal(p)
	r, err := DB.UpsertPayment.Query(p.ID, string(jsonByte))
	r.Close()
	if err != nil {
		log.Print(err)
		return err
	}
	return nil
}

func (p *Payment) Update() (err error) {
	jsonByte, _ := json.Marshal(p)
	_, err = DB.UpdatePayment.Exec(p.ID, string(jsonByte))
	if err != nil {
		log.Print(err)
		return err
	}
	return nil
}

func FindPaymentsByVersionID(id string) (p []*Payment, err error) {
	r, err := DB.FindPaymentsByVersionID.Query(id)
	if err != nil {
		return nil, err
	}
	defer r.Close()
	return dbRowsToPayments(r)
}

func FindPaymentByID(id string) (p *Payment, err error) {
	var s string
	err = DB.FindPaymentByID.QueryRow(id).Scan(&s)
	if err != nil {
		return nil, err
	}
	json.Unmarshal([]byte(s), &p)
	return p, nil
}

func dbRowsToPayments(r *sql.Rows) (payments []*Payment, err error) {
	for r.Next() {
		var s string
		err = r.Scan(&s)
		if err != nil {
			return nil, err
		}
		var p *Payment
		json.Unmarshal([]byte(s), &p)
		payments = append(payments, p)
	}
	return payments, nil
}

func (p *Payment) UnmarshalJSON(bytes []byte) (err error) {
	var py *payment
	err = json.Unmarshal(bytes, &py)
	if err != nil {
		return err
	}

	if py.ID == "" {
		id, _ := uuid.NewV4()
		py.ID = id.String()
	}

	p.ID = py.ID
	p.VersionID = py.VersionID
	p.Title = py.Title
	p.DateExpected = py.DateExpected
	p.PaymentItems = py.PaymentItems
	p.LastAction = py.LastAction
	p.IsDeposit = py.IsDeposit
	p.AmountDue = py.AmountDue
	p.AmountPaid = py.AmountPaid
	p.Number = py.Number
	return nil
}

func (p *Payment) SetAsPaid() {
	p.AmountPaid = p.AmountDue
}
