# Introduction

This is an unofficial Go API wrapper for the Whistle smart pet collar.
This wrapper handles authentication and interacting with all known endpoints.

Currently, the wrapper only focuses on reading from the API endpoints,
though this may change in the future. Most of the endpoints do support
standard CRUD actions, however.

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
  <summary>With Bearer Token</summary>

  If you already have a bearer token,
  you can instantiate a new wrapper via

  ```go
  whistle, err := whistle.InitializeBearer("API_TOKEN_HERE")
  ```

  This is useful for cases where you want to reduce overhead on page reload.
  You should ideally use this method as often as possible.
</details>

<details>
  <summary>With API Key</summary>
  **Note**: I believe this is deprecated and should not be used.
  The mobile application uses HTTP bearer, and this may be removed unpredictably.

  If you already have an API key (`X-Whistle-AuthToken`),
  you can instantiate a new wrapper via

  ```go
  whistle, err := whistle.InitializeToken("API_TOKEN_HERE")
  ```

  This is useful for cases where you want to reduce overhead on page reload.
  You should ideally use this method as often as possible.
</details>

<details>
  <summary>With Email/Password</summary>

  If you don't have an active API key, but have credentials that work on the <https://Whistle.com>
  mobile app or on <https://app.Whistle.com>, you can instantiate a new wrapper via

  ```go
  whistle, err := whistle.Initialize("EMAIL", "PASSWORD")
  ```

</details>

<details>
  <summary>Manually</summary>

  In the event that you have an advanced need, you may also
  initialize the wrapper directly. You only need `email`/`password`,
  `email`/`refresh_token`, `token`, or `bearer`, but never all 4 options together.

  If you provide a `email` and `password` or `email` and `refresh_token`,
  a HTTP bearer will automatically be requested and stored on your first API query.

  ```go
    client := whistle.Client{
      email: "ABC", // Option 1
      password: "XYZ", // Option 1-1
      refreshToken: "XYZ", // Option 1-2
      token: "123", // Option 2
      bearer: "abc12932", // Option 3
      Timeout: 3000,
      Env: whistle.ProdEnv, // Or: whistle.StagingEnv
      UserAgent: "Custom User Agent",
    }
  ```

</details>

## Methods

**Important note**: The Whistle.com API REQUIRES a `Accept: application/vnd.whistle.com.v4+json`
header to be present in almost ALL REQUESTS otherwise it will return 404.
Occasionally a endpoint (usually a deprecated one) will accept `application/json`.

### Users

This section covers all implementations relating to the REST API surrounding users
(`/api/users`).

<details>
  <summary>DEPRECATED: Users()</summary>

  Get information about the currently authenticated user.
  This does NOT provide information about all associated users.

  ```go
  // ...
  q := client.Users()

  q.StatusCode // "200"
  q.Error // nil

  fmt.Println(q.Response) // {CreatedAt, ..., Username}
  // ...
  ```

</details>

<details>
  <summary>Me()</summary>

  Returns information about the authenticated user.

  ```go
  // ...
  q := client.Me()

  q.StatusCode // "200"
  q.Error // nil

  fmt.Println(q.Response.User) // {CreatedAt, ..., Username}
  // ...
  ```

</details>

<details>
  <summary>CheckEmail(email string)</summary>

  Used to check if an email exists within the database.

  HTTP 404 - Non existing

  HTTP 204 - User exists

  ```go
  // ...
  q := client.CheckEmail("abc@gmail.com")

  fmt.Println(q.Response) // true = exists, false = non-existing
  // ...
  ```

</details>

<details>
  <summary>InvitationCodes(code string)</summary>

  List information about a invitation code. Used during the Whistle App invite process.

  ```go
  // ...
  q := client.InvitationCodes("code123")

  q.StatusCode // "200"
  q.Error // nil

  fmt.Println(q.Response) // {pet: ...}
  // ...
  ```

</details>

<details>
  <summary>ApplicationState()</summary>

  Get information about the current application state.
  Current usage unknown.

  ```go
  // ...
  q := client.ApplicationState()

  q.StatusCode // "200"
  q.Error // nil

  fmt.Println(q.Response.ApplicationState) // {...}
  // ...
  ```

</details>

