// Package option provides functional options for configuring the MyOrganization
// SDK client and individual API requests.
//
// Options fall into three categories:
//
// # Authentication
//
// Exactly one authentication option must be provided when creating a client:
//
//   - [WithClientCredentials] / [WithClientCredentialsAndAudience] — OAuth2 client credentials (M2M)
//   - [WithPrivateKeyJWT] / [WithPrivateKeyJWTAndAudience] — Private Key JWT assertion
//   - [WithTokenSource] — custom [golang.org/x/oauth2.TokenSource]
//   - [WithStaticToken] / WithToken — static bearer token
//
// # HTTP Configuration
//
//   - WithBaseURL — override the default base URL
//   - WithHTTPClient — supply a custom [net/http.Client]
//   - WithHTTPHeader — attach additional headers to every request
//   - WithMaxAttempts — configure retry attempts (default: 2)
//   - WithQueryParameters — add query parameters
//   - WithBodyProperties — add extra JSON body fields
//
// # Telemetry
//
//   - [WithNoAuth0ClientInfo] — disable the Auth0-Client header
//   - [WithAuth0ClientEnvEntry] — add custom entries to the Auth0-Client env map
//
// # Debug
//
//   - [WithDebug] — enable HTTP request/response debug logging
//
// Options can be applied at the client level (affecting all requests). Some
// options may also be passed to individual API methods for per-request
// overrides. Authentication options based on client credentials, private-key
// JWT, or a [golang.org/x/oauth2.TokenSource] are only used when constructing
// a client and are ignored when passed per-request; per-request authentication
// overrides must use [WithToken] or [WithStaticToken].
package option
