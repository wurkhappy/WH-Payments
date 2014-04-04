package DB

import (
	"database/sql"
	_ "github.com/bmizerany/pq"
	// "log"
)

var SavePayment *sql.Stmt
var UpdatePayment *sql.Stmt
var FindPaymentsByVersionID *sql.Stmt
var FindPaymentByID *sql.Stmt

func CreateStatements() {
	var err error
	SavePayment, err = DB.Prepare("INSERT INTO payment(id, data) VALUES($1, $2)")
	if err != nil {
		panic(err)
	}

	UpdatePayment, err = DB.Prepare("UPDATE payment SET data = $2 WHERE id = $1")
	if err != nil {
		panic(err)
	}

	FindPaymentsByVersionID, err = DB.Prepare("SELECT data FROM payment WHERE data->>'versionID' = $1")
	if err != nil {
		panic(err)
	}

	FindPaymentByID, err = DB.Prepare("SELECT data FROM payment WHERE id = $1")
	if err != nil {
		panic(err)
	}
}
