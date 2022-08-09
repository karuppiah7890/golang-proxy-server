package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestProxy(t *testing.T) {
	t.Run("return bad request for request methods other than GET", func(t *testing.T) {
		h := handler{}

		methods := []string{http.MethodDelete, http.MethodHead,
			http.MethodOptions, http.MethodPatch, http.MethodPost,
			http.MethodPut, http.MethodTrace, http.MethodConnect}

		for _, method := range methods {
			t.Run(fmt.Sprintf("return bad request for %v request method", method), func(t *testing.T) {
				req := httptest.NewRequest(method, "localhost:8080/golang.org/x/mod/@v/v0.2.0.mod", nil)
				resRecord := httptest.NewRecorder()
				h.ServeHTTP(resRecord, req)
				code := resRecord.Result().StatusCode
				expectedCode := 400
				if code != expectedCode {
					t.Logf("expected response status code to be %v but got %v for request method %v", expectedCode, code, method)
					t.Fail()
				}
			})
		}
	})

	t.Run("return OK for GET request method", func(t *testing.T) {
		h := handler{}
		resRecord := httptest.NewRecorder()
		req := httptest.NewRequest("GET", "localhost:8080/golang.org/x/mod/@v/v0.2.0.mod", nil)
		h.ServeHTTP(resRecord, req)
		code := resRecord.Result().StatusCode
		expectedCode := 200
		if code != expectedCode {
			t.Logf("expected response status code to be %v but got %v", expectedCode, code)
			t.Fail()
		}
	})
}
