package game

// write test for game.go

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

// TestTopup is a test function for Topup function
func TestTopup(t *testing.T) {
	handler := NewGameHandler()

	testCases := []struct {
		name         string
		requestBody  RequestBody
		expectedCode int
		expectedBody string
	}{
		{
			name: "TestTopup",
			requestBody: RequestBody{
				Game:    "Mobile Legends",
				GamerID: "GYUTDTE",
				Points:  20,
			},
			expectedCode: http.StatusOK,
			expectedBody: `{"game":"Mobile Legends","gamerID":"GYUTDTE","points":20}`,
		},
		{
			name: "TestTopupWithEmptyGame",
			requestBody: RequestBody{
				Game:    "",
				GamerID: "GYUTDTE",
				Points:  20,
			},
			expectedCode: http.StatusBadRequest,
			expectedBody: `{"error":"Key: 'RequestBody.Game' Error:Field validation for 'Game' failed on the 'required' tag"}`,
		},
		{
			name: "TestTopupWithEmptyGamerID",
			requestBody: RequestBody{
				Game:    "Mobile Legends",
				GamerID: "",
				Points:  20,
			},
			expectedCode: http.StatusBadRequest,
			expectedBody: `{"error":"Key: 'RequestBody.GamerID' Error:Field validation for 'GamerID' failed on the 'required' tag"}`,
		},
		{
			name: "TestTopupWithEmptyPoints",
			requestBody: RequestBody{
				Game:    "Mobile Legends",
				GamerID: "GYUTDTE",
				Points:  0,
			},
			expectedCode: http.StatusBadRequest,
			expectedBody: `{"error":"Key: 'RequestBody.Points' Error:Field validation for 'Points' failed on the 'required' tag"}`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create a new HTTP POST request with the JSON body
			jsonBody, err := json.Marshal(tc.requestBody)
			if err != nil {
				t.Fatalf("Error marshalling JSON: %v", err)
			}
			req, err := http.NewRequest("POST", "/", bytes.NewBuffer(jsonBody))
			if err != nil {
				t.Fatalf("Error creating HTTP request: %v", err)
			}

			// Create a new HTTP response recorder
			rr := httptest.NewRecorder()

			// Create a new Gin router
			router := gin.Default()

			// Define a route to handle POST requests
			router.POST("/", handler.Topup)

			// Perform the request
			router.ServeHTTP(rr, req)

			// Check the status code
			if status := rr.Code; status != tc.expectedCode {
				t.Errorf("Handler returned wrong status code: got %v want %v", status, tc.expectedCode)
			}

			// Check the response body
			if rr.Body.String() != tc.expectedBody {
				t.Errorf("Handler returned unexpected body: got %v want %v", rr.Body.String(), tc.expectedBody)
			}
		})
	}
}
