package auth0

import (
	"errors"
	"strings"

	core "github.com/auth0/myorganization-go/core"
)

// SanitizeDomain strips protocol prefixes and trailing slashes from the domain.
// It loops to handle edge cases like "http://http://example.com".
func SanitizeDomain(domain string) string {
	d := domain
	for {
		trimmed := strings.TrimPrefix(d, "https://")
		trimmed = strings.TrimPrefix(trimmed, "http://")
		if trimmed == d {
			break
		}
		d = trimmed
	}
	d = strings.TrimRight(d, "/")
	return d
}

// DeriveURLs computes the base URL, token URL, and audience from a sanitized domain.
func DeriveURLs(sanitizedDomain string, options *core.RequestOptions) (baseURL, tokenURL, audience string) {
	baseURL = options.BaseURL
	if baseURL == "" {
		baseURL = "https://" + sanitizedDomain + "/my-org"
	}
	tokenURL = "https://" + sanitizedDomain + "/oauth/token"
	audience = options.Audience
	if audience == "" {
		audience = "https://" + sanitizedDomain + "/my-org/"
	}
	return
}

// ValidateOptions checks that domain and auth options are valid.
func ValidateOptions(domain string, options *core.RequestOptions) error {
	if domain == "" {
		return errors.New("auth0: domain must not be empty")
	}

	// Detect auth modes by their unique discriminating fields.
	// ClientID is shared by client credentials and private key JWT, so it's not
	// used as a discriminator. ClientSecret uniquely identifies client credentials;
	// PrivateKeyPEM uniquely identifies private key JWT.
	hasClientCredentials := options.ClientSecret != ""
	hasPrivateKeyJWT := options.PrivateKeyPEM != ""
	hasStaticToken := options.Token != ""
	hasTokenSource := options.TokenSource != nil

	// ClientID alone (without ClientSecret or PrivateKeyPEM) looks like
	// an incomplete client credentials config.
	hasOrphanClientID := options.ClientID != "" && !hasClientCredentials && !hasPrivateKeyJWT

	// Count how many auth modes are active.
	count := 0
	if hasClientCredentials || hasOrphanClientID {
		count++
	}
	if hasPrivateKeyJWT {
		count++
	}
	if hasStaticToken {
		count++
	}
	if hasTokenSource {
		count++
	}

	if count > 1 {
		return errors.New("auth0: only one authentication mode may be used (client credentials, private key JWT, token source, or static token)")
	}
	if count == 0 {
		return errors.New("auth0: must provide either client credentials (WithClientCredentials), private key JWT (WithPrivateKeyJWT), a token source (WithTokenSource), or a static token (WithStaticToken/WithToken)")
	}
	if hasClientCredentials || hasOrphanClientID {
		if options.ClientID == "" {
			return errors.New("auth0: client credentials: client ID must not be empty")
		}
		if options.ClientSecret == "" {
			return errors.New("auth0: client credentials: client secret must not be empty")
		}
	}
	if hasPrivateKeyJWT {
		if options.ClientID == "" {
			return errors.New("auth0: private key JWT: client ID must not be empty")
		}
		if options.SigningAlgorithm == "" {
			return errors.New("auth0: private key JWT: signing algorithm must not be empty")
		}
	}
	return nil
}
