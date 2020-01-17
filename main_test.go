package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

/*
TestGoodOutput - validate it generates correct output
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
