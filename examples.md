# Examples

This document provides practical examples for common use cases of the Auth0 Go SDK.

## Table of Contents

- [Authentication](#authentication)
  - [Client Credentials (M2M)](#client-credentials-m2m)
  - [Client Credentials with Custom Audience](#client-credentials-with-custom-audience)
  - [Private Key JWT](#private-key-jwt)
  - [Private Key JWT with Custom Audience](#private-key-jwt-with-custom-audience)
  - [Custom Token Source](#custom-token-source)
  - [Static Token](#static-token)
- [Organization](#organization)
- [Managing Domains](#managing-domains)
  - [List Domains with Pagination](#list-domains-with-pagination)
  - [Create a Domain](#create-a-domain)
  - [Delete a Domain](#delete-a-domain)
- [Managing Clients](#managing-clients)
  - [List Clients](#list-clients)
  - [Create a Client](#create-a-client)
- [Organization Details](#organization-details)
  - [Get Organization Details](#get-organization-details)
  - [Update Organization Details](#update-organization-details)
- [Raw Responses](#raw-responses)
- [Error Handling](#error-handling)
- [Explicit Null Values](#explicit-null-values)
- [Configuration Options](#configuration-options)
  - [Custom HTTP Client](#custom-http-client)
  - [Retries](#retries)
  - [Custom Headers](#custom-headers)
  - [Debug Logging](#debug-logging)
  - [Telemetry](#telemetry)

---

## Authentication

### Client Credentials (M2M)

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
        "mytenant.auth0.com",
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

### Client Credentials with Custom Audience

```go
c, err := client.New(
    "mytenant.auth0.com",
    option.WithClientCredentialsAndAudience(
        context.Background(),
        "<YOUR_CLIENT_ID>",
        "<YOUR_CLIENT_SECRET>",
        "https://custom-api.example.com/",
    ),
)
```

### Private Key JWT

```go
import (
    "context"
    "os"

    "github.com/auth0/myorganization-go/client"
    "github.com/auth0/myorganization-go/option"
)

privateKeyPEM, _ := os.ReadFile("private_key.pem")

c, err := client.New(
    "mytenant.auth0.com",
    option.WithPrivateKeyJWT(
        context.Background(),
        "<YOUR_CLIENT_ID>",
        string(privateKeyPEM),
        "RS256",
    ),
)
```

Supported signing algorithms: RS256, RS384, RS512, PS256, PS384, PS512, ES256, ES384, ES512.

### Private Key JWT with Custom Audience

```go
c, err := client.New(
    "mytenant.auth0.com",
    option.WithPrivateKeyJWTAndAudience(
        context.Background(),
        "<YOUR_CLIENT_ID>",
        string(privateKeyPEM),
        "RS256",
        "https://custom-api.example.com/",
    ),
)
```

### Custom Token Source

Provide your own `oauth2.TokenSource` for full control over how tokens are obtained, cached, and refreshed.

```go
import (
    "github.com/auth0/myorganization-go/client"
    "github.com/auth0/myorganization-go/option"
    "golang.org/x/oauth2"
)

c, err := client.New(
    "mytenant.auth0.com",
    option.WithTokenSource(
        oauth2.StaticTokenSource(&oauth2.Token{
            AccessToken: "<YOUR_TOKEN>",
            TokenType:   "Bearer",
        }),
    ),
)
```

### Static Token

Use when you already have a bearer token.

```go
c, err := client.New(
    "mytenant.auth0.com",
    option.WithStaticToken("<YOUR_API_TOKEN>"),
)
```

---

## Organization

Use `WithOrganization` when your application has no [default organization](https://auth0.com/docs/manage-users/organizations/organizations-for-m2m-applications/configure-your-application-for-m2m-access#set-default-organization-for-an-application) configured, or when it has access to multiple organizations and the SDK needs to know which one to target. Without this option and no default organization, token requests will fail with an **"An organization is required"** error.

```go
c, err := client.New(
    "mytenant.auth0.com",
    option.WithClientCredentials(ctx, clientID, clientSecret),
    option.WithOrganization("org_abc123"), // organization ID or name
)
```

---

## Managing Domains

### List Domains with Pagination

```go
import (
    "context"
    "fmt"
    "net/url"

    "github.com/auth0/myorganization-go/option"
)

ctx := context.Background()

// Fetch the first page.
response, err := c.Organization.Domains.List(ctx)
if err != nil {
    log.Fatal(err)
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
        log.Fatal(err)
    }
    for _, domain := range response.OrganizationDomains {
        fmt.Println(domain.GetDomain())
    }
}
```

### Create a Domain

```go
import myorganization "github.com/auth0/myorganization-go"

domain, err := c.Organization.Domains.Create(
    ctx,
    &myorganization.CreateOrganizationDomainRequestContent{
        Domain: "example.com",
    },
)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Created domain: %s (ID: %s)\n", domain.GetDomain(), domain.GetDomainId())
```

### Delete a Domain

```go
err := c.Organization.Domains.Delete(ctx, "<DOMAIN_ID>")
if err != nil {
    log.Fatal(err)
}
```

---

## Managing Clients

### List Clients

```go
clients, err := c.Organization.Clients.List(ctx)
if err != nil {
    log.Fatal(err)
}
for _, cl := range clients.Clients {
    fmt.Println(cl.GetClientId())
}
```

### Create a Client

```go
import myorganization "github.com/auth0/myorganization-go"

client, err := c.Organization.Clients.Create(
    ctx,
    &myorganization.CreateClientRequestContent{
        Name:        "my-client",
        Description: myorganization.String("A description"),
        AppType:     myorganization.String("non_interactive"),
    },
)
if err != nil {
    log.Fatal(err)
}
```

---

## Organization Details

### Get Organization Details

```go
details, err := c.OrganizationDetails.Get(ctx)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Org: %s (%s)\n", details.GetDisplayName(), details.GetName())
```

### Update Organization Details

```go
import myorganization "github.com/auth0/myorganization-go"

updated, err := c.OrganizationDetails.Update(ctx, &myorganization.OrgDetails{
    DisplayName: myorganization.String("My Organization"),
})
if err != nil {
    log.Fatal(err)
}
```

---

## Raw Responses

Access the full HTTP response (status code, headers, and parsed body) via `WithRawResponse`.

```go
raw, err := c.Organization.Domains.WithRawResponse.List(ctx)
if err != nil {
    log.Fatal(err)
}

fmt.Println("Status:", raw.StatusCode)
fmt.Println("Rate limit remaining:", raw.Header.Get("X-RateLimit-Remaining"))

for _, domain := range raw.Body.OrganizationDomains {
    fmt.Println(domain.GetDomain())
}
```

---

## Error Handling

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

    log.Fatal(err)
}
```

Available error types: `BadRequestError` (400), `UnauthorizedError` (401), `ForbiddenError` (403), `NotFoundError` (404), `ConflictError` (409), `TooManyRequestsError` (429).

---

## Explicit Null Values

Send explicit `null` in JSON payloads using `Set*` methods.

```go
import myorganization "github.com/auth0/myorganization-go"

details := &myorganization.OrgDetails{}

// Sends {"display_name": null} instead of omitting the field.
details.SetDisplayName(nil)

// Sends {"display_name": "My Org"}.
details.SetDisplayName(myorganization.String("My Org"))
```

---

## Configuration Options

### Custom HTTP Client

```go
import (
    "net/http"
    "time"

    "github.com/auth0/myorganization-go/client"
    "github.com/auth0/myorganization-go/option"
)

c, err := client.New(
    "mytenant.auth0.com",
    option.WithStaticToken("<TOKEN>"),
    option.WithHTTPClient(&http.Client{Timeout: 10 * time.Second}),
)
```

### Retries

The SDK retries on 408, 429, and 5XX with exponential backoff. Default: 2 attempts.

```go
c, err := client.New(
    "mytenant.auth0.com",
    option.WithStaticToken("<TOKEN>"),
    option.WithMaxAttempts(5),
)
```

### Custom Headers

```go
response, err := c.Organization.Domains.Create(
    ctx,
    request,
    option.WithHTTPHeader(http.Header{
        "X-Request-Id": {"abc-123"},
    }),
)
```

### Debug Logging

Log every HTTP request and response. Sensitive headers are redacted automatically.

```go
c, err := client.New(
    "mytenant.auth0.com",
    option.WithStaticToken("<TOKEN>"),
    option.WithDebug(true),
)
```

### Telemetry

Disable the `Auth0-Client` header:

```go
c, err := client.New(
    "mytenant.auth0.com",
    option.WithStaticToken("<TOKEN>"),
    option.WithNoAuth0ClientInfo(),
)
```

Add custom entries to the `Auth0-Client` header's `env` map:

```go
c, err := client.New(
    "mytenant.auth0.com",
    option.WithStaticToken("<TOKEN>"),
    option.WithAuth0ClientEnvEntry("myapp", "1.0.0"),
)
```
