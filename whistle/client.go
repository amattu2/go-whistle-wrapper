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
	// API Email
	email string

	// API Password
	password string

	// API Token
	token string

	// Environment
	// ProdEnv or StagingEnv
	Env string

	// Timeout configures the request timeout
	Timeout time.Duration

	// UserAgent is the User-Agent header to send with each request
	UserAgent string
}

// Initialize creates a new client with the given email and password.
//
// email account email address
// password account password
func Initialize(email string, password string) *Client {
	return &Client{
		email:     email,
		password:  password,
		Timeout:   10 * time.Second,
		Env:       ProdEnv,
		UserAgent: "Mozilla/5.0 (X11; Linux x86_64)",
	}
}

// InitializeToken creates a new client with the given token.
// Can be used to restore a session.
//
// token API token
func InitializeToken(token string) *Client {
	return &Client{
		token:     token,
		Timeout:   10 * time.Second,
		Env:       ProdEnv,
		UserAgent: "Mozilla/5.0 (X11; Linux x86_64)",
	}
}

// Get performs a GET request to the given path
//
// path API path
// headers HTTP headers
func (c *Client) Get(path string, headers map[string]string) (*http.Response, error) {
	// Initialize the client
	client := http.Client{}
	client.Timeout = c.Timeout

	// Initialize the request
	request, err := http.NewRequest("GET", fmt.Sprintf("%s/%s", c.Env, path), nil)
	if err != nil {
		return nil, err
	}

	// Add headers
	request.Header.Set("User-Agent", c.UserAgent)
	request.Header.Set("Referer", "https://app.whistle.com/")
	request.Header.Set("X-Whistle-AuthToken", c.getToken())
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Accept-Language", "en-US")
	for key, value := range headers {
		request.Header.Set(key, value)
	}

	return client.Do(request)
}

// Post performs a POST request to the given path
//
// path API path
// headers HTTP headers
// body HTTP body (JSON)
func (c *Client) Post(path string, headers map[string]string, body map[string]string) (*http.Response, error) {
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
	request.Header.Set("User-Agent", c.UserAgent)
	request.Header.Set("Referer", "https://app.whistle.com/")
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Accept-Language", "en-US")
	for key, value := range headers {
		request.Header.Set(key, value)
	}

	return client.Do(request)
}

// getToken returns the API token from a user's email and password
func (c *Client) getToken() string {
	// If token is empty, login and get token
	if (c.token) == "" {
		data := map[string]string{
			"email":    c.email,
			"password": c.password,
		}

		resp, err := c.Post("api/tokens.json", nil, data)
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

// Users returns a list of information about the authenticated user's account
func (c Client) Users() (*UsersResponse, error) {
	// Get data
	resp, err := c.Get("api/users.json", nil)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP error: %d", resp.StatusCode)
	}

	defer resp.Body.Close()

	// Parse json response
	body, _ := io.ReadAll(resp.Body)
	result := UsersResponse{}
	json.Unmarshal(body, &result)

	return &result, nil
}

func (c Client) Notifications() (*NotificationsResponse, error) {
	// Get data
	resp, err := c.Get("api/notifications.json", nil)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP error: %d", resp.StatusCode)
	}

	defer resp.Body.Close()

	// Parse json response
	body, _ := io.ReadAll(resp.Body)
	result := NotificationsResponse{}
	json.Unmarshal(body, &result)

	return &result, nil
}

// Device returns information about a specific device registered
//
// serialNumber Device ID or serial number
func (c Client) Device(serialNumber string) (*[]Device, error) {
	// Get data
	resp, err := c.Get(fmt.Sprintf("api/devices/%s.json", serialNumber), nil)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP error: %d", resp.StatusCode)
	}

	defer resp.Body.Close()

	// Parse json response
	body, _ := io.ReadAll(resp.Body)
	result := []Device{}
	json.Unmarshal(body, &result)

	return &result, nil
}

// Dogs returns a list of information about the authenticated user's dogs
func (c Client) Dogs() (*[]Dog, error) {
	// Get data
	resp, err := c.Get("api/dogs.json", nil)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP error: %d", resp.StatusCode)
	}

	defer resp.Body.Close()

	// Parse json response
	body, _ := io.ReadAll(resp.Body)
	result := []Dog{}
	json.Unmarshal(body, &result)

	return &result, nil
}

// Dog returns information about a specific dog
//
// dogId Dog ID
func (c Client) Dog(dogId string) (*Dog, error) {
	// Get data
	resp, err := c.Get(fmt.Sprintf("api/dogs/%s.json", dogId), nil)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP error: %d", resp.StatusCode)
	}

	defer resp.Body.Close()

	// Parse json response
	body, _ := io.ReadAll(resp.Body)
	result := Dog{}
	json.Unmarshal(body, &result)

	return &result, nil
}

// Highlights returns a list of highlights for a specific dog
//
// dogId Dog ID
// highlightType Highlight type (TBD)
func (c Client) Highlights(dogId string, highlightType string) (*[]Highlight, error) {
	// Get data
	resp, err := c.Get(fmt.Sprintf("api/dogs/%s/highlights.json?type=%s", dogId, highlightType), nil)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP error: %d", resp.StatusCode)
	}

	defer resp.Body.Close()

	// Parse json response
	body, _ := io.ReadAll(resp.Body)
	result := []Highlight{}
	json.Unmarshal(body, &result)

	return &result, nil
}

