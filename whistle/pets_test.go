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

package whistle_test

import (
	"net/http"
	"testing"

	"github.com/amattu2/go-whistle-wrapper/utils"
	"github.com/amattu2/go-whistle-wrapper/whistle"
	"github.com/go-playground/assert/v2"
)

func TestPets(t *testing.T) {
	t.Parallel()

	c := whistle.InitializeBearer(utils.GetEnv("WHISTLE_BEARER", ""))

	r := c.Pets()

	assert.Equal(t, http.StatusOK, r.StatusCode)

	if len(r.Response.Pets) <= 0 {
		t.Error("Expected at least one pet, got 0")
	}
	if r.Response.Pets[0].ID == 0 {
		t.Error("Expected pet ID to be greater than 0, got 0")
	}

	assert.NotEqual(t, "", r.Response.Pets[0].Name)
}

func TestTransfers(t *testing.T) {
	t.Parallel()

	c := whistle.InitializeBearer(utils.GetEnv("WHISTLE_BEARER", ""))

	r := c.PetTransfers()

	assert.Equal(t, http.StatusOK, r.StatusCode)
	assert.NotEqual(t, nil, r.Response.Transfers)

	if len(r.Response.Transfers) <= 0 {
		t.Error("Expected at least one pets, got 0")
	}
	if r.Response.Transfers[0].Pet.ID == 0 {
		t.Error("Expected pet ID to be greater than 0, got 0")
	}

	assert.NotEqual(t, nil, r.Response.Transfers[0].Pet.Name)
}

func TestPet(t *testing.T) {
	t.Parallel()

	c := whistle.InitializeBearer(utils.GetEnv("WHISTLE_BEARER", ""))

	r := c.Pet(utils.GetEnv("WHISTLE_PET_ID", ""))

	assert.Equal(t, http.StatusOK, r.StatusCode)
	assert.NotEqual(t, 0, r.Response.Pet.ID)
	assert.NotEqual(t, "", r.Response.Pet.Name)
}

func TestPetOwners(t *testing.T) {
	t.Parallel()

	c := whistle.InitializeBearer(utils.GetEnv("WHISTLE_BEARER", ""))

	r := c.PetOwners(utils.GetEnv("WHISTLE_PET_ID", ""))

	assert.Equal(t, http.StatusOK, r.StatusCode)

	if len(r.Response.Owners) <= 0 {
		t.Error("Expected at least one owner, got 0")
	}

	assert.NotEqual(t, 0, r.Response.Owners[0].ID)
	assert.NotEqual(t, "", r.Response.Owners[0].FirstName)
	assert.NotEqual(t, "", r.Response.Owners[0].LastName)
}

func TestPetWhereabouts(t *testing.T) {
	t.Parallel()

	c := whistle.InitializeBearer(utils.GetEnv("WHISTLE_BEARER", ""))

	r := c.PetWhereabouts(utils.GetEnv("WHISTLE_PET_ID", ""), "2022-12-01", "2023-03-03")

	assert.Equal(t, http.StatusOK, r.StatusCode)

	if len(r.Response.Locations) <= 0 {
		t.Error("Expected at least one location, got 0")
	}
	if len(r.Response.Places) <= 0 {
		t.Error("Expected at least one place, got 0")
	}

	assert.NotEqual(t, "", r.Response.Locations[0].Reason)
	assert.NotEqual(t, "", r.Response.Places[0].Address)
	assert.NotEqual(t, "", r.Response.Places[0].Name)
	assert.NotEqual(t, 0, r.Response.Places[0].ID)
}

func TestPetLocationsRecent(t *testing.T) {
	t.Parallel()

	c := whistle.InitializeBearer(utils.GetEnv("WHISTLE_BEARER", ""))

	r := c.PetLocationsRecent(utils.GetEnv("WHISTLE_PET_ID", ""))

	assert.Equal(t, http.StatusOK, r.StatusCode)

	if len(r.Response.Locations) <= 0 {
		t.Skip("Expected at least one location, got 0")
	}

	assert.NotEqual(t, "", r.Response.Locations[0].Reason)
	assert.NotEqual(t, nil, r.Response.Locations[0].Latitude)
	assert.NotEqual(t, nil, r.Response.Locations[0].Longitude)
}

func TestPetAchievements(t *testing.T) {
	t.Parallel()

	c := whistle.InitializeBearer(utils.GetEnv("WHISTLE_BEARER", ""))

	r := c.PetAchievements(utils.GetEnv("WHISTLE_PET_ID", ""))

	assert.Equal(t, http.StatusOK, r.StatusCode)

	if len(r.Response.Achievements) <= 0 {
		t.Error("Expected at least one achievement, got 0")
	}

	assert.NotEqual(t, 0, r.Response.Achievements[0].ID)
	assert.NotEqual(t, "nil", r.Response.Achievements[0].Title)
	assert.NotEqual(t, "", r.Response.Achievements[0].Description)
}

func TestPetStatistics(t *testing.T) {
	t.Skip("Cannot test due to dependence on changing states")
}

func TestPetDailies(t *testing.T) {
	t.Skip("Cannot test due to dependence on changing states")
}

func TestPetDaily(t *testing.T) {
	t.Skip("Cannot test due to dependence on changing states")
}

func TestPetDailyItems(t *testing.T) {
	t.Skip("Cannot test due to dependence on changing states")
}

func TestPetHealthTrends(t *testing.T) {
	t.Parallel()

	c := whistle.InitializeBearer(utils.GetEnv("WHISTLE_BEARER", ""))

	r := c.PetHealthTrends(utils.GetEnv("WHISTLE_PET_ID", ""))

	assert.Equal(t, http.StatusOK, r.StatusCode)

	if len(r.Response.Trends) <= 0 {
		t.Error("Expected at least one trend, got 0")
	}
	if len(r.Response.Trends[0].StatusThresholds) <= 0 {
		t.Error("Expected at least one trend status threshold, got 0")
	}

	assert.NotEqual(t, "", r.Response.Trends[0].Type)
}

func TestPetHealthGraphs(t *testing.T) {
	t.Skip("Cannot test due to dependence on changing states")
}

func TestPetNutritionPortions(t *testing.T) {
	t.Skip("Cannot test due to dependence on changing states")
}

func TestPetFoodPortions(t *testing.T) {
	t.Skip("Cannot test due to dependence on changing states")
}

func TestPetTask(t *testing.T) {
	t.Skip("TBD. No taskId is currently known.")
}

func TestPetTaskOccurrence(t *testing.T) {
	t.Skip("TBD. No taskId is currently known.")
}
