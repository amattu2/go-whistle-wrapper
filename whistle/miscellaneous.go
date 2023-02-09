/*
 * Produced: Sat Feb 04 2023
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
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type NotificationsResponse struct {
	Items []Notification `json:"items"`
}

type ReverseGeocodeResponse struct {
	Description GeocodeDescription `json:"description"`
	QueryLat    string             `json:"query_latitude"`
	QueryLon    string             `json:"query_longitude"`
}

type GeocodeDescription struct {
	Address string `json:"address"`
	City    string `json:"place"`
	Region  string `json:"region"`
	Country string `json:"country"`
}

type AdventureCategoriesResponse struct {
}

type PetFood struct {
	ID   int    `json:"id"`
	Name string `json:"name"`

	// Not present in all responses
	FoodPortion string `json:"food_portion"`
	Unit        string `json:"unit"`
}

type Notification struct {
	Error string             `json:"error"`
	Items []NotificationItem `json:"items"`
}

type NotificationItem struct {
	// Not present in all responses
	Actor NotificationItemActor `json:"actor"`

	Message          string                `json:"message"`
	Target           NotificationItemActor `json:"target"`
	CreatedAt        string                `json:"created_at"`
	Unread           bool                  `json:"unread"`
	NotificationType string                `json:"notification_type"`
}

type NotificationItemActor struct {
	Type string `json:"type"`

	// An abbreviated version of the pet
	Value Pet `json:"value"`
}

type Place struct {
	ID            int         `json:"id"`
	Name          string      `json:"name"`
	Address       string      `json:"address"`
	Latitude      string      `json:"latitude"`
	Longitude     string      `json:"longitude"`
	RadiusMeters  float64     `json:"radius_meters"`
	Shape         string      `json:"shape"`
	Outline       []LatLon    `json:"outline"`
	CreatedByUser bool        `json:"created_by_user"`
	PetIds        []int       `json:"pet_ids"`
	WifiNetwork   WifiNetwork `json:"wifi_network"`
}

type LatLon struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type Coupon struct {
}

// Notifications returns a list of the pending notifications for the user.
func (c Client) Notifications() *HttpResponse[NotificationsResponse] {
	resp, err := c.get("api/notifications", nil, true)
	if err != nil || resp.StatusCode != http.StatusOK {
		return &HttpResponse[NotificationsResponse]{
			StatusCode: resp.StatusCode,
			Error:      err,
			Raw:        resp,
		}
	}

	defer resp.Body.Close()

	// Parse json response
	body, _ := io.ReadAll(resp.Body)
	result := NotificationsResponse{}
	json.Unmarshal(body, &result)

	return &HttpResponse[NotificationsResponse]{
		StatusCode: resp.StatusCode,
		Response:   result,
		Raw:        resp,
	}
}

// PetFoods lists the pet foods by food type (dog_treat, dog_food)
func (c Client) PetFoods(foodType string) *HttpResponse[[]PetFood] {
	resp, err := c.get(fmt.Sprintf("api/pet_foods?type=%s", foodType), nil, true)
	if err != nil || resp.StatusCode != http.StatusOK {
		return &HttpResponse[[]PetFood]{
			StatusCode: resp.StatusCode,
			Error:      err,
			Raw:        resp,
		}
	}

	defer resp.Body.Close()

	// Parse json response
	body, _ := io.ReadAll(resp.Body)
	result := []PetFood{}
	json.Unmarshal(body, &result)

	return &HttpResponse[[]PetFood]{
		StatusCode: resp.StatusCode,
		Response:   result,
		Raw:        resp,
	}
}

// ReverseGeocode returns the best address guess of a given latitude and longitude
func (c Client) ReverseGeocode(lat string, lon string) *HttpResponse[ReverseGeocodeResponse] {
	resp, err := c.get(fmt.Sprintf("api/reverse_geocode?latitude=%s&longitude=%s", lat, lon), nil, true)
	if err != nil || resp.StatusCode != http.StatusOK {
		return &HttpResponse[ReverseGeocodeResponse]{
			StatusCode: resp.StatusCode,
			Error:      err,
			Raw:        resp,
		}
	}

	defer resp.Body.Close()

	// Parse json response
	body, _ := io.ReadAll(resp.Body)
	result := ReverseGeocodeResponse{}
	json.Unmarshal(body, &result)

	return &HttpResponse[ReverseGeocodeResponse]{
		StatusCode: resp.StatusCode,
		Response:   result,
		Raw:        resp,
	}
}

// Places returns a list of places tied to the current user
func (c Client) Places() *HttpResponse[[]Place] {
	resp, err := c.get("api/places", nil, true)
	if err != nil || resp.StatusCode != http.StatusOK {
		return &HttpResponse[[]Place]{
			StatusCode: resp.StatusCode,
			Error:      err,
			Raw:        resp,
		}
	}

	defer resp.Body.Close()

	// Parse json response
	body, _ := io.ReadAll(resp.Body)
	result := []Place{}
	json.Unmarshal(body, &result)

	return &HttpResponse[[]Place]{
		StatusCode: resp.StatusCode,
		Response:   result,
		Raw:        resp,
	}
}

// AdventureCategories returns a list of adventure categories
func (c Client) AdventureCategories() *HttpResponse[AdventureCategoriesResponse] {
	// Get data
	resp, err := c.get("api/adventures/categories", nil, true)
	if err != nil || (resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent) {
		return &HttpResponse[AdventureCategoriesResponse]{
			StatusCode: resp.StatusCode,
			Error:      err,
			Raw:        resp,
		}
	}

	defer resp.Body.Close()

	// Parse json response
	body, _ := io.ReadAll(resp.Body)
	result := AdventureCategoriesResponse{}
	json.Unmarshal(body, &result)

	return &HttpResponse[AdventureCategoriesResponse]{
		StatusCode: resp.StatusCode,
		Response:   result,
		Raw:        resp,
	}
}
