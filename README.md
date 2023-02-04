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
  whistle := whistle.InitializeBearer("API_TOKEN_HERE")
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
to be present in ALL REQUESTS otherwise it will return 404.

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

- [X] Switch to Bearer usage
- [ ] Implement all of the newly discovered endpoints (see Thunder Client)
- [ ] Test cases
- [ ] Handle API key expiration or renewal
- [ ] Finish README.md usage docs
- [ ] Fill in missing struct definitions
- [ ] Add support for the async event pusher service (See "realtime_channel")
