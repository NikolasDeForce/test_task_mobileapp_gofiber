package queries

import (
	"log"
	"mobileapp_go_fiber/app/models"
	"mobileapp_go_fiber/platform/db"
	"time"
)

func GetAllUsers() ([]models.User, error) {
	db, err := db.ConnectPostgres()
	if err != nil {
		db.Close()
		return []models.User{}, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM users \n")
	if err != nil {
		return []models.User{}, err
	}
	users := []models.User{}

	var c1, c10 int
	var c2 time.Time
	var c3, c4, c5, c6, c7, c8, c9 string

	for rows.Next() {
		err = rows.Scan(&c1, &c2, &c3, &c4, &c5, &c6, &c7, &c8, &c9, &c10)
		temp := models.User{c1, c2, c3, c4, c5, c6, c7, c8, c9, c10}
		users = append(users, temp)
	}

	return users, nil
}

func GetUserBalance(phonenumber string) ([]models.User, error) {
	db, err := db.ConnectPostgres()
	if err != nil {
		db.Close()
		return []models.User{}, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT balance, phonenumber FROM users WHERE Phonenumber = $1 \n", phonenumber)
	if err != nil {
		return []models.User{}, err
	}
	users := []models.User{}

	var c1, c10 int
	var c2 time.Time
	var c3, c4, c5, c6, c7, c8, c9 string

	for rows.Next() {
		err = rows.Scan(&c1, &c2, &c3, &c4, &c5, &c6, &c7, &c8, &c9, &c10)
		temp := models.User{c1, c2, c3, c4, c5, c6, c7, c8, c9, c10}
		users = append(users, temp)
	}

	return users, nil
}

func InsertUser(u models.User) error {
	db, err := db.ConnectPostgres()
	if err != nil {
		db.Close()
		return err
	}
	defer db.Close()

	if IsUserValid(u) {
		log.Println("User", u.Email, "already exist!")
		return err
	}

	stmt, err := db.Prepare("INSERT INTO users(Id, Created_At, Jwt, Fname, Email, Phonenumber, Password, Gender, Birthday, Balance) values($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)")
	if err != nil {
		return err
	}

	stmt.Exec(u.ID, u.CreatedAt, u.JWTToken, u.FName, u.Email, u.Phonenumber, u.Password, u.Gender, u.Birthday, u.Balance)
	return nil
}

func IsUserValid(u models.User) bool {
	db, err := db.ConnectPostgres()
	if err != nil {
		db.Close()
		return false
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM users WHERE Email = $1 \n", u.Email)
	if err != nil {
		return false
	}

	temp := models.User{}
	var c1, c10 int
	var c2 time.Time
	var c3, c4, c5, c6, c7, c8, c9 string

	for rows.Next() {
		err = rows.Scan(&c1, &c2, &c3, &c4, &c5, &c6, &c7, &c8, &c9, &c10)
		if err != nil {
			return false
		}
		temp = models.User{c1, c2, c3, c4, c5, c6, c7, c8, c9, c10}
	}
	if u.Email == temp.Email {
		return true
	}

	return false
}

func UpdateUserJWTToken(u models.User) error {
	db, err := db.ConnectPostgres()
	if err != nil {
		db.Close()
		return err
	}
	defer db.Close()

	if !IsUserValid(u) {
		log.Println("User", u.Email, "not found!")
		return err
	}

	stmt, err := db.Prepare("UPDATE users SET Jwt = $2 WHERE Id = $1")
	if err != nil {
		return err
	}

	stmt.Exec(u.ID, u.JWTToken)
	return nil
}

func UpdateUser(u models.User) error {
	db, err := db.ConnectPostgres()
	if err != nil {
		db.Close()
		return err
	}
	defer db.Close()

	if !IsUserValid(u) {
		log.Println("User", u.Email, "not found!")
		return err
	}

	stmt, err := db.Prepare("UPDATE users SET Fname = $2, Email = $3, Gender = $4, Birthday = $5 WHERE Id = $1")
	if err != nil {
		return err
	}

	stmt.Exec(u.ID, u.FName, u.Gender, u.Email, u.Birthday)
	return nil
}

func UpdateBalance(u models.User) error {
	db, err := db.ConnectPostgres()
	if err != nil {
		db.Close()
		return err
	}
	defer db.Close()

	if !IsUserValid(u) {
		log.Println("User", u.Email, "not found!")
		return err
	}

	stmt, err := db.Prepare("UPDATE users SET Balance = $2 WHERE Phonenumber = $1")
	if err != nil {
		return err
	}

	stmt.Exec(u.Phonenumber, u.Balance)
	return nil
}

func FindUserId(id int) (models.User, error) {
	db, err := db.ConnectPostgres()
	if err != nil {
		log.Println("Cannot connect to PostreSQL!")
		db.Close()
		return models.User{}, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM users WHERE Id = $1 \n", id)
	if err != nil {
		log.Println("Query:", err)
		return models.User{}, err
	}
	defer rows.Close()

	u := models.User{}

	var c1, c10 int
	var c2 time.Time
	var c3, c4, c5, c6, c7, c8, c9 string

	for rows.Next() {
		err := rows.Scan(&c1, &c2, &c3, &c4, &c5, &c6, &c7, &c8, &c9, &c10)
		if err != nil {
			log.Println(err)
			return models.User{}, err
		}

		u = models.User{c1, c2, c3, c4, c5, c6, c7, c8, c9, c10}
	}

	return u, nil
}

func FindUserEmail(email string) (models.User, error) {
	db, err := db.ConnectPostgres()
	if err != nil {
		log.Println("Cannot connect to PostreSQL!")
		db.Close()
		return models.User{}, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM users WHERE Email = $1 \n", email)
	if err != nil {
		log.Println("Query:", err)
		return models.User{}, err
	}
	defer rows.Close()

	u := models.User{}

	var c1, c10 int
	var c2 time.Time
	var c3, c4, c5, c6, c7, c8, c9 string

	for rows.Next() {
		err := rows.Scan(&c1, &c2, &c3, &c4, &c5, &c6, &c7, &c8, &c9, &c10)
		if err != nil {
			log.Println(err)
			return models.User{}, err
		}

		u = models.User{c1, c2, c3, c4, c5, c6, c7, c8, c9, c10}
	}

	return u, nil
}

func FindUserToken(jwt string) (models.User, error) {
	db, err := db.ConnectPostgres()
	if err != nil {
		log.Println("Cannot connect to PostreSQL!")
		db.Close()
		return models.User{}, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM users WHERE Jwt = $1 \n", jwt)
	if err != nil {
		log.Println("Query:", err)
		return models.User{}, err
	}
	defer rows.Close()

	u := models.User{}

	var c1, c10 int
	var c2 time.Time
	var c3, c4, c5, c6, c7, c8, c9 string

	for rows.Next() {
		err := rows.Scan(&c1, &c2, &c3, &c4, &c5, &c6, &c7, &c8, &c9, &c10)
		if err != nil {
			log.Println(err)
			return models.User{}, err
		}

		u = models.User{c1, c2, c3, c4, c5, c6, c7, c8, c9, c10}
	}

	return u, nil
}

func IsJwtValid(u models.User) bool {
	db, err := db.ConnectPostgres()
	if err != nil {
		db.Close()
		return false
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM users WHERE Jwt = $1 \n", u.JWTToken)
	if err != nil {
		return false
	}

	temp := models.User{}
	var c1, c10 int
	var c2 time.Time
	var c3, c4, c5, c6, c7, c8, c9 string

	for rows.Next() {
		err = rows.Scan(&c1, &c2, &c3, &c4, &c5, &c6, &c7, &c8, &c9, &c10)
		if err != nil {
			return false
		}
		temp = models.User{c1, c2, c3, c4, c5, c6, c7, c8, c9, c10}
	}
	if u.JWTToken == temp.JWTToken {
		return true
	}

	return false
}
