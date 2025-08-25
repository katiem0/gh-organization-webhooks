package data

import (
	"bytes"
	"context"
	"io"
	"net/http"
)

// MockRESTClient for testing REST API calls
type MockRESTClient struct {
	RequestFunc func(method string, path string, body io.Reader) (*http.Response, error)
}

func (m *MockRESTClient) Request(method string, path string, body io.Reader) (*http.Response, error) {
	if m.RequestFunc != nil {
		return m.RequestFunc(method, path, body)
	}
	return &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(bytes.NewReader([]byte("{}"))),
	}, nil
}

func (m *MockRESTClient) RequestWithContext(ctx context.Context, method string, path string, body io.Reader) (*http.Response, error) {
	return m.Request(method, path, body)
}

// This function is required for the interface compatibility
func (m *MockRESTClient) GraphQL(query string, variables map[string]interface{}, result interface{}) error {
	return nil
}

// This function is required for the interface compatibility
func (m *MockRESTClient) Do(req *http.Request) (*http.Response, error) {
	return &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(bytes.NewReader([]byte(`[]`))),
	}, nil
}

// This function is required for the interface compatibility
func (m *MockRESTClient) BuildRequestURL(path string) string {
	return path
}
