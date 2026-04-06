package auth0

import (
	"testing"

	core "github.com/auth0/myorganization-go/core"
)

// FuzzSanitizeDomain verifies SanitizeDomain never panics on arbitrary input
// and always strips protocol prefixes and trailing slashes.
func FuzzSanitizeDomain(f *testing.F) {
	f.Add("mytenant.auth0.com")
	f.Add("https://mytenant.auth0.com")
	f.Add("http://mytenant.auth0.com")
	f.Add("https://mytenant.auth0.com/")
	f.Add("https://mytenant.auth0.com///")
	f.Add("")
	f.Add("https://")
	f.Add("http://")
	f.Add("///")
	f.Add("https://a]b[c.com/")

	f.Fuzz(func(t *testing.T, domain string) {
		result := SanitizeDomain(domain)

		// Must never contain protocol prefix.
		if len(result) >= 8 && result[:8] == "https://" {
			t.Errorf("SanitizeDomain(%q) still contains https:// prefix: %q", domain, result)
		}
		if len(result) >= 7 && result[:7] == "http://" {
			t.Errorf("SanitizeDomain(%q) still contains http:// prefix: %q", domain, result)
		}

		// Must never end with a slash.
		if len(result) > 0 && result[len(result)-1] == '/' {
			t.Errorf("SanitizeDomain(%q) has trailing slash: %q", domain, result)
		}
	})
}

// FuzzDeriveURLs verifies DeriveURLs never panics and produces well-formed URLs.
func FuzzDeriveURLs(f *testing.F) {
	f.Add("mytenant.auth0.com", "", "")
	f.Add("example.com", "https://custom.example.com/api", "https://custom-api.example.com/")
	f.Add("", "", "")
	f.Add("a]b[c.com", "", "")

	f.Fuzz(func(t *testing.T, domain, baseURL, audience string) {
		opts := &core.RequestOptions{
			BaseURL:  baseURL,
			Audience: audience,
		}
		b, tok, aud := DeriveURLs(domain, opts)

		// Token URL must always contain the domain.
		if domain != "" && len(tok) == 0 {
			t.Errorf("DeriveURLs(%q, ...) returned empty tokenURL", domain)
		}

		// If custom baseURL was given, it should be returned as-is.
		if baseURL != "" && b != baseURL {
			t.Errorf("DeriveURLs(%q, ...) did not preserve custom baseURL: got %q, want %q", domain, b, baseURL)
		}

		// If custom audience was given, it should be returned as-is.
		if audience != "" && aud != audience {
			t.Errorf("DeriveURLs(%q, ...) did not preserve custom audience: got %q, want %q", domain, aud, audience)
		}
	})
}

// FuzzValidateOptions verifies ValidateOptions never panics.
func FuzzValidateOptions(f *testing.F) {
	f.Add("mytenant.auth0.com", "id", "secret", "", "tok")
	f.Add("", "", "", "", "")
	f.Add("example.com", "id", "", "key-pem", "")
	f.Add("example.com", "", "", "", "my-token")

	f.Fuzz(func(_ *testing.T, domain, clientID, clientSecret, privateKeyPEM, token string) {
		opts := &core.RequestOptions{
			ClientID:      clientID,
			ClientSecret:  clientSecret,
			PrivateKeyPEM: privateKeyPEM,
			Token:         token,
		}
		if privateKeyPEM != "" {
			opts.SigningAlgorithm = "RS256"
		}
		// We only care that it doesn't panic.
		_ = ValidateOptions(domain, opts)
	})
}
