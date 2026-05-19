package controller

import (
	"bytes"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddEnroll(t *testing.T) {
	// api endpoint to test
	url := "http://localhost:8080/enroll"

	// user data to send to server
	var jsonStr = []byte(`{"stdid": 100, "cid": "CSC101"}`)

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
	assert.Equal(t, http.StatusCreated, res.StatusCode)

	// verify the response data
	assert.JSONEq(t, `{"status": "student enrolled"}`, string(data))
}

func TestDeleteEnroll(t *testing.T) {
	// api endpoint to test
	url := "http://localhost:8080/enroll/100/CSC101"

	// create an http request
	req, _ := http.NewRequest("DELETE", url, nil)

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
	assert.JSONEq(t, `{"status": "Entrollment deleted"}`, string(data))
}