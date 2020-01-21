package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHTTPClient(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, `{"name":"Ravi","Company":"TW"}`)
	}))
	defer ts.Close()
	user, err := HTTPClient(ts.URL)
	require.NoError(t, err, "Unexpected error")
	assert.Equal(t, "Ravi", user.Name, "Invalid user Name")
	assert.Equal(t, "TW", user.Company, "Invalid company Name")
}

func TestHTTPClientWhenGetError(t *testing.T) {
	_, err := HTTPClient("https://random-url")
	assert.EqualError(t, err, "request error: Get https://random-url: dial tcp: lookup random-url: no such host", "Invalid error message")
}

func TestHTTPClientWhenReadError(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, `{"name":"Ravi","Company":"TW"}`)
	}))
	defer ts.Close()
	// TODO:: Make sure ioutils.ReadAll fails
	user, err := HTTPClient(ts.URL)
	require.NoError(t, err, "Unexpected error")
	assert.Equal(t, "Ravi", user.Name, "Invalid user Name")
	assert.Equal(t, "TW", user.Company, "Invalid company Name")
}

func TestHTTPClientWhenUnmarshalError(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, `MALFORMED`)
	}))
	defer ts.Close()
	_, err := HTTPClient(ts.URL)
	assert.EqualError(t, err, "umnarshal error: invalid character 'M' looking for beginning of value", "Invalid company Name")
}

