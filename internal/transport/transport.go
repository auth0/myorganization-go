package transport

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	myorganization "github.com/auth0/myorganization-go"
	"github.com/auth0/myorganization-go/internal/telemetry"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

// ChainConfig holds the parameters needed to build the full transport chain.
type ChainConfig struct {
	Base              http.RoundTripper
	ClientID          string
	ClientSecret      string
	TokenURL          string
	Audience          string
	Organization      string
	OAuthContext      context.Context
	PrivateKeyPEM     string
	SigningAlgorithm  string
	TokenSource       oauth2.TokenSource
	NoAuth0ClientInfo bool
	Auth0ClientEnv    map[string]string
	Debug             bool
}

// BuildChain constructs the layered RoundTripper chain:
//
//	base → UserAgent → Auth0ClientInfo (unless disabled) → oauth2 (if token source, private key JWT, or M2M)
func BuildChain(cfg ChainConfig) (http.RoundTripper, error) {
	base := cfg.Base
	if base == nil {
		base = http.DefaultTransport
	}

	base = DebugTransport(base, cfg.Debug)

	var chain http.RoundTripper = &UserAgentTransport{
		Base:      base,
		UserAgent: fmt.Sprintf("MyOrganization-Go/%s", myorganization.SDKVersion),
	}

	if !cfg.NoAuth0ClientInfo {
		info := telemetry.BuildAuth0ClientInfo(cfg.Auth0ClientEnv)
		headerValue, err := telemetry.EncodeAuth0ClientInfo(info)
		if err != nil {
			return nil, fmt.Errorf("auth0: failed to encode telemetry header: %w", err)
		}
		chain = &Auth0ClientInfoTransport{
			Base:        chain,
			HeaderValue: headerValue,
		}
	}

	oauthCtx := cfg.OAuthContext
	if oauthCtx == nil {
		oauthCtx = context.Background()
	}

	switch {
	case cfg.TokenSource != nil:
		// Custom token source authentication.
		chain = &oauth2.Transport{
			Source: cfg.TokenSource,
			Base:   chain,
		}
	case cfg.PrivateKeyPEM != "":
		// Private Key JWT authentication.
		src, err := NewPrivateKeyJwtTokenSource(PrivateKeyJwtConfig{
			ClientID:         cfg.ClientID,
			PrivateKeyPEM:    cfg.PrivateKeyPEM,
			SigningAlgorithm: cfg.SigningAlgorithm,
			TokenURL:         cfg.TokenURL,
			Audience:         cfg.Audience,
			Organization:     cfg.Organization,
			Context:          oauthCtx,
		})
		if err != nil {
			return nil, err
		}
		chain = &oauth2.Transport{
			Source: src,
			Base:   chain,
		}
	case cfg.ClientID != "" && cfg.ClientSecret != "":
		// Client credentials (M2M) authentication.
		endpointParams := url.Values{
			"audience": {cfg.Audience},
		}
		if cfg.Organization != "" {
			endpointParams.Set("organization", cfg.Organization)
		}
		oauthCfg := &clientcredentials.Config{
			ClientID:       cfg.ClientID,
			ClientSecret:   cfg.ClientSecret,
			TokenURL:       cfg.TokenURL,
			EndpointParams: endpointParams,
		}
		chain = &oauth2.Transport{
			Source: oauthCfg.TokenSource(oauthCtx),
			Base:   chain,
		}
	}

	return chain, nil
}

// UserAgentTransport sets the User-Agent header on every outgoing request.
type UserAgentTransport struct {
	Base      http.RoundTripper
	UserAgent string
}

// RoundTrip sets the User-Agent header and delegates to the base transport.
func (t *UserAgentTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	r := req.Clone(req.Context())
	r.Header.Set("User-Agent", t.UserAgent)
	return t.Base.RoundTrip(r)
}

// Auth0ClientInfoTransport sets the Auth0-Client telemetry header on every outgoing request.
type Auth0ClientInfoTransport struct {
	Base        http.RoundTripper
	HeaderValue string // pre-computed base64-encoded value
}

// RoundTrip sets the Auth0-Client header and delegates to the base transport.
func (t *Auth0ClientInfoTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	r := req.Clone(req.Context())
	r.Header.Set("Auth0-Client", t.HeaderValue)
	return t.Base.RoundTrip(r)
}

// HTTPClient is a minimal interface matching core.HTTPClient.
type HTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}

// CloneHTTPClient returns a shallow clone of the provided HTTP client.
// If c is nil, a shallow clone of http.DefaultClient is returned. If c is an
// *http.Client, a shallow clone of c is returned. Otherwise, a new *http.Client
// is returned whose Transport wraps c via adapterTransport.
func CloneHTTPClient(c HTTPClient) *http.Client {
	if c == nil {
		clone := *http.DefaultClient
		return &clone
	}
	if hc, ok := c.(*http.Client); ok {
		clone := *hc
		return &clone
	}
	return &http.Client{
		Transport: &adapterTransport{client: c},
	}
}

// adapterTransport adapts an HTTPClient to http.RoundTripper by delegating to
// its Do method. Note: Do may follow redirects or apply cookie jars, which
// differs from pure RoundTripper semantics. In practice this path is only used
// when the caller provides a custom core.HTTPClient that is not an *http.Client.
type adapterTransport struct {
	client HTTPClient
}

func (t *adapterTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	return t.client.Do(req)
}
