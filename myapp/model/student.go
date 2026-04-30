package model

import (
	"fmt"
	"myapp/dataStore/postgres"
)

type Student struct {
	StdId     int64  `json:"stdid"`
	FirstName string `json:"fname"`
	LastName  string `json:"lname"`
	Email     string `json:"email"`
}

const queryInsertUser = "INSERT INTO student(stdid, firstname, lastname, email) VALUES ($1, $2, $3, $4)"
const queryGetUser = "SELECT * FROM student WHERE stdid=$1"
const queryUpdateUser = "UPDATE student SET stdid=$1, firstname=$2, lastname=$3, email=$4 WHERE stdid=$5 RETURNING stdid"
const queryDeleteUser = "DELETE from student WHERE stdid=$1 RETURNING stdid"

// responsible for interacting with database
// stud -> s
func (s *Student) Create() error {
	_, err := postgres.Db.Exec(queryInsertUser, s.StdId, s.FirstName, s.LastName, s.Email)
	fmt.Println(err)
	return err
}

func (s *Student) Read() error {
	row := postgres.Db.QueryRow(queryGetUser, s.StdId)
	return row.Scan(&s.StdId, &s.FirstName, &s.LastName, &s.Email)
}

func (s *Student) Update(old_stdID int64) error {
	row := postgres.Db.QueryRow(queryUpdateUser, s.StdId, s.FirstName, s.LastName, s.Email, old_stdID)
	return row.Scan(&s.StdId)
}

func (s *Student) Delete() error {
	return postgres.Db.QueryRow(queryDeleteUser, s.StdId).Scan(&s.StdId)
}

func GetAllStudents() ([]Student, error) {
	rows, getErr := postgres.Db.Query("SELECT * FROM student;")
	if getErr != nil {
		return nil, getErr
	}
	// create a slice of type student
	students := []Student{}

	for rows.Next() {
		var s Student
		dbErr := rows.Scan(&s.StdId, &s.FirstName, &s.LastName, &s.Email)
		if dbErr != nil {
			return nil, dbErr
		}
		students = append(students, s)
	}
	rows.Close()
	return students, nil
}
