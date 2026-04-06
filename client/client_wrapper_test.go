package client

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"runtime"
	"strings"
	"testing"

	myorganization "github.com/auth0/myorganization-go"
	"github.com/auth0/myorganization-go/internal/auth0"
	"github.com/auth0/myorganization-go/internal/telemetry"
	option "github.com/auth0/myorganization-go/option"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/oauth2"
)

// --- Domain Sanitization ---

func TestSanitizeDomain(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"mytenant.auth0.com", "mytenant.auth0.com"},
		{"https://mytenant.auth0.com", "mytenant.auth0.com"},
		{"http://mytenant.auth0.com", "mytenant.auth0.com"},
		{"https://mytenant.auth0.com/", "mytenant.auth0.com"},
		{"https://mytenant.auth0.com///", "mytenant.auth0.com"},
		{"mytenant.auth0.com/", "mytenant.auth0.com"},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			assert.Equal(t, tt.expected, auth0.SanitizeDomain(tt.input))
		})
	}
}

func TestDerivedURLs(t *testing.T) {
	tests := []struct {
		domain      string
		baseURL     string
		tokenURL    string
		audienceURL string
	}{
		{
			domain:      "mytenant.auth0.com",
			baseURL:     "https://mytenant.auth0.com/my-org",
			tokenURL:    "https://mytenant.auth0.com/oauth/token",
			audienceURL: "https://mytenant.auth0.com/my-org/",
		},
		{
			domain:      "https://mytenant.auth0.com/",
			baseURL:     "https://mytenant.auth0.com/my-org",
			tokenURL:    "https://mytenant.auth0.com/oauth/token",
			audienceURL: "https://mytenant.auth0.com/my-org/",
		},
	}
	for _, tt := range tests {
		t.Run(tt.domain, func(t *testing.T) {
			d := auth0.SanitizeDomain(tt.domain)
			assert.Equal(t, tt.baseURL, "https://"+d+"/my-org")
			assert.Equal(t, tt.tokenURL, "https://"+d+"/oauth/token")
			assert.Equal(t, tt.audienceURL, "https://"+d+"/my-org/")
		})
	}
}

// --- Telemetry ---

func TestBuildAuth0ClientInfo(t *testing.T) {
	info := telemetry.BuildAuth0ClientInfo(nil)
	assert.Equal(t, myorganization.SDKName, info.Name)
	assert.Equal(t, myorganization.SDKVersion, info.Version)
	assert.Equal(t, runtime.Version(), info.Env["go"])
}

func TestBuildAuth0ClientInfoWithCustomEnv(t *testing.T) {
	custom := map[string]string{"myapp": "1.0.0"}
	info := telemetry.BuildAuth0ClientInfo(custom)
	assert.Equal(t, "1.0.0", info.Env["myapp"])
	assert.Equal(t, runtime.Version(), info.Env["go"])
}

func TestEncodeAuth0ClientInfo(t *testing.T) {
	info := telemetry.Auth0ClientInfo{
		Name:    "myorganization-go",
		Version: "0.0.1",
		Env:     map[string]string{"go": "go1.21"},
	}
	encoded, err := telemetry.EncodeAuth0ClientInfo(info)
	require.NoError(t, err)

	decoded, err := base64.StdEncoding.DecodeString(encoded)
	require.NoError(t, err)

	var roundTrip telemetry.Auth0ClientInfo
	err = json.Unmarshal(decoded, &roundTrip)
	require.NoError(t, err)

	assert.Equal(t, info.Name, roundTrip.Name)
	assert.Equal(t, info.Version, roundTrip.Version)
	assert.Equal(t, info.Env, roundTrip.Env)
}

// --- Option Validation ---

func TestNew_EmptyDomain(t *testing.T) {
	_, err := New("", option.WithStaticToken("tok"))
	require.Error(t, err)
	assert.Contains(t, err.Error(), "domain must not be empty")
}

func TestNew_NoAuth(t *testing.T) {
	_, err := New("mytenant.auth0.com")
	require.Error(t, err)
	assert.Contains(t, err.Error(), "must provide either client credentials")
}

