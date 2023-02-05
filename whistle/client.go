/*
 * Produced: Fri Feb 03 2023
 * Author: Alec M.
 * GitHub: https://amattu.com/links/github
 * Copyright: (C) 2023 Alec M.
 * License: License GNU Affero General Public License v3.0
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package whistle

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	ProdEnv    = "https://app.whistle.com"
	StagingEnv = "https://app-staging.whistle.com"
)

type Client struct {
	// API Credentials
	email, password string

	// API Token/Bearer
	token, bearer string

	// Environment (ProdEnv or StagingEnv)
	Env string

	// Timeout configures the request timeout
	Timeout time.Duration

	// UserAgent is the User-Agent header to send with each request
	UserAgent string
}

type HttpResponse[T interface{}] struct {
	// HTTP Status Code
	StatusCode int `json:"status_code"`

	// Embedded Error Message
	Error error `json:"error"`

	// Embedded Unmarshalled Response Struct
	Response T `json:"response"`

	// Embedded HTTP Response
	Raw *http.Response `json:"raw"`
}

type TokenResponse struct {
	Success  bool     `json:"success"`
	Token    string   `json:"token"`
	Messages []string `json:"messages"`
}

type BearerResponse struct {
	AuthToken    string `json:"auth_token"`
	RefreshToken string `json:"refresh_token"`
	User         User   `json:"user"`
}

type Error struct {
	Message string `json:"message"`
	Code    string `json:"code"`
	Field   string `json:"field"`
}

type User struct {
	CreatedAt            string            `json:"created_at"`
	CurrentUser          bool              `json:"current_user"`
	Email                string            `json:"email"`
	FirstName            string            `json:"first_name"`
	ID                   int               `json:"id"`
	LastName             string            `json:"last_name"`
	ProfilePhotoUrl      string            `json:"profile_photo_url"`
	ProfilePhotoUrlSizes map[string]string `json:"profile_photo_url_sizes"`
	RealtimeChannel      RealtimeChannel   `json:"realtime_channel"`
	Searchable           bool              `json:"searchable"`
	SendMarketingEmails  bool              `json:"send_marketing_emails"`
	UserActivations      string            `json:"user_activations"`
	UserType             string            `json:"user_type"`
	Username             string            `json:"username"`
}

type RealtimeChannel struct {
	Channel string `json:"channel"`
	Service string `json:"service"`
}

// Initialize creates a new client with email and password credentials.
func Initialize(email string, password string) *Client {
	if email == "" || password == "" {
		panic("valid email and password are required")
	}

	return &Client{
		email:     email,
		password:  password,
		Timeout:   10 * time.Second,
		Env:       ProdEnv,
		UserAgent: "Mozilla/5.0 (X11; Linux x86_64)",
	}
}

// InitializeToken creates a new client with an existing API token.
//
// Deprecated: Use username/password or bearer token instead
func InitializeToken(token string) *Client {
	if token == "" {
		panic("valid API token is required")
	}

	return &Client{
		token:     token,
		Timeout:   10 * time.Second,
		Env:       ProdEnv,
		UserAgent: "Mozilla/5.0 (X11; Linux x86_64)",
	}
}

// InitializeBearer creates a new client with an existing HTTP bearer token.
func InitializeBearer(bearer string) *Client {
	if bearer == "" {
		panic("valid http bearer is required")
	}

	return &Client{
		bearer:    bearer,
		Timeout:   10 * time.Second,
		Env:       ProdEnv,
		UserAgent: "Mozilla/5.0 (X11; Linux x86_64)",
	}
}

// addDefaultHeaders adds default headers to a Whistle API request
func (c *Client) addDefaultHeaders(request *http.Request, addAuth bool) {
	// Add headers
	request.Header.Set("User-Agent", c.UserAgent)
	request.Header.Set("Referer", "https://app.whistle.com/")
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/vnd.whistle.com.v4+json")
	request.Header.Set("Accept-Language", "en-US")

	// Add authorization
	if addAuth {
		if c.token != "" {
			request.Header.Set("X-Whistle-AuthToken", c.getToken())
		} else {
			request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.getBearer()))
		}
	}
}

// get makes a HTTP GET request to the Whistle API
func (c *Client) get(path string, headers map[string]string, addAuth bool) (*http.Response, error) {
	// Initialize the client
	client := http.Client{}
	client.Timeout = c.Timeout

	// Initialize the request
	request, err := http.NewRequest("GET", fmt.Sprintf("%s/%s", c.Env, path), nil)
	if err != nil {
		return nil, err
	}

	// Add headers
	c.addDefaultHeaders(request, addAuth)
	for key, value := range headers {
		request.Header.Set(key, value)
	}

	return client.Do(request)
}

// post makes a HTTP POST request to the Whistle API
func (c *Client) post(path string, headers map[string]string, body map[string]string, addAuth bool) (*http.Response, error) {
	// Initialize the client
	client := http.Client{}
	client.Timeout = c.Timeout

	// Initialize the request
	jsonData, _ := json.Marshal(body)
	request, err := http.NewRequest("POST", fmt.Sprintf("%s/%s", c.Env, path), bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	// Add headers
	c.addDefaultHeaders(request, addAuth)
	for key, value := range headers {
		request.Header.Set(key, value)
	}

	return client.Do(request)
}

// getToken returns the API token if it exists, otherwise it will login and return the token
//
// Deprecated: Use getBearer() instead
func (c *Client) getToken() string {
	// If token is empty, login and get token
	if c.token == "" && c.email != "" && c.password != "" {
		data := map[string]string{
			"email":    c.email,
			"password": c.password,
		}

		resp, err := c.post("api/tokens", nil, data, false)
		if err != nil {
			panic(err)
		}
		if resp.StatusCode != http.StatusOK {
			panic(fmt.Errorf("auth failed with HTTP error: %d", resp.StatusCode))
		}

		defer resp.Body.Close()

		// Parse json response
		body, _ := io.ReadAll(resp.Body)
		result := TokenResponse{}
		json.Unmarshal(body, &result)

		if !result.Success {
			panic("Failed to get token")
		}

		c.token = result.Token
	}

	// Return API token
	return c.token
}

// getBearer returns the HTTP bearer if it exists, otherwise it will login and return the bearer
func (c *Client) getBearer() string {
	// If bearer is empty, login and get bearer
	if c.bearer == "" && c.email != "" && c.password != "" {
		data := map[string]string{
			"email":    c.email,
			"password": c.password,
		}

		resp, err := c.post("api/login", nil, data, false)
		if err != nil {
			panic(err)
		}
		if resp.StatusCode != http.StatusCreated {
			panic(fmt.Errorf("auth failed with HTTP error: %d", resp.StatusCode))
		}

		defer resp.Body.Close()

		// Parse json response
		body, _ := io.ReadAll(resp.Body)
		result := BearerResponse{}
		json.Unmarshal(body, &result)

		if result.AuthToken == "" {
			panic("Failed to get bearer")
		}

		c.bearer = result.AuthToken
	}

	// Return HTTP Bearer
	return c.bearer
}
