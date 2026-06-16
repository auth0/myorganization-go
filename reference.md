# Reference
## OrganizationDetails
<details><summary><code>client.OrganizationDetails.Get() -> myorganization.GetOrganizationDetailsResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve details for this Organization, including display name and branding options. To learn more about Auth0 Organizations, read [Organizations](https://auth0.com/docs/manage-users/organizations).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.OrganizationDetails.Get(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.OrganizationDetails.Update(request) -> myorganization.UpdateOrganizationDetailsResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Update details for this Organization, such as display name and branding options. To learn more about Auth0 Organizations, read [Organizations](https://auth0.com/docs/manage-users/organizations).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &myorganization.OrgDetails{
        Name: myorganization.String(
            "testorg",
        ),
        DisplayName: myorganization.String(
            "Test Organization",
        ),
        Branding: &myorganization.OrgBranding{
            LogoURL: myorganization.String(
                "https://example.com/logo.png",
            ),
            Colors: &myorganization.OrgBrandingColors{
                Primary: "#000000",
                PageBackground: "#FFFFFF",
            },
        },
    }
client.OrganizationDetails.Update(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**request:** `myorganization.UpdateOrganizationDetailsRequestContent` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Organization Configuration
<details><summary><code>client.Organization.Configuration.Get() -> *myorganization.GetConfigurationResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve the My Organization API configuration. Returns only the `connection_deletion_behavior` and `allowed_strategies`. Identifier attributes such as `user_attribute_profile_id` and `connection_profile_id` are not included. Cache this information, as it does not change frequently.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Organization.Configuration.Get(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Organization Domains
<details><summary><code>client.Organization.Domains.List() -> *myorganization.ListOrganizationDomainsResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve a list of all pending and verified domains for this Organization.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &myorganization.ListOrganizationDomainsRequestParameters{
        From: myorganization.String(
            "from",
        ),
        Take: myorganization.Int(
            1,
        ),
    }
client.Organization.Domains.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**from:** `*string` — An optional cursor from which to start the selection (exclusive).
    
</dd>
</dl>

<dl>
<dd>

**take:** `*int` — Number of results per page. Defaults to 50.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Organization.Domains.Create(request) -> myorganization.CreateOrganizationDomainResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create a new domain for this Organization.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &myorganization.CreateOrganizationDomainRequestContent{
        Domain: "acme.com",
    }
client.Organization.Domains.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**domain:** `myorganization.OrgDomainName` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Organization.Domains.Get(DomainID) -> myorganization.GetOrganizationDomainResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve details of a domain specified by ID for this Organization.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Organization.Domains.Get(
        context.TODO(),
        "domain_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**domainID:** `myorganization.OrgDomainID` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Organization.Domains.Delete(DomainID) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Remove a domain specified by ID from this Organization.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Organization.Domains.Delete(
        context.TODO(),
        "domain_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**domainID:** `myorganization.OrgDomainID` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Organization IdentityProviders
<details><summary><code>client.Organization.IdentityProviders.List() -> *myorganization.ListIdentityProvidersResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve a list of all Identity Providers for this Organization.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Organization.IdentityProviders.List(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Organization.IdentityProviders.Create(request) -> myorganization.CreateIdentityProviderResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create a new Identity Provider for this Organization.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &myorganization.IdpKnownRequest{
        IdpAdfsRequest: &myorganization.IdpAdfsRequest{
            Name: "oidcIdp",
            Domains: []string{
                "mydomain.com",
            },
            DisplayName: myorganization.String(
                "OIDC IdP",
            ),
            ShowAsButton: myorganization.Bool(
                true,
            ),
            AssignMembershipOnLogin: myorganization.Bool(
                false,
            ),
            IsEnabled: myorganization.Bool(
                true,
            ),
            Options: &myorganization.IdpAdfsOptionsRequest{
                IdpAdfsOptionsRequestAdfsServer: &myorganization.IdpAdfsOptionsRequestAdfsServer{},
            },
        },
    }
client.Organization.IdentityProviders.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**request:** `myorganization.CreateIdentityProviderRequestContent` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Organization.IdentityProviders.Get(IdpID) -> myorganization.GetIdentityProviderResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve details of an Identity Provider specified by ID for this Organization.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Organization.IdentityProviders.Get(
        context.TODO(),
        "idp_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**idpID:** `myorganization.IdpID` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Organization.IdentityProviders.Delete(IdpID) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Delete an Identity Provider specified by ID from this Organization. This will remove the association and delete the underlying Identity Provider. Members will no longer be able to authenticate using this Identity Provider.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Organization.IdentityProviders.Delete(
        context.TODO(),
        "idp_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**idpID:** `myorganization.IdpID` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Organization.IdentityProviders.Update(IdpID, request) -> myorganization.UpdateIdentityProviderResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Update the details of an Identity Provider specified by ID for this Organization.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &myorganization.IdpUpdateKnownRequest{
        IdpAdfsUpdateRequest: &myorganization.IdpAdfsUpdateRequest{
            DisplayName: myorganization.String(
                "OIDC IdP",
            ),
            ShowAsButton: myorganization.Bool(
                true,
            ),
            AssignMembershipOnLogin: myorganization.Bool(
                false,
            ),
            IsEnabled: myorganization.Bool(
                true,
            ),
            Options: &myorganization.IdpAdfsOptionsRequest{
                IdpAdfsOptionsRequestAdfsServer: &myorganization.IdpAdfsOptionsRequestAdfsServer{},
            },
        },
    }
client.Organization.IdentityProviders.Update(
        context.TODO(),
        "idp_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**idpID:** `myorganization.IdpID` 
    
</dd>
</dl>

<dl>
<dd>

**request:** `myorganization.UpdateIdentityProviderRequestContent` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Organization.IdentityProviders.UpdateAttributes(IdpID, request) -> myorganization.GetIdentityProviderResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Refresh the attribute mapping for an Identity Provider specified by ID for this Organization. Mappings are reset to the admin-defined defaults.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := map[string]any{
        "key": "value",
    }
client.Organization.IdentityProviders.UpdateAttributes(
        context.TODO(),
        "idp_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**idpID:** `myorganization.IdpID` 
    
</dd>
</dl>

<dl>
<dd>

**request:** `map[string]any` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Organization.IdentityProviders.Detach(IdpID) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Remove an Identity Provider specified by ID from this Organization. This only removes the association; the underlying Identity Provider is not deleted. Members will no longer be able to authenticate using this Identity Provider.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Organization.IdentityProviders.Detach(
        context.TODO(),
        "idp_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**idpID:** `myorganization.IdpID` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Organization Members
<details><summary><code>client.Organization.Members.List() -> *myorganization.ListOrganizationMembersResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve a list of all members for this Organization. The `roles` field is only included for each member when the token also carries the `read:my_org:member_roles` scope; without that scope the `roles` field is omitted from the response.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &myorganization.ListOrganizationMembersRequestParameters{
        Fields: myorganization.String(
            "fields",
        ),
        IncludeFields: myorganization.Bool(
            true,
        ),
        From: myorganization.String(
            "from",
        ),
        Take: myorganization.Int(
            1,
        ),
    }
client.Organization.Members.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**fields:** `*string` — Comma-separated list of fields to include or exclude (based on value provided for include_fields) in the result. Leave empty to retrieve all fields.
    
</dd>
</dl>

<dl>
<dd>

**includeFields:** `*bool` — Whether specified fields are to be included (true) or excluded (false). Defaults to true
    
</dd>
</dl>

<dl>
<dd>

**from:** `*string` — An optional cursor from which to start the selection (exclusive).
    
</dd>
</dl>

<dl>
<dd>

**take:** `*int` — Number of results per page. Defaults to 50.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Organization.Members.Get(UserID) -> myorganization.GetOrganizationMemberResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve details of a member specified by user ID for this Organization.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &myorganization.GetOrganizationMemberRequestParameters{
        Fields: myorganization.String(
            "fields",
        ),
        IncludeFields: myorganization.Bool(
            true,
        ),
    }
client.Organization.Members.Get(
        context.TODO(),
        "user_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userID:** `myorganization.OrgMemberID` 
    
</dd>
</dl>

<dl>
<dd>

**fields:** `*string` — Comma-separated list of fields to include or exclude (based on value provided for include_fields) in the result. Leave empty to retrieve all fields.
    
</dd>
</dl>

<dl>
<dd>

**includeFields:** `*bool` — Whether specified fields are to be included (true) or excluded (false). Defaults to true
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Organization Memberships
<details><summary><code>client.Organization.Memberships.DeleteMemberships(request) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Remove one member from this Organization. The underlying user account is not deleted.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &myorganization.DeleteOrganizationMembershipsRequestParameters{
        Members: []myorganization.OrgMemberID{
            "auth0|1234567890",
        },
    }
client.Organization.Memberships.DeleteMemberships(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**members:** `[]myorganization.OrgMemberID` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Organization Invitations
<details><summary><code>client.Organization.Invitations.List() -> *myorganization.ListMembersInvitationsResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve a list of all member invitations for this Organization.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &myorganization.ListMemberInvitationsRequestParameters{
        Fields: myorganization.String(
            "fields",
        ),
        IncludeFields: myorganization.Bool(
            true,
        ),
        From: myorganization.String(
            "from",
        ),
        Take: myorganization.Int(
            1,
        ),
        Sort: myorganization.String(
            "sort",
        ),
    }
client.Organization.Invitations.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**fields:** `*string` — Comma-separated list of fields to include or exclude (based on value provided for include_fields) in the result. Leave empty to retrieve all fields. Note: you cannot filter on ticket_id and this value will only be returned when fields are not filtered.
    
</dd>
</dl>

<dl>
<dd>

**includeFields:** `*bool` — Whether specified fields are to be included (true) or excluded (false). Defaults to true
    
</dd>
</dl>

<dl>
<dd>

**from:** `*string` — An optional cursor from which to start the selection (exclusive).
    
</dd>
</dl>

<dl>
<dd>

**take:** `*int` — Number of results per page. Defaults to 50.
    
</dd>
</dl>

<dl>
<dd>

**sort:** `*string` — Field to sort by. Use field:order where order is 1 for ascending and -1 for descending. Defaults to created_at:-1
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Organization.Invitations.Create(request) -> myorganization.CreateMemberInvitationResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create one or more member invitations for this Organization. If an active invitation already exists for a user, generating a new invitation will automatically revoke any outstanding invitations for that user. Roles specified in the payload will be granted to the user upon acceptance of the invitation.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &myorganization.CreateMemberInvitationRequestContent{
        Invitees: []*myorganization.CreateMemberInvitationInvitee{
            &myorganization.CreateMemberInvitationInvitee{
                Email: "user@example.com",
                Roles: []myorganization.RoleID{
                    "rol_0000000000000001",
                },
            },
        },
        Inviter: &myorganization.MemberInvitationInviter{
            Name: myorganization.String(
                "Allison the Admin",
            ),
        },
        IdentityProviderID: myorganization.String(
            "con_2CZPv6IY0gWzDaQJ",
        ),
        TTLSec: myorganization.Int(
            3600,
        ),
    }
client.Organization.Invitations.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**auth0CustomDomain:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**invitees:** `[]*myorganization.CreateMemberInvitationInvitee` 
    
</dd>
</dl>

<dl>
<dd>

**inviter:** `*myorganization.MemberInvitationInviter` 
    
</dd>
</dl>

<dl>
<dd>

**identityProviderID:** `*string` — Identity provider identifier.
    
</dd>
</dl>

<dl>
<dd>

**ttlSec:** `*int` — Number of seconds for which the invitation is valid before expiration. If unspecified or set to 0, this value defaults to 604800 seconds (7 days). Max value: 2592000 seconds (30 days).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Organization.Invitations.Get(InvitationID) -> myorganization.GetMemberInvitationResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve details of a member invitation specified by ID for this Organization.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &myorganization.GetMemberInvitationRequestParameters{
        Fields: myorganization.String(
            "fields",
        ),
        IncludeFields: myorganization.Bool(
            true,
        ),
    }
client.Organization.Invitations.Get(
        context.TODO(),
        "invitation_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**invitationID:** `myorganization.InvitationID` 
    
</dd>
</dl>

<dl>
<dd>

**fields:** `*string` — Comma-separated list of fields to include or exclude (based on value provided for include_fields) in the result. Leave empty to retrieve all fields. Note: you cannot filter on ticket_id and this value will only be returned when fields are not filtered.
    
</dd>
</dl>

<dl>
<dd>

**includeFields:** `*bool` — Whether specified fields are to be included (true) or excluded (false). Defaults to true
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Organization.Invitations.Delete(InvitationID) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Revoke a member invitation specified by ID for this Organization.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Organization.Invitations.Delete(
        context.TODO(),
        "invitation_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**invitationID:** `myorganization.InvitationID` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Organization Roles
<details><summary><code>client.Organization.Roles.List() -> *myorganization.ListRolesResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve the list of roles available for binding to members and invitations for this Organization. Only roles made visible to this Organization by the Tenant Admin are returned.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &myorganization.ListRolesRequestParameters{
        From: myorganization.String(
            "from",
        ),
        Take: myorganization.Int(
            1,
        ),
        Name: myorganization.String(
            "name",
        ),
    }
client.Organization.Roles.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**from:** `*string` — An optional cursor from which to start the selection (exclusive).
    
</dd>
</dl>

<dl>
<dd>

**take:** `*int` — Number of results per page. Defaults to 50.
    
</dd>
</dl>

<dl>
<dd>

**name:** `*string` — An optional filter on the name (case-insensitive).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Organization Configuration IdentityProviders
<details><summary><code>client.Organization.Configuration.IdentityProviders.Get() -> myorganization.GetIdpConfigurationResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve the [Connection Profile](https://auth0.com/docs/authenticate/enterprise-connections/connection-profile) for this application. You should cache this information as it does not change frequently.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Organization.Configuration.IdentityProviders.Get(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Organization Domains Verify
<details><summary><code>client.Organization.Domains.Verify.Create(DomainID) -> myorganization.StartOrganizationDomainVerificationResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Initiate the verification process for a domain specified by ID for this Organization.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Organization.Domains.Verify.Create(
        context.TODO(),
        "domain_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**domainID:** `myorganization.OrgDomainID` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Organization Domains IdentityProviders
<details><summary><code>client.Organization.Domains.IdentityProviders.Get(DomainID) -> *myorganization.ListDomainIdentityProvidersResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve the list of Identity Providers associated with a domain specified by ID for this Organization.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Organization.Domains.IdentityProviders.Get(
        context.TODO(),
        "domain_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**domainID:** `myorganization.OrgDomainID` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Organization IdentityProviders Domains
<details><summary><code>client.Organization.IdentityProviders.Domains.Create(IdpID, request) -> *myorganization.CreateIdpDomainResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Associate a domain with an Identity Provider specified by ID for this Organization. The domain must be claimed and verified.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &myorganization.CreateIdpDomainRequestContent{
        Domain: "my-domain.com",
    }
client.Organization.IdentityProviders.Domains.Create(
        context.TODO(),
        "idp_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**idpID:** `myorganization.IdpID` 
    
</dd>
</dl>

<dl>
<dd>

**domain:** `myorganization.OrgDomainName` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Organization.IdentityProviders.Domains.Delete(IdpID, Domain) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Remove a domain specified by name from an Identity Provider specified by ID for this Organization.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Organization.IdentityProviders.Domains.Delete(
        context.TODO(),
        "idp_id",
        "domain",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**idpID:** `myorganization.IdpID` 
    
</dd>
</dl>

<dl>
<dd>

**domain:** `myorganization.OrgDomainName` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Organization IdentityProviders Provisioning
<details><summary><code>client.Organization.IdentityProviders.Provisioning.Get(IdpID) -> *myorganization.GetIDPProvisioningConfigResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve the Provisioning Configuration for an Identity Provider specified by ID for this Organization.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Organization.IdentityProviders.Provisioning.Get(
        context.TODO(),
        "idp_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**idpID:** `myorganization.IdpID` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Organization.IdentityProviders.Provisioning.Create(IdpID) -> *myorganization.CreateIDPProvisioningConfigResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create a new Provisioning Configuration for an Identity Provider specified by ID for this Organization.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Organization.IdentityProviders.Provisioning.Create(
        context.TODO(),
        "idp_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**idpID:** `myorganization.IdpID` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Organization.IdentityProviders.Provisioning.Delete(IdpID) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Delete the Provisioning Configuration for an Identity Provider specified by ID for this Organization.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Organization.IdentityProviders.Provisioning.Delete(
        context.TODO(),
        "idp_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**idpID:** `myorganization.IdpID` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Organization.IdentityProviders.Provisioning.UpdateAttributes(IdpID, request) -> *myorganization.GetIDPProvisioningConfigResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Refresh the attribute mapping for the Provisioning Configuration of an Identity Provider specified by ID for this Organization. Mappings are reset to the admin-defined defaults.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := map[string]any{
        "key": "value",
    }
client.Organization.IdentityProviders.Provisioning.UpdateAttributes(
        context.TODO(),
        "idp_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**idpID:** `myorganization.IdpID` 
    
</dd>
</dl>

<dl>
<dd>

**request:** `map[string]any` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Organization IdentityProviders Provisioning SCIMTokens
<details><summary><code>client.Organization.IdentityProviders.Provisioning.SCIMTokens.List(IdpID) -> *myorganization.ListIdpProvisioningSCIMTokensResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve a list of [SCIM tokens](https://auth0.com/docs/authenticate/protocols/scim/configure-inbound-scim#scim-endpoints-and-tokens) for the Provisioning Configuration of an Identity Provider specified by ID for this Organization.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Organization.IdentityProviders.Provisioning.SCIMTokens.List(
        context.TODO(),
        "idp_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**idpID:** `myorganization.IdpID` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Organization.IdentityProviders.Provisioning.SCIMTokens.Create(IdpID, request) -> myorganization.CreateIdpProvisioningSCIMTokenResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create a new SCIM token for the Provisioning Configuration of an Identity Provider specified by ID for this Organization.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &myorganization.CreateIdpProvisioningSCIMTokenRequestContent{
        TokenLifetime: myorganization.Int(
            86400,
        ),
    }
client.Organization.IdentityProviders.Provisioning.SCIMTokens.Create(
        context.TODO(),
        "idp_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**idpID:** `myorganization.IdpID` 
    
</dd>
</dl>

<dl>
<dd>

**tokenLifetime:** `*int` — Lifetime of the token in seconds. Do not set for non-expiring tokens.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Organization.IdentityProviders.Provisioning.SCIMTokens.Delete(IdpID, IdpSCIMTokenID) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Revoke a SCIM token specified by token ID for the Provisioning Configuration of an Identity Provider specified by ID for this Organization.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Organization.IdentityProviders.Provisioning.SCIMTokens.Delete(
        context.TODO(),
        "idp_id",
        "idp_scim_token_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**idpID:** `myorganization.IdpID` 
    
</dd>
</dl>

<dl>
<dd>

**idpSCIMTokenID:** `myorganization.IdpProvisioningSCIMTokenID` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Organization Members Roles
<details><summary><code>client.Organization.Members.Roles.List(UserID) -> *myorganization.GetOrganizationMemberRolesResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve a list of roles assigned to a member specified by ID for this Organization.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &myorganization.ListOrgMemberRolesRequestParameters{
        From: myorganization.String(
            "from",
        ),
        Take: myorganization.Int(
            1,
        ),
    }
client.Organization.Members.Roles.List(
        context.TODO(),
        "user_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userID:** `myorganization.OrgMemberID` 
    
</dd>
</dl>

<dl>
<dd>

**from:** `*string` — An optional cursor from which to start the selection (exclusive).
    
</dd>
</dl>

<dl>
<dd>

**take:** `*int` — Number of results per page. Defaults to 50.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Organization.Members.Roles.Assign(UserID, request) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Assign roles to a member specified by ID for this Organization.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &myorganization.OrganizationMemberRolesChangeRequestContent{
        RoleIDs: []myorganization.RoleID{
            "rol_SO2j0sFo9NFa3F9w",
        },
    }
client.Organization.Members.Roles.Assign(
        context.TODO(),
        "user_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userID:** `myorganization.OrgMemberID` 
    
</dd>
</dl>

<dl>
<dd>

**request:** `*myorganization.OrganizationMemberRolesChangeRequestContent` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Organization.Members.Roles.Unassign(UserID, request) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Remove roles from a member specified by ID for this Organization.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &myorganization.OrganizationMemberRolesChangeRequestContent{
        RoleIDs: []myorganization.RoleID{
            "rol_SO2j0sFo9NFa3F9w",
        },
    }
client.Organization.Members.Roles.Unassign(
        context.TODO(),
        "user_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userID:** `myorganization.OrgMemberID` 
    
</dd>
</dl>

<dl>
<dd>

**request:** `*myorganization.OrganizationMemberRolesChangeRequestContent` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