func TestNew_BothAuthModes(t *testing.T) {
	_, err := New("mytenant.auth0.com",
		option.WithClientCredentials(context.Background(), "id", "secret"),
		option.WithStaticToken("tok"),
	)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "only one authentication mode")
}

func TestNew_ClientCredentials_MissingClientID(t *testing.T) {
	_, err := New("mytenant.auth0.com",
		option.WithClientCredentials(context.Background(), "", "secret"),
	)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "client ID must not be empty")
}

func TestNew_ClientCredentials_MissingClientSecret(t *testing.T) {
	_, err := New("mytenant.auth0.com",
		option.WithClientCredentials(context.Background(), "id", ""),
	)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "client secret must not be empty")
}

func TestNew_StaticToken(t *testing.T) {
	c, err := New("mytenant.auth0.com", option.WithStaticToken("my-token"))
	require.NoError(t, err)
	assert.NotNil(t, c)
	assert.NotNil(t, c.Organization)
	assert.NotNil(t, c.OrganizationDetails)
}

func TestNew_WithToken(t *testing.T) {
	c, err := New("mytenant.auth0.com", option.WithToken("my-token"))
	require.NoError(t, err)
	assert.NotNil(t, c)
}

// --- Transport Chain Headers ---

func TestTransportHeaders_StaticToken(t *testing.T) {
	var capturedHeaders http.Header
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		capturedHeaders = r.Header.Clone()
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	c, err := New("mytenant.auth0.com",
		option.WithStaticToken("my-token"),
		option.WithBaseURL(server.URL),
		option.WithHTTPClient(&http.Client{}),
	)
	require.NoError(t, err)

	req, err := http.NewRequest("GET", server.URL+"/test", nil)
	require.NoError(t, err)

	resp, err := c.options.HTTPClient.Do(req)
	require.NoError(t, err)
	defer func() { _ = resp.Body.Close() }()

	// Verify transport-set headers.
	assert.Equal(t, "MyOrganization-Go/"+myorganization.SDKVersion, capturedHeaders.Get("User-Agent"))

	auth0Client := capturedHeaders.Get("Auth0-Client")
	assert.NotEmpty(t, auth0Client)

	decoded, err := base64.StdEncoding.DecodeString(auth0Client)
	require.NoError(t, err)

	var info telemetry.Auth0ClientInfo
	err = json.Unmarshal(decoded, &info)
	require.NoError(t, err)
	assert.Equal(t, myorganization.SDKName, info.Name)
	assert.Equal(t, myorganization.SDKVersion, info.Version)
	assert.Equal(t, runtime.Version(), info.Env["go"])

	// Verify X-Fern-* headers are NOT present.
	assert.Empty(t, capturedHeaders.Get("X-Fern-Language"))
	assert.Empty(t, capturedHeaders.Get("X-Fern-SDK-Name"))
	assert.Empty(t, capturedHeaders.Get("X-Fern-SDK-Version"))
}

func TestTransportHeaders_NoAuth0ClientInfo(t *testing.T) {
	var capturedHeaders http.Header
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		capturedHeaders = r.Header.Clone()
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	c, err := New("mytenant.auth0.com",
		option.WithStaticToken("my-token"),
		option.WithBaseURL(server.URL),
		option.WithNoAuth0ClientInfo(),
	)
	require.NoError(t, err)

	req, err := http.NewRequest("GET", server.URL+"/test", nil)
	require.NoError(t, err)

	resp, err := c.options.HTTPClient.Do(req)
	require.NoError(t, err)
	defer func() { _ = resp.Body.Close() }()

	assert.Equal(t, "MyOrganization-Go/"+myorganization.SDKVersion, capturedHeaders.Get("User-Agent"))
	assert.Empty(t, capturedHeaders.Get("Auth0-Client"))
}

