package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewTime(test *testing.T) {
	time := time.Now()
	actual := newTime(time)
	expected := timeStruct{time}

	assert.Equal(test, expected, actual)
}

func TestNewTimeNow(test *testing.T) {
	actual := newTimeNow()
	expected := time.Now()

	assert.WithinDuration(test, expected, actual.Time, time.Second)
}

func TestHealthCheck(t *testing.T) {
	router := setupRouter()

	req1, req1Err := http.NewRequest("GET", "/api/v1/healthcheck", nil)
	req2, req2Err := http.NewRequest("GET", "/api/v1/categories", nil)
	req3, req3Err := http.NewRequest("POST", "/api/v1/category", nil)

	assert.NoError(t, req1Err)
	assert.NoError(t, req2Err)
	assert.NoError(t, req3Err)

	expectedTime := time.Now()
	newRecordReq1 := httptest.NewRecorder()
	router.ServeHTTP(newRecordReq1, req1)
	newRecordReq2 := httptest.NewRecorder()
	router.ServeHTTP(newRecordReq2, req2)
	newRecordReq3 := httptest.NewRecorder()
	router.ServeHTTP(newRecordReq3, req3)

	var actualTime timeStruct
	req1Err = json.Unmarshal(newRecordReq1.Body.Bytes(), &actualTime)

	assert.Equal(t, http.StatusOK, newRecordReq1.Code)
	assert.Equal(t, "application/json; charset=utf-8", newRecordReq1.Header().Get("Content-Type"))
	assert.NoError(t, req1Err)
	assert.WithinDuration(t, expectedTime, actualTime.Time, time.Second)

	assert.Equal(t, http.StatusOK, newRecordReq2.Code)
	assert.Equal(t, http.StatusBadRequest, newRecordReq3.Code)
}
