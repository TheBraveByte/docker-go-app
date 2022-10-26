package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_homeHandler(t *testing.T) {
	wr := httptest.NewRecorder()
	rq := httptest.NewRequest("GET", "/", nil)
	homeHandler(wr, rq)
	resp := wr.Result()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("want statuscode %v got statuscode %v", http.StatusOK, resp.StatusCode)
	}
}
