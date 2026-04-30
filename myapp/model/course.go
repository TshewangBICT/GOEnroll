package model

import (
	"fmt"
	"myapp/dataStore/postgres"
)

type Course struct {
	CourseId   int64  `json:"cid"`
	CourseName string `json:"coursename"`
}

const queryInsertCourse = "INSERT INTO course(cid, coursename) VALUES ($1, $2)"
const queryGetCourse = "SELECT * FROM course WHERE cid=$1"
const queryUpdateCourse = "UPDATE course SET cid=$1, coursename=$2 WHERE cid=$3 RETURNING cid"
const queryDeleteCourse = "DELETE from course WHERE cid=$1 RETURNING cid"

// responsible for interacting with database
// course -> c
func (c Course) Create() error {
	_, err := postgres.Db.Exec(queryInsertCourse, c.CourseId, c.CourseName)
	fmt.Println(err)
	return err
}

func (c *Course) Read() error {
	row := postgres.Db.QueryRow(queryGetCourse, c.CourseId)
	return row.Scan(&c.CourseId, &c.CourseName)
}

func (c *Course) Update(old_cID int64) error {
	row := postgres.Db.QueryRow(queryUpdateCourse, c.CourseId, c.CourseName, old_cID)
	return row.Scan(&c.CourseId)
}

func (c *Course) Delete() error {
	return postgres.Db.QueryRow(queryDeleteCourse, c.CourseId).Scan(&c.CourseId)
}

func GetAllCourses() ([]Course, error) {
	rows, getErr := postgres.Db.Query("SELECT * FROM course;")
	if getErr != nil {
		return nil, getErr
	}
	// create a slice of type student
	courses := []Course{}

	for rows.Next() {
		var c Course
		dbErr := rows.Scan(&c.CourseId, &c.CourseName)
		if dbErr != nil {
			return nil, dbErr
		}
		courses = append(courses, c)
	}
	rows.Close()
	return courses, nil
}