package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

/*
TestGoodOutput - validate it generates correct output with good input
*/
func TestGoodOutput(t *testing.T) {
	inputFile, err := os.Open("ex_input.xml")
	if err != nil {
		t.Fatal(err)
	}

	expectedJSON, err := ioutil.ReadFile("./ex_output.json")
	if err != nil {
		t.Fatal(err)
	}

	request, err := http.NewRequest("POST", "/process", inputFile)
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(processRequestHandler)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(responseRecorder, request)

	// Check the status code is what we expect.
	if responseRecorder.Code != http.StatusOK {
		t.Errorf("Server returned incorrect status: %v", responseRecorder.Code)
	}

	// Check the response body is what we expect.
	if responseRecorder.Body.String() != string(expectedJSON) {
		t.Errorf("Server returned incorrect body: %v", responseRecorder.Body.String())
	}
}

/*
TestEmptyOutput - valid payload with zero orders should still return 200
*/
func TestEmptyOutput(t *testing.T) {
	inputReader := strings.NewReader("<orderList></orderList>")

	request, err := http.NewRequest("POST", "/process", inputReader)
	if err != nil {
		t.Fatal(err)
	}

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(processRequestHandler)
	handler.ServeHTTP(responseRecorder, request)

	// Check the status code is what we expect.
	if responseRecorder.Code != http.StatusOK {
		t.Errorf("Server returned incorrect status: %v", responseRecorder.Code)
	}
}

/*
TestWrongMethod - try a GET, get a 405
*/
func TestWrongMethod(t *testing.T) {
	request, err := http.NewRequest("GET", "/process", nil)
	if err != nil {
		t.Fatal(err)
	}

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(processRequestHandler)
	handler.ServeHTTP(responseRecorder, request)
	if responseRecorder.Code != http.StatusMethodNotAllowed {
		t.Errorf("Server returned incorrect status: %v", responseRecorder.Code)
	}
}

/*
TestMissingFields - make sure gaps in data don't cause the server to choke
*/
func TestMissingFields(t *testing.T) {
	inputReader := strings.NewReader("<orderList><order><id>a</id></order></orderList>")

	request, err := http.NewRequest("POST", "/process", inputReader)
	if err != nil {
		t.Fatal(err)
	}

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(processRequestHandler)
	handler.ServeHTTP(responseRecorder, request)
	if responseRecorder.Code != http.StatusOK {
		t.Errorf("Server returned incorrect status: %v", responseRecorder.Code)
	}
}
