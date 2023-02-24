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

package whistle_test

import (
	"net/http"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestNotifications(t *testing.T) {
	t.Skip("Skipping test due to API changes")

	// t.Parallel()

	// resp := c.Notifications()

	// assert.Equal(t, http.StatusOK, resp.StatusCode)
	// assert.Equal(t, nil, resp.Error)
	// assert.NotEqual(t, resp.Response, nil)
	// assert.NotEqual(t, resp.Response.Items, nil)
}

func TestPetFoodsDogFood(t *testing.T) {
	t.Parallel()

	resp := c.PetFoods("dog_food")

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, nil, resp.Error)
	assert.NotEqual(t, resp.Response, nil)

	if len(resp.Response) <= 0 {
		t.Errorf("Expected at least one food, got %d", len(resp.Response))
	}
	if (resp.Response)[0].ID == 0 {
		t.Errorf("Expected valid food ID, got %d", (resp.Response)[0].ID)
	}
	if (resp.Response)[0].Name == "" {
		t.Errorf("Expected valid food name, got %s", (resp.Response)[0].Name)
	}
}

func TestPetFoodsDogTreat(t *testing.T) {
	t.Parallel()

	resp := c.PetFoods("dog_treat")

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, nil, resp.Error)
	assert.NotEqual(t, resp.Response, nil)

	if len(resp.Response) <= 0 {
		t.Errorf("Expected at least one treat, got %d", len(resp.Response))
	}
	if (resp.Response)[0].ID == 0 {
		t.Errorf("Expected valid food ID, got %d", (resp.Response)[0].ID)
	}
	if (resp.Response)[0].Name == "" {
		t.Errorf("Expected valid food name, got %s", (resp.Response)[0].Name)
	}
}

func TestPetFoodsInvalidType(t *testing.T) {
	t.Parallel()

	resp := c.PetFoods("rhino_treat")

	assert.Equal(t, nil, resp.Error)
	assert.Equal(t, http.StatusUnprocessableEntity, resp.StatusCode)
}

func TestReverseGeocode(t *testing.T) {
	t.Parallel()

	// https://www.google.com/maps/place/37%C2%B046'06.9%22N+92%C2%B017'10.5%22W
	resp := c.ReverseGeocode("37.768578", "-92.286243")

	assert.Equal(t, nil, resp.Error)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.NotEqual(t, resp.Response.Description.Address, nil) // Exact address is not guaranteed
	assert.Equal(t, resp.Response.Description.City, "Laquey")
	assert.Equal(t, resp.Response.Description.Region, "Missouri")
	assert.Equal(t, resp.Response.Description.Country, "United States")
}

func TestPlaces(t *testing.T) {
	t.Parallel()

	resp := c.Places()

	assert.Equal(t, nil, resp.Error)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.NotEqual(t, len(resp.Response[0].PetIds), 0)
	if len(resp.Response) <= 0 {
		t.Errorf("Expected at least one place, got %d", len(resp.Response))
	}
	if len(resp.Response[0].Outline) <= 0 {
		t.Skip("Expected at least one outline bound set, got none")
	}
}

func TestAdventureCategories(t *testing.T) {
	t.Parallel()

	resp := c.AdventureCategories()

	assert.Equal(t, nil, resp.Error)
	assert.Equal(t, http.StatusNoContent, resp.StatusCode)
}
