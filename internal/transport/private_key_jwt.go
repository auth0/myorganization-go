package transport

import (
	"context"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/lestrrat-go/jwx/v3/jwa"
	"github.com/lestrrat-go/jwx/v3/jwk"
	"github.com/lestrrat-go/jwx/v3/jwt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

// allowedAlgorithms is the set of signing algorithms supported for Private Key JWT.
var allowedAlgorithms = map[string]jwa.SignatureAlgorithm{
	"RS256": jwa.RS256(),
	"RS384": jwa.RS384(),
	"RS512": jwa.RS512(),
	"PS256": jwa.PS256(),
	"PS384": jwa.PS384(),
	"PS512": jwa.PS512(),
	"ES256": jwa.ES256(),
	"ES384": jwa.ES384(),
	"ES512": jwa.ES512(),
}

// PrivateKeyJwtConfig holds the configuration for Private Key JWT authentication.
type PrivateKeyJwtConfig struct {
	ClientID         string
	PrivateKeyPEM    string
	SigningAlgorithm string
	TokenURL         string
	Audience         string
	Organization     string
	Context          context.Context
}

// privateKeyJwtTokenSource implements oauth2.TokenSource by creating a fresh
// JWT client assertion on each Token() call and exchanging it via the
// client_credentials grant with client_assertion parameters.
type privateKeyJwtTokenSource struct {
	cfg PrivateKeyJwtConfig
	key jwk.Key
	alg jwa.SignatureAlgorithm
}

// NewPrivateKeyJwtTokenSource creates an oauth2.TokenSource that uses Private Key JWT
// for authentication. The returned source caches tokens and reuses them until expiry.
func NewPrivateKeyJwtTokenSource(cfg PrivateKeyJwtConfig) (oauth2.TokenSource, error) {
	alg, ok := allowedAlgorithms[strings.ToUpper(cfg.SigningAlgorithm)]
	if !ok {
		return nil, fmt.Errorf("auth0: unsupported signing algorithm %q; supported: RS256, RS384, RS512, PS256, PS384, PS512, ES256, ES384, ES512", cfg.SigningAlgorithm)
	}

	key, err := parsePrivateKey([]byte(cfg.PrivateKeyPEM))
	if err != nil {
		return nil, fmt.Errorf("auth0: failed to parse private key: %w", err)
	}

	src := &privateKeyJwtTokenSource{
		cfg: cfg,
		key: key,
		alg: alg,
	}

	// Wrap with ReuseTokenSource so tokens are cached until expiry.
	return oauth2.ReuseTokenSource(nil, src), nil
}

// Token creates a fresh JWT assertion and exchanges it for an access token.
func (s *privateKeyJwtTokenSource) Token() (*oauth2.Token, error) {
	assertion, err := s.createClientAssertion()
	if err != nil {
		return nil, fmt.Errorf("auth0: failed to create client assertion: %w", err)
	}

	ctx := s.cfg.Context
	if ctx == nil {
		ctx = context.Background()
	}

	endpointParams := url.Values{
		"audience":              {s.cfg.Audience},
		"client_assertion":      {assertion},
		"client_assertion_type": {"urn:ietf:params:oauth:client-assertion-type:jwt-bearer"},
	}
	if s.cfg.Organization != "" {
		endpointParams.Set("organization", s.cfg.Organization)
	}
	oauthCfg := clientcredentials.Config{
		ClientID:       s.cfg.ClientID,
		TokenURL:       s.cfg.TokenURL,
		EndpointParams: endpointParams,
	}

	return oauthCfg.TokenSource(ctx).Token()
}

// createClientAssertion builds and signs a JWT assertion for the client_assertion parameter.
func (s *privateKeyJwtTokenSource) createClientAssertion() (string, error) {
	now := time.Now()

	// Use token URL base (scheme + host) as audience for the assertion.
	aud := audienceFromTokenURL(s.cfg.TokenURL)

	token, err := jwt.NewBuilder().
		Issuer(s.cfg.ClientID).
		Subject(s.cfg.ClientID).
		Audience([]string{aud}).
		JwtID(uuid.New().String()).
		IssuedAt(now).
		NotBefore(now).
		Expiration(now.Add(2 * time.Minute)).
		Build()
	if err != nil {
		return "", fmt.Errorf("failed to build JWT: %w", err)
	}

	signed, err := jwt.Sign(token, jwt.WithKey(s.alg, s.key))
	if err != nil {
		return "", fmt.Errorf("failed to sign JWT: %w", err)
	}

	return string(signed), nil
}

// parsePrivateKey parses a PEM-encoded private key into a jwk.Key.
func parsePrivateKey(pemData []byte) (jwk.Key, error) {
	key, err := jwk.ParseKey(pemData, jwk.WithPEM(true))
	if err != nil {
		return nil, err
	}
	return key, nil
}

// audienceFromTokenURL extracts the base URL (scheme + host) from a token URL.
func audienceFromTokenURL(tokenURL string) string {
	u, err := url.Parse(tokenURL)
	if err != nil {
		return tokenURL
	}
	return u.Scheme + "://" + u.Host + "/"
}