<details>
  <summary>DEPRECATED: CreditCard()</summary>

  Get information about the current credit card on file.
  Does not return the actual card number.

  ```go
  // ...
  q := client.CreditCard()

  q.StatusCode // "200"
  q.Error // nil

  fmt.Println(q.Response)  // {CardType, ..., ZipCode}
  // ...
  ```

</details>

<details>
  <summary>Subscriptions()</summary>

  Get a list of subscriptions tied to an account, along with
  any Partner subscriptions.

  ```go
  // ...
  q := client.Subscriptions()

  q.StatusCode // "200"
  q.Error // nil

  fmt.Println(q.Response) // {Subscriptions: ..., PartnerServices: ...}
  // ...
  ```

</details>

<details>
  <summary>CancellationPreview()</summary>

  Current usage unknown.

  ```go
  // ...
  q := client.CancellationPreview()

  q.StatusCode // "200"
  q.Error // nil

  fmt.Println(q.Response) // TBD
  // ...
  ```

</details>

<details>
  <summary>CancellationReasons()</summary>

  Returns a list of reasons to cancel a subscription.

  ```go
  // ...
  q := client.CancellationReasons()

  q.StatusCode // "200"
  q.Error // nil

  fmt.Println(q.Response) // {cancellation_reasons: [{id: 123, ...}, ...]}
  // ...
  ```

</details>

### Devices

This portion of the document outlines the implementations of the smart collar
REST api endpoints.

<details>
  <summary>Device(deviceId string)</summary>

  Provides information about the specified smart collar device.

  ```go
  // ...
  q := client.Device("serial_num")

  q.StatusCode // "200"
  q.Error // nil

  fmt.Println(q.Response) // {device: {model_id: ..., ..., has_gps: true, ...}
  // ...
  ```

</details>

<details>
  <summary>DeviceActivation(deviceId string)</summary>

  Provides information about the specified device activation status

  ```go
  // ...
  q := client.DeviceActivation("serial_num")

  q.StatusCode // "200"
  q.Error // nil

  fmt.Println(q.Response) // TBD
  // ...
  ```

</details>

<details>
  <summary>DevicePlans(deviceId string)</summary>

  Provides information about the specified device plans

  ```go
  // ...
  q := client.DevicePlans("serial_num")

  q.StatusCode // "200"
  q.Error // nil

  fmt.Println(q.Response) // {paid_through: "", plans: [ ... ] }
  // ...
  ```

</details>

<details>
  <summary>DeviceSubscription(deviceId string)</summary>

  Provides information about the specified device subscription status

  ```go
  // ...
  q := client.DeviceSubscription("serial_num")

  q.StatusCode // "200"
  q.Error // nil

  fmt.Println(q.Response) // {id: 123, ..., plan: {...}}
  // ...
  ```

</details>

<details>
  <summary>DeviceSubscriptionPreview(deviceId string, planId string)</summary>

  Current usage unknown

  ```go
  // ...
  q := client.DeviceSubscriptionPreview("serial_num", "abc")

  q.StatusCode // "200"
  q.Error // nil

  fmt.Println(q.Response) // TBD
  // ...
  ```

</details>

<details>
  <summary>DeviceUpgradePreview(deviceId string)</summary>

  Current usage unknown

  ```go
  // ...
  q := client.DeviceUpgradePreview("serial_num")

  q.StatusCode // "200"
  q.Error // nil

  fmt.Println(q.Response) // TBD
  // ...
  ```

</details>

<details>
  <summary>DeviceWifiNetworks(deviceId string)</summary>

  Provides a listing of all connected networks associated with a device.

  ```go
  // ...
  q := client.DeviceWifiNetworks("serial_num")

  q.StatusCode // "200"
  q.Error // nil

  fmt.Println(q.Response) // [{id: ..., ssid: "xyz"}, ...]
  // ...
  ```

</details>

### Breeds

This section related to all of the endpoints (currently only 1)
relating to animal breeds.

<details>
  <summary>Breeds(animal string)</summary>

  Provides a list of breeds given the current animal species.
  Known options are `dogs` or `cats`

  ```go
  // ...
  q := client.Breeds("dogs")

  q.StatusCode // "200"
  q.Error // nil

  fmt.Println(q.Response.Breeds) // [{ID: 123, Name: "German Shepherd", ...}, ...]
  // ...
  ```

