package main

import (
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/foo", nil)
	w := httptest.NewRecorder()
	handler(w, req)

	resp := w.Result()

	if resp.StatusCode != 200 {
		t.Errorf("resp.StatusCode is not 200. got = %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Fatal(err)
	}

	if string(body) != "Hello World, foo!" {
		t.Errorf("resp.body is not 'Hello World, foo!'. got = %s", string(body))
	}

}