func TestTransportHeaders_CustomEnvEntry(t *testing.T) {
	var capturedHeaders http.Header
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		capturedHeaders = r.Header.Clone()
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	c, err := New("mytenant.auth0.com",
		option.WithStaticToken("my-token"),
		option.WithBaseURL(server.URL),
		option.WithAuth0ClientEnvEntry("myapp", "2.0.0"),
	)
	require.NoError(t, err)

	req, err := http.NewRequest("GET", server.URL+"/test", nil)
	require.NoError(t, err)

	resp, err := c.options.HTTPClient.Do(req)
	require.NoError(t, err)
	defer func() { _ = resp.Body.Close() }()

	auth0Client := capturedHeaders.Get("Auth0-Client")
	decoded, err := base64.StdEncoding.DecodeString(auth0Client)
	require.NoError(t, err)

	var info telemetry.Auth0ClientInfo
	err = json.Unmarshal(decoded, &info)
	require.NoError(t, err)
	assert.Equal(t, "2.0.0", info.Env["myapp"])
	assert.Equal(t, runtime.Version(), info.Env["go"])
}

// --- M2M Client Credentials Flow ---

func TestNew_ClientCredentials_M2MFlow(t *testing.T) {
	tokenServer := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		bodyStr := string(body)
		assert.Contains(t, bodyStr, "grant_type=client_credentials")
		assert.Contains(t, bodyStr, "audience=")

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"access_token":"test-access-token","token_type":"Bearer","expires_in":86400}`))
	}))
	defer tokenServer.Close()

	var capturedHeaders http.Header
	apiServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		capturedHeaders = r.Header.Clone()
		w.WriteHeader(http.StatusOK)
	}))
	defer apiServer.Close()

	tokenHost := strings.TrimPrefix(tokenServer.URL, "https://")
	ctx := context.WithValue(context.Background(), oauth2.HTTPClient, tokenServer.Client())

	c, err := New(tokenHost,
		option.WithClientCredentialsAndAudience(ctx, "test-client-id", "test-client-secret", "https://test-audience/"),
		option.WithBaseURL(apiServer.URL),
	)
	require.NoError(t, err)

	req, err := http.NewRequest("GET", apiServer.URL+"/test", nil)
	require.NoError(t, err)

	resp, err := c.options.HTTPClient.Do(req)
	require.NoError(t, err)
	defer func() { _ = resp.Body.Close() }()

	// Verify oauth2 transport set Authorization.
	assert.Equal(t, "Bearer test-access-token", capturedHeaders.Get("Authorization"))
	// Verify transport chain headers.
	assert.Equal(t, "MyOrganization-Go/"+myorganization.SDKVersion, capturedHeaders.Get("User-Agent"))
	assert.NotEmpty(t, capturedHeaders.Get("Auth0-Client"))
}

// --- M2M with default audience (empty string → derived from domain) ---

func TestNew_ClientCredentials_DefaultAudience(t *testing.T) {
	tokenServer := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		bodyStr := string(body)
		// Default audience should be https://{domain}/my-org/
		assert.Contains(t, bodyStr, "audience=https")
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"access_token":"tok","token_type":"Bearer","expires_in":3600}`))
	}))
	defer tokenServer.Close()

	tokenHost := strings.TrimPrefix(tokenServer.URL, "https://")
	ctx := context.WithValue(context.Background(), oauth2.HTTPClient, tokenServer.Client())

	apiServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer apiServer.Close()

	// No audience → uses default.
	c, err := New(tokenHost,
		option.WithClientCredentials(ctx, "id", "secret"),
		option.WithBaseURL(apiServer.URL),
	)
	require.NoError(t, err)

	req, _ := http.NewRequest("GET", apiServer.URL+"/test", nil)
	resp, err := c.options.HTTPClient.Do(req)
	require.NoError(t, err)
	defer func() { _ = resp.Body.Close() }()
}

// --- Custom Base URL ---

func TestNew_CustomBaseURL(t *testing.T) {
	c, err := New("mytenant.auth0.com",
		option.WithStaticToken("tok"),
		option.WithBaseURL("https://custom.example.com/api"),
	)
	require.NoError(t, err)
	assert.Equal(t, "https://custom.example.com/api", c.baseURL)
}

