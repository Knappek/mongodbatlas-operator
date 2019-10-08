package test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	knappekv1alpha1 "github.com/Knappek/mongodbatlas-operator/pkg/apis/knappek/v1alpha1"
	"github.com/stretchr/testify/assert"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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

// CreateAtlasProject returns a standard atlas project
func CreateAtlasProject(projectName string, projectID string, namespace string, organizationID string) *knappekv1alpha1.MongoDBAtlasProject {
	return &knappekv1alpha1.MongoDBAtlasProject{
		ObjectMeta: metav1.ObjectMeta{
			Name:      projectName,
			Namespace: namespace,
		},
		Spec: knappekv1alpha1.MongoDBAtlasProjectSpec{
			OrgID: organizationID,
		},
		Status: knappekv1alpha1.MongoDBAtlasProjectStatus{
			ID:           projectID,
			Name:         projectName,
			OrgID:        organizationID,
			Created:      "2016-07-14T14:19:33Z",
			ClusterCount: 0,
		},
	}
}
