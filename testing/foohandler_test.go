package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHnadleGetFooRR(t *testing.T) {
	rr := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodGet, "", nil)
	if err != nil {
		t.Error(err)
	}
	handlGetFoo(rr, req)
	if rr.Result().StatusCode != http.StatusOK {
		t.Errorf("expected 200 but got %d", rr.Result().StatusCode)
	}
	defer rr.Result().Body.Close()
	expected := "FOO"
	b, err := io.ReadAll(rr.Result().Body)
	if err != nil {
		t.Error(err)
	}
	if string(b) != expected {
		t.Errorf("expected %s but got %s", expected, string(b))

	}

}

func TestHnadleGetFoo(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(handlGetFoo))
	resp, err := http.Get(server.URL)
	if err != nil {
		t.Error(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected 200 but got %d", resp.StatusCode)
	}
	defer resp.Body.Close()
	expected := "FOO"
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}
	if string(b) != expected {
		t.Errorf("expected %s but got %s", expected, string(b))

	}
}