func TestNew_DefaultBaseURL(t *testing.T) {
	c, err := New("mytenant.auth0.com",
		option.WithStaticToken("tok"),
	)
	require.NoError(t, err)
	assert.Equal(t, "https://mytenant.auth0.com/my-org", c.baseURL)
}

// --- Custom Organization ---

func TestNew_WithOrganization_SentInTokenRequest(t *testing.T) {
	tokenServer := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		bodyStr := string(body)
		assert.Contains(t, bodyStr, "organization=org_abc123")
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"access_token":"tok","token_type":"Bearer","expires_in":3600}`))
	}))
	defer tokenServer.Close()

	tokenHost := strings.TrimPrefix(tokenServer.URL, "https://")
	ctx := context.WithValue(context.Background(), oauth2.HTTPClient, tokenServer.Client())

	apiServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer apiServer.Close()

	c, err := New(tokenHost,
		option.WithClientCredentials(ctx, "id", "secret"),
		option.WithOrganization("org_abc123"),
		option.WithBaseURL(apiServer.URL),
	)
	require.NoError(t, err)

	// Base URL is unchanged — organization does not affect it.
	assert.Equal(t, apiServer.URL, c.baseURL)

	// Trigger a request so the token endpoint is called.
	req, _ := http.NewRequest("GET", apiServer.URL+"/test", nil)
	resp, err := c.options.HTTPClient.Do(req)
	require.NoError(t, err)
	defer func() { _ = resp.Body.Close() }()
}

func TestNew_WithoutOrganization_NotSentInTokenRequest(t *testing.T) {
	tokenServer := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		bodyStr := string(body)
		assert.NotContains(t, bodyStr, "organization=")
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"access_token":"tok","token_type":"Bearer","expires_in":3600}`))
	}))
	defer tokenServer.Close()

	tokenHost := strings.TrimPrefix(tokenServer.URL, "https://")
	ctx := context.WithValue(context.Background(), oauth2.HTTPClient, tokenServer.Client())

	apiServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer apiServer.Close()

	c, err := New(tokenHost,
		option.WithClientCredentials(ctx, "id", "secret"),
		option.WithBaseURL(apiServer.URL),
	)
	require.NoError(t, err)

	req, _ := http.NewRequest("GET", apiServer.URL+"/test", nil)
	resp, err := c.options.HTTPClient.Do(req)
	require.NoError(t, err)
	defer func() { _ = resp.Body.Close() }()
}

// --- Custom Audience ---

