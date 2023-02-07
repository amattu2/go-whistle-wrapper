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
	"strings"
)

type InvitationCodeResponse struct {
	Errors []Error `json:"errors"`
	Pet    Pet     `json:"pet"`
}

type ApplicationStateResponse struct {
	ApplicationState map[string]string `json:"application_state"`
}

type SubscriptionsResponse struct {
	Subscriptions   []Subscription   `json:"subscriptions"`
	PartnerServices []PartnerService `json:"partner_services"`
}

type CancellationReasonsResponse struct {
	Errors              []Error              `json:"errors"`
	CancellationReasons []CancellationReason `json:"cancellation_reasons"`
}

type CancellationReason struct {
	ID          int    `json:"id"`
	ShortName   string `json:"short_name"`
	Description string `json:"description"`
}

type CancellationPreviewResponse struct {
	Errors []Error `json:"errors"`
}

type UsersResponse struct {
	CreatedAt              string               `json:"created_at"`
	CurrentUser            bool                 `json:"current_user"`
	Dogs                   []Dog                `json:"dogs"`
	Email                  string               `json:"email"`
	FirstName              string               `json:"first_name"`
	Friends                []Friends            `json:"friends"`
	HasUnreadNotifications bool                 `json:"has_unread_notifications"`
	ID                     string               `json:"id"`
	LastName               string               `json:"last_name"`
	Name                   string               `json:"name"`
	NotificationSettings   NotificationSettings `json:"notification_settings"`
	ProfilePhotoUrl        string               `json:"profile_photo_url"`
	ProfilePhotoSizes      map[string]string    `json:"profile_photo_sizes"`
	RealtimeChannel        RealtimeChannel      `json:"realtime_channel"`
	Searchable             bool                 `json:"searchable"`
	SendMarketingEmails    bool                 `json:"send_marketing_emails"`
	UserType               string               `json:"user_type"`
	Username               string               `json:"username"`
	UserActivations        []UserActivation     `json:"user_activations"`
}

type UserActivation struct {
	DeviceSerial string          `json:"device_serial"`
	Status       string          `json:"status"`
	Events       map[string]bool `json:"events"`
}

type MeResponse struct {
	User UsersResponse `json:"user"`
}

type NotificationSettings struct {
	EmailCategories       map[string]bool `json:"email_categories"`
	PhoneNumber           PhoneNumber     `json:"phone_number"`
	PushCategories        map[string]bool `json:"push_categories"`
	SecondaryEmails       []string        `json:"secondary_emails"`
	SecondaryPhoneNumbers []string        `json:"secondary_phone_numbers"`
	SendEmail             bool            `json:"send_email"`
	SendSMS               bool            `json:"send_sms"`
	SMSCategories         map[string]bool `json:"sms_categories"`
}

type PhoneNumber struct {
	ID       int    `json:"id"`
	Primary  string `json:"primary"`
	Number   string `json:"number"`
	Verified bool   `json:"verified"`
}

type CreditCard struct {
	CardType        string `json:"card_type"`
	ExpirationMonth int    `json:"expiration_month"`
	ExpirationYear  int    `json:"expiration_year"`
	LastFour        string `json:"last4"`
	ZipCode         string `json:"zip"`
}

type Subscription struct {
	ID                      string `json:"id"`
	CanceledAt              string `json:"canceled_at"`
	CancellationEffectiveOn string `json:"cancellation_effective_on"`
	CancelAtEndOfContract   bool   `json:"cancel_at_end_of_contract"`
	User                    User   `json:"user"`
	PaidThrough             string `json:"paid_through"`
	Plan                    Plan   `json:"plan"`
	Status                  string `json:"status"`
	Legacy                  bool   `json:"legacy"`
	PetId                   string `json:"pet_id"`
	Coupon                  Coupon `json:"coupon"`
}

type PartnerService struct {
}

type Dog struct {
}

type Friends struct {
}

// Users returns information about the current user
//
// Deprecated: Use Me() instead
func (c Client) Users() *HttpResponse[UsersResponse] {
	resp, err := c.get("api/users", nil, true)
	if err != nil || resp.StatusCode != http.StatusOK {
		return &HttpResponse[UsersResponse]{
			StatusCode: resp.StatusCode,
			Error:      err,
			Raw:        resp,
		}
	}

	defer resp.Body.Close()

	// Parse json response
	body, _ := io.ReadAll(resp.Body)
	result := UsersResponse{}
	json.Unmarshal(body, &result)

	return &HttpResponse[UsersResponse]{
		StatusCode: resp.StatusCode,
		Response:   result,
		Raw:        resp,
	}
}

// Me returns information about the current user
func (c Client) Me() *HttpResponse[MeResponse] {
	resp, err := c.get("api/users/me", nil, true)
	if err != nil || resp.StatusCode != http.StatusOK {
		return &HttpResponse[MeResponse]{
			StatusCode: resp.StatusCode,
			Error:      err,
			Raw:        resp,
		}
	}

	defer resp.Body.Close()

	// Parse json response
	body, _ := io.ReadAll(resp.Body)
	result := MeResponse{}
	json.Unmarshal(body, &result)

	return &HttpResponse[MeResponse]{
		StatusCode: resp.StatusCode,
		Response:   result,
		Raw:        resp,
	}
}

