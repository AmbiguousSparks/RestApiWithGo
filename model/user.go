package model

import (
	"ApiUsers/database"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type User struct {
	ID       string `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
	Birthday string `json:"birthday,omitempty"`
	Photo    string `json:"photo,omitempty"`
	Adm      string `json:"adm,omitempty"`
	Gender   string `json:"gender,omitempty"`
	Password string `json:"password,omitempty"`
	Country  string `json:"country,omitempty"`
}

func (r *User) UpdateNoPhoto() bool {
	//atualizar o usuário sem mexer na foto, usada pra quando não atualizou a foto
	query := "update users set name = ?, email = ?, birthday = ?, "
	query += "adm = ?, gender = ?, password = ?, country = ? where id = ? limit 1"

	stmt, err := connection().Prepare(query)
	checkErr(err)

	res, _ := stmt.Exec(r.Name, r.Email, r.Birthday, r.Adm, r.Gender, r.Password, r.Country, r.ID)

	stmt.Close()
	connection().Close()
	if res == nil {
		return false
	}

	return true
}

func (r *User) UpdatePhoto() bool {
	//atualizar o usuário e mexer na foto, usada pra quando atualizou a foto
	query := "update users set name = ?, email = ?, birthday = ?, "
	query += "adm = ?, gender = ?, password = ?, country = ?, photo = ? where id = ? limit 1"

	stmt, err := connection().Prepare(query)
	checkErr(err)

	res, _ := stmt.Exec(r.Name, r.Email, r.Birthday, r.Adm, r.Gender, r.Password, r.Country, r.Photo, r.ID)

	stmt.Close()
	connection().Close()
	if res == nil {
		return false
	}

	return true
}

func (r *User) Save() bool {
	query := "insert users set name = ?, email = ?, birthday = ?, "
	query += "photo = ?, adm = ?, gender = ?, password = ?, country = ?"
	stmt, err := connection().Prepare(query)
	checkErr(err)

	res, _ := stmt.Exec(r.Name, r.Email, r.Birthday, r.Photo, r.Adm, r.Gender, r.Password, r.Country)

	checkErr(err)
	stmt.Close()
	connection().Close()

	if res == nil {
		return false
	}
	return true
}

func (r *User) Delete() bool {
	query := "delete from users where id = ?"
	stmt, err := connection().Prepare(query)
	checkErr(err)

	res, _ := stmt.Exec(r.ID)

	stmt.Close()
	connection().Close()

	if res == nil {
		return false
	}
	return true
}

func (r *User) ReadOne() {
	query := "Select name, photo, email, birthday, country, password, adm, gender from users where id = ?"
	err := connection().QueryRow(query, r.ID).Scan(&r.Name, &r.Photo, &r.Email, &r.Birthday, &r.Country,
		&r.Password, &r.Adm, &r.Gender)

	connection().Close()
	checkErr(err)
}

func checkErr(error error) {
	if error != nil {
		log.Fatal(error)
	}
}

func (r *User) All() []User {
	query := "Select * from users"
	rows, err := connection().Query(query)

	checkErr(err)

	var users []User

	for rows.Next() {
		var user User

		err := rows.Scan(&user.ID, &user.Name, &user.Gender, &user.Birthday,
			&user.Email, &user.Photo, &user.Password, &user.Adm, &user.Country)

		checkErr(err)

		users = append(users, user)
	}
	rows.Close()
	connection().Close()
	return users

}

func connection() *sql.DB {
	return database.GetConnection()
}
