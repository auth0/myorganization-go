// Package transport builds the layered HTTP round-tripper chain used by the
// MyOrganization SDK client.
//
// The chain applies middleware in the following order:
//
//	base transport → Debug (if enabled) → User-Agent → Auth0-Client telemetry → OAuth2 (if applicable)
//
// Three OAuth2 token acquisition strategies are supported:
//
//   - Client credentials — uses [golang.org/x/oauth2/clientcredentials]
//   - Private Key JWT — signs a JWT assertion and exchanges it for an access token
//   - Custom token source — caller-provided [golang.org/x/oauth2.TokenSource]
//
// This package is internal and not intended for direct use by SDK consumers.
package transport
