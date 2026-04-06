// Package client provides the top-level [MyOrganization] client for the
// Auth0 MyOrganization API.
//
// # Creating a Client
//
// Use [New] to create a client with Auth0 domain-based configuration,
// automatic token management, and telemetry:
//
//	c, err := client.New(
//	    "mytenant.auth0.com",
//	    option.WithClientCredentials(ctx, clientID, clientSecret),
//	)
//
// # Resource Access
//
// The client exposes API resources through nested sub-clients:
//
//	c.OrganizationDetails                                  // Organization details
//	c.Organization.Domains                                 // Domain management
//	c.Organization.Domains.Verify                          // Domain verification
//	c.Organization.Domains.IdentityProviders               // Domain identity providers
//	c.Organization.Clients                                 // Client (application) management
//	c.Organization.ClientGrants                            // Client grant management
//	c.Organization.IdentityProviders                       // Identity provider management
//	c.Organization.IdentityProviders.Domains               // IdP domain mappings
//	c.Organization.IdentityProviders.Provisioning          // IdP provisioning
//	c.Organization.IdentityProviders.Provisioning.ScimTokens // SCIM token management
//	c.Organization.Configuration.APIs                      // API configuration
//	c.Organization.Configuration.IdentityProviders         // IdP configuration
//
// Each sub-client also exposes a WithRawResponse field for accessing full
// HTTP response details (status code, headers, and body).
package client
