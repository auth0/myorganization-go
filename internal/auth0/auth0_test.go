package auth0

import (
	"testing"

	core "github.com/auth0/myorganization-go/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// --- SanitizeDomain ---

func TestSanitizeDomain_PlainDomain(t *testing.T) {
	assert.Equal(t, "mytenant.auth0.com", SanitizeDomain("mytenant.auth0.com"))
}

func TestSanitizeDomain_HTTPSPrefix(t *testing.T) {
	assert.Equal(t, "mytenant.auth0.com", SanitizeDomain("https://mytenant.auth0.com"))
}

func TestSanitizeDomain_HTTPPrefix(t *testing.T) {
	assert.Equal(t, "mytenant.auth0.com", SanitizeDomain("http://mytenant.auth0.com"))
}

func TestSanitizeDomain_TrailingSlash(t *testing.T) {
	assert.Equal(t, "mytenant.auth0.com", SanitizeDomain("mytenant.auth0.com/"))
}

func TestSanitizeDomain_HTTPSAndTrailingSlash(t *testing.T) {
	assert.Equal(t, "mytenant.auth0.com", SanitizeDomain("https://mytenant.auth0.com/"))
}

func TestSanitizeDomain_MultipleTrailingSlashes(t *testing.T) {
	assert.Equal(t, "mytenant.auth0.com", SanitizeDomain("https://mytenant.auth0.com///"))
}

// --- DeriveURLs ---

func TestDeriveURLs_Defaults(t *testing.T) {
	opts := &core.RequestOptions{}
	baseURL, tokenURL, audience := DeriveURLs("mytenant.auth0.com", opts)
	assert.Equal(t, "https://mytenant.auth0.com/my-org", baseURL)
	assert.Equal(t, "https://mytenant.auth0.com/oauth/token", tokenURL)
	assert.Equal(t, "https://mytenant.auth0.com/my-org/", audience)
}

func TestDeriveURLs_CustomBaseURL(t *testing.T) {
	opts := &core.RequestOptions{BaseURL: "https://custom.example.com/api"}
	baseURL, tokenURL, audience := DeriveURLs("mytenant.auth0.com", opts)
	assert.Equal(t, "https://custom.example.com/api", baseURL)
	assert.Equal(t, "https://mytenant.auth0.com/oauth/token", tokenURL)
	assert.Equal(t, "https://mytenant.auth0.com/my-org/", audience)
}

func TestDeriveURLs_CustomAudience(t *testing.T) {
	opts := &core.RequestOptions{Audience: "https://custom-api.example.com/"}
	baseURL, tokenURL, audience := DeriveURLs("mytenant.auth0.com", opts)
	assert.Equal(t, "https://mytenant.auth0.com/my-org", baseURL)
	assert.Equal(t, "https://mytenant.auth0.com/oauth/token", tokenURL)
	assert.Equal(t, "https://custom-api.example.com/", audience)
}

func TestDeriveURLs_CustomBaseURLAndAudience(t *testing.T) {
	opts := &core.RequestOptions{
		BaseURL:  "https://custom.example.com/api",
		Audience: "https://custom-api.example.com/",
	}
	baseURL, tokenURL, audience := DeriveURLs("mytenant.auth0.com", opts)
	assert.Equal(t, "https://custom.example.com/api", baseURL)
	assert.Equal(t, "https://mytenant.auth0.com/oauth/token", tokenURL)
	assert.Equal(t, "https://custom-api.example.com/", audience)
}


// --- ValidateOptions ---

func TestValidateOptions_EmptyDomain(t *testing.T) {
	opts := &core.RequestOptions{Token: "tok"}
	err := ValidateOptions("", opts)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "domain must not be empty")
}

func TestValidateOptions_NoAuth(t *testing.T) {
	opts := &core.RequestOptions{}
	err := ValidateOptions("mytenant.auth0.com", opts)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "must provide either client credentials")
}

func TestValidateOptions_BothAuthModes(t *testing.T) {
	opts := &core.RequestOptions{
		ClientID:     "id",
		ClientSecret: "secret",
		Token:        "tok",
	}
	err := ValidateOptions("mytenant.auth0.com", opts)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "only one authentication mode")
}

func TestValidateOptions_ClientCredentials_Valid(t *testing.T) {
	opts := &core.RequestOptions{
		ClientID:     "id",
		ClientSecret: "secret",
	}
	err := ValidateOptions("mytenant.auth0.com", opts)
	assert.NoError(t, err)
}

func TestValidateOptions_ClientCredentials_MissingClientID(t *testing.T) {
	opts := &core.RequestOptions{
		ClientSecret: "secret",
	}
	err := ValidateOptions("mytenant.auth0.com", opts)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "client ID must not be empty")
}

func TestValidateOptions_ClientCredentials_MissingClientSecret(t *testing.T) {
	opts := &core.RequestOptions{
		ClientID: "id",
	}
	err := ValidateOptions("mytenant.auth0.com", opts)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "client secret must not be empty")
}

func TestValidateOptions_StaticToken_Valid(t *testing.T) {
	opts := &core.RequestOptions{Token: "my-token"}
	err := ValidateOptions("mytenant.auth0.com", opts)
	assert.NoError(t, err)
}
