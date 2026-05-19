package model

import (
	"myapp/dataStore/postgres"
)

type Enroll struct {
	StdId         int64  `json:stdid`
	CourseID      string `json:"cid"`
	Date_Enrolled string `json:"date"`
}

const queryEnrollStd = "INSERT INTO enroll (std_id, course_id, date_enrolled) VALUES ($1, $2, $3);"

const queryGetAllEnrollments = `
	SELECT std_id, course_id, date_enrolled 
	FROM enroll 
	ORDER BY date_enrolled DESC`

const queryDeleteEnrollment = "DELETE FROM enroll WHERE std_id=$1 AND course_id=$2"

func (e *Enroll) EnrollStud() error {
	_, err := postgres.Db.Exec(queryEnrollStd, e.StdId, e.CourseID, e.Date_Enrolled)
	return err
}

func GetAllEnrollments() ([]Enroll, error) {
	rows, err := postgres.Db.Query(queryGetAllEnrollments)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var enrollments []Enroll
	for rows.Next() {
		var e Enroll
		err := rows.Scan(&e.StdId, &e.CourseID, &e.Date_Enrolled)
		if err != nil {
			return nil, err
		}
		enrollments = append(enrollments, e)
	}
	return enrollments, nil
}

func DeleteEnrollment(stdId int64, courseId string) error {
	_, err := postgres.Db.Exec(queryDeleteEnrollment, stdId, courseId)
	return err
}
