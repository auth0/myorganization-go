package transport

import (
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/lestrrat-go/jwx/v3/jwa"
	"github.com/lestrrat-go/jwx/v3/jwt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/oauth2"
)

// generateRSAKeyPEM generates a fresh RSA private key in PEM format for testing.
func generateRSAKeyPEM(t *testing.T) string {
	t.Helper()
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	require.NoError(t, err)
	der := x509.MarshalPKCS1PrivateKey(key)
	block := &pem.Block{Type: "RSA PRIVATE KEY", Bytes: der}
	return string(pem.EncodeToMemory(block))
}

// generateECKeyPEM generates a fresh EC private key in PEM format for testing.
func generateECKeyPEM(t *testing.T, curve elliptic.Curve) string {
	t.Helper()
	key, err := ecdsa.GenerateKey(curve, rand.Reader)
	require.NoError(t, err)
	der, err := x509.MarshalECPrivateKey(key)
	require.NoError(t, err)
	block := &pem.Block{Type: "EC PRIVATE KEY", Bytes: der}
	return string(pem.EncodeToMemory(block))
}

// --- Algorithm Validation ---

func TestNewPrivateKeyJwtTokenSource_UnsupportedAlgorithm(t *testing.T) {
	keyPEM := generateRSAKeyPEM(t)
	_, err := NewPrivateKeyJwtTokenSource(PrivateKeyJwtConfig{
		ClientID:         "test-client",
		PrivateKeyPEM:    keyPEM,
		SigningAlgorithm: "HS256",
		TokenURL:         "https://example.auth0.com/oauth/token",
		Audience:         "https://example.auth0.com/my-org/",
	})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "unsupported signing algorithm")
}

func TestNewPrivateKeyJwtTokenSource_AllSupportedAlgorithms(t *testing.T) {
	rsaKeyPEM := generateRSAKeyPEM(t)
	ecP256KeyPEM := generateECKeyPEM(t, elliptic.P256())
	ecP384KeyPEM := generateECKeyPEM(t, elliptic.P384())
	ecP521KeyPEM := generateECKeyPEM(t, elliptic.P521())

	tests := []struct {
		alg    string
		keyPEM string
	}{
		{"RS256", rsaKeyPEM},
		{"RS384", rsaKeyPEM},
		{"RS512", rsaKeyPEM},
		{"PS256", rsaKeyPEM},
		{"PS384", rsaKeyPEM},
		{"PS512", rsaKeyPEM},
		{"ES256", ecP256KeyPEM},
		{"ES384", ecP384KeyPEM},
		{"ES512", ecP521KeyPEM},
	}

	for _, tt := range tests {
		t.Run(tt.alg, func(t *testing.T) {
			src, err := NewPrivateKeyJwtTokenSource(PrivateKeyJwtConfig{
				ClientID:         "test-client",
				PrivateKeyPEM:    tt.keyPEM,
				SigningAlgorithm: tt.alg,
				TokenURL:         "https://example.auth0.com/oauth/token",
				Audience:         "https://example.auth0.com/my-org/",
			})
			require.NoError(t, err)
			assert.NotNil(t, src)
		})
	}
}

// --- Key Parsing ---

func TestNewPrivateKeyJwtTokenSource_InvalidKey(t *testing.T) {
	_, err := NewPrivateKeyJwtTokenSource(PrivateKeyJwtConfig{
		ClientID:         "test-client",
		PrivateKeyPEM:    "not-a-valid-pem-key",
		SigningAlgorithm: "RS256",
		TokenURL:         "https://example.auth0.com/oauth/token",
		Audience:         "https://example.auth0.com/my-org/",
	})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "failed to parse private key")
}

// --- Client Assertion JWT ---

func TestCreateClientAssertion_Claims(t *testing.T) {
	keyPEM := generateRSAKeyPEM(t)

	src := &privateKeyJwtTokenSource{
		cfg: PrivateKeyJwtConfig{
			ClientID:         "my-client-id",
			PrivateKeyPEM:    keyPEM,
			SigningAlgorithm: "RS256",
			TokenURL:         "https://mytenant.auth0.com/oauth/token",
			Audience:         "https://mytenant.auth0.com/my-org/",
		},
		alg: jwa.RS256(),
	}

	// Parse the key for the source.
	key, err := parsePrivateKey([]byte(keyPEM))
	require.NoError(t, err)
	src.key = key

	assertion, err := src.createClientAssertion()
	require.NoError(t, err)
	assert.NotEmpty(t, assertion)

	// Parse and verify claims (without signature verification).
	parsed, err := jwt.ParseInsecure([]byte(assertion))
	require.NoError(t, err)

	iss, ok := parsed.Issuer()
	assert.True(t, ok)
	assert.Equal(t, "my-client-id", iss)

	sub, ok := parsed.Subject()
	assert.True(t, ok)
	assert.Equal(t, "my-client-id", sub)

	aud, ok := parsed.Audience()
	assert.True(t, ok)
	assert.Contains(t, aud, "https://mytenant.auth0.com/")

	jti, ok := parsed.JwtID()
	assert.True(t, ok)
	assert.NotEmpty(t, jti)

	iat, ok := parsed.IssuedAt()
	assert.True(t, ok)
	assert.False(t, iat.IsZero())

	nbf, ok := parsed.NotBefore()
	assert.True(t, ok)
	assert.False(t, nbf.IsZero())

	exp, ok := parsed.Expiration()
	assert.True(t, ok)
	assert.False(t, exp.IsZero())
	assert.True(t, exp.After(iat))
}

