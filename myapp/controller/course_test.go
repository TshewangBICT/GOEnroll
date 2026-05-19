package controller

import (
	"bytes"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddCourse(t *testing.T) {
	// api endpoint to test
	url := "http://localhost:8080/course/add"

	// user data to send to server
	var jsonStr = []byte(`{"cid": "CSC104", "coursename": "Programming"}`)

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
	assert.JSONEq(t, `{"status": "Course added"}`, string(data))
}

func TestGetCourse(t *testing.T) {
	// api endpoint to test
	url := "http://localhost:8080/course/CSC104"

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
	assert.JSONEq(t, `{"cid": "CSC104", "coursename": "Programming"}`, string(data))
}

func TestDeleteCourse(t *testing.T) {
	// api endpoint to test
	url := "http://localhost:8080/course/CSC104"

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
	assert.JSONEq(t, `{"status": "Course Deleted"}`, string(data))
}