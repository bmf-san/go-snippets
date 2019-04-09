package main

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestSetFlash(t *testing.T) {
	recorder := httptest.NewRecorder()

	SetFlash(recorder, "test", []byte("test message"))

	request := &http.Request{Header: http.Header{"Cookie": recorder.HeaderMap["Set-Cookie"]}}

	cookie, err := request.Cookie("test")
	if err != nil {
		t.Errorf("%v", err)
	}

	t.Errorf("%v", cookie)
}

func TestGetFlash(t *testing.T) {
	recorder := httptest.NewRecorder()

	expected := []byte("test message")
	SetFlash(recorder, "test", expected)

	request := &http.Request{Header: http.Header{"Cookie": recorder.HeaderMap["Set-Cookie"]}}

	actual, err := GetFlash(recorder, request, "test")
	if err != nil {
		t.Errorf("%v", err)
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("actual: %v expected: %v\n", actual, expected)
	}
}