// CheckEmail checks the provided email address to see if it is already in use
func (c Client) CheckEmail(email string) *HttpResponse[bool] {
	email = strings.ReplaceAll(email, "@", "%40")
	email = strings.ReplaceAll(email, ".", "%2E")

	resp, err := c.get(fmt.Sprintf("api/users/emails/%s", email), nil, true)

	if err != nil {
		return &HttpResponse[bool]{
			Error: err,
			Raw:   resp,
		}
	}

	defer resp.Body.Close()

	if http.StatusNoContent == resp.StatusCode {
		return &HttpResponse[bool]{
			StatusCode: resp.StatusCode,
			Response:   true,
			Raw:        resp,
		}
	} else {
		return &HttpResponse[bool]{
			StatusCode: resp.StatusCode,
			Response:   false,
			Raw:        resp,
		}
	}
}

// InvitationCodes returns the pet information for the provided invitation code
func (c Client) InvitationCodes(code string) *HttpResponse[InvitationCodeResponse] {
	resp, err := c.get(fmt.Sprintf("api/users/invitation_codes/%s", code), nil, true)
	if err != nil || resp.StatusCode != http.StatusOK {
		return &HttpResponse[InvitationCodeResponse]{
			StatusCode: resp.StatusCode,
			Error:      err,
			Raw:        resp,
		}
	}

	defer resp.Body.Close()

	// Parse json response
	body, _ := io.ReadAll(resp.Body)
	result := InvitationCodeResponse{}
	json.Unmarshal(body, &result)

	return &HttpResponse[InvitationCodeResponse]{
		StatusCode: resp.StatusCode,
		Response:   result,
		Raw:        resp,
	}
}

// ApplicationState provides information about the current application state
func (c Client) ApplicationState() *HttpResponse[ApplicationStateResponse] {
	resp, err := c.get("api/users/application_state", nil, true)
	if err != nil || resp.StatusCode != http.StatusOK {
		return &HttpResponse[ApplicationStateResponse]{
			StatusCode: resp.StatusCode,
			Error:      err,
			Raw:        resp,
		}
	}

	defer resp.Body.Close()

	// Parse json response
	body, _ := io.ReadAll(resp.Body)
	result := ApplicationStateResponse{}
	json.Unmarshal(body, &result)

	return &HttpResponse[ApplicationStateResponse]{
		StatusCode: resp.StatusCode,
		Response:   result,
		Raw:        resp,
	}
}

// CreditCard provides information about the current user's credit card on file
//
// Deprecated: Unknown replacement.
func (c Client) CreditCard() *HttpResponse[CreditCard] {
	resp, err := c.get("api/users/credit_card", map[string]string{"Accept": "application/json"}, true)
	if err != nil || resp.StatusCode != http.StatusOK {
		return &HttpResponse[CreditCard]{
			StatusCode: resp.StatusCode,
			Error:      err,
			Raw:        resp,
		}
	}

	defer resp.Body.Close()

	// Parse json response
	body, _ := io.ReadAll(resp.Body)
	result := CreditCard{}
	json.Unmarshal(body, &result)

	return &HttpResponse[CreditCard]{
		StatusCode: resp.StatusCode,
		Response:   result,
		Raw:        resp,
	}
}

// Subscriptions provides a listing of the current user's subscriptions
func (c Client) Subscriptions() *HttpResponse[SubscriptionsResponse] {
	resp, err := c.get("api/users/subscriptions", nil, true)
	if err != nil || resp.StatusCode != http.StatusOK {
		return &HttpResponse[SubscriptionsResponse]{
			StatusCode: resp.StatusCode,
			Error:      err,
			Raw:        resp,
		}
	}

	defer resp.Body.Close()

	// Parse json response
	body, _ := io.ReadAll(resp.Body)
	result := SubscriptionsResponse{}
	json.Unmarshal(body, &result)

	return &HttpResponse[SubscriptionsResponse]{
		StatusCode: resp.StatusCode,
		Response:   result,
		Raw:        resp,
	}
}

// Todo: Figure out what this does
func (c Client) CancellationPreview(subId string) *HttpResponse[CancellationPreviewResponse] {
	resp, err := c.get(fmt.Sprintf("api/users/subscriptions/%s/cancellation/preview", subId), nil, true)
	if err != nil || resp.StatusCode != http.StatusOK {
		return &HttpResponse[CancellationPreviewResponse]{
			StatusCode: resp.StatusCode,
			Error:      err,
			Raw:        resp,
		}
	}

	defer resp.Body.Close()

	// Parse json response
	body, _ := io.ReadAll(resp.Body)
	result := CancellationPreviewResponse{}
	json.Unmarshal(body, &result)

	return &HttpResponse[CancellationPreviewResponse]{
		StatusCode: resp.StatusCode,
		Response:   result,
		Raw:        resp,
	}
}

// CancellationReasons returns a list of reasons why a user may be cancelling their subscription
func (c Client) CancellationReasons(subId string) *HttpResponse[CancellationReasonsResponse] {
	resp, err := c.get(fmt.Sprintf("api/users/subscriptions/%s/cancellation/preview", subId), nil, true)
	if err != nil || resp.StatusCode != http.StatusOK {
		return &HttpResponse[CancellationReasonsResponse]{
			StatusCode: resp.StatusCode,
			Error:      err,
			Raw:        resp,
		}
	}

	defer resp.Body.Close()

	// Parse json response
	body, _ := io.ReadAll(resp.Body)
	result := CancellationReasonsResponse{}
	json.Unmarshal(body, &result)

	return &HttpResponse[CancellationReasonsResponse]{
		StatusCode: resp.StatusCode,
		Response:   result,
		Raw:        resp,
	}
}
