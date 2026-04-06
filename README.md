![Go SDK for Auth0 MyOrganization](https://cdn.auth0.com/website/sdks/banners/myorganization-go-banner.png)

<div align="center">

[![GoDoc](https://pkg.go.dev/badge/github.com/auth0/myorganization-go.svg)](https://pkg.go.dev/github.com/auth0/myorganization-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/auth0/myorganization-go?style=flat-square)](https://goreportcard.com/report/github.com/auth0/myorganization-go)
[![Release](https://img.shields.io/github/v/release/auth0/myorganization-go?include_prereleases&style=flat-square)](https://github.com/auth0/myorganization-go/releases)
[![License](https://img.shields.io/github/license/auth0/myorganization-go.svg?style=flat-square)](https://github.com/auth0/myorganization-go/blob/main/LICENSE)
[![Build Status](https://img.shields.io/github/actions/workflow/status/auth0/myorganization-go/ci.yml?branch=main&style=flat-square)](https://github.com/auth0/myorganization-go/actions?query=branch%3Amain)
[![Codecov](https://img.shields.io/codecov/c/github/auth0/myorganization-go/main?style=flat-square)](https://codecov.io/gh/auth0/myorganization-go/tree/main)

📚 [Documentation](#documentation) • 🚀 [Getting Started](#getting-started) • 💬 [Feedback](#feedback)

</div>

---

## Documentation

- [Godoc](https://pkg.go.dev/github.com/auth0/myorganization-go) - explore the Go SDK documentation.
- [Docs site](https://www.auth0.com/docs) — explore our docs site and learn more about Auth0.
- [Examples](./examples.md) - Practical usage examples for all SDK features.
- [API Reference](./reference.md) - Complete API reference documentation.

## Getting Started

### Requirements

This library follows the [same support policy as Go](https://go.dev/doc/devel/release#policy). The last two major Go releases are actively supported and compatibility issues will be fixed. While you may find that older versions of Go may work, we will not actively test and fix compatibility issues with these versions.

- Go 1.25+

### Installation

```shell
go get github.com/auth0/myorganization-go
```

### Usage

#### Client Credentials (M2M)

Use `option.WithClientCredentials` for machine-to-machine authentication via the OAuth2 client credentials grant. The SDK automatically fetches, caches, and refreshes access tokens.

```go
package main

import (
    "context"
    "fmt"
    "log"

    "github.com/auth0/myorganization-go/client"
    "github.com/auth0/myorganization-go/option"
)

func main() {
    c, err := client.New(
        "<YOUR_AUTH0_DOMAIN>",                  // e.g. "mytenant.auth0.com"
        option.WithClientCredentials(
            context.Background(),
            "<YOUR_CLIENT_ID>",
            "<YOUR_CLIENT_SECRET>",
        ),
    )
    if err != nil {
        log.Fatalf("failed to create client: %v", err)
    }

    details, err := c.OrganizationDetails.Get(context.Background())
    if err != nil {
        log.Fatalf("failed to get organization details: %v", err)
    }
    fmt.Println(details)
}
```

> **Note**
> The domain is sanitized automatically (`https://` prefix and trailing slashes are stripped).

> The default audience is `https://{domain}/my-org/`. To specify a custom audience, use `WithClientCredentialsAndAudience`:

```go
c, err := client.New(
    "<YOUR_AUTH0_DOMAIN>",
    option.WithClientCredentialsAndAudience(
        context.Background(),
        "<YOUR_CLIENT_ID>",
        "<YOUR_CLIENT_SECRET>",
        "https://custom-api.example.com/",
    ),
)
```

#### Private Key JWT

Use `option.WithPrivateKeyJWT` for authentication using a signed JWT assertion instead of a client secret. The SDK creates a JWT signed with your private key, then exchanges it for an access token via the `client_credentials` grant with `client_assertion`.

Supported signing algorithms: RS256, RS384, RS512, PS256, PS384, PS512, ES256, ES384, ES512.

```go
import (
    "context"
    "os"

    "github.com/auth0/myorganization-go/client"
    "github.com/auth0/myorganization-go/option"
)

privateKeyPEM, _ := os.ReadFile("private_key.pem")

c, err := client.New(
    "<YOUR_AUTH0_DOMAIN>",
    option.WithPrivateKeyJWT(
        context.Background(),
        "<YOUR_CLIENT_ID>",
        string(privateKeyPEM),
        "RS256",
    ),
)
```

> The default audience is `https://{domain}/my-org/`. To specify a custom audience, use `WithPrivateKeyJWTAndAudience`:

```go
c, err := client.New(
    "<YOUR_AUTH0_DOMAIN>",
    option.WithPrivateKeyJWTAndAudience(
        context.Background(),
        "<YOUR_CLIENT_ID>",
        string(privateKeyPEM),
        "RS256",
        "https://custom-api.example.com/",
    ),
)
```

#### Custom Token Source

Use `option.WithTokenSource` to provide your own `oauth2.TokenSource` implementation. This gives you full control over how access tokens are obtained, cached, and refreshed.

```go
import (
    "github.com/auth0/myorganization-go/client"
    "github.com/auth0/myorganization-go/option"
    "golang.org/x/oauth2"
)

c, err := client.New(
    "<YOUR_AUTH0_DOMAIN>",
    option.WithTokenSource(
        oauth2.StaticTokenSource(&oauth2.Token{
            AccessToken: "<YOUR_TOKEN>",
            TokenType:   "Bearer",
        }),
    ),
)
```

#### Static Token

Use `option.WithStaticToken` (or `option.WithToken`) when you already have a bearer token:

```go
c, err := client.New(
    "<YOUR_AUTH0_DOMAIN>",
    option.WithStaticToken("<YOUR_API_TOKEN>"),
)
```

### Pagination

List endpoints return a cursor-based paginated response. Each response includes a `Next` field — when non-nil, pass its value as the `next` query parameter to retrieve the next page.

```go
import (
    "context"
    "fmt"
    "net/url"

    myorganization "github.com/auth0/myorganization-go"
    "github.com/auth0/myorganization-go/option"
)

ctx := context.Background()

// Fetch the first page.
response, err := c.Organization.Domains.List(ctx)
if err != nil {
    return err
}

for _, domain := range response.OrganizationDomains {
    fmt.Println(domain.GetDomain())
}

// Fetch subsequent pages using the cursor.
for response.Next != nil {
    response, err = c.Organization.Domains.List(ctx,
        option.WithQueryParameters(url.Values{
            "next": {*response.Next},
        }),
    )
    if err != nil {
        return err
    }
    for _, domain := range response.OrganizationDomains {
        fmt.Println(domain.GetDomain())
    }
}
```

### Request Options

Options can be applied at the client level (affecting all requests) or per-request:

```go
// Client-level options (applied to every request).
c, err := client.New(
    "<YOUR_AUTH0_DOMAIN>",
    option.WithClientCredentials(ctx, clientID, clientSecret),
    option.WithMaxAttempts(3),
    option.WithHTTPClient(&http.Client{Timeout: 10 * time.Second}),
)

// Per-request overrides.
response, err := c.Organization.Domains.Create(
    ctx,
    request,
    option.WithToken("<override-token>"),         // Override auth for one call
    option.WithMaxAttempts(5),                    // More retries for this call
    option.WithHTTPHeader(http.Header{            // Extra headers
        "X-Request-Id": {"abc-123"},
    }),
)
```

Available options:

| Option | Scope | Description |
|---|---|---|
| **Authentication** | | |
| `WithClientCredentials(ctx, id, secret)` | Client | OAuth2 client credentials (M2M) |
| `WithClientCredentialsAndAudience(ctx, id, secret, aud)` | Client | Client credentials with custom audience |
| `WithPrivateKeyJWT(ctx, id, pem, alg)` | Client | Private Key JWT assertion |
| `WithPrivateKeyJWTAndAudience(ctx, id, pem, alg, aud)` | Client | Private Key JWT with custom audience |
| `WithTokenSource(source)` | Client | Custom `oauth2.TokenSource` |
| `WithStaticToken(token)` | Both | Static bearer token (alias for `WithToken`) |
| `WithToken(token)` | Both | Set a bearer token |
| **HTTP Configuration** | | |
| `WithBaseURL(url)` | Both | Override the base URL |
| `WithHTTPClient(client)` | Both | Provide a custom `http.Client` |
| `WithHTTPHeader(header)` | Both | Set additional HTTP headers |
| `WithMaxAttempts(n)` | Both | Set max retry attempts |
| `WithQueryParameters(params)` | Both | Add query parameters |
| `WithBodyProperties(props)` | Both | Add extra JSON body properties |
| **Organization** | | |
| `WithOrganization(org)` | Client | Set the organization name or ID for token requests (required when no default org is configured or the app has access to multiple orgs) |
| **Telemetry** | | |
| `WithNoAuth0ClientInfo()` | Client | Disable the `Auth0-Client` header |
| `WithAuth0ClientEnvEntry(key, value)` | Client | Add custom entries to the `Auth0-Client` env map |
| **Debug** | | |
| `WithDebug(true)` | Client | Enable HTTP request/response debug logging |

> **Scope**: *Client* = only when creating the client. *Both* = client-level or per-request override. Auth options based on client credentials, private key JWT, or token source are ignored when passed per-request. Note: when the client is configured with these OAuth-based auth modes, the underlying `oauth2.Transport` manages the `Authorization` header, so per-request `WithToken`/`WithStaticToken` cannot override it. Per-request token overrides only work when no OAuth-based auth is configured on the client.

### Raw Responses

Every resource client exposes a `WithRawResponse` field that returns the full HTTP response (status code, headers, and parsed body):

```go
raw, err := c.Organization.Domains.WithRawResponse.List(ctx)
if err != nil {
    return err
}

fmt.Println("Status:", raw.StatusCode)
fmt.Println("Rate limit remaining:", raw.Header.Get("X-RateLimit-Remaining"))

for _, domain := range raw.Body.OrganizationDomains {
    fmt.Println(domain.GetDomain())
}
```

### Error Handling

API calls that return non-success status codes return typed errors from the root package. These are compatible with `errors.Is` and `errors.As`:

```go
import (
    "errors"
    "fmt"

    myorganization "github.com/auth0/myorganization-go"
    "github.com/auth0/myorganization-go/core"
)

response, err := c.Organization.Domains.Create(ctx, request)
if err != nil {
    // Match a specific error type.
    var notFound *myorganization.NotFoundError
    if errors.As(err, &notFound) {
        fmt.Println(notFound.Body)
        return
    }

    var unauthorized *myorganization.UnauthorizedError
    if errors.As(err, &unauthorized) {
        fmt.Println(unauthorized.Body)
        return
    }

    // Match any API error to inspect status code.
    var apiErr *core.APIError
    if errors.As(err, &apiErr) {
        fmt.Println("Status:", apiErr.StatusCode)
        return
    }

    return err
}
```

Available error types:

| Type | Status Code | Description |
|---|---|---|
| `BadRequestError` | 400 | Invalid request body |
| `UnauthorizedError` | 401 | Token missing, invalid, or expired |
| `ForbiddenError` | 403 | Insufficient scope |
| `NotFoundError` | 404 | Resource not found |
| `ConflictError` | 409 | Resource already exists |
| `TooManyRequestsError` | 429 | Rate limit exceeded |

### Explicit Null Values

By default, fields with zero/nil values are omitted from the JSON payload. To explicitly send `null` for a field, use the `Set*` methods:

```go
details := &myorganization.OrgDetails{}

// This sends {"display_name": null} instead of omitting the field entirely.
details.SetDisplayName(nil)

// This sends {"display_name": "My Org"}.
details.SetDisplayName(myorganization.String("My Org"))

_, err := c.OrganizationDetails.Update(ctx, details)
```

### Pointer Helpers

The SDK provides helper functions for creating pointers to primitive values, useful when constructing request bodies with optional fields:

```go
import myorganization "github.com/auth0/myorganization-go"

request := &myorganization.CreateClientRequestContent{
    Name:        "my-client",
    Description: myorganization.String("A description"),  // *string
    AppType:     myorganization.String("non_interactive"),
}
```

Available helpers: `String`, `Bool`, `Int`, `Int64`, `Float64`, `Time`, `UUID`, and more.

### Retries

The SDK automatically retries requests with exponential backoff on the following status codes:

- `408` (Timeout)
- `429` (Too Many Requests)
- `5XX` (Internal Server Errors)

The default retry limit is 2 attempts. The `Retry-After` header is respected when present. Configure via `option.WithMaxAttempts`:

```go
c, err := client.New(
    "<YOUR_AUTH0_DOMAIN>",
    option.WithStaticToken("<TOKEN>"),
    option.WithMaxAttempts(5),
)
```

### Timeouts

Use the standard `context` library to set per-request timeouts:

```go
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

response, err := c.Organization.Domains.Create(ctx, request)
```

### Base URL

The base URL defaults to `https://{domain}/my-org`. Override it with `option.WithBaseURL`:

```go
c, err := client.New(
    "<YOUR_AUTH0_DOMAIN>",
    option.WithStaticToken("<TOKEN>"),
    option.WithBaseURL("https://custom.example.com/api"),
)
```

### Organization

Use `option.WithOrganization` when your application has no [default organization](https://auth0.com/docs/manage-users/organizations/organizations-for-m2m-applications/configure-your-application-for-m2m-access#set-default-organization-for-an-application) configured, or when it has access to multiple organizations and the SDK needs to know which one to target. Without this option and no default organization, token requests will fail with an **"An organization is required"** error.

The value is sent as the `organization` parameter in OAuth2 token requests (client credentials and private key JWT flows).

```go
c, err := client.New(
    "<YOUR_AUTH0_DOMAIN>",
    option.WithClientCredentials(ctx, clientID, clientSecret),
    option.WithOrganization("org_abc123"), // organization ID or name
)
```

### Debug Logging

Use `option.WithDebug` to enable HTTP request/response debug logging. When enabled, every outgoing request and incoming response is logged via `log.Printf`. Sensitive headers (`Authorization`, `Cookie`, etc.) are redacted automatically.

```go
c, err := client.New(
    "<YOUR_AUTH0_DOMAIN>",
    option.WithStaticToken("<TOKEN>"),
    option.WithDebug(true),
)
```

### Telemetry

The SDK sends an `Auth0-Client` header on every request containing the SDK name, version, and Go runtime version (base64-encoded JSON). The `User-Agent` header is set to `MyOrganization-Go/{version}`.

To **disable** the `Auth0-Client` header:

```go
c, err := client.New(
    "<YOUR_AUTH0_DOMAIN>",
    option.WithStaticToken("<TOKEN>"),
    option.WithNoAuth0ClientInfo(),
)
```

To **add custom entries** to the `Auth0-Client` header's `env` map (useful for framework-level telemetry):

```go
c, err := client.New(
    "<YOUR_AUTH0_DOMAIN>",
    option.WithStaticToken("<TOKEN>"),
    option.WithAuth0ClientEnvEntry("myapp", "1.0.0"),
)
```

## Feedback

### Contributing

We appreciate feedback and contribution to this repo! Before you get started, please see the following:

- [Auth0's General Contribution Guidelines](https://github.com/auth0/open-source-template/blob/master/GENERAL-CONTRIBUTING.md)
- [Auth0's Code of Conduct Guidelines](https://github.com/auth0/open-source-template/blob/master/CODE-OF-CONDUCT.md)

While we value open-source contributions to this SDK, this library is generated programmatically. Additions made directly to this library would have to be moved over to our generation code, otherwise they would be overwritten upon the next generated release. Feel free to open a PR as a proof of concept, but know that we will not be able to merge it as-is. We suggest opening an issue first to discuss with us!

### Raise an Issue

To provide feedback or report a bug, please [raise an issue on our issue tracker](https://github.com/auth0/myorganization-go/issues).

### Vulnerability Reporting

Please do not report security vulnerabilities on the public GitHub issue tracker. The [Responsible Disclosure Program](https://auth0.com/responsible-disclosure-policy) details the procedure for disclosing security issues.

---

<p align="center">
  <picture>
    <source media="(prefers-color-scheme: light)" srcset="https://cdn.auth0.com/website/sdks/logos/auth0_light_mode.png" width="150">
    <source media="(prefers-color-scheme: dark)" srcset="https://cdn.auth0.com/website/sdks/logos/auth0_dark_mode.png" width="150">
    <img alt="Auth0 Logo" src="https://cdn.auth0.com/website/sdks/logos/auth0_light_mode.png" width="150">
  </picture>
</p>

<p align="center">Auth0 is an easy to implement, adaptable authentication and authorization platform.<br />To learn more check out <a href="https://auth0.com/why-auth0">Why Auth0?</a></p>

<p align="center">This project is licensed under the Apache-2.0 license. See the <a href="./LICENSE"> LICENSE</a> file for more info.</p>