func TestNew_CustomAudience(t *testing.T) {
	tokenServer := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		bodyStr := string(body)
		assert.Contains(t, bodyStr, "audience=https%3A%2F%2Fcustom-api.example.com%2F")
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"access_token":"tok","token_type":"Bearer","expires_in":3600}`))
	}))
	defer tokenServer.Close()

	tokenHost := strings.TrimPrefix(tokenServer.URL, "https://")

	apiServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer apiServer.Close()

	ctx := context.WithValue(context.Background(), oauth2.HTTPClient, tokenServer.Client())

	c, err := New(tokenHost,
		option.WithClientCredentialsAndAudience(ctx, "id", "secret", "https://custom-api.example.com/"),
		option.WithBaseURL(apiServer.URL),
	)
	require.NoError(t, err)

	req, _ := http.NewRequest("GET", apiServer.URL+"/test", nil)
	resp, err := c.options.HTTPClient.Do(req)
	require.NoError(t, err)
	defer func() { _ = resp.Body.Close() }()
}

// --- Max Attempts ---

func TestNew_WithMaxAttempts(t *testing.T) {
	c, err := New("mytenant.auth0.com",
		option.WithStaticToken("tok"),
		option.WithMaxAttempts(5),
	)
	require.NoError(t, err)
	assert.NotNil(t, c)
}

// --- Domain Variants Produce Same Base URL ---

func TestNew_DomainVariants(t *testing.T) {
	variants := []string{
		"mytenant.auth0.com",
		"https://mytenant.auth0.com",
		"https://mytenant.auth0.com/",
	}
	for _, domain := range variants {
		t.Run(domain, func(t *testing.T) {
			c, err := New(domain, option.WithStaticToken("tok"))
			require.NoError(t, err)
			assert.Equal(t, "https://mytenant.auth0.com/my-org", c.baseURL)
		})
	}
}

// --- Debug Mode ---

func TestNew_WithDebug(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"ok":true}`))
	}))
	defer server.Close()

	var buf bytes.Buffer
	log.SetOutput(&buf)
	log.SetFlags(0)
	defer log.SetOutput(os.Stderr)
	defer log.SetFlags(log.LstdFlags)

	c, err := New("mytenant.auth0.com",
		option.WithStaticToken("my-secret-token"),
		option.WithBaseURL(server.URL),
		option.WithDebug(true),
	)
	require.NoError(t, err)

	req, err := http.NewRequest("GET", server.URL+"/test", nil)
	require.NoError(t, err)
	// Simulate the Authorization header that the generated Fern caller adds.
	req.Header.Set("Authorization", "Bearer my-secret-token")

	resp, err := c.options.HTTPClient.Do(req)
	require.NoError(t, err)
	defer func() { _ = resp.Body.Close() }()

	output := buf.String()

	// Debug output should contain request and response sections.
	assert.Contains(t, output, "---[ REQUEST ]---")
	assert.Contains(t, output, "---[ RESPONSE ]---")
	assert.Contains(t, output, "GET /test")

	// The token set via WithStaticToken must be redacted, not logged in plain text.
	assert.NotContains(t, output, "my-secret-token")
	assert.Contains(t, output, "[REDACTED]")

	// The transport chain adds User-Agent and Auth0-Client, which should be visible.
	assert.Contains(t, output, "User-Agent: MyOrganization-Go/"+myorganization.SDKVersion)
}

func TestNew_WithDebugDisabled(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	var buf bytes.Buffer
	log.SetOutput(&buf)
	log.SetFlags(0)
	defer log.SetOutput(os.Stderr)
	defer log.SetFlags(log.LstdFlags)

	c, err := New("mytenant.auth0.com",
		option.WithStaticToken("my-token"),
		option.WithBaseURL(server.URL),
		option.WithDebug(false),
	)
	require.NoError(t, err)

	req, err := http.NewRequest("GET", server.URL+"/test", nil)
	require.NoError(t, err)

	resp, err := c.options.HTTPClient.Do(req)
	require.NoError(t, err)
	defer func() { _ = resp.Body.Close() }()

	// No debug output should be produced.
	assert.Empty(t, buf.String())
}

// --- End-to-end: headers through generated client path ---

func TestE2E_HeadersThroughGeneratedClient(t *testing.T) {
	var capturedHeaders http.Header
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		capturedHeaders = r.Header.Clone()
		// Return a valid JSON response so the generated client can parse it.
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{}`))
	}))
	defer server.Close()

	c, err := New("mytenant.auth0.com",
		option.WithStaticToken("my-secret-token"),
		option.WithBaseURL(server.URL),
	)
	require.NoError(t, err)

	// Call through the actual generated client (OrganizationDetails.Get).
	// The request goes: generated RawClient → Caller → newRequest (sets ToHeader headers) → HTTP Client (transport chain).
	_, _ = c.OrganizationDetails.Get(context.Background())

	// Transport-layer headers (set by RoundTripper chain).
	assert.Equal(t, "MyOrganization-Go/"+myorganization.SDKVersion, capturedHeaders.Get("User-Agent"))
	assert.NotEmpty(t, capturedHeaders.Get("Auth0-Client"))

	// Application-layer headers (set by generated code via ToHeader → Caller).
	assert.Equal(t, "Bearer my-secret-token", capturedHeaders.Get("Authorization"))
	assert.NotEmpty(t, capturedHeaders.Get("Content-Type"))

	// X-Fern headers should NOT be present.
	assert.Empty(t, capturedHeaders.Get("X-Fern-Language"))
	assert.Empty(t, capturedHeaders.Get("X-Fern-SDK-Name"))
	assert.Empty(t, capturedHeaders.Get("X-Fern-SDK-Version"))
}
