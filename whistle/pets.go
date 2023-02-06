/*
 * Produced: Sat Feb 05 2023
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

type PetsResponse struct {
	Pets []Pet `json:"pets"`
}

type Pet struct {
}

type TransfersResponse struct {
	Transfers []Transfer `json:"transfers"`
}

type Transfer struct {
}

type PetOwnersResponse struct {
	Errors []Error `json:"errors"`
}

type PetOwner struct {
}

type PetWhereaboutsResponse struct {
	Errors []Error `json:"errors"`
}

type PetLocationsRecentResponse struct {
	Errors []Error `json:"errors"`
}

type PetAchievementsResponse struct {
	Errors []Error `json:"errors"`
}

type PetStatisticsResponse struct {
	Errors []Error `json:"errors"`
}

type PetDailiesResponse struct {
	Errors []Error `json:"errors"`
}

type PetDailyResponse struct {
	Errors []Error `json:"errors"`
}

type PetDailyItemsResponse struct {
	Errors []Error `json:"errors"`
}

type PetHealthTrendsResponse struct {
	Errors []Error `json:"errors"`
}

type PetHealthGraphsResponse struct {
	Errors []Error `json:"errors"`
}

type PetNutritionPortionsResponse struct {
	Errors []Error `json:"errors"`
}

type PetFoodPortionsResponse struct {
	Errors []Error `json:"errors"`
}

type PetTaskResponse struct {
	Errors []Error `json:"errors"`
}

type PetTaskOccurrenceResponse struct {
	Errors []Error `json:"errors"`
}

// Pets returns a list of pets owned by the user.
func (c Client) Pets() *HttpResponse[PetsResponse] {
	resp, err := c.get("api/pets", nil, true)
	if err != nil || resp.StatusCode != http.StatusOK {
		return &HttpResponse[PetsResponse]{
			StatusCode: resp.StatusCode,
			Error:      err,
			Raw:        resp,
		}
	}

	defer resp.Body.Close()

	// Parse json response
	body, _ := io.ReadAll(resp.Body)
	result := PetsResponse{}
	json.Unmarshal(body, &result)

	return &HttpResponse[PetsResponse]{
		StatusCode: resp.StatusCode,
		Response:   result,
		Raw:        resp,
	}
}

// Transfers returns a list of pet transfers
func (c Client) PetTransfers() *HttpResponse[TransfersResponse] {
	resp, err := c.get("api/pets", nil, true)
	if err != nil || resp.StatusCode != http.StatusOK {
		return &HttpResponse[TransfersResponse]{
			StatusCode: resp.StatusCode,
			Error:      err,
			Raw:        resp,
		}
	}

	defer resp.Body.Close()

	// Parse json response
	body, _ := io.ReadAll(resp.Body)
	result := TransfersResponse{}
	json.Unmarshal(body, &result)

	return &HttpResponse[TransfersResponse]{
		StatusCode: resp.StatusCode,
		Response:   result,
		Raw:        resp,
	}
}

// Pet returns detailed information about a user's pet.
func (c Client) Pet(petId string) *HttpResponse[Pet] {
	resp, err := c.get("api/pets/"+petId, nil, true)
	if err != nil || resp.StatusCode != http.StatusOK {
		return &HttpResponse[Pet]{
			StatusCode: resp.StatusCode,
			Error:      err,
			Raw:        resp,
		}
	}

	defer resp.Body.Close()

	// Parse json response
	body, _ := io.ReadAll(resp.Body)
	result := Pet{}
	json.Unmarshal(body, &result)

	return &HttpResponse[Pet]{
		StatusCode: resp.StatusCode,
		Response:   result,
		Raw:        resp,
	}
}

// PetOwners returns a list of users who own a pet.
func (c Client) PetOwners(petId string) *HttpResponse[PetOwnersResponse] {
	resp, err := c.get("api/pets/"+petId+"/owners", nil, true)
	if err != nil || resp.StatusCode != http.StatusOK {
		return &HttpResponse[PetOwnersResponse]{
			StatusCode: resp.StatusCode,
			Error:      err,
			Raw:        resp,
		}
	}

	defer resp.Body.Close()

	// Parse json response
	body, _ := io.ReadAll(resp.Body)
	result := PetOwnersResponse{}
	json.Unmarshal(body, &result)

	return &HttpResponse[PetOwnersResponse]{
		StatusCode: resp.StatusCode,
		Response:   result,
		Raw:        resp,
	}
}

// PetWhereabouts returns information about a pet's current location.
func (c Client) PetWhereabouts(petId string) *HttpResponse[PetWhereaboutsResponse] {
	resp, err := c.get("api/pets/"+petId+"/whereabouts", nil, true)
	if err != nil || resp.StatusCode != http.StatusOK {
		return &HttpResponse[PetWhereaboutsResponse]{
			StatusCode: resp.StatusCode,
			Error:      err,
			Raw:        resp,
		}
	}

	defer resp.Body.Close()

	// Parse json response
	body, _ := io.ReadAll(resp.Body)
	result := PetWhereaboutsResponse{}
	json.Unmarshal(body, &result)

	return &HttpResponse[PetWhereaboutsResponse]{
		StatusCode: resp.StatusCode,
		Response:   result,
		Raw:        resp,
	}
}

// PetLocationsRecent provides a list of recent tracking locations for a pet
func (c Client) PetLocationsRecent(petId string) *HttpResponse[PetLocationsRecentResponse] {
	resp, err := c.get("api/pets/"+petId+"/locations/recent", nil, true)
	if err != nil || resp.StatusCode != http.StatusOK {
		return &HttpResponse[PetLocationsRecentResponse]{
			StatusCode: resp.StatusCode,
			Error:      err,
			Raw:        resp,
		}
	}

	defer resp.Body.Close()

	// Parse json response
	body, _ := io.ReadAll(resp.Body)
	result := PetLocationsRecentResponse{}
	json.Unmarshal(body, &result)

	return &HttpResponse[PetLocationsRecentResponse]{
		StatusCode: resp.StatusCode,
		Response:   result,
		Raw:        resp,
	}
}

// PetAchievements returns a list of achievements for a pet.
func (c Client) PetAchievements(petId string) *HttpResponse[PetAchievementsResponse] {
	resp, err := c.get("api/pets/"+petId+"/achievements", nil, true)
	if err != nil || resp.StatusCode != http.StatusOK {
		return &HttpResponse[PetAchievementsResponse]{
			StatusCode: resp.StatusCode,
			Error:      err,
			Raw:        resp,
		}
	}

	defer resp.Body.Close()

	// Parse json response
	body, _ := io.ReadAll(resp.Body)
	result := PetAchievementsResponse{}
	json.Unmarshal(body, &result)

	return &HttpResponse[PetAchievementsResponse]{
		StatusCode: resp.StatusCode,
		Response:   result,
		Raw:        resp,
	}
}

// PetStatistics returns statistics statistical insights about a pet.
func (c Client) PetStatistics(petId string) *HttpResponse[PetStatisticsResponse] {
	resp, err := c.get("api/pets/"+petId+"/stats", nil, true)
	if err != nil || resp.StatusCode != http.StatusOK {
		return &HttpResponse[PetStatisticsResponse]{
			StatusCode: resp.StatusCode,
			Error:      err,
			Raw:        resp,
		}
	}

	defer resp.Body.Close()

	// Parse json response
	body, _ := io.ReadAll(resp.Body)
	result := PetStatisticsResponse{}
	json.Unmarshal(body, &result)

	return &HttpResponse[PetStatisticsResponse]{
		StatusCode: resp.StatusCode,
		Response:   result,
		Raw:        resp,
	}
}

// PetDailies returns a list of daily activities for a pet.
func (c Client) PetDailies(petId string) *HttpResponse[PetDailiesResponse] {
	resp, err := c.get("api/pets/"+petId+"/dailies", nil, true)
	if err != nil || resp.StatusCode != http.StatusOK {
		return &HttpResponse[PetDailiesResponse]{
			StatusCode: resp.StatusCode,
			Error:      err,
			Raw:        resp,
		}
	}

	defer resp.Body.Close()

	// Parse json response
	body, _ := io.ReadAll(resp.Body)
	result := PetDailiesResponse{}
	json.Unmarshal(body, &result)

	return &HttpResponse[PetDailiesResponse]{
		StatusCode: resp.StatusCode,
		Response:   result,
		Raw:        resp,
	}
}

// PetDaily returns information about a pet's daily activity on the specified day.
func (c Client) PetDaily(petId string, dailyId string) *HttpResponse[PetDailyResponse] {
	resp, err := c.get("api/pets/"+petId+"/dailies/"+dailyId, nil, true)
	if err != nil || resp.StatusCode != http.StatusOK {
		return &HttpResponse[PetDailyResponse]{
			StatusCode: resp.StatusCode,
			Error:      err,
			Raw:        resp,
		}
	}

	defer resp.Body.Close()

	// Parse json response
	body, _ := io.ReadAll(resp.Body)
	result := PetDailyResponse{}
	json.Unmarshal(body, &result)

	return &HttpResponse[PetDailyResponse]{
		StatusCode: resp.StatusCode,
		Response:   result,
		Raw:        resp,
	}
}

// PetDailyItems returns a item breakdown of a pet's daily activity on the specified day.
func (c Client) PetDailyItems(petId string, dailyId string) *HttpResponse[PetDailyItemsResponse] {
	resp, err := c.get("api/pets/"+petId+"/dailies/"+dailyId+"/daily_items", nil, true)
	if err != nil || resp.StatusCode != http.StatusOK {
		return &HttpResponse[PetDailyItemsResponse]{
			StatusCode: resp.StatusCode,
			Error:      err,
			Raw:        resp,
		}
	}

	defer resp.Body.Close()

	// Parse json response
	body, _ := io.ReadAll(resp.Body)
	result := PetDailyItemsResponse{}
	json.Unmarshal(body, &result)

	return &HttpResponse[PetDailyItemsResponse]{
		StatusCode: resp.StatusCode,
		Response:   result,
		Raw:        resp,
	}
}

// PetHealthTrends returns health trend information about a pet.
func (c Client) PetHealthTrends(petId string) *HttpResponse[PetHealthTrendsResponse] {
	resp, err := c.get("api/pets/"+petId+"/health/trends", nil, true)
	if err != nil || resp.StatusCode != http.StatusOK {
		return &HttpResponse[PetHealthTrendsResponse]{
			StatusCode: resp.StatusCode,
			Error:      err,
			Raw:        resp,
		}
	}

	defer resp.Body.Close()

	// Parse json response
	body, _ := io.ReadAll(resp.Body)
	result := PetHealthTrendsResponse{}
	json.Unmarshal(body, &result)

	return &HttpResponse[PetHealthTrendsResponse]{
		StatusCode: resp.StatusCode,
		Response:   result,
		Raw:        resp,
	}
}

// PetHealthGraphs returns graphical information about a pet's health based on the specified trend
func (c Client) PetHealthGraphs(petId string, trend string, days int) *HttpResponse[PetHealthGraphsResponse] {
	resp, err := c.get(fmt.Sprintf("api/pets/%s/health/graphs/%s?num_of_days=%d", petId, trend, days), nil, true)
	if err != nil || resp.StatusCode != http.StatusOK {
		return &HttpResponse[PetHealthGraphsResponse]{
			StatusCode: resp.StatusCode,
			Error:      err,
			Raw:        resp,
		}
	}

	defer resp.Body.Close()

	// Parse json response
	body, _ := io.ReadAll(resp.Body)
	result := PetHealthGraphsResponse{}
	json.Unmarshal(body, &result)

	return &HttpResponse[PetHealthGraphsResponse]{
		StatusCode: resp.StatusCode,
		Response:   result,
		Raw:        resp,
	}
}

// PetNutritionPortions returns information about suggested food portions for a pet.
func (c Client) PetNutritionPortions(petId string) *HttpResponse[PetNutritionPortionsResponse] {
	resp, err := c.get("api/pets/"+petId+"/nutrition/v2/suggested_portions", nil, true)
	if err != nil || resp.StatusCode != http.StatusOK {
		return &HttpResponse[PetNutritionPortionsResponse]{
			StatusCode: resp.StatusCode,
			Error:      err,
			Raw:        resp,
		}
	}

	defer resp.Body.Close()

	// Parse json response
	body, _ := io.ReadAll(resp.Body)
	result := PetNutritionPortionsResponse{}
	json.Unmarshal(body, &result)

	return &HttpResponse[PetNutritionPortionsResponse]{
		StatusCode: resp.StatusCode,
		Response:   result,
		Raw:        resp,
	}
}

// PetFoodPortions returns information about food portions for a pet.
//
// Deprecated: Use PetNutritionPortions instead
func (c Client) PetFoodPortions(petId string) *HttpResponse[PetFoodPortionsResponse] {
	resp, err := c.get("api/pets/"+petId+"/pet_food_portions", nil, true)
	if err != nil || resp.StatusCode != http.StatusOK {
		return &HttpResponse[PetFoodPortionsResponse]{
			StatusCode: resp.StatusCode,
			Error:      err,
			Raw:        resp,
		}
	}

	defer resp.Body.Close()

	// Parse json response
	body, _ := io.ReadAll(resp.Body)
	result := PetFoodPortionsResponse{}
	json.Unmarshal(body, &result)

	return &HttpResponse[PetFoodPortionsResponse]{
		StatusCode: resp.StatusCode,
		Response:   result,
		Raw:        resp,
	}
}

// PetTask returns detailed information about the specified task for a pet.
func (c Client) PetTask(petId string, taskId string) *HttpResponse[PetTaskResponse] {
	resp, err := c.get("api/pets/"+petId+"/tasks/"+taskId, nil, true)
	if err != nil || resp.StatusCode != http.StatusOK {
		return &HttpResponse[PetTaskResponse]{
			StatusCode: resp.StatusCode,
			Error:      err,
			Raw:        resp,
		}
	}

	defer resp.Body.Close()

	// Parse json response
	body, _ := io.ReadAll(resp.Body)
	result := PetTaskResponse{}
	json.Unmarshal(body, &result)

	return &HttpResponse[PetTaskResponse]{
		StatusCode: resp.StatusCode,
		Response:   result,
		Raw:        resp,
	}
}

// PetTaskOccurrence returns information about the occurrence type (e.g. incomplete)
func (c Client) PetTaskOccurrence(petId string, occurrenceType string) *HttpResponse[PetTaskOccurrenceResponse] {
	resp, err := c.get("api/pets/"+petId+"/task_occurrences/?type="+occurrenceType, nil, true)
	if err != nil || resp.StatusCode != http.StatusOK {
		return &HttpResponse[PetTaskOccurrenceResponse]{
			StatusCode: resp.StatusCode,
			Error:      err,
			Raw:        resp,
		}
	}

	defer resp.Body.Close()

	// Parse json response
	body, _ := io.ReadAll(resp.Body)
	result := PetTaskOccurrenceResponse{}
	json.Unmarshal(body, &result)

	return &HttpResponse[PetTaskOccurrenceResponse]{
		StatusCode: resp.StatusCode,
		Response:   result,
		Raw:        resp,
	}
}