</details>

### Pets

These are the operations relating to pets (cats/dogs/etc).

<details>
  <summary>Pets()</summary>

  Returns a populated array of objects describing a Pet belonging to
  the authenticated user.

  ```go
  // ...
  q := client.Pets()

  q.StatusCode // "200"
  q.Error // nil

  fmt.Println(q.Response.Pets) // {ID: 135, Name: "Baker", ...}
  // ...
  ```

</details>

<details>
  <summary>PetTransfers()</summary>

  Returns an array of pets that qualify for a transfer.
  Unsure of the current usage.

  ```go
  // ...
  q := client.PetTransfers()

  q.StatusCode // "200"
  q.Error // nil

  fmt.Println(q.Response.Transfers) // [{ID: 123, Name: "Fido", ...}, ...]
  // ...
  ```

</details>

<details>
  <summary>Pet(petId string)</summary>

  Returns detailed information about a specific pet.

  ```go
  // ...
  q := client.Pet("petid123")

  q.StatusCode // "200"
  q.Error // nil

  fmt.Println(q.Response.Pet) // {ID: 123, ..., Name: "Fido"}
  // ...
  ```

</details>

<details>
  <summary>PetOwners(petId string)</summary>

  Returns an array of people that are tied to a pet as owners.

  ```go
  // ...
  q := client.PetOwners("pet1233")

  q.StatusCode // "200"
  q.Error // nil

  fmt.Println(q.Response.Owners) // [{Name: "amattu2", ..., Email: "xyz@gmail.com"}]
  // ...
  ```

</details>

<details>
  <summary>PetWhereabouts(petId string, startDate string, endDate string)</summary>

  Returns informations about a pet's historical locations.
  Based on start/end dates.
  Provides locations and known places.

  ```go
  // ...
  q := client.PetWhereabouts("pet321", "2022-03-03", "2024-01-01")

  q.StatusCode // "200"
  q.Error // nil

  fmt.Println(q.Response) // {Locations: [...], Places: [...]}
  // ...
  ```

</details>

<details>
  <summary>PetLocationsRecent(petId string)</summary>

  Similar to PetWhereabouts, this returns detailed locations
  about where a pet has been as of recent.

  ```go
  // ...
  q := client.PetLocationsRecent("3892821")

  q.StatusCode // "200"
  q.Error // nil

  fmt.Println(q.Response.Locations) // [{...}]
  // ...
  ```

</details>

<details>
  <summary>PetAchievements(petId string)</summary>

  Returns a list of achievements that a pet CAN make.
  The achievements indicate whether or not that goal
  has been met.

  ```go
  // ...
  q := client.PetAchievements("3828111")

  q.StatusCode // "200"
  q.Error // nil

  fmt.Println(q.Response.Achievements) // [{ID: 8382, Name: "1 Week Streak"}]
  // ...
  ```

</details>

<details>
  <summary>PetStatistics(petId string)</summary>

  Returns analytical insights about a pet.

  ```go
  // ...
  q := client.PetStatistics("12345")

  q.StatusCode // "200"
  q.Error // nil

  fmt.Println(q.Response.Statistics) // {AverageMinutesActive: 0, ...}
  // ...
  ```

</details>

<details>
  <summary>PetDailies(petId string)</summary>

  Returns high-level information about a pet's daily activities

  ```go
  // ...
  q := client.PetDailies("12345")

  q.StatusCode // "200"
  q.Error // nil

  fmt.Println(q.Response.Dailies) // [{DayNumber: 93381, ..., UpdatedAt: "..."}]
  // ...
  ```

</details>

<details>
  <summary>PetDaily(petId string, dailyId string)</summary>

  Returns detailed information about a particular pet's daily activity

  ```go
  // ...
  q := client.PetDaily("1234", "938191")

  q.StatusCode // "200"
  q.Error // nil

  fmt.Println(q.Response.Daily) // [{DayNumber: 938191, ...}]
  // ...
  ```

</details>

<details>
  <summary>PetDailyItems(petId string, dailyId string)</summary>

  Returns very low-level, and highly-detailed breakdown of a pet's
  daily activity.

  ```go
  // ...
  q := client.PetDailyItems("1234", "938191")

  q.StatusCode // "200"
  q.Error // nil

  fmt.Println(q.Response.DailyItems) // [...]
  // ...
  ```

