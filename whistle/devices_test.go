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
	"github.com/go-playground/assert/v2"
)

func TestDevice(t *testing.T) {
	t.Parallel()

	// Get data
	resp := c.Device(utils.GetEnv("WHISTLE_DEVICE_ID", ""))
	device := resp.Response.Device

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, resp.Error, nil)
	assert.NotEqual(t, resp, nil)
	assert.Equal(t, device.SerialNumber, utils.GetEnv("WHISTLE_DEVICE_ID", ""))
	assert.Equal(t, device.ModelId, "W05")
}

func TestDeviceActivation(t *testing.T) {
	t.Parallel()

	// Get data
	resp := c.DeviceActivationCheck(utils.GetEnv("WHISTLE_DEVICE_ID", ""))

	assert.Equal(t, http.StatusUnprocessableEntity, resp.StatusCode)
	assert.Equal(t, resp.Error, nil)
}

func TestDevicePlans(t *testing.T) {
	t.Skip("Cannot test due to dependence on changing states")
}

func TestDeviceSubscription(t *testing.T) {
	t.Skip("Cannot test due to dependence on changing states")
}

func TestDeviceSubscriptionPreview(t *testing.T) {
	t.Skip("Cannot test due to dependence on changing states")
}

func TestDeviceUpgradePreview(t *testing.T) {
	t.Skip("Cannot test due to dependence on changing states")
}

func TestDeviceWifiNetworks(t *testing.T) {
	t.Skip("Cannot test due to dependence on changing states")
}
