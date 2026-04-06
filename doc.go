// Package myorganization provides a Go client for the Auth0 MyOrganization API.
//
// This SDK handles authentication, automatic token management, retries with
// exponential backoff, and telemetry out of the box.
//
// # Client Initialization
//
// Create a client using [client.New] with your Auth0 domain and one of the
// supported authentication options:
//
//	c, err := client.New(
//	    "mytenant.auth0.com",
//	    option.WithClientCredentials(ctx, clientID, clientSecret),
//	)
//
// Four authentication modes are supported (mutually exclusive):
//
//   - Client Credentials — [option.WithClientCredentials]
//   - Private Key JWT — [option.WithPrivateKeyJWT]
//   - Custom Token Source — [option.WithTokenSource]
//   - Static Token — [option.WithStaticToken]
//
// # Resource Hierarchy
//
// The top-level client exposes API resources through a nested structure:
//
//	c.OrganizationDetails          // Get / Update organization details
//	c.Organization.Domains         // List, Create, Get, Delete domains
//	c.Organization.Clients         // List, Create, Delete clients
//	c.Organization.ClientGrants    // Create, Delete client grants
//	c.Organization.IdentityProviders  // Manage identity providers
//	c.Organization.Configuration   // APIs and identity provider config
//
// # Error Handling
//
// All API methods return typed errors that can be inspected with [errors.As]:
//
//	var notFound *myorganization.NotFoundError
//	if errors.As(err, &notFound) {
//	    fmt.Println(notFound.Body)
//	}
//
// Available error types: [BadRequestError], [UnauthorizedError],
// [ForbiddenError], [NotFoundError], [ConflictError], [TooManyRequestsError].
//
// # Pointer Helpers
//
// Helper functions like [String], [Bool], [Int], and [Time] create pointers
// to primitive values for use with optional request fields.
//
// # Explicit Null Values
//
// Use Set* methods on request types to explicitly include null values in the
// JSON payload instead of omitting the field:
//
//	details := &myorganization.OrgDetails{}
//	details.SetDisplayName(nil) // sends {"display_name": null}
package myorganization
