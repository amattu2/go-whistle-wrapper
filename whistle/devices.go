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
	Error  string `json:"error"`
	Device Device `json:"device"`
}

type Device struct {
	ModelId              string                 `json:"model_id"`
	SerialNumber         string                 `json:"serial_number"`
	LastCheckIn          string                 `json:"last_check_in"`
	FirmwareVersion      string                 `json:"firmware_version"`
	BatteryLevel         int                    `json:"battery_level"`
	BatteryStatus        string                 `json:"battery_status"`
	PendingLocate        bool                   `json:"pending_locate"`
	TrackingStatus       string                 `json:"tracking_status"`
	HasGPS               bool                   `json:"has_gps"`
	RequiresSubscription bool                   `json:"requires_subscription"`
	FlashlightStatus     string                 `json:"flashlight_status"`
	PartnerRecord        string                 `json:"partner_record"`
	BundledSubscription  bool                   `json:"bundled_subscription"`
	DeviceConfig         map[string]interface{} `json:"device_configs"`
	BatteryStats         BatteryStats           `json:"battery_stats"`
}

type BatteryStats struct {
	BatteryDaysLeft         int     `json:"battery_days_left"`
	BatteryDrainLast24Hours int     `json:"battery_drain_last_24_hours"`
	TotalBatteryLifeDays    float64 `json:"total_battery_life_days"`
	PriorUsageMinutes       map[string]interface{}
}

type DeviceActivationResponse struct {
	Error string `json:"error"`
}

type DevicePlansResponse struct {
	Error       string `json:"error"`
	PaidThrough string `json:"paid_through"`
	Plans       []Plan `json:"plans"`
}

type Plan struct {
	ID                    string  `json:"id"`
	Name                  string  `json:"name"`
	PlanType              string  `json:"plan_type"`
	Interval              string  `json:"interval"`
	IntervalCount         int     `json:"interval_count"`
	FullAmount            float64 `json:"full_amount"`
	MonthlyAmount         float64 `json:"monthly_amount"`
	Currency              string  `json:"currency"`
	RiskFreeDays          int     `json:"risk_free_days"`
	SavePercent           float64 `json:"save_percent"`
	ShortName             string  `json:"short_name"`
	CurrentPlan           bool    `json:"current_plan"`
	DefaultPlan           bool    `json:"default_plan"`
	RequiresContract      bool    `json:"requires_contract"`
	ContractInterval      string  `json:"contract_interval"`
	ContractIntervalCount int     `json:"contract_interval_count"`
	TrialPeriod           int     `json:"trial_period"`
	TrialPeriodUnit       string  `json:"trial_period_unit"`
}

type DeviceSubscriptionResponse struct {
	Error          string       `json:"error"`
	Subscription   Subscription `json:"subscription"`
	PartnerService string       `json:"partner_service"`
}

type DeviceSubscriptionPreviewResponse struct {
	Error string `json:"error"`
}

type DeviceUpgradePreviewResponse struct {
	Error string `json:"error"`
}

type DeviceWifiNetworksResponse struct {
	WifiNetworks []WifiNetwork `json:"wifi_networks"`
}

type WifiNetwork struct {
	ID      int    `json:"id"`
	SSID    string `json:"ssid"`
	Name    string `json:"name"`
	PlaceId int    `json:"place_id"`
	PetIds  []int  `json:"pet_ids"`
}

// Device gets detailed information about a smart collar device by deviceId
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

// DeviceActivation returns detailed information about device activation status by deviceId
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

// DevicePlans provides the available plans for a device by deviceId
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

// DeviceSubscription returns detailed information about device subscription by deviceId
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

// DeviceSubscriptionPreview gets information about device subscription renewal by deviceId and planId
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

// DeviceUpgradePreview returns information about device upgrade by deviceId
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

// DeviceWifiNetworks returns information about Wifi networks a device has connected to
func (c Client) DeviceWifiNetworks(deviceId string) *HttpResponse[DeviceWifiNetworksResponse] {
	resp, err := c.get(fmt.Sprintf("api/devices/%s/wifi_networks", deviceId), nil, true)
	if err != nil || resp.StatusCode != http.StatusOK {
		return &HttpResponse[DeviceWifiNetworksResponse]{
			StatusCode: resp.StatusCode,
			Error:      err,
			Raw:        resp,
		}
	}

	defer resp.Body.Close()

	// Parse json response
	body, _ := io.ReadAll(resp.Body)
	result := DeviceWifiNetworksResponse{}
	json.Unmarshal(body, &result)

	return &HttpResponse[DeviceWifiNetworksResponse]{
		StatusCode: resp.StatusCode,
		Response:   result,
		Raw:        resp,
	}
}
