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
  initialize the wrapper directly. You only need `email`/`password`, `token`,
  or `bearer`, but never all 3 together.

  ```go
    client := whistle.Client{
      email: "ABC", // Option 1
      password: "XYZ", // Option 1
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
header to be present in ALL REQUESTS otherwise it will return 404.

### Users

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
  <summary>InvitationCodes(code string)</summary>

  Current usage unknown.

  ```go
  // ...
  q := client.InvitationCodes("code123")

  q.StatusCode // "200"
  q.Error // nil

  fmt.Println(q.Response) // TBD
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

  Current usage unknown.

  ```go
  // ...
  q := client.CancellationReasons()

  q.StatusCode // "200"
  q.Error // nil

  fmt.Println(q.Response) // TBD
  // ...
  ```

</details>

### Devices

Todo

### Breeds

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

Todo

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
  <summary>ReverseGeocode(...)</summary>

  Current usage unknown.

  ```go
  // ...
  q := client.ReverseGeocode()

  q.StatusCode // "200"
  q.Error // nil

  fmt.Println(q.Response) // TBD
  // ...
  ```

</details>

<details>
  <summary>Places()</summary>

  Returns a list of saved (?) places tied to a user account.

  ```go
  // ...
  q := client.Places()

  q.StatusCode // "200"
  q.Error // nil

  fmt.Println(q.Response) // TBD
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

  fmt.Println(q.Response) // TBD
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

- [X] Switch to Bearer usage
- [X] Implement all of the newly discovered endpoints (see Thunder Client)
- [X] Test cases
- [ ] ~~Handle API key expiration or renewal~~
- [ ] Finish README.md usage documentation
- [ ] Add support for the async event pusher service (See `realtime_channel`)
