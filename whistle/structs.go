/*
 * Produced: Fri Feb 03 2023
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

// ----------------------------------------------
// Nested Structure Definitions
// ----------------------------------------------

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

type RealtimeChannel struct {
	Channel string `json:"channel"`
	Service string `json:"service"`
}

type CreditCard struct {
	CardType        string `json:"card_type"`
	ExpirationMonth int    `json:"expiration_month"`
	ExpirationYear  int    `json:"expiration_year"`
	LastFour        string `json:"last4"`
	ZipCode         string `json:"zip"`
}

type Dog struct {
}

type Friends struct {
}

type Device struct {
}

type Daily struct {
}

type Timeline struct {
}

type Highlight struct {
}

type Goal struct {
}

type Average struct {
}

type DailyTotal struct {
}

type Notification struct {
}

type User struct {
	CreatedAt            string            `json:"created_at"`
	CurrentUser          bool              `json:"current_user"`
	Email                string            `json:"email"`
	FirstName            string            `json:"first_name"`
	ID                   string            `json:"id"`
	LastName             string            `json:"last_name"`
	ProfilePhotoUrl      string            `json:"profile_photo_url"`
	ProfilePhotoUrlSizes map[string]string `json:"profile_photo_url_sizes"`
	RealtimeChannel      RealtimeChannel   `json:"realtime_channel"`
	Searchable           bool              `json:"searchable"`
	SendMarketingEmails  bool              `json:"send_marketing_emails"`
	UserActivations      string            `json:"user_activations"`
	UserType             string            `json:"user_type"`
	Username             string            `json:"username"`
}
