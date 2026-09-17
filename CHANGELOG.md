# Change Log

## [v2.0.0](https://github.com/auth0/myorganization-go/tree/v2.0.0) (2026-09-17)

This is a major release that adds Organization user stores and member invitation roles, introduces bulk invitation revocation, and expands filtering on several list endpoints. It also changes the Go module path to `github.com/auth0/myorganization-go/v2`, so it requires updating your imports. Please review the breaking changes and the migration guide below before upgrading.

### Installation

```shell
go get github.com/auth0/myorganization-go/v2
```

**Requirements:** Go 1.25+

### Breaking Changes

- **Module path is now `github.com/auth0/myorganization-go/v2`.** All import paths must include the `/v2` suffix.
- **`Organization.Invitations` moved to the `organization/invitations/client` package.** Access through `client.Organization.Invitations` is unchanged, but any direct import of the old `organization/invitations` package path must be updated.
- **`Organization.Invitations.Delete` now revokes multiple invitations at once.** It calls `POST /delete-member-invitations` instead of `DELETE /member-invitations/{invitationId}`. The `invitationID` argument is removed and the method now takes `*DeleteMemberInvitationsRequestContent`, which carries an `Invitations []InvitationID` list.
- **`Organization.Domains.IdentityProviders.Get` was renamed to `List`.** The endpoint (`GET /domains/{domainId}/identity-providers`) and behavior are unchanged.
- **`Organization.IdentityProviders.List` now requires a request argument.** It takes `*ListOrganizationIdentityProvidersRequestParameters`, which supports the `member_access_level` and `is_enabled` filters.
- **`Organization.Members.Roles.Unassign` now calls `POST /members/{userId}/unassign-roles`** instead of `DELETE /members/{userId}/roles`.

### Features

- **User Stores** - List the user stores associated with an Organization via `Organization.UserStores.List`, with `member_access_level` and `is_enabled` filters.
- **Invitation Roles** - List the roles assigned to a member invitation via `Organization.Invitations.Roles.List`.
- **List totals** - `include_totals` filter added to `Organization.Members.List` and `Organization.Invitations.List` through the new `IncludeTotals` field (defaults to `false`).
- **Identity Provider access levels** - `access_level` and `member_access_level` fields added to the enterprise Identity Provider request and response types, backed by the new `OrganizationAccessLevelEnum` and `OrganizationMemberAccessLevelEnum` types.
- **Invitation routing** - `user_store_id` field added to `CreateMemberInvitationRequestContent` (`SetUserStoreID`).

### Fixes

- Reorder the `LICENSE` file so the canonical Apache 2.0 text comes first, which lets GitHub correctly detect the project as Apache-2.0 licensed.

### Maintenance

- Bump the Go toolchain to `go1.25.14` to clear five standard library advisories reported by `govulncheck` (GO-2026-6218, GO-2026-6090, GO-2026-5972, GO-2026-5856, GO-2026-5026).
- Pin `govulncheck` to `v1.7.0` in the Security workflow for compatibility with the pinned Go version.

### Migration Guide

**1. Update your import paths to `/v2`**

```shell
go get github.com/auth0/myorganization-go/v2
```

```go
// Before
import (
    "github.com/auth0/myorganization-go/client"
    "github.com/auth0/myorganization-go/option"
)

// After
import (
    "github.com/auth0/myorganization-go/v2/client"
    "github.com/auth0/myorganization-go/v2/option"
)
```

**2. Revoke invitations in bulk**

```go
// Before: one invitation by ID
err := client.Organization.Invitations.Delete(ctx, "uinv_0000000000000001")

// After: one or more invitations by ID
err := client.Organization.Invitations.Delete(ctx, &myorganization.DeleteMemberInvitationsRequestContent{
    Invitations: []myorganization.InvitationID{
        "uinv_0000000000000001",
        "uinv_0000000000000002",
    },
})
```

**3. Rename domain Identity Providers `Get` to `List`**

```go
// Before
resp, err := client.Organization.Domains.IdentityProviders.Get(ctx, "domain_id")

// After
resp, err := client.Organization.Domains.IdentityProviders.List(ctx, "domain_id")
```

**4. Pass request parameters to `IdentityProviders.List`**

```go
// Before
resp, err := client.Organization.IdentityProviders.List(ctx)

// After (pass nil for no filters)
resp, err := client.Organization.IdentityProviders.List(ctx, &myorganization.ListOrganizationIdentityProvidersRequestParameters{
    IsEnabled: myorganization.Bool(true),
})
```

**5. Member role unassignment now uses POST**

No code change is required. `Organization.Members.Roles.Unassign` keeps the same signature; only the underlying HTTP method and path changed.

## [v1.0.0](https://github.com/auth0/myorganization-go/tree/v1.0.0) (2026-06-16)

This is the first stable release of the Auth0 MyOrganization Go SDK. It promotes `v1.0.0-beta.0` to a stable `v1.0.0`, with new APIs for managing organization members, roles, memberships, and invitations.

### Features

- **Members** - List and get organization members with cursor-based pagination via `Organization.Members`
- **Member Roles** - List, assign, and unassign roles for a member via `Organization.Members.Roles`
- **Roles** - List organization roles via `Organization.Roles`
- **Memberships** - Remove organization memberships via `Organization.Memberships`
- **Invitations** - Full lifecycle (list, create, get, delete) for member invitations via `Organization.Invitations`
- Member-specific typed error codes in `organization/members/error_codes.go`

### Maintenance

- Bump the Go toolchain to `go1.25.11` to clear four standard library advisories reported by `govulncheck` (GO-2026-5039, GO-2026-5037, GO-2026-4971, GO-2026-4918)

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
