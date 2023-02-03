# Introduction

This is an unofficial Go API wrapper for the Whistle smart pet collar.
This wrapper handles authentication and interacting with all known endpoints.

If you have discovered endpoints not listed here, please open a PR or submit an issue.

# Usage

## ThunderClient / Postman

If you're only here for the (previously) undocumented API endpoints provided by <https://Whistle.com>,
check out the [Thunder Tests](/.vscode/thunder-tests/) folder.
This folder contains the HTTP request collection supported by the Whistle V3 API.

## Setup

Install the wrapper within your project via `Go Get`

```bash
go get github.com/amattu2/go-whistle-wrapper
```

## Initialization

The wrapper exposes two ways of instantiating a client.

<details>
  <summary>With API Key</summary>

  If you already have an API key (`X-Whistle-AuthToken`),
  you can instantiate a new wrapper via

  ```go
  whistle := whistle.InitializeToken("API_TOKEN_HERE")
  ```

  This is useful for cases where you want to reduce overhead on page reload.
  You should ideally use this method as often as possible.
</details>

<details>
  <summary>With Email/Password</summary>

  If you don't have an active API key, but have credentials that work on the <https://Whistle.com>
  mobile app or on <https://app.Whistle.com>, you can instantiate a new wrapper via

  ```go
  whistle := whistle.Initialize("EMAIL", "PASSWORD")
  ```

</details>

<details>
  <summary>Manually</summary>

  In the event that you have an advanced need, you may also
  initialize the wrapper directly.

  ```go
    client := whistle.Client{
      email: "ABC",
      password: "XYZ",
      token: "123", // Not required if email/pass are passed
      Timeout: 3000,
      Env: whistle.ProdEnv, // Or: whistle.StagingEnv
      UserAgent: "Custom User Agent",
    }
  ```

</details>

## Methods

<details>
  <summary>Users</summary>

</details>

<details>
  <summary>Notifications</summary>

</details>

<details>
  <summary>Device</summary>

</details>

<details>
  <summary>Dogs</summary>

</details>

<details>
  <summary>Dog(dogId)</summary>

</details>

<details>
  <summary>Highlights(dogId, type)</summary>

</details>

<details>
  <summary>Dailies(dogId, limit)</summary>

</details>

<details>
  <summary>Daily(dogId, dailyId)</summary>

</details>

<details>
  <summary>Timeline(dogId, timelineId)</summary>

</details>

<details>
  <summary>UsersPresent(dogId)</summary>

</details>

<details>
  <summary>Goals(dogId)</summary>

</details>

<details>
  <summary>Averages(dogId)</summary>

</details>

<details>
  <summary>DailyTotals(dogId, startDate)</summary>

</details>

<details>
  <summary>UsersCreditCard</summary>

</details>

# Credits

The following resources were used to compile this API wrapper.

- API Endpoint Reference <https://github.com/aolney/WhistleAPI-DOTNET/blob/master/WhistleAPI-DOTNET-Fsharp.ipynb>
- API Reference <https://github.com/martzcodes/node-whistle>
- API Reference <https://community.smartthings.com/t/beta-release-whistle-3-pet-tracker-presence-and-battery-dth/156031>
- Design Reference <https://www.reddit.com/r/golang/comments/d8m5a5/advice_for_creating_an_api_wrapper>
- Implementation Reference <https://github.com/ovh/go-ovh>

# To-Do

These are the remaining action items of this project.
Some of them are on hold until I have a device to test with.

- [ ] Test cases
- [ ] Handle API key expiration or renewal
- [ ] Finish README.md usage docs
- [ ] Fill in missing struct definitions
- [ ] Figure out what `highlightType` options are available
- [ ] Add support for the async event pusher service

<details>
  <summary>Confirm all of these are implemented</summary>

  Some of these may not actually work, all of them are `GET` requests

  ```java
  devices/{serial_number}
  devices/{serial_number}/activation
  devices/{serial_number}/plans
  devices/{serial_number}/subscription
  devices/{serial_number}/subscription/previews/{plan_id}
  devices/{old_device_serial_number}/upgrade/preview
  devices/{serial_number}/wifi_networks
  users/me
  users/emails/{email}
  users/invitation_codes/{invitation_code}
  users/application_state
  users/referral_code
  users/subscriptions
  pets/{pet_id}/owners
  pets/{id}
  pets/{id}/whereabouts
  pets/{pet_id}/achievements
  pets/{id}/dailies/{dayNumber}/daily_items
  pets/{pet_id}/task_occurrences?q=complete
  pets/{id}/dailies
  pets/{pet_id}/health/trends
  pets/{pet_id}/health/graphs/eating_events
  pets/{pet_id}/health/graphs/sleeping
  pets/{pet_id}/health/graphs/{trend_type}
  pets/{pet_id}/task_occurrences?q=incomplete
  pets/{id}/dailies/{dayNumber}
  pets/{pet_id}/nutrition/v2/suggested_portions
  pets/{pet_id}/task_occurrences?q=overdue
  pets/{pet_id}/pet_food_portions
  pets/{id}/locations/recent_trackings
  pets/{pet_id}/tasks/{task_id}
  pets/{id}/stats
  pets/transfers
  pets/{pet_id}/task_occurrences?q=upcoming
  pets/{id}/whereabouts
  adventures/categories
  adventures/poi
  adventures/poi/{poi_id}
  partners/banfield/pets/{client_id}
  subscriptions/{subscription_id}/cancellation/preview
  subscriptions/{subscription_id}/cancellation/reasons
  breeds/cats
  breeds/dogs
  pet_foods?type=dog_food
  health_conditions/dogs
  pet_foods?type=dog_treat
  /api/users/me/firmware_updates
  coupons/{coupon_id}
  partners/{partner_id}/account
  performance_settings
  pets
  places
  partners/vca/devices/{serial_number}
  reverse_geocode
  ```

</details>
