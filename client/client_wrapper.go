package client

import (
	core "github.com/auth0/myorganization-go/core"
	"github.com/auth0/myorganization-go/internal/auth0"
	"github.com/auth0/myorganization-go/internal/transport"
	option "github.com/auth0/myorganization-go/option"
)

// New creates a new MyOrganization client with Auth0 authentication and telemetry.
// It accepts the Auth0 tenant domain and options using the same option.RequestOption
// type as NewWithOptions, plus the new auth/telemetry options.
//
// Domain is sanitized automatically (https:// prefix and trailing slashes stripped).
//
// Authentication modes (mutually exclusive):
//   - option.WithClientCredentials(ctx, clientID, clientSecret) — M2M OAuth2
//   - option.WithPrivateKeyJWT(ctx, clientID, privateKeyPEM, alg) — Private Key JWT
//   - option.WithTokenSource(tokenSource) — custom oauth2.TokenSource
//   - option.WithStaticToken(token) or option.WithToken(token) — static bearer token
func New(domain string, opts ...option.RequestOption) (*MyOrganization, error) {
	options := core.NewRequestOptions(opts...)

	if err := auth0.ValidateOptions(domain, options); err != nil {
		return nil, err
	}

	sanitizedDomain := auth0.SanitizeDomain(domain)
	baseURL, tokenURL, audience := auth0.DeriveURLs(sanitizedDomain, options)

	httpClient := transport.CloneHTTPClient(options.HTTPClient)

	chain, err := transport.BuildChain(transport.ChainConfig{
		Base:              httpClient.Transport,
		ClientID:          options.ClientID,
		ClientSecret:      options.ClientSecret,
		TokenURL:          tokenURL,
		Audience:          audience,
		Organization:      options.Organization,
		OAuthContext:      options.OAuthContext,
		PrivateKeyPEM:     options.PrivateKeyPEM,
		SigningAlgorithm:  options.SigningAlgorithm,
		TokenSource:       options.TokenSource,
		NoAuth0ClientInfo: options.NoAuth0ClientInfo,
		Auth0ClientEnv:    options.Auth0ClientEnv,
		Debug:             options.Debug,
	})
	if err != nil {
		return nil, err
	}
	httpClient.Transport = chain

	fernOpts := []option.RequestOption{
		option.WithHTTPClient(httpClient),
		option.WithBaseURL(baseURL),
	}
	if options.MaxAttempts > 0 {
		fernOpts = append(fernOpts, option.WithMaxAttempts(options.MaxAttempts))
	}
	if options.Token != "" {
		fernOpts = append(fernOpts, option.WithToken(options.Token))
	}
	if len(options.HTTPHeader) > 0 {
		fernOpts = append(fernOpts, option.WithHTTPHeader(options.HTTPHeader))
	}

	return NewWithOptions(fernOpts...), nil
}