// Dailies returns a list of daily details for a specific pet
//
// dogId Dog ID
// limit Number of dailies to return
func (c Client) Dailies(dogId string, limit int) (*[]Daily, error) {
	// Get data
	resp, err := c.Get(fmt.Sprintf("api/dogs/%s/dailies.json?count=%d", dogId, limit), nil)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP error: %d", resp.StatusCode)
	}

	defer resp.Body.Close()

	// Parse json response
	body, _ := io.ReadAll(resp.Body)
	result := []Daily{}
	json.Unmarshal(body, &result)

	return &result, nil
}

// Daily returns information about a specific daily
//
// dogId Dog ID
// dailyId Daily ID
func (c Client) Daily(dogId string, dailyId string) (*Daily, error) {
	// Get data
	resp, err := c.Get(fmt.Sprintf("api/dogs/%s/dailies/%s.json", dogId, dailyId), nil)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP error: %d", resp.StatusCode)
	}

	defer resp.Body.Close()

	// Parse json response
	body, _ := io.ReadAll(resp.Body)
	result := Daily{}
	json.Unmarshal(body, &result)

	return &result, nil
}

// Timelines returns a list of timeline events for a specific dog
//
// dogId Dog ID
// timelineId Timeline ID
func (c Client) Timeline(dogId string, timelineId string) (*Timeline, error) {
	// Get data
	resp, err := c.Get(fmt.Sprintf("api/dogs/%s/timelines/%s.json", dogId, timelineId), nil)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP error: %d", resp.StatusCode)
	}

	defer resp.Body.Close()

	// Parse json response
	body, _ := io.ReadAll(resp.Body)
	result := Timeline{}
	json.Unmarshal(body, &result)

	return &result, nil
}

// func (c Client) Statistics(dogId string, statType string) (*StatisticsResponse, error) {
// 	// Get data
// 	resp, err := c.Get(fmt.Sprintf("api/dogs/%s/stats.json?type=%s", dogId, statType), nil)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if resp.StatusCode != http.StatusOK {
// 		return nil, fmt.Errorf("HTTP error: %d", resp.StatusCode)
// 	}

// 	defer resp.Body.Close()

// 	// Parse json response
// 	body, _ := io.ReadAll(resp.Body)
// 	result := StatisticsResponse{}
// 	json.Unmarshal(body, &result)

// 	return &result, nil
// }

// UsersPresent returns a list of users present with a specific dog
//
// dogId Dog ID
//
// Note: Unsure if this endpoint actually does anything
func (c Client) UsersPresent(dogId string) (*UsersPresentResponse, error) {
	// Get data
	resp, err := c.Get(fmt.Sprintf("api/dogs/%s/stats/users_present.json", dogId), nil)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP error: %d", resp.StatusCode)
	}

	defer resp.Body.Close()

	// Parse json response
	body, _ := io.ReadAll(resp.Body)
	result := UsersPresentResponse{}
	json.Unmarshal(body, &result)

	return &result, nil
}

// Goals returns a list of pre-set goals for a specific dog
//
// dogId Dog ID
func (c Client) Goals(dogId string) (*[]Goal, error) {
	// Get data
	resp, err := c.Get(fmt.Sprintf("api/dogs/%s/stats/goals.json", dogId), nil)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP error: %d", resp.StatusCode)
	}

	defer resp.Body.Close()

	// Parse json response
	body, _ := io.ReadAll(resp.Body)
	result := []Goal{}
	json.Unmarshal(body, &result)

	return &result, nil
}

// Averages returns a list of statistical averages for a specific dog
//
// dogId Dog ID
func (c Client) Averages(dogId string) (*[]Average, error) {
	// Get data
	resp, err := c.Get(fmt.Sprintf("api/dogs/%s/stats/averages.json", dogId), nil)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP error: %d", resp.StatusCode)
	}

	defer resp.Body.Close()

	// Parse json response
	body, _ := io.ReadAll(resp.Body)
	result := []Average{}
	json.Unmarshal(body, &result)

	return &result, nil
}

// DailyTotals returns a list of daily totals for a specific dog
//
// dogId Dog ID
// startDate Start date for the daily total events
func (c Client) DailyTotals(dogId string, startDate time.Time) (*[]DailyTotal, error) {
	// Get data
	resp, err := c.Get(fmt.Sprintf("api/dogs/%s/stats/daily_totals.json?start_time=%s", dogId, startDate.Format("Y-m-d")), nil)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP error: %d", resp.StatusCode)
	}

	defer resp.Body.Close()

	// Parse json response
	body, _ := io.ReadAll(resp.Body)
	result := []DailyTotal{}
	json.Unmarshal(body, &result)

	return &result, nil
}

// UsersCreditCard returns the abbreviated credit card details for the current user
func (c Client) UsersCreditCard() (*CreditCard, error) {
	// Get data
	resp, err := c.Get("api/users/credit_card.json", nil)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP error: %d", resp.StatusCode)
	}

	defer resp.Body.Close()

	// Parse json response
	body, _ := io.ReadAll(resp.Body)
	result := CreditCard{}
	json.Unmarshal(body, &result)

	return &result, nil
}
