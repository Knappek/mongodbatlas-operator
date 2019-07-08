package test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Server returns an http Client, ServeMux, and Server. The client proxies
// requests to the server and handlers can be registered on the mux to handle
// requests. The caller must close the test server.
func Server() (*http.Client, *http.ServeMux, *httptest.Server) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	transport := &RewriteTransport{&http.Transport{
		Proxy: func(req *http.Request) (*url.URL, error) {
			return url.Parse(server.URL)
		},
	}}
	client := &http.Client{Transport: transport}
	return client, mux, server
}

// RewriteTransport rewrites https requests to http to avoid TLS cert issues
// during testing.
type RewriteTransport struct {
	Transport http.RoundTripper
}

// RoundTrip rewrites the request scheme to http and calls through to the
// composed RoundTripper or if it is nil, to the http.DefaultTransport.
func (t *RewriteTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.URL.Scheme = "http"
	if t.Transport == nil {
		return http.DefaultTransport.RoundTrip(req)
	}
	return t.Transport.RoundTrip(req)
}

// AssertMethod test that the Request has the expected http method
func AssertMethod(t *testing.T, expectedMethod string, req *http.Request) {
	assert.Equal(t, expectedMethod, req.Method)
}

// AssertReqJSON tests that the Request has the expected key values pairs json
// encoded in its Body
func AssertReqJSON(t *testing.T, expected map[string]interface{}, req *http.Request) {
	var reqJSON interface{}
	err := json.NewDecoder(req.Body).Decode(&reqJSON)
	if err != nil {
		t.Errorf("error decoding request JSON %v", err)
	}
	assert.Equal(t, expected, reqJSON)
}
