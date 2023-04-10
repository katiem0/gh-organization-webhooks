package data

import (
	"fmt"
	"io"
	"time"

	"github.com/cli/go-gh/pkg/api"
)

type Getter interface {
	GetOrganizationWebhooks(owner string) ([]byte, error)
}

type APIGetter struct {
	restClient api.RESTClient
}

func NewAPIGetter(restClient api.RESTClient) *APIGetter {
	return &APIGetter{
		restClient: restClient,
	}
}

func (g *APIGetter) GetOrganizationWebhooks(owner string) ([]byte, error) {
	url := fmt.Sprintf("orgs/%s/hooks", owner)

	resp, err := g.restClient.Request("GET", url, nil)
	defer resp.Body.Close()
	responseData, err := io.ReadAll(resp.Body)
	return responseData, err
}

type Webhook struct {
	HookType  string    `json:"type"`
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Active    bool      `json:"active"`
	Events    []string  `json:"events"`
	Config    Config    `json:"config"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}
type Config struct {
	ContentType string `json:"content_type"`
	InsecureSSL string `json:"insecure_ssl"`
	Secret      string `json:"secret"`
	Url         string `json:"url"`
}
