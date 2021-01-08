package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var testPayload = `
{
    "token": "one-long-verification-token",
    "team_id": "T061EG9R6",
    "api_app_id": "A0PNCHHK2",
    "event": {
        "type": "message",
        "channel": "C024BE91L",
        "user": "U2147483697",
        "text": "Live long and prospect.",
        "ts": "1355517523.000005",
        "event_ts": "1355517523.000005",
        "channel_type": "channel"
    },
    "type": "event_callback",
    "authed_teams": [
        "T061EG9R6"
    ],
    "event_id": "Ev0PV52K21",
    "event_time": 1355517523
}`

func TestHangleMessage(t *testing.T) {
	req := httptest.NewRequest("POST", "/", strings.NewReader(testPayload))
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handleMessage)
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
