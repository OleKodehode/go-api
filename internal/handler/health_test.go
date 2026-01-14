package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"
)
// t is a pointer to a struct from the testing package
func TestHealth(t *testing.T) {
	tests := []struct {
		name	string
		method	string
		path	string
		expectedStatus	int
		expectedBody	string
	}{
		{
			name: "GET /HEALTH - Correct method/path",
			method: http.MethodGet,
			path:	"/v1/health",
			expectedStatus: http.StatusOK,
			expectedBody:	"OK - API Server running\n",
		},
		{
			name: "Wrong Method",
			method: http.MethodPost,
			path: "/v1/health",
			expectedStatus: http.StatusMethodNotAllowed,
			expectedBody: "",	// ServeMux returns 405
		},
		{
			name: "Wrong Path",
			method: http.MethodGet,
			path: "/wrong",
			expectedStatus: http.StatusNotFound,
		},
	}
	// Note for self; tt = test table, test this, test target. 
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, tt.path, nil)
			// rr, response recorder - Pretends to be a responsewriter
			rr := httptest.NewRecorder()

			mux := http.NewServeMux()
			mux.HandleFunc("GET /v1/health", Health)
			mux.ServeHTTP(rr, req)

			if status := rr.Code; status != tt.expectedStatus {
				t.Errorf("Handler returned wrong status code: got %v want %v", status, tt.expectedStatus)
			}

			if tt.expectedBody != "" {
				if body := rr.Body.String(); body != tt.expectedBody {
					t.Errorf("handler returned unexpected body: got %q want %q", body, tt.expectedBody)
				}
			}
		})
	}
}