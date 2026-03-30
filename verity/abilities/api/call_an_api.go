package api

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"sync"

	"github.com/nchursin/verity-bdd/verity/abilities"
)

// CallAnAPI enables an actor to make HTTP requests to APIs
type CallAnAPI interface {
	abilities.Ability
	// SendRequest sends an HTTP request and stores the response
	SendRequest(req *http.Request, ctx context.Context) (*http.Response, error)
	// LastResponse returns the most recent response
	LastResponse() *http.Response
	// SetBaseURL sets the base URL for subsequent requests
	SetBaseURL(baseURL string) error
	// GetBaseURL returns the current base URL
	GetBaseURL() string
}

// callAnAPI implements the CallAnAPI interface
type callAnAPI struct {
	client       *http.Client
	baseURL      string
	lastResponse *http.Response
	mutex        sync.RWMutex
}

// Using creates a new CallAnAPI ability with the given HTTP client
func Using(client *http.Client) CallAnAPI {
	if client == nil {
		client = http.DefaultClient
	}

	return &callAnAPI{
		client:  client,
		baseURL: "",
	}
}

// CallAnApiAt creates a new CallAnAPI ability with the given base URL.
// Паника, если baseURL невалиден, чтобы ошибка surfaced сразу.
func CallAnApiAt(baseURL string) CallAnAPI {
	ab := Using(http.DefaultClient)
	if err := ab.SetBaseURL(baseURL); err != nil {
		panic(fmt.Errorf("CallAnApiAt invalid base URL: %w", err))
	}
	return ab
}

// SendRequest sends an HTTP request and stores the response
func (c *callAnAPI) SendRequest(req *http.Request, ctx context.Context) (*http.Response, error) {
	// Apply base URL if request URL is relative
	c.mutex.RLock()
	baseURL := c.baseURL
	c.mutex.RUnlock()

	if baseURL != "" && req.URL != nil && !req.URL.IsAbs() {
		parsedBaseURL, err := validateBaseURL(baseURL)
		if err != nil {
			return nil, err
		}

		req.URL = parsedBaseURL.ResolveReference(req.URL)
	}

	req = req.WithContext(ctx)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("HTTP request failed: %w", err)
	}

	// Store the response for later retrieval
	c.mutex.Lock()
	if c.lastResponse != nil && c.lastResponse.Body != nil {
		_ = c.lastResponse.Body.Close()
	}
	c.lastResponse = resp
	c.mutex.Unlock()

	return resp, nil
}

// LastResponse returns the most recent response
func (c *callAnAPI) LastResponse() *http.Response {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.lastResponse
}

// SetBaseURL sets the base URL for subsequent requests
func (c *callAnAPI) SetBaseURL(baseURL string) error {
	parsed, err := validateBaseURL(baseURL)
	if err != nil {
		return err
	}

	c.mutex.Lock()
	c.baseURL = parsed.String()
	c.mutex.Unlock()
	return nil
}

// GetBaseURL returns the current base URL
func (c *callAnAPI) GetBaseURL() string {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.baseURL
}

func validateBaseURL(baseURL string) (*url.URL, error) {
	parsedBaseURL, err := url.Parse(baseURL)
	if err != nil {
		return nil, fmt.Errorf("invalid base URL: %w", err)
	}

	if parsedBaseURL.Scheme == "" || parsedBaseURL.Host == "" {
		return nil, fmt.Errorf("invalid base URL: scheme and host are required")
	}

	return parsedBaseURL, nil
}
