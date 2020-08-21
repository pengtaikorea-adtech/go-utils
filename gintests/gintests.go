package gintests

import (
	"fmt"
	"testing"

	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

// HTTPTesterFunc tester function prototype
type HTTPTesterFunc func(t *testing.T, rec *httptest.ResponseRecorder, req *http.Request)

// Assume http assert
type Assume struct {
	Request *http.Request
	Expects []HTTPTesterFunc
}

// TestRequest - send request then run asserts
func TestRequest(t *testing.T, engine *gin.Engine, asserts ...Assume) {
	rec := httptest.NewRecorder()

	for _, a := range asserts {
		// run request
		engine.ServeHTTP(rec, a.Request)
		//
		for _, exp := range a.Expects {
			exp(t, rec, a.Request)
		}
	}
}

// ExpectStatusOK - expecting http.status = 200
var ExpectStatusOK = ExpectStatusIs(http.StatusOK)

// ExpectStatusExists - http.status != 404
var ExpectStatusExists = ExpectStatusNot(http.StatusNotFound)

const logFormat = "[gin-test] [%s] %s\n\t>>%s"

func ef(t *testing.T, req *http.Request, message string) {
	t.Errorf(logFormat, req.Method, req.URL.Path, message)
}
func fail(t *testing.T, req *http.Request, message string) {
	t.Fatalf("!!"+logFormat, req.Method, req.URL.Path, message)
}

// ExpectStatusIs - check statuscode
func ExpectStatusIs(statusCode int) HTTPTesterFunc {
	return func(t *testing.T, rec *httptest.ResponseRecorder, req *http.Request) {
		if rec.Code != statusCode {
			ef(t, req, fmt.Sprintf("Expecting status %d but %d", statusCode, rec.Code))
		}
	}
}

// ExpectStatusNot - pass if the status code is not
func ExpectStatusNot(statusCode int) HTTPTesterFunc {
	return func(t *testing.T, rec *httptest.ResponseRecorder, req *http.Request) {
		if rec.Code == statusCode {
			ef(t, req, fmt.Sprintf("Expecting status not %d but it is", statusCode))
		}
	}
}

// ExpectHeader - expecting response has header.
// Put "" as val, when to test existance only
func ExpectHeader(key string, val string) HTTPTesterFunc {
	return func(t *testing.T, rec *httptest.ResponseRecorder, req *http.Request) {
		if v, ok := rec.Header()[key]; ok {
			if 0 < len(val) && val != v {
				ef(t, req, fmt.Sprintf("Expecting Header[%s]=%s but %s", key, val, v))
			}
		} else {
			ef(t, req, fmt.Sprintf("Expecting Header[%s] not exists", key))
		}
	}
}

// ExpectResponse - expecting response has cookie
// Put "" as val, when to test existance only.
func ExpectResponse(tester func(ts, rune)) HTTPTesterFunc {
	return func(t *testing.T, rec *httptest.ResponseRecorder, req *http.Request) {
		if payload, _, err := rec.Body.ReadRune(); err == nil {
			tester(t, payload)
		} else {
			fail(t, req, fmt.Sprintf("Request failed"))
		}
	}
}
