package controller

import (
	"bytes"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddStudent(t *testing.T) {
	// api endpoint to test
	url := "http://localhost:8080/student/add"

	// user data to send to server
	var jsonStr = []byte(`{"stdid": 104, "fname": "Karma", "lname": "Dorji", "email": "karma3@gmail1.com"}`)

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
	assert.JSONEq(t, `{"status": "Student added"}`, string(data))
}

func TestGetStudent(t *testing.T) {
	// api endpoint to test
	url := "http://localhost:8080/student/103"

	// create a client
	client := &http.Client{}

	// send api request
	res, err := client.Get(url)

	if err != nil {
		panic(err)
	}

	// check the data in response body
	data, _ := io.ReadAll(res.Body)

	// verifying the status code in response
	assert.Equal(t, http.StatusOK, res.StatusCode)

	// verify the response data
	assert.JSONEq(t, `{"stdid": 103, "fname": "Karma", "lname": "Dorji", "email": "karma2@gmail1.com"}`, string(data))
}

func TestDeleteStudent(t *testing.T) {
	// api endpoint to test
	url := "http://localhost:8080/student/103"

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
	assert.JSONEq(t, `{"status": "Student Deleted"}`, string(data))
}
