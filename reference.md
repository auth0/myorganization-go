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

Retrieve details for an Organization.
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

Update the details of a specific Organization, such as display name and branding options.
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
                "http://example.com/logo.png",
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

Retrieve the configuration for the /my-org API. This will return all stored client information with the exception of attributes that are identifiers. Identifier attributes will be given their own endpoint that will return the full object. This will give the components all of the information they will need to be successful. The SDK provider for the components should manage fetching and caching this information for all components.
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

Lists all domains pending and verified for an organization.
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
client.Organization.Domains.List(
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

<details><summary><code>client.Organization.Domains.Create(request) -> myorganization.CreateOrganizationDomainResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create a new domain for an organization.
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

Retrieve a domain for an organization.
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

Remove a domain from this organization.
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

List the identity providers associated with this organization.
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

Create an identity provider associated with this organization.
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

Retrieve the details for one particular identity-provider.
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

Delete an identity provider from this organization.
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

Update an identity provider associated with this organization.
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

Triggers a refresh of attribute mappings on the identity provider by overriding it with the admin defined defaults. The endpoint doesn't accept any body parameters.
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

Delete underlying identity provider from this organization.
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

## Organization ClientGrants
<details><summary><code>client.Organization.ClientGrants.Create(request) -> *myorganization.CreateClientGrantResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create a new client grant for the provided client.
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
request := &myorganization.CreateClientGrantRequestContent{
        ClientID: "client_id",
        Audience: "audience",
        Scope: []string{
            "scope",
        },
    }
client.Organization.ClientGrants.Create(
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

**clientID:** `string` — ID of the client.
    
</dd>
</dl>

<dl>
<dd>

**audience:** `string` — The audience (API identifier) of this client grant.
    
</dd>
</dl>

<dl>
<dd>

**scope:** `[]string` — Scopes allowed for this client grant.
    
</dd>
</dl>

<dl>
<dd>

**subjectType:** `*myorganization.SubjectTypeEnum` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Organization Clients
<details><summary><code>client.Organization.Clients.List() -> *myorganization.ListClientsResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Lists all API clients associated with the developer organization. Clients are used to obtain credentials for programmatic API access.
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
client.Organization.Clients.List(
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

<details><summary><code>client.Organization.Clients.Create(request) -> *myorganization.CreateClientResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create a new API client for the developer organization. The client can be used to obtain access tokens for calling tenant APIs.
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
request := &myorganization.CreateClientRequestContent{
        Name: "name",
    }
client.Organization.Clients.Create(
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

**name:** `myorganization.ClientName` 
    
</dd>
</dl>

<dl>
<dd>

**description:** `*myorganization.ClientDescription` 
    
</dd>
</dl>

<dl>
<dd>

**appType:** `*myorganization.ClientAppTypeEnum` 
    
</dd>
</dl>

<dl>
<dd>

**tokenEndpointAuthMethod:** `*myorganization.ClientTokenEndpointAuthMethodEnum` 
    
</dd>
</dl>

<dl>
<dd>

**grantTypes:** `*myorganization.ClientGrantTypes` 
    
</dd>
</dl>

<dl>
<dd>

**jwtConfiguration:** `*myorganization.ClientJwtConfiguration` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Organization.Clients.Delete(ClientID) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Delete an API client from the organization.
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
client.Organization.Clients.Delete(
        context.TODO(),
        "client_id",
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

**clientID:** `myorganization.ClientID` — The ID of the client
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Organization Configuration APIs
<details><summary><code>client.Organization.Configuration.APIs.Get() -> *myorganization.GetAllowedAPIsResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve the list of allowed APIs/resource servers for this organization.
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
client.Organization.Configuration.APIs.Get(
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

## Organization Configuration IdentityProviders
<details><summary><code>client.Organization.Configuration.IdentityProviders.Get() -> myorganization.GetIdpConfigurationResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve the connection profile for the application. This will give the components all of the information they will need to be successful. The SDK provider for the components should manage fetching and caching this information for all components.
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

Get a verification text and start the domain verification process for a particular domain.
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

Retrieve the list of identity providers that have a specific organization domain alias.
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

Add a domain to the identity provider's list of domains for [Home Realm Discovery (HRD)](https://auth0.com/docs/get-started/architecture-scenarios/business-to-business/authentication#home-realm-discovery). The domain passed must be claimed and verified by this organization.
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

Remove a domain from an identity provider.
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

Retrieve the Provisioning configuration for this identity provider.
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

Create the Provisioning configuration for this identity provider.
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

Delete the Provisioning configuration for an identity provider.
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

Triggers a refresh of attribute mappings on the provisioning configuration by overriding it with the admin defined defaults. The endpoint doesn't accept any body parameters.
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

List the Provisioning SCIM tokens for this identity provider.
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

Create a Provisioning SCIM token for this identity provider.
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

Delete a Provisioning SCIM configuration for an identity provider.
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
