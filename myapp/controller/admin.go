package controller

import (
	"encoding/json"
	"myapp/model"
	"myapp/utils/httpResp"
	"net/http"
	"time"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	var admin model.Admin

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&admin); err != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	r.Body.Close()
	saveErr := admin.Create()

	if saveErr != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, saveErr.Error())
		return
	}
	httpResp.RespondWithJSON(w, http.StatusCreated, map[string]string{"Status": "Admin added"})

}

func Login(w http.ResponseWriter, r *http.Request) {

	var admin model.Admin
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&admin); err != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	r.Body.Close()
	getErr := admin.Get()

	if getErr != nil {
		httpResp.RespondWithError(w, http.StatusInternalServerError, getErr.Error())
	} else {
		// create a cookies
		cookie := http.Cookie{
			Name:    "my-cookie",
			Value:   "my-value",
			Expires: time.Now().Add(30 * time.Minute),
			Secure:  false,
		}

		// send cookie back to client
		http.SetCookie(w, &cookie)

		httpResp.RespondWithJSON(w, http.StatusOK, map[string]string{"Status": "Login Successful"})
	}
}

func LogOut(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:     "my-cookie",
		Value:    "",
		Path:     "/",
		Expires:  time.Now().Add(-1 * time.Hour),
		MaxAge:   -1,
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)
	httpResp.RespondWithJSON(w, http.StatusOK, map[string]string{"Status": "Logged Out"})
}

// helper function to verify cookie
func VerifyCookie(w http.ResponseWriter, r *http.Request) bool {
	cookie, err := r.Cookie("my-cookie")
	if err != nil {
		if err == http.ErrNoCookie {
			httpResp.RespondWithError(w, http.StatusUnauthorized, err.Error())
			return false
		}
		httpResp.RespondWithError(w, http.StatusUnauthorized, err.Error())
		return false
	}
	// verify the cookie value
	if cookie.Value != "my-value" {
		httpResp.RespondWithError(w, http.StatusUnauthorized, "Cookie value does not match")
		return false
	}
	return true
}
