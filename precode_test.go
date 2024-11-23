package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=5&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusOK, responseRecorder.Result().StatusCode)
	expectedCount := 4
	actualCount := len(strings.Split(responseRecorder.Body.String(), ","))
	require.Equal(t, expectedCount, actualCount)
}

func TestMainHandlerReturnStatusOKAndBodyNotEmpty(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=3&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusOK, responseRecorder.Result().StatusCode)
	require.NotEmpty(t, responseRecorder.Body)
}

func TestMainHandlerCityNotSupported(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=1&city=ulyanovsk", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusBadRequest, responseRecorder.Result().StatusCode)

	expectedBody := "wrong city value"
	actualBody := responseRecorder.Body.String()
	require.Equal(t, expectedBody, actualBody)
}
