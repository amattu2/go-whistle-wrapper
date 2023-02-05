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
}

type AdventureCategoriesResponse struct {
}

type PetFood struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Notification struct {
}

type Place struct {
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

// Todo: Figure out what arguments are required
func (c Client) ReverseGeocode() *HttpResponse[ReverseGeocodeResponse] {
	resp, err := c.get("api/reverse_geocode", nil, true)
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
