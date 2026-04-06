package option

import (
	"context"

	core "github.com/auth0/myorganization-go/core"
	"golang.org/x/oauth2"
)

// WithClientCredentials configures OAuth2 client credentials authentication.
// The SDK will automatically fetch, cache, and refresh access tokens
// using the client_credentials grant type.
//
// The audience defaults to https://{domain}/my-org/.
// To specify a custom audience, use WithClientCredentialsAndAudience instead.
func WithClientCredentials(ctx context.Context, clientID, clientSecret string) *core.ClientCredentialsOption {
	return &core.ClientCredentialsOption{
		Context:      ctx,
		ClientID:     clientID,
		ClientSecret: clientSecret,
	}
}

// WithClientCredentialsAndAudience configures OAuth2 client credentials authentication
// with a custom audience. The SDK will automatically fetch, cache, and refresh
// access tokens using the client_credentials grant type.
func WithClientCredentialsAndAudience(ctx context.Context, clientID, clientSecret, audience string) *core.ClientCredentialsOption {
	return &core.ClientCredentialsOption{
		Context:      ctx,
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Audience:     audience,
	}
}

// WithPrivateKeyJWT configures Private Key JWT authentication.
// The SDK will create a signed JWT assertion and exchange it for an access token
// via the OAuth2 client_credentials grant with client_assertion.
//
// The audience defaults to https://{domain}/my-org/.
// To specify a custom audience, use WithPrivateKeyJWTAndAudience instead.
//
// Supported signing algorithms: RS256, RS384, RS512, PS256, PS384, PS512, ES256, ES384, ES512.
func WithPrivateKeyJWT(ctx context.Context, clientID, privateKeyPEM, signingAlgorithm string) *core.PrivateKeyJwtOption {
	return &core.PrivateKeyJwtOption{
		Context:          ctx,
		ClientID:         clientID,
		PrivateKeyPEM:    privateKeyPEM,
		SigningAlgorithm: signingAlgorithm,
	}
}

// WithPrivateKeyJWTAndAudience configures Private Key JWT authentication with a custom audience.
// The SDK will create a signed JWT assertion and exchange it for an access token
// via the OAuth2 client_credentials grant with client_assertion.
//
// Supported signing algorithms: RS256, RS384, RS512, PS256, PS384, PS512, ES256, ES384, ES512.
func WithPrivateKeyJWTAndAudience(ctx context.Context, clientID, privateKeyPEM, signingAlgorithm, audience string) *core.PrivateKeyJwtOption {
	return &core.PrivateKeyJwtOption{
		Context:          ctx,
		ClientID:         clientID,
		PrivateKeyPEM:    privateKeyPEM,
		SigningAlgorithm: signingAlgorithm,
		Audience:         audience,
	}
}

// WithTokenSource configures a custom oauth2.TokenSource for authentication.
// The SDK will use the provided token source to obtain access tokens.
// This allows full control over how tokens are fetched, cached, and refreshed.
//
// The token source is mutually exclusive with other authentication modes
// (client credentials, private key JWT, and static token).
func WithTokenSource(tokenSource oauth2.TokenSource) *core.TokenSourceOption {
	return &core.TokenSourceOption{
		TokenSource: tokenSource,
	}
}

// WithStaticToken configures a static bearer token for authentication.
// This is an alias for WithToken — provided for symmetry with WithClientCredentials.
func WithStaticToken(token string) *core.TokenOption {
	return &core.TokenOption{
		Token: token,
	}
}

// WithNoAuth0ClientInfo disables the Auth0-Client telemetry header.
func WithNoAuth0ClientInfo() *core.NoAuth0ClientInfoOption {
	return &core.NoAuth0ClientInfoOption{}
}

// WithAuth0ClientEnvEntry adds a custom key-value entry to the Auth0-Client
// header's env map (e.g., WithAuth0ClientEnvEntry("myapp", "1.0.0")).
func WithAuth0ClientEnvEntry(key, value string) *core.Auth0ClientEnvEntryOption {
	return &core.Auth0ClientEnvEntryOption{
		Key:   key,
		Value: value,
	}
}

// WithOrganization sets the organization name or ID. When provided, this
// value is sent as the "organization" parameter in OAuth2 token requests
// (client credentials and private key JWT flows).
//
// It does not affect base URL or audience derivation.
func WithOrganization(organization string) *core.OrganizationOption {
	return &core.OrganizationOption{
		Organization: organization,
	}
}

// WithDebug enables HTTP request/response debug logging. When enabled, the SDK
// logs every outgoing request and incoming response via log.Printf. Sensitive
// headers (Authorization, Cookie, etc.) are redacted automatically.
func WithDebug(debug bool) *core.DebugOption {
	return &core.DebugOption{
		Debug: debug,
	}
}
