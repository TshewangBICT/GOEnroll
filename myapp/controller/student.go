package controller

import (
	"database/sql"
	"encoding/json"
	"myapp/model"
	"myapp/utils/httpResp"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func AddStudent(w http.ResponseWriter, r *http.Request) {
	// if !VerifyCookie(w, r) {
	// 	return
	// }
	// create a varibale of type student to store student info
	var stud model.Student

	// extract the data from the request body sent by client
	jsonObj := json.NewDecoder(r.Body)

	// store json data in the stud variable, converting jason obj to struct, go object
	err := jsonObj.Decode(&stud)

	if err != nil {
		// sending the response back to client
		// converting map to json obj
		httpResp.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	r.Body.Close()

	// no error
	// call model and pass student info
	dbErr := stud.Create()

	if dbErr != nil {
		// sending the response back to client
		// converting map to json obj
		httpResp.RespondWithError(w, http.StatusInternalServerError, dbErr.Error())
		return
	}
	// no error
	//w.Write([]byte("Successfully stored"))
	// converting map to json obj
	httpResp.RespondWithJSON(w, http.StatusCreated, map[string]string{"status": "Student added"})
}

// helper function to convert string to int
func getUserId(userID string) (int64, error) {
	intID, err := strconv.ParseInt(userID, 10, 64)
	return intID, err
}

func GetStud(w http.ResponseWriter, r *http.Request) {
	// if !VerifyCookie(w, r) {
	// 	return
	// }

	myMap := mux.Vars(r)
	stdid := myMap["sid"]
	stdID, idErr := getUserId(stdid)

	if idErr != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, idErr.Error())
		return
	}

	// no error
	var sDetail model.Student
	sDetail = model.Student{StdId: stdID}

	// pass student data to model
	getErr := sDetail.Read()
	if getErr != nil {
		switch getErr {
		case sql.ErrNoRows:
			httpResp.RespondWithError(w, http.StatusNotFound, getErr.Error())
		default:
			httpResp.RespondWithError(w, http.StatusInternalServerError, getErr.Error())
		}
		return
	}
	httpResp.RespondWithJSON(w, http.StatusOK, sDetail)
}

func UpdateStud(w http.ResponseWriter, r *http.Request) {
	// if !VerifyCookie(w, r) {
	// 	return
	// }

	oldSID := mux.Vars(r)["sid"]
	old_stdID, idErr := getUserId(oldSID)

	if idErr != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, idErr.Error())
		return
	}

	var stud model.Student
	// extract the data from the request body sent by client
	jsonObj := json.NewDecoder(r.Body)

	// store json data in the stud variable, converting jason obj to struct, go object
	err := jsonObj.Decode(&stud)

	if err != nil {
		// sending the response back to client
		// converting map to json obj
		httpResp.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	r.Body.Close()

	updateErr := stud.Update(old_stdID)

	if updateErr != nil {
		switch updateErr {
		case sql.ErrNoRows:
			httpResp.RespondWithError(w, http.StatusNotFound, updateErr.Error())

		default:
			httpResp.RespondWithError(w, http.StatusInternalServerError, updateErr.Error())
		}
	} else {
		httpResp.RespondWithJSON(w, http.StatusOK, stud)
	}
}

func DeleteStud(w http.ResponseWriter, r *http.Request) {
	// if !VerifyCookie(w, r) {
	// 	return
	// }

	// processing the request
	sid := mux.Vars(r)["sid"]
	stdID, idErr := getUserId(sid)

	if idErr != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, idErr.Error())
		return
	}

	// option1
	stud := model.Student{StdId: stdID}
	delErr := stud.Delete()

	if delErr != nil {
		switch delErr {
		case sql.ErrNoRows:
			httpResp.RespondWithError(w, http.StatusNotFound, delErr.Error())

		default:
			httpResp.RespondWithError(w, http.StatusInternalServerError, delErr.Error())
		}
	} else {
		httpResp.RespondWithJSON(w, http.StatusOK, map[string]string{"status": "Student Deleted"})
	}
}

func GetAllStuds(w http.ResponseWriter, r *http.Request) {
	students, getErr := model.GetAllStudents()
	if getErr != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, getErr.Error())
		return
	}
	httpResp.RespondWithJSON(w, http.StatusOK, students)
}
