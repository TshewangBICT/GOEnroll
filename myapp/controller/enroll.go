package controller

import (
	"encoding/json"
	"fmt"
	"myapp/model"
	"myapp/utils/date"
	"myapp/utils/httpResp"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

// handler function to handle add enrollmwent
func Enroll(w http.ResponseWriter, r *http.Request) {
	if !VerifyCookie(w, r) {
		return
	}

	var e model.Enroll

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&e); err != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	current_date := date.GetDate()

	e.Date_Enrolled = current_date

	// Pass e to the model
	saveErr := e.EnrollStud()
	fmt.Println(e)

	if saveErr != nil {
		if strings.Contains(saveErr.Error(), "duplicate key") {
			httpResp.RespondWithError(w, http.StatusForbidden, saveErr.Error())
		} else {
			httpResp.RespondWithError(w, http.StatusInternalServerError, saveErr.Error())
		}
	} else {
		httpResp.RespondWithJSON(w, http.StatusCreated, map[string]string{"status": "student enrolled"})
	}
}

func GetAllEnrollments(w http.ResponseWriter, r *http.Request) {
	if !VerifyCookie(w, r) {
		return
	}

	enrollments, err := model.GetAllEnrollments()
	if err != nil {
		httpResp.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	httpResp.RespondWithJSON(w, http.StatusOK, enrollments)
}

func DeleteEnrollment(w http.ResponseWriter, r *http.Request) {
	if !VerifyCookie(w, r) {
		return
	}

	vars := mux.Vars(r)
	sid := vars["sid"]
	cid := vars["cid"]

	stdId, err := strconv.ParseInt(sid, 10, 64)
	if err != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, "Invalid Student ID")
		return
	}

	err = model.DeleteEnrollment(stdId, cid)
	if err != nil {
		httpResp.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	httpResp.RespondWithJSON(w, http.StatusOK, map[string]string{"status": "Entrollment deleted"})
}
