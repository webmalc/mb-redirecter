package main

import (
	"fmt"
	"net/http"
	"webmalc/mb-redirector/common/logger"

	resty "github.com/go-resty/resty/v2"
)

// The API object.
type API struct {
	config *Config
	client *resty.Client
	logger *logger.Logger
}

func (s *API) getClients(email, login, alias string) *Clients {
	clients := &Clients{}
	if email == "" && login == "" {
		return &Clients{}
	}
	url := s.config.APIUrl
	if email != "" {
		url += "?email=" + email
	} else if login != "" {
		url += "?login=" + login
	} else if alias != "" {
		url += "?alias=" + login
	}

	s.logger.Infof("Getting clients from %s", url)

	resp, err := s.client.R().
		SetResult(clients).
		SetHeader("Accept", "application/json").
		SetHeader("Authorization", "Token "+s.config.APIToken).
		Get(url)

	if err != nil {
		return &Clients{}
	}
	if resp.StatusCode() != http.StatusOK {
		return &Clients{}
	}

	return clients
}

// GetClientDomain returns client domain.
func (s *API) GetClientDomain(emailLogin string) string {
	if emailLogin == "" {
		return ""
	}
	clients := s.getClients(emailLogin, "", "")
	if len(clients.Results) == 0 {
		clients = s.getClients("", emailLogin, "")
	}
	if len(clients.Results) == 0 {
		clients = s.getClients("", "", emailLogin)
	}
	if len(clients.Results) == 0 {
		return ""
	}
	domain := clients.Results[0].URL
	if domain != "" {
		return domain
	}
	domain = clients.Results[0].LoginAlias
	if domain != "" {
		return fmt.Sprintf(s.config.BaseURL, domain)
	}
	domain = clients.Results[0].Login
	if domain != "" {
		return fmt.Sprintf(s.config.BaseURL, domain)
	}

	return ""
}

// NewAPI creates a new API.
func NewAPI(config *Config, log *logger.Logger) *API {
	maxRetries := 3

	return &API{
		config: config,
		logger: log,
		client: resty.New().SetRetryCount(maxRetries),
	}
}
