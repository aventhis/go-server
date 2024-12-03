package apiserver

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAPIServer_HandleHello(t *testing.T) {
	s := New(NewConfig())
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/hello", nil)
	s.handelHello().ServeHTTP(rec, req)
	assert.Equal(t, rec.Body.String(), "Hello World")
}
