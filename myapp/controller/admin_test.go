package controller

import (
	"bytes"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Admin existing test 1
func TestAdmLogin(t *testing.T) {
	// api endpoint to test
	url := "http://localhost:8080/login"

	// user data to send to server
	var jsonStr = []byte(`{"email": "12250163.gcit@rub.edu.bt", "password": "pass"}`)

	// create an http request
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	// create a client
	client := &http.Client{}

	// send api request
	res, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	// check the data in response body
	data, _ := io.ReadAll(res.Body)

	// verifying the status code in response
	assert.Equal(t, http.StatusOK, res.StatusCode)

	// verify the response data
	assert.JSONEq(t, `{"Status": "Login Successful"}`, string(data))
}

// Admin does not exist test 2
func TestAdmNotExist(t *testing.T) {
	// api endpoint to test
	url := "http://localhost:8080/login"

	// user data to send to server
	var jsonStr = []byte(`{"email": "12250163.gcit1@rub.edu.bt", "password": "pass"}`)

	// create an http request
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	// create a client
	client := &http.Client{}

	// send api request
	res, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	// check the data in response body
	data, _ := io.ReadAll(res.Body)

	// verifying the status code in response
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode)

	// verify the response data
	assert.JSONEq(t, `{"error": "sql: no rows in result set"}`, string(data))
}
