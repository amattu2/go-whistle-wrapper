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
	"time"
)

type PetsResponse struct {
	Pets []Pet `json:"pets"`
}

type PetResponse struct {
	Pet Pet `json:"pet"`
}

type Pet struct {
	ID                   int               `json:"id"`
	Name                 string            `json:"name"`
	ProfilePhotoUrlSizes map[string]string `json:"profile_photo_url_sizes"`
	RealtimeChannel      RealtimeChannel   `json:"realtime_channel"`
	SubscriptionStatus   string            `json:"subscription_status"`
	PartnerServiceStatus string            `json:"partner_service_status"`
	Device               Device            `json:"device"`
	ActivitySummary      ActivitySummary   `json:"activity_summary"`
	LastLocation         Location          `json:"last_location"`
	Profile              PetProfile        `json:"profile"`
}

type ActivitySummary struct {
	ActiveSummaryStartDate      string       `json:"active_summary_start_date"`
	ActivityEnabled             bool         `json:"activity_enabled"`
	CurrentStreak               int          `json:"current_streak"`
	CurrentMinutesActive        int          `json:"current_minutes_active"`
	CurrentMinutesRest          int          `json:"current_minutes_rest"`
	SimilarPetsMinutesActive    float64      `json:"similar_pets_minutes_active"`
	SimilarPetsMinutesRest      float64      `json:"similar_pets_minutes_rest"`
	SuggestedActivityRangeLower float64      `json:"suggested_activity_range_lower"`
	SuggestedActivityRangeUpper float64      `json:"suggested_activity_range_upper"`
	CurrentActivityGoal         ActivityGoal `json:"current_activity_goal"`
	UpcomingActivityGoal        ActivityGoal `json:"upcoming_activity_goal"`
}

type ActivityGoal struct {
	Minutes   int           `json:"minutes"`
	StartedAt string        `json:"started_at"`
	TimeZone  time.Location `json:"time_zone"`
}

type PetProfile struct {
	Breed                      Breed   `json:"breed"`
	DateOfBirth                string  `json:"date_of_birth"`
	AgeInMonths                int     `json:"age_in_months"`
	AgeInYears                 int     `json:"age_in_years"`
	TimeZoneName               string  `json:"time_zone_name"`
	Weight                     float64 `json:"weight"`
	WeightType                 string  `json:"weight_type"`
	Species                    string  `json:"species"`
	OverdueTaskOccurrenceCount int     `json:"overdue_task_occurrence_count"`
	IsFixed                    bool    `json:"is_fixed"`
	BodyConditionScore         float64 `json:"body_condition_score"`
	PetFood                    PetFood `json:"pet_food"`
}

type TransfersResponse struct {
	Transfers []TransferPet `json:"transfers"`
}

type TransferPet struct {
	Pet Pet `json:"pet"`
}

type PetOwnersResponse struct {
	Errors []Error    `json:"errors"`
	Owners []PetOwner `json:"owners"`
}

type PetOwner struct {
	ID                   int               `json:"id"`
	FirstName            string            `json:"first_name"`
	LastName             string            `json:"last_name"`
	CurrentUser          bool              `json:"current_user"`
	Searchable           bool              `json:"searchable"`
	ProfilePhotoUrlSizes map[string]string `json:"profile_photo_url_sizes"`
	Email                string            `json:"email"`
}

type PetWhereaboutsResponse struct {
	Errors    []Error    `json:"errors"`
	Locations []Location `json:"locations"`
	Places    []Place    `json:"places"`
}

type Location struct {
	Latitude          float64 `json:"latitude"`
	Longitude         float64 `json:"longitude"`
	Timestamp         string  `json:"timestamp"`
	UncertaintyMeters float64 `json:"uncertainty_meters"`
	Reason            string  `json:"reason"`

	// Does not exist on all responses.
	Place       Place              `json:"place"`
	Description GeocodeDescription `json:"description"`
}

type PetLocationsRecentResponse struct {
	Errors    []Error    `json:"errors"`
	Locations []Location `json:"locations"`
}

type PetAchievementsResponse struct {
	Errors       []Error          `json:"errors"`
	Achievements []PetAchievement `json:"achievements"`
}

type PetAchievement struct {
	ID                  int               `json:"id"`
	EarnedAchievementId int               `json:"earned_achievement_id"`
	Actionable          string            `json:"actionable"`
	Title               string            `json:"title"`
	ShortName           string            `json:"short_name"`
	Description         string            `json:"description"`
	Type                string            `json:"type"`
	BackgroundColor     string            `json:"background_color"`
	StrokeColor         string            `json:"stroke_color"`
	BadgeImages         map[string]string `json:"badge_images"`
	TemplateType        string            `json:"template_type"`
	TemplateProperties  map[string]string `json:"template_properties"`
	Earned              bool              `json:"earned"`
	EarnedTimestamp     string            `json:"earned_timestamp"`
	TypeProperties      map[string]string `json:"type_properties"`
}

type PetStatisticsResponse struct {
	Errors     []Error         `json:"errors"`
	Statistics []PetStatistics `json:"stats"`
}

type PetStatistics struct {
	AverageMinutesActive float64 `json:"average_minutes_active"`
	AverageMinutesRest   float64 `json:"average_minutes_rest"`
	AverageCalories      float64 `json:"average_calories"`
	AverageDistance      float64 `json:"average_distance"`
	DistanceUnit         string  `json:"distance_units"`
	CurrentStreak        int     `json:"current_streak"`
	LongestStreak        int     `json:"longest_streak"`
	MostActiveDay        any     `json:"most_active_day"` // Todo: Validate this type
}

