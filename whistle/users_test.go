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

	"github.com/amattu2/go-whistle-wrapper/utils"
	"github.com/amattu2/go-whistle-wrapper/whistle"
	"github.com/go-playground/assert/v2"
)

func TestUsers(t *testing.T) {
	t.Parallel()

	c := whistle.InitializeBearer(utils.GetEnv("WHISTLE_BEARER", ""))

	resp := c.Users()

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, resp.Error, nil)
	assert.NotEqual(t, resp.Response, nil)
	assert.Equal(t, resp.Response.CurrentUser, true)
}

func TestMe(t *testing.T) {
	t.Parallel()

	c := whistle.InitializeBearer(utils.GetEnv("WHISTLE_BEARER", ""))

	resp := c.Me()

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, resp.Error, nil)
	assert.NotEqual(t, resp.Response, nil)
	assert.NotEqual(t, resp.Response.User, nil)
	assert.Equal(t, resp.Response.User.CurrentUser, true)
}

func TestCheckEmailExisting(t *testing.T) {
	t.Parallel()

	c := whistle.InitializeBearer(utils.GetEnv("WHISTLE_BEARER", ""))

	resp := c.CheckEmail("admin@whistle.com")

	assert.Equal(t, http.StatusNoContent, resp.StatusCode) // Email exists
	assert.Equal(t, resp.Error, nil)
	assert.Equal(t, resp.Response, true)
}

func TestCheckEmailNonExisting(t *testing.T) {
	t.Parallel()

	c := whistle.InitializeBearer(utils.GetEnv("WHISTLE_BEARER", ""))

	resp := c.CheckEmail("thisuserwillneverexisthopefully19283201@whistle.com")

	assert.Equal(t, http.StatusNotFound, resp.StatusCode) // Email does not exist
	assert.Equal(t, resp.Error, nil)
	assert.Equal(t, resp.Response, false)
}

func TestInvitationCodes(t *testing.T) {
	t.Skip("TBD: No valid invitation codes known")
}

func TestApplicationState(t *testing.T) {
	t.Parallel()

	c := whistle.InitializeBearer(utils.GetEnv("WHISTLE_BEARER", ""))

	resp := c.ApplicationState()

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, resp.Error, nil)
	assert.NotEqual(t, resp.Response, nil)
	assert.NotEqual(t, resp.Response.ApplicationState, nil)

	t.Skip("TBD: No known states to test against")
}

func TestCreditCard(t *testing.T) {
	t.Parallel()

	c := whistle.InitializeBearer(utils.GetEnv("WHISTLE_BEARER", ""))

	resp := c.CreditCard()

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, resp.Error, nil)
	assert.NotEqual(t, resp.Response.LastFour, nil)

	t.Skip("TBD: No credit cards tied to account to test")
}

func TestSubscriptions(t *testing.T) {
	t.Parallel()

	c := whistle.InitializeBearer(utils.GetEnv("WHISTLE_BEARER", ""))

	resp := c.Subscriptions()

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, resp.Error, nil)
	assert.NotEqual(t, resp.Response, nil)
	assert.NotEqual(t, resp.Response.Subscriptions, nil)
	assert.NotEqual(t, resp.Response.PartnerServices, nil)

	t.Skip("TBD: No subscriptions tied to account to test")
}
