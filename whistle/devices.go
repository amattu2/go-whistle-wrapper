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

type DeviceResponse struct {
	Error string `json:"error"`
}

type DeviceActivationResponse struct {
	Error string `json:"error"`
}

type DevicePlansResponse struct {
	Error string `json:"error"`
}

type DeviceSubscriptionResponse struct {
	Error string `json:"error"`
}

type DeviceSubscriptionPreviewResponse struct {
	Error string `json:"error"`
}

type DeviceUpgradePreviewResponse struct {
	Error string `json:"error"`
}

// Get detailed information about a device by deviceId
func (c Client) Device(deviceId string) *HttpResponse[DeviceResponse] {
	resp, err := c.get(fmt.Sprintf("api/devices/%s", deviceId), nil, true)
	if err != nil || resp.StatusCode != http.StatusOK {
		return &HttpResponse[DeviceResponse]{
			StatusCode: resp.StatusCode,
			Error:      err,
			Raw:        resp,
		}
	}

	defer resp.Body.Close()

	// Parse json response
	body, _ := io.ReadAll(resp.Body)
	result := DeviceResponse{}
	json.Unmarshal(body, &result)

	return &HttpResponse[DeviceResponse]{
		StatusCode: resp.StatusCode,
		Response:   result,
		Raw:        resp,
	}
}

// Get detailed information about device activation status by deviceId
func (c Client) DeviceActivation(deviceId string) *HttpResponse[DeviceActivationResponse] {
	resp, err := c.get(fmt.Sprintf("api/devices/%s/activation", deviceId), nil, true)
	if err != nil || resp.StatusCode != http.StatusOK {
		return &HttpResponse[DeviceActivationResponse]{
			StatusCode: resp.StatusCode,
			Error:      err,
			Raw:        resp,
		}
	}

	defer resp.Body.Close()

	// Parse json response
	body, _ := io.ReadAll(resp.Body)
	result := DeviceActivationResponse{}
	json.Unmarshal(body, &result)

	return &HttpResponse[DeviceActivationResponse]{
		StatusCode: resp.StatusCode,
		Response:   result,
		Raw:        resp,
	}
}

// Get available plans for a device by deviceId
func (c Client) DevicePlans(deviceId string) *HttpResponse[DevicePlansResponse] {
	resp, err := c.get(fmt.Sprintf("api/devices/%s/plans", deviceId), nil, true)
	if err != nil || resp.StatusCode != http.StatusOK {
		return &HttpResponse[DevicePlansResponse]{
			StatusCode: resp.StatusCode,
			Error:      err,
			Raw:        resp,
		}
	}

	defer resp.Body.Close()

	// Parse json response
	body, _ := io.ReadAll(resp.Body)
	result := DevicePlansResponse{}
	json.Unmarshal(body, &result)

	return &HttpResponse[DevicePlansResponse]{
		StatusCode: resp.StatusCode,
		Response:   result,
		Raw:        resp,
	}
}

// Get detailed information about device subscription by deviceId
func (c Client) DeviceSubscription(deviceId string) *HttpResponse[DeviceSubscriptionResponse] {
	resp, err := c.get(fmt.Sprintf("api/devices/%s/subscription", deviceId), nil, true)
	if err != nil || resp.StatusCode != http.StatusOK {
		return &HttpResponse[DeviceSubscriptionResponse]{
			StatusCode: resp.StatusCode,
			Error:      err,
			Raw:        resp,
		}
	}

	defer resp.Body.Close()

	// Parse json response
	body, _ := io.ReadAll(resp.Body)
	result := DeviceSubscriptionResponse{}
	json.Unmarshal(body, &result)

	return &HttpResponse[DeviceSubscriptionResponse]{
		StatusCode: resp.StatusCode,
		Response:   result,
		Raw:        resp,
	}
}

// Get information about device subscription renewal by deviceId and planId
func (c Client) DeviceSubscriptionPreview(deviceId string, planId string) *HttpResponse[DeviceSubscriptionPreviewResponse] {
	resp, err := c.get(fmt.Sprintf("api/devices/%s/subscription/previews/%s", deviceId, planId), nil, true)
	if err != nil || resp.StatusCode != http.StatusOK {
		return &HttpResponse[DeviceSubscriptionPreviewResponse]{
			StatusCode: resp.StatusCode,
			Error:      err,
			Raw:        resp,
		}
	}

	defer resp.Body.Close()

	// Parse json response
	body, _ := io.ReadAll(resp.Body)
	result := DeviceSubscriptionPreviewResponse{}
	json.Unmarshal(body, &result)

	return &HttpResponse[DeviceSubscriptionPreviewResponse]{
		StatusCode: resp.StatusCode,
		Response:   result,
		Raw:        resp,
	}
}

// Get information about device upgrade by deviceId
func (c Client) DeviceUpgradePreview(deviceId string) *HttpResponse[DeviceUpgradePreviewResponse] {
	resp, err := c.get(fmt.Sprintf("api/devices/%s/upgrade/preview", deviceId), nil, true)
	if err != nil || resp.StatusCode != http.StatusOK {
		return &HttpResponse[DeviceUpgradePreviewResponse]{
			StatusCode: resp.StatusCode,
			Error:      err,
			Raw:        resp,
		}
	}

	defer resp.Body.Close()

	// Parse json response
	body, _ := io.ReadAll(resp.Body)
	result := DeviceUpgradePreviewResponse{}
	json.Unmarshal(body, &result)

	return &HttpResponse[DeviceUpgradePreviewResponse]{
		StatusCode: resp.StatusCode,
		Response:   result,
		Raw:        resp,
	}
}