type PetDailiesResponse struct {
	Errors  []Error `json:"errors"`
	Dailies []Daily `json:"dailies"`
}

type PetDailyResponse struct {
	Errors []Error `json:"errors"`
	Daily  Daily   `json:"daily"`
}

type Daily struct {
	ActivityGoal  int     `json:"activity_goal"`
	DayNumber     int     `json:"day_number"`
	Excluded      bool    `json:"excluded"`
	MinutesActive int     `json:"minutes_active"`
	MinutesRest   int     `json:"minutes_rest"`
	Calories      float64 `json:"calories"`
	Distance      float64 `json:"distance"`
	DistanceUnits string  `json:"distance_units"`
	Timestamp     string  `json:"timestamp"`
	UpdatedAt     string  `json:"updated_at"`

	// Only present in PetDailyResponse
	BarChart18Min   []int  `json:"bar_chart_18_min"`
	BarChart3Min    []int  `json:"bar_chart_3_min"`
	HourlyActivity  []int  `json:"hourly_activity"`
	CurrentStreak   int    `json:"current_streak"`
	StreakDayNumber int    `json:"streak_day_number"`
	Date            string `json:"date"`
	LastUpdatedAt   string `json:"last_updated_at"`
}

type PetDailyItemsResponse struct {
	Errors     []Error     `json:"errors"`
	DailyItems []DailyItem `json:"daily_items"`
}

type DailyItem struct {
	Type      string          `json:"type"`
	Title     string          `json:"title"`
	Data      []DailyItemData `json:"data"`
	StartTime string          `json:"start_time"`
	EndTime   string          `json:"end_time"`
	TimeZone  time.Location   `json:"time_zone"`
}

type DailyItemData struct {
	ID                 int      `json:"id"`
	Category           string   `json:"category"`
	MinActivity        int      `json:"min_activity"`
	MinRest            int      `json:"min_rest"`
	Calories           float64  `json:"calories"`
	Distance           float64  `json:"distance"`
	DistanceUnits      string   `json:"distance_units"`
	OverrideEventTypes []string `json:"override_event_types"`
	StaticMapUrl       string   `json:"static_map_url"`
}

type PetHealthTrendsResponse struct {
	Errors       []Error       `json:"errors"`
	PetId        int           `json:"pet_id"`
	HealthReport string        `json:"health_report"`
	LastUpdated  string        `json:"last_updated"`
	Trends       []HealthTrend `json:"trends"`
}

type HealthTrend struct {
	Type             string              `json:"type"`
	Title            string              `json:"title"`
	Status           string              `json:"status"`
	Metrics          []any               `json:"metrics"` // TBD describe this type
	StatusThresholds []map[string]string `json:"status_thresholds"`
}

type PetHealthGraphsResponse struct {
	Errors           []Error             `json:"errors"`
	PetId            int                 `json:"pet_id"`
	StartDate        string              `json:"start_date"`
	NumOfDays        int                 `json:"num_of_days"`
	Score            any                 `json:"score"` // TBD describe this type
	Unit             string              `json:"unit"`
	Status           string              `json:"status"`
	Data             []any               `json:"data"` // TBD describe this type
	StatusThresholds []map[string]string `json:"status_thresholds"`
}

type PetNutritionPortionsResponse struct {
	Errors               []Error `json:"errors"`
	PetFoodPortions      []any   `json:"pet_food_portions"` // TBD describe this type
	SuggestedCalories    float64 `json:"suggested_calories"`
	AverageCalories      float64 `json:"average_calories"`
	AverageMinutesActive float64 `json:"average_minutes_active"`
	Treats               []any   `json:"treats"` // TBD describe this type
}

type PetFoodPortionsResponse struct {
	Errors          []Error `json:"errors"`
	PetFoodPortions []any   `json:"pet_food_portions"` // TBD describe this type
}

type PetTaskResponse struct {
	Errors []Error `json:"errors"`
}

type PetTaskOccurrenceResponse struct {
	Errors          []Error   `json:"errors"`
	PetId           int       `json:"pet_id"`
	TaskOccurrences []PetTask `json:"task_occurrences"`
}

type PetTask struct {
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
	resp, err := c.get("api/pets/transfers", nil, true)
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
func (c Client) Pet(petId string) *HttpResponse[PetResponse] {
	resp, err := c.get("api/pets/"+petId, nil, true)
	if err != nil || resp.StatusCode != http.StatusOK {
		return &HttpResponse[PetResponse]{
			StatusCode: resp.StatusCode,
			Error:      err,
			Raw:        resp,
		}
	}

	defer resp.Body.Close()

	// Parse json response
	body, _ := io.ReadAll(resp.Body)
	result := PetResponse{}
	json.Unmarshal(body, &result)

	return &HttpResponse[PetResponse]{
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

// PetWhereabouts returns information about a pet's location history.
func (c Client) PetWhereabouts(petId string, startDate string, endDate string) *HttpResponse[PetWhereaboutsResponse] {
	resp, err := c.get(fmt.Sprintf("api/pets/%s/whereabouts?start_time=%s&end_time=%s", petId, startDate, endDate), nil, true)
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
	resp, err := c.get("api/pets/"+petId+"/locations/recent_trackings", nil, true)
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
