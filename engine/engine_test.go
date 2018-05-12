package engine

import (
	"testing"
	"net/http"
	"io/ioutil"
	"net/http/httptest"
)

var routeTests = []struct {
	route	string
	status	int
	body	string
}{
	{"/TestCheck", 200, "Hi /TestCheck"},
	{"/Foo", 200, "Hi /Foo"},
	{"/Foo/Bar", 200, "Hi /Foo/Bar"},
}

func TestEngineHandler (t *testing.T) {
	server := httptest.NewServer(&EngineHandler{})
	defer server.Close()

	for _, tt := range routeTests {
		response, err := http.Get(server.URL + tt.route)
		if err != nil {
			t.Errorf("Unexpected error %v", err)
		}
		if (response.StatusCode != tt.status) {
			t.Errorf("Expected status %d, got %d", tt.status, response.StatusCode)
		}
		bodyBytes, _ := ioutil.ReadAll(response.Body)
		bodyString := string(bodyBytes)
		response.Body.Close()
		if (bodyString != tt.body) {
			t.Errorf("Expected body %s, got %s", tt.body, bodyString)
		}
	}
}

func TestEngine (t *testing.T) {
	go NewEngine()

	for _, tt := range routeTests {
		response, err := http.Get("http://localhost:8080" + tt.route)
		if err != nil {
			t.Errorf("Unexpected error %v", err)
		}
		if (response.StatusCode != tt.status) {
			t.Errorf("Expected status %d, got %d", tt.status, response.StatusCode)
		}
		bodyBytes, _ := ioutil.ReadAll(response.Body)
		bodyString := string(bodyBytes)
		response.Body.Close()
		if (bodyString != tt.body) {
			t.Errorf("Expected body %s, got %s", tt.body, bodyString)
		}
	}
}