</details>

<details>
  <summary>PetHealthTrends(petId string)</summary>

  Provides health trend information about the specified pet.

  ```go
  // ...
  q := client.PetHealthTrends("12345")

  q.StatusCode // "200"
  q.Error // nil

  fmt.Println(q.Response.Trends) // {...}
  // ...
  ```

</details>

<details>
  <summary>PetHealthTrends(petId string, trend string, days int)</summary>

  Provides data to generate a graph for the specified health trend.
  Days limits the number of observations to include.

  ```go
  // ...
  q := client.PetHealthTrends("1234", "sleeping", 7)

  q.StatusCode // "200"
  q.Error // nil

  fmt.Println(q.Response) // {Data: [...], PetId: 1234, ...}
  // ...
  ```

</details>

<details>
  <summary>PetNutritionPortions(petId string)</summary>

  Returns the suggested food portions for the given pet.

  ```go
  // ...
  q := client.PetNutritionPortions("1234")

  q.StatusCode // "200"
  q.Error // nil

  fmt.Println(q.Response) // {Treats: ..., SuggestedCalories: 0.0}
  // ...
  ```

</details>

<details>
  <summary>DEPRECATED: PetFoodPortions(petId string)</summary>

  Returns the suggested food portions for the given pet.
  Replaced by the above method.

  ```go
  // ...
  q := client.PetFoodPortions("1234")

  q.StatusCode // "200"
  q.Error // nil

  fmt.Println(q.Response.PetFoodPortions) // ...
  // ...
  ```

</details>

<details>
  <summary>PetTask(petId string, taskId string)</summary>

  Returns information about the pet's task.

  ```go
  // ...
  q := client.PetTask("1234", "35")

  q.StatusCode // "200"
  q.Error // nil

  fmt.Println(q.Response) // ...
  // ...
  ```

</details>

<details>
  <summary>PetTaskOccurrence(petId string, occurrenceType string)</summary>

  Current usage unknown.

  ```go
  // ...
  q := client.PetTaskOccurrence("1234", "incomplete")

  q.StatusCode // "200"
  q.Error // nil

  fmt.Println(q.Response) // ...
  // ...
  ```

</details>

### Miscellaneous

These are operations not categorized by another API route.

<details>
  <summary>Notifications()</summary>

  Returns an array of unread notifications for the current user.

  ```go
  // ...
  q := client.Notifications()

  q.StatusCode // "200"
  q.Error // nil

  fmt.Println(q.Response) // {Items: [...]}
  // ...
  ```

</details>

<details>
  <summary>PetFoods(foodType string)</summary>

  Returns a list of pet foods given the food type.
  Known options are `dog_treat`, `dog_food`. Cat variant does not work.

  ```go
  // ...
  q := client.PetFoods("dog_food")

  q.StatusCode // "200"
  q.Error // nil

  fmt.Println(q.Response) // [{ID: 321, Name: "Purina XXX"}, ...]
  // ...
  ```

</details>

<details>
  <summary>ReverseGeocode(latitude string, longitude string)</summary>

  Decode latitude and longitude to a physical address.

  ```go
  // ...
  q := client.ReverseGeocode("LAT", "LON")

  q.StatusCode // "200"
  q.Error // nil

  fmt.Println(q.Response.Description) // {address: ..., region: ..., etc}
  // ...
  ```

</details>

<details>
  <summary>Places()</summary>

  Returns a list of saved places tied to a user account.

  ```go
  // ...
  q := client.Places()

  q.StatusCode // "200"
  q.Error // nil

  fmt.Println(q.Response) // [ {address: "123 ABC Lane", ..., id: 123}, ...]
  // ...
  ```

</details>

<details>
  <summary>AdventureCategories()</summary>

  Returns a list of adventure categories.
  Current usage unknown.

  ```go
  // ...
  q := client.AdventureCategories()

  q.StatusCode // "200"
  q.Error // nil

  fmt.Println(q.Response) // []
  // ...
  ```

</details>

# Requirements

- Go 1.18+ (Required for Generics)
- <https://whistle.com> account

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

- [ ] Add support for the async event pusher service (See `realtime_channel`)
- [ ] Check into the WiFi adding endpoint
