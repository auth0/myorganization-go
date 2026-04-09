# Change Log

## [v1.0.0-beta.0](https://github.com/auth0/myorganization-go/tree/v1.0.0-beta.0) (2026-04-09)

This is the first beta release of the Auth0 MyOrganization Go SDK, providing a fully-featured Go client for managing Auth0 Organizations.

### Installation

```shell
go get github.com/auth0/myorganization-go
```

**Requirements:** Go 1.25+

### Usage

```go
import (
    "context"

    "github.com/auth0/myorganization-go/client"
    "github.com/auth0/myorganization-go/option"
)

c, err := client.New(
    "mytenant.auth0.com",
    option.WithClientCredentials(
        context.Background(),
        "<YOUR_CLIENT_ID>",
        "<YOUR_CLIENT_SECRET>",
    ),
)
```

### Authentication

- **Client Credentials (M2M)** with automatic token caching and refresh via `option.WithClientCredentials`
- **Private Key JWT** using signed JWT assertions (RS256, PS256, ES256, and more) via `option.WithPrivateKeyJWT`
- **Custom Token Source** for full control over token management via `option.WithTokenSource`
- **Static Token** for pre-existing bearer tokens via `option.WithStaticToken`

### Supported APIs

- **Organization Details** - Get and update organization details, display name, and branding
- **Organization Configuration** - Retrieve API configuration and connection profile settings
- **Domains** - List, create, get, delete, and verify organization domains with cursor-based pagination
- **Identity Providers** - Full lifecycle management including create, update, delete, detach, and attribute refresh
- **Identity Provider Domains** - Associate and remove verified domains from identity providers
- **Provisioning** - Manage provisioning configurations and SCIM tokens (list, create, revoke)

### SDK Features

- Automatic retry with exponential backoff on 408, 429, and 5XX (default: 2 attempts, configurable via `option.WithMaxAttempts`)
- Configurable streaming buffer size via `option.WithMaxStreamBufSize`
- Typed error handling with `BadRequestError`, `UnauthorizedError`, `ForbiddenError`, `NotFoundError`, and `TooManyRequestsError`
- Raw HTTP response access via `WithRawResponse` for status codes, headers, and rate limit info
- Explicit null values in JSON payloads using `Set*` methods
- Request-level option overrides for per-call configuration
- Debug logging with automatic sensitive header redaction via `option.WithDebug`
- Telemetry via `Auth0-Client` header (opt-out with `option.WithNoAuth0ClientInfo`)

For usage examples, see [examples.md](./examples.md). For full API reference, see [reference.md](./reference.md).
