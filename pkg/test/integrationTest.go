package test

import (
	"fmt"
	"bytes"
	"testing"
	"net/http"
	"net/http/httptest"

	"github.com/stretchr/testify/assert"
	"github.com/gin-gonic/gin"
)

type TestCase struct {
	Name	string
	Method	string
	Route	string
	Body	string
	Headers	map[string]string
	WantStatus	int
	WantResp	string
}

func Run(t *testing.T, router *gin.Engine, tc TestCase) {
	t.Run(tc.Name, func(t *testing.T) {
		req, _ := http.NewRequest(tc.Method, tc.Route, bytes.NewBufferString(tc.Body))
		if tc.Headers != nil && len(tc.Headers) > 0 {
			for k, v := range tc.Headers {
				req.Header.Add(k, v)
			}
		}

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, tc.WantStatus, w.Code, fmt.Sprintf("want status code %d, got %d instead", tc.WantStatus, w.Code))
		if tc.WantResp != "" {
			assert.JSONEq(t, tc.WantResp, w.Body.String(), "response mismatch")
		}
	})
}