package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"myapp/model"
	"myapp/utils/httpResp"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func AddCourse(w http.ResponseWriter, r *http.Request) {
	// create a varibale of type course to store student info
	var course model.Course

	// extract the data from the request body sent by client
	jsonObj := json.NewDecoder(r.Body)

	// store json data in the course variable, converting jason obj to struct, go object
	err := jsonObj.Decode(&course)

	if err != nil {
		// sending the response back to client
		// converting map to json obj
		httpResp.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	r.Body.Close()

	// no error
	// call model and pass student info
	dbErr := course.Create()

	if dbErr != nil {
		// sending the response back to client
		// converting map to json obj
		httpResp.RespondWithError(w, http.StatusInternalServerError, dbErr.Error())
		return
	}
	// no error
	//w.Write([]byte("Successfully stored"))
	// converting map to json obj
	httpResp.RespondWithJSON(w, http.StatusCreated, map[string]string{"status": "Course added"})
}

// helper function to convert string to int
func getCourseId(courseID string) (int64, error) {
	intID, err := strconv.ParseInt(courseID, 10, 64)
	return intID, err
}

func GetCourse(w http.ResponseWriter, r *http.Request) {
	myMap := mux.Vars(r)
	cid := myMap["cid"]
	cID, idErr := getCourseId(cid)

	if idErr != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, idErr.Error())
		return
	}

	// no error
	var cDetail model.Course
	cDetail = model.Course{CourseId: cID}

	// pass course data to model
	getErr := cDetail.Read()
	if getErr != nil {
		switch getErr {
		case sql.ErrNoRows:
			fmt.Println("m here")
			httpResp.RespondWithError(w, http.StatusNotFound, getErr.Error())
		default:
			httpResp.RespondWithError(w, http.StatusInternalServerError, getErr.Error())
		}
		return
	}
	httpResp.RespondWithJSON(w, http.StatusOK, cDetail)
}

func UpdateCourse(w http.ResponseWriter, r *http.Request) {
	oldCID := mux.Vars(r)["cid"]
	old_cID, idErr := getUserId(oldCID)

	if idErr != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, idErr.Error())
		return
	}

	var course model.Course
	// extract the data from the request body sent by client
	jsonObj := json.NewDecoder(r.Body)

	// store json data in the stud variable, converting jason obj to struct, go object
	err := jsonObj.Decode(&course)

	if err != nil {
		// sending the response back to client
		// converting map to json obj
		httpResp.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	r.Body.Close()

	updateErr := course.Update(old_cID)

	if updateErr != nil {
		switch updateErr {
		case sql.ErrNoRows:
			httpResp.RespondWithError(w, http.StatusNotFound, updateErr.Error())

		default:
			httpResp.RespondWithError(w, http.StatusInternalServerError, updateErr.Error())
		}
	} else {
		httpResp.RespondWithJSON(w, http.StatusOK, course)
	}
}

func DeleteCourse(w http.ResponseWriter, r *http.Request) {
	// processing the request
	cid := mux.Vars(r)["cid"]
	cID, idErr := getCourseId(cid)

	if idErr != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, idErr.Error())
		return
	}

	// option1
	stud := model.Course{CourseId: cID}
	delErr := stud.Delete()

	if delErr != nil {
		switch delErr {
		case sql.ErrNoRows:
			httpResp.RespondWithError(w, http.StatusNotFound, delErr.Error())

		default:
			httpResp.RespondWithError(w, http.StatusInternalServerError, delErr.Error())
		}
	} else {
		httpResp.RespondWithJSON(w, http.StatusOK, map[string]string{"status": "Course Deleted"})
	}
}

func GetAllCourses(w http.ResponseWriter, r *http.Request) {
	courses, getErr := model.GetAllCourses()
	if getErr != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, getErr.Error())
		return
	}
	httpResp.RespondWithJSON(w, http.StatusOK, courses)
}
