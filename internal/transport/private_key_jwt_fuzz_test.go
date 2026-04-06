package transport

import (
	"testing"
)

// FuzzAudienceFromTokenURL verifies audienceFromTokenURL never panics
// and always returns a non-empty string for non-empty input.
func FuzzAudienceFromTokenURL(f *testing.F) {
	f.Add("https://mytenant.auth0.com/oauth/token")
	f.Add("https://example.com/oauth/token")
	f.Add("https://example.com:8443/oauth/token")
	f.Add("")
	f.Add("not-a-url")
	f.Add("://broken")
	f.Add("https://")
	f.Add("ftp://host/path")

	f.Fuzz(func(t *testing.T, tokenURL string) {
		result := audienceFromTokenURL(tokenURL)

		// Should never panic (the test itself verifies this).
		// For non-empty input, result should be non-empty.
		if tokenURL != "" && result == "" {
			t.Errorf("audienceFromTokenURL(%q) returned empty string", tokenURL)
		}
	})
}

// FuzzNewPrivateKeyJwtTokenSource_InvalidInputs verifies the constructor
// never panics on arbitrary key/algorithm combinations.
func FuzzNewPrivateKeyJwtTokenSource_InvalidInputs(f *testing.F) {
	f.Add("test-client", "not-a-key", "RS256", "https://example.com/oauth/token", "https://example.com/")
	f.Add("", "", "", "", "")
	f.Add("client", "-----BEGIN RSA PRIVATE KEY-----\ngarbage\n-----END RSA PRIVATE KEY-----", "ES256", "https://a.com/t", "https://a.com/")

	f.Fuzz(func(_ *testing.T, clientID, keyPEM, alg, tokenURL, audience string) {
		// We only care that this never panics.
		_, _ = NewPrivateKeyJwtTokenSource(PrivateKeyJwtConfig{
			ClientID:         clientID,
			PrivateKeyPEM:    keyPEM,
			SigningAlgorithm: alg,
			TokenURL:         tokenURL,
			Audience:         audience,
		})
	})
}