// --- Audience from Token URL ---

func TestAudienceFromTokenURL(t *testing.T) {
	tests := []struct {
		tokenURL string
		expected string
	}{
		{"https://mytenant.auth0.com/oauth/token", "https://mytenant.auth0.com/"},
		{"https://example.com/oauth/token", "https://example.com/"},
		{"https://example.com:8443/oauth/token", "https://example.com:8443/"},
	}
	for _, tt := range tests {
		t.Run(tt.tokenURL, func(t *testing.T) {
			assert.Equal(t, tt.expected, audienceFromTokenURL(tt.tokenURL))
		})
	}
}

// --- Token Exchange (end-to-end with mock token server) ---

func TestPrivateKeyJwtTokenSource_TokenExchange(t *testing.T) {
	keyPEM := generateRSAKeyPEM(t)

	var capturedBody string
	tokenServer := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		capturedBody = string(body)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(map[string]interface{}{
			"access_token": "test-access-token",
			"token_type":   "Bearer",
			"expires_in":   86400,
		})
	}))
	defer tokenServer.Close()

	// Use the TLS server's client for the OAuth context.
	ctx := context.WithValue(context.Background(), oauth2.HTTPClient, tokenServer.Client())

	src, err := NewPrivateKeyJwtTokenSource(PrivateKeyJwtConfig{
		ClientID:         "my-client-id",
		PrivateKeyPEM:    keyPEM,
		SigningAlgorithm: "RS256",
		TokenURL:         tokenServer.URL + "/oauth/token",
		Audience:         "https://mytenant.auth0.com/my-org/",
		Context:          ctx,
	})
	require.NoError(t, err)

	token, err := src.Token()
	require.NoError(t, err)
	assert.Equal(t, "test-access-token", token.AccessToken)
	assert.Equal(t, "Bearer", token.TokenType)

	// Verify the token request contained the expected parameters.
	assert.Contains(t, capturedBody, "client_assertion=")
	assert.Contains(t, capturedBody, "client_assertion_type=urn")
	assert.Contains(t, capturedBody, "audience=")
	assert.Contains(t, capturedBody, "grant_type=client_credentials")
}

// --- BuildChain with Private Key JWT ---

func TestBuildChain_PrivateKeyJWT(t *testing.T) {
	keyPEM := generateRSAKeyPEM(t)

	// Mock token server.
	tokenServer := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]interface{}{
			"access_token": "pk-jwt-token",
			"token_type":   "Bearer",
			"expires_in":   3600,
		})
	}))
	defer tokenServer.Close()

	var capturedHeaders http.Header
	apiServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		capturedHeaders = r.Header.Clone()
		w.WriteHeader(http.StatusOK)
	}))
	defer apiServer.Close()

	ctx := context.WithValue(context.Background(), oauth2.HTTPClient, tokenServer.Client())

	chain, err := BuildChain(ChainConfig{
		ClientID:         "my-client-id",
		PrivateKeyPEM:    keyPEM,
		SigningAlgorithm: "RS256",
		TokenURL:         tokenServer.URL + "/oauth/token",
		Audience:         "https://mytenant.auth0.com/my-org/",
		OAuthContext:     ctx,
	})
	require.NoError(t, err)

	client := &http.Client{Transport: chain}
	req, _ := http.NewRequest("GET", apiServer.URL+"/test", nil)
	resp, err := client.Do(req)
	require.NoError(t, err)
	defer func() { _ = resp.Body.Close() }()

	assert.Equal(t, "Bearer pk-jwt-token", capturedHeaders.Get("Authorization"))
	assert.NotEmpty(t, capturedHeaders.Get("Auth0-Client"))
}

// --- Case Insensitivity for Algorithm ---

func TestNewPrivateKeyJwtTokenSource_AlgorithmCaseInsensitive(t *testing.T) {
	keyPEM := generateRSAKeyPEM(t)
	src, err := NewPrivateKeyJwtTokenSource(PrivateKeyJwtConfig{
		ClientID:         "test-client",
		PrivateKeyPEM:    keyPEM,
		SigningAlgorithm: "rs256",
		TokenURL:         "https://example.auth0.com/oauth/token",
		Audience:         "https://example.auth0.com/my-org/",
	})
	require.NoError(t, err)
	assert.NotNil(t, src)
}
