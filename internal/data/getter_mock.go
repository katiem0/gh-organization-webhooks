package data

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

// MockAPIGetter is a mock implementation of the Getter interface
type MockAPIGetter struct {
	OrganizationWebhooksData []byte
	ShouldReturnError        bool
	ErrorMessage             string
	CreatedWebhooks          []CreatedWebhook
	ShouldReturnResponse     bool
	ResponseBody             []byte
}

// NewMockAPIGetter creates a new mock API getter
func NewMockAPIGetter() *MockAPIGetter {
	return &MockAPIGetter{}
}

// GetOrganizationWebhooks mocks retrieving organization webhooks
func (m *MockAPIGetter) GetOrganizationWebhooks(owner string) ([]byte, error) {
	if m.ShouldReturnError {
		return nil, fmt.Errorf(m.ErrorMessage)
	}
	return m.OrganizationWebhooksData, nil
}

func (m *MockAPIGetter) GetSourceOrganizationWebhooks(owner string) ([]byte, error) {
	if m.ShouldReturnError {
		return nil, fmt.Errorf(m.ErrorMessage)
	}
	if m.ShouldReturnResponse {
		return m.ResponseBody, nil
	}
	// Return an empty array by default
	return []byte(`[]`), nil
}

func (m *MockAPIGetter) CreateWebhookList(data [][]string) []CreatedWebhook {
	if len(m.CreatedWebhooks) > 0 {
		return m.CreatedWebhooks
	}

	// Default implementation
	webhooks := []CreatedWebhook{}
	// Skip header row
	for i, row := range data {
		if i == 0 {
			continue
		}

		if len(row) < 9 {
			continue
		}

		webhook := CreatedWebhook{
			Name:   row[2],
			Active: row[3] == "true",
			Events: strings.Split(row[4], ";"),
			Config: Config{
				ContentType: row[5],
				InsecureSSL: row[6],
				Secret:      row[7],
				Url:         row[8],
			},
		}
		webhooks = append(webhooks, webhook)
	}
	return webhooks
}

func (m *MockAPIGetter) CreateOrganizationWebhook(owner string, data io.Reader) error {
	if m.ShouldReturnError {
		return fmt.Errorf(m.ErrorMessage)
	}
	return nil
}

// TestAPIGetterWrapper wraps a MockRESTClient with the APIGetter interface
type TestAPIGetterWrapper struct {
	MockClient *MockRESTClient
}

func NewAPIGetterWithMockREST(client *MockRESTClient) *TestAPIGetterWrapper {
	return &TestAPIGetterWrapper{
		MockClient: client,
	}
}

// GetOrganizationWebhooks implementation for TestAPIGetterWrapper
func (t *TestAPIGetterWrapper) GetOrganizationWebhooks(owner string) ([]byte, error) {
	url := fmt.Sprintf("orgs/%s/hooks", owner)
	resp, err := t.MockClient.Request("GET", url, nil)
	if err != nil {
		return nil, err
	}
	responseData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return responseData, nil
}

func (t *TestAPIGetterWrapper) CreateWebhookList(data [][]string) []CreatedWebhook {
	var webhookList []CreatedWebhook
	// Implementation similar to MockAPIGetter.CreateWebhookList
	return webhookList
}

func (t *TestAPIGetterWrapper) CreateOrganizationWebhook(owner string, data io.Reader) error {
	url := fmt.Sprintf("orgs/%s/hooks", owner)

	// Check if we can read from the data reader before making the request
	// This simulates checking the request body for errors
	if data != nil {
		// Try to peek at the data without consuming it
		// If it's a bytes.Reader, we can check its size
		if br, ok := data.(*bytes.Reader); ok {
			if br.Len() == 0 {
				return fmt.Errorf("empty or already consumed request body")
			}
		}
	}

	resp, err := t.MockClient.Request("POST", url, data)
	if err != nil {
		return err
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			// Since this is in a testing helper, we can't return the error
			// but we can log it or handle it appropriately
			fmt.Printf("Error closing response body: %v\n", err)
		}
	}()

	// Check for non-successful status codes
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, string(bodyBytes))
	}

	return nil
}
