package model

import "myapp/dataStore/postgres"

type Admin struct {
	FirstName string
	Lastname  string
	Email     string
	Password  string
}

const queryInsertAdmin = "INSERT INTO admin(firstname, lastname, email, password) VALUES ($1, $2, $3, $4);"
const queryGetAdmin = "SELECT email, password FROM admin WHERE email=$1 and password=$2;"

func (adm *Admin) Create() error {
	_, err := postgres.Db.Exec(queryInsertAdmin, adm.FirstName, adm.Lastname, adm.Email, adm.Password)
	return err
}

func (adm *Admin) Get() error {
	row := postgres.Db.QueryRow(queryGetAdmin, adm.Email, adm.Password)
	return row.Scan(&adm.Email, &adm.Password)
}
