package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMainHandlerWhenOk(t *testing.T) {
	totalCount := 2
	urlReq := fmt.Sprintf("/cafe?count=%d&city=moscow", totalCount)
	req := httptest.NewRequest("GET", urlReq, nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	status := responseRecorder.Code
	body := responseRecorder.Body

	require.Equal(t, status, http.StatusOK)
	require.NotNil(t, body)
}

func TestMainHandlerWhenCityNotSupported(t *testing.T) {
	city := "Arkhangelsk"
	urlReq := fmt.Sprintf("/cafe?count=2&city=%s", city)
	req := httptest.NewRequest("GET", urlReq, nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	status := responseRecorder.Code
	err, _ := io.ReadAll(responseRecorder.Body)

	require.Equal(t, status, http.StatusBadRequest)
	require.Equal(t, string(err), "wrong city value")

}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 5
	urlReqMore := fmt.Sprintf("/cafe?count=%d&city=moscow", totalCount)
	urlReqMax := fmt.Sprintf("/cafe?count=%d&city=moscow", 4)

	reqMore := httptest.NewRequest("GET", urlReqMore, nil)
	reqMax := httptest.NewRequest("GET", urlReqMax, nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, reqMore)

	body := responseRecorder.Body

	handler.ServeHTTP(responseRecorder, reqMax)

	bodyMax := responseRecorder.Body

	require.Equal(t, body, bodyMax)
}
