# MyOrganization Go SDK Migration Guide

This document covers the changes required to move between major versions of myorganization-go.

- [Migrating from v1 to v2](#migrating-from-v1-to-v2)

---

# Migrating from v1 to v2

**Please review this section thoroughly to understand the changes required to migrate from myorganization-go v1 to myorganization-go v2.**

## Overview

v2 keeps the client initialization, package layout, and option pattern from v1. The one change every application must make is the module import path, which moves from `github.com/auth0/myorganization-go` to `github.com/auth0/myorganization-go/v2` as required for a Go major version. Beyond that, the breaking changes are focused on a small number of methods where signatures changed to support new behavior, such as bulk invitation revocation and additional list filters. Most applications will only need to update their import paths and the specific call sites listed below.

## v2 Breaking Changes

- [Module Import Path](#module-import-path)
- [Invitations Package Location](#invitations-package-location)
- [Bulk Invitation Revocation](#bulk-invitation-revocation)
- [Domain Identity Providers List](#domain-identity-providers-list)
- [Identity Providers List Parameters](#identity-providers-list-parameters)
- [Member Role Unassignment](#member-role-unassignment)

### Module Import Path

The module path changes from `github.com/auth0/myorganization-go` to `github.com/auth0/myorganization-go/v2`. Update your `go get` command and every import in your codebase.

<table>
<tr>
<th>v1</th>
<th>v2</th>
</tr>
<tr>
<td>

```go
// go get github.com/auth0/myorganization-go

import (
    "github.com/auth0/myorganization-go/client"
    "github.com/auth0/myorganization-go/option"
)
```

</td>
<td>

```go
// go get github.com/auth0/myorganization-go/v2

import (
    "github.com/auth0/myorganization-go/v2/client"
    "github.com/auth0/myorganization-go/v2/option"
)
```

</td>
</tr>
</table>

The quickest way to update an existing codebase is a find-and-replace of `github.com/auth0/myorganization-go` with `github.com/auth0/myorganization-go/v2` across your Go files, followed by `go mod tidy`.

### Invitations Package Location

The invitations resource moved from the `organization/invitations` package to `organization/invitations/client`. If you access invitations through the top-level client (`client.Organization.Invitations`), no change is required. Only update code that imports the invitations package directly.

<table>
<tr>
<th>v1</th>
<th>v2</th>
</tr>
<tr>
<td>

```go
import invitations "github.com/auth0/myorganization-go/organization/invitations"
```

</td>
<td>

```go
import invitations "github.com/auth0/myorganization-go/v2/organization/invitations/client"
```

</td>
</tr>
</table>

### Bulk Invitation Revocation

`Organization.Invitations.Delete` now revokes one or more invitations in a single call. It calls `POST /delete-member-invitations` instead of `DELETE /member-invitations/{invitationId}`. The `invitationID` argument is removed and the method now takes `*DeleteMemberInvitationsRequestContent`, which carries an `Invitations []InvitationID` list.

<table>
<tr>
<th>v1</th>
<th>v2</th>
</tr>
<tr>
<td>

```go
err := client.Organization.Invitations.Delete(
    ctx,
    "uinv_0000000000000001",
)
```

</td>
<td>

```go
err := client.Organization.Invitations.Delete(
    ctx,
    &myorganization.DeleteMemberInvitationsRequestContent{
        Invitations: []myorganization.InvitationID{
            "uinv_0000000000000001",
            "uinv_0000000000000002",
        },
    },
)
```

</td>
</tr>
</table>

### Domain Identity Providers List

`Organization.Domains.IdentityProviders.Get` was renamed to `List`. The endpoint (`GET /domains/{domainId}/identity-providers`) and behavior are unchanged.

<table>
<tr>
<th>v1</th>
<th>v2</th>
</tr>
<tr>
<td>

```go
resp, err := client.Organization.Domains.IdentityProviders.Get(
    ctx,
    "domain_id",
)
```

</td>
<td>

```go
resp, err := client.Organization.Domains.IdentityProviders.List(
    ctx,
    "domain_id",
)
```

</td>
</tr>
</table>

### Identity Providers List Parameters

`Organization.IdentityProviders.List` now requires a request argument, `*ListOrganizationIdentityProvidersRequestParameters`, which supports the `member_access_level` and `is_enabled` filters. Pass `nil` when you do not need any filters.

<table>
<tr>
<th>v1</th>
<th>v2</th>
</tr>
<tr>
<td>

```go
resp, err := client.Organization.IdentityProviders.List(ctx)
```

</td>
<td>

```go
resp, err := client.Organization.IdentityProviders.List(
    ctx,
    &myorganization.ListOrganizationIdentityProvidersRequestParameters{
        IsEnabled: myorganization.Bool(true),
    },
)

// or, when no filters are needed:
resp, err := client.Organization.IdentityProviders.List(ctx, nil)
```

</td>
</tr>
</table>

### Member Role Unassignment

`Organization.Members.Roles.Unassign` now calls `POST /members/{userId}/unassign-roles` instead of `DELETE /members/{userId}/roles`. The method signature is unchanged, so no code change is required.

<table>
<tr>
<th>v1</th>
<th>v2</th>
</tr>
<tr>
<td>

```go
err := client.Organization.Members.Roles.Unassign(
    ctx,
    "user_id",
    request,
)
```

</td>
<td>

```go
// No code change required. Same call, updated HTTP method and path.
err := client.Organization.Members.Roles.Unassign(
    ctx,
    "user_id",
    request,
)
```

</td>
</tr>
</table>
