package queries

import (
	"mobileapp_go_fiber/app/models"
	"mobileapp_go_fiber/platform/db"
	"time"
)

func GetUserTransactions(id int) ([]models.Transaction, error) {
	db, err := db.ConnectPostgres()
	if err != nil {
		db.Close()
		return []models.Transaction{}, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM transactions WHERE Iduser = $1 \n", id)
	if err != nil {
		return []models.Transaction{}, err
	}
	trc := []models.Transaction{}

	var c1, c2, c5 int
	var c3 time.Time
	var c4 string

	for rows.Next() {
		err = rows.Scan(&c1, &c2, &c3, &c4, &c5)
		temp := models.Transaction{c1, c2, c3, c4, c5}
		trc = append(trc, temp)
	}

	return trc, nil
}

func InsertTransaction(t models.Transaction) error {
	db, err := db.ConnectPostgres()
	if err != nil {
		db.Close()
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO transactions(Id, Iduser, Created_At, Phonenumber, Summary) values($1, $2, $3, $4, $5)")
	if err != nil {
		return err
	}

	stmt.Exec(t.ID, t.IDUser, t.CreatedAt, t.Phonenumber, t.Summary)
	return nil
}
