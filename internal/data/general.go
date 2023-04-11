package data

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/cli/go-gh/pkg/api"
	"go.uber.org/zap"
	"golang.org/x/term"
)

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

type CreatedWebhook struct {
	Name   string   `json:"name"`
	Active bool     `json:"active"`
	Events []string `json:"events"`
	Config Config   `json:"config"`
}

type Getter interface {
	GetOrganizationWebhooks(owner string) ([]byte, error)
	CreateWebhookList(data [][]string) []Webhook
	CreateOrganizationWebhook(owner string, data []byte) error
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
	if err != nil {
		log.Printf("Body read error, %v", err)
	}
	defer resp.Body.Close()
	responseData, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Body read error, %v", err)
	}
	return responseData, err
}

func (g *APIGetter) CreateWebhookList(data [][]string) []CreatedWebhook {
	// convert csv lines to array of structs
	var webhookList []CreatedWebhook
	var hook CreatedWebhook
	for _, each := range data[1:] {
		hook.Name = each[2]
		hook.Active, _ = strconv.ParseBool(each[3])
		hook.Events = strings.Split(each[4], ";")
		hook.Config.ContentType = each[5]
		hook.Config.InsecureSSL = each[6]
		hook.Config.Secret = each[7]
		hook.Config.Url = each[8]
		webhookList = append(webhookList, hook)
	}
	return webhookList
}

func (g *APIGetter) CreateOrganizationWebhook(owner string, data io.Reader) error {
	url := fmt.Sprintf("orgs/%s/hooks", owner)

	resp, err := g.restClient.Request("POST", url, data)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	return err
}

func GetSourceOrganizationWebhooks(owner string, g *APIGetter) ([]byte, error) {
	url := fmt.Sprintf("orgs/%s/hooks", owner)
	zap.S().Debugf("Reading in hooks from %v", url)
	resp, err := g.restClient.Request("GET", url, nil)
	if err != nil {
		log.Printf("Body read error, %v", err)
	}
	defer resp.Body.Close()
	responseData, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Body read error, %v", err)
	}
	return responseData, err
}

// The entered password will not be displayed on the screen
func SensitivePrompt(label string) string {
	var s string
	for {
		fmt.Fprint(os.Stderr, label+" ")
		pw, _ := term.ReadPassword(int(syscall.Stdin))
		s = string(pw)
		if s != "" {
			break
		}
	}
	fmt.Println()
	return s
}
