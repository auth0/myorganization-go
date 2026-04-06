package transport

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	myorganization "github.com/auth0/myorganization-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// --- UserAgentTransport ---

func TestUserAgentTransport_SetsHeader(t *testing.T) {
	var capturedUA string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		capturedUA = r.Header.Get("User-Agent")
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	transport := &UserAgentTransport{
		Base:      http.DefaultTransport,
		UserAgent: "TestAgent/1.0",
	}
	client := &http.Client{Transport: transport}

	req, _ := http.NewRequest("GET", server.URL, nil)
	resp, err := client.Do(req)
	require.NoError(t, err)
	defer func() { _ = resp.Body.Close() }()

	assert.Equal(t, "TestAgent/1.0", capturedUA)
}

func TestUserAgentTransport_OverridesExistingHeader(t *testing.T) {
	var capturedUA string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		capturedUA = r.Header.Get("User-Agent")
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	transport := &UserAgentTransport{
		Base:      http.DefaultTransport,
		UserAgent: "TestAgent/1.0",
	}
	client := &http.Client{Transport: transport}

	req, _ := http.NewRequest("GET", server.URL, nil)
	req.Header.Set("User-Agent", "OldAgent/0.1")
	resp, err := client.Do(req)
	require.NoError(t, err)
	defer func() { _ = resp.Body.Close() }()

	assert.Equal(t, "TestAgent/1.0", capturedUA)
}

func TestUserAgentTransport_DoesNotModifyOriginalRequest(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	transport := &UserAgentTransport{
		Base:      http.DefaultTransport,
		UserAgent: "TestAgent/1.0",
	}
	client := &http.Client{Transport: transport}

	req, _ := http.NewRequest("GET", server.URL, nil)
	req.Header.Set("User-Agent", "OriginalAgent")
	resp, err := client.Do(req)
	require.NoError(t, err)
	defer func() { _ = resp.Body.Close() }()

	assert.Equal(t, "OriginalAgent", req.Header.Get("User-Agent"))
}

// --- Auth0ClientInfoTransport ---

func TestAuth0ClientInfoTransport_SetsHeader(t *testing.T) {
	var capturedHeader string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		capturedHeader = r.Header.Get("Auth0-Client")
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	transport := &Auth0ClientInfoTransport{
		Base:        http.DefaultTransport,
		HeaderValue: "test-encoded-value",
	}
	client := &http.Client{Transport: transport}

	req, _ := http.NewRequest("GET", server.URL, nil)
	resp, err := client.Do(req)
	require.NoError(t, err)
	defer func() { _ = resp.Body.Close() }()

	assert.Equal(t, "test-encoded-value", capturedHeader)
}

func TestAuth0ClientInfoTransport_DoesNotModifyOriginalRequest(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	transport := &Auth0ClientInfoTransport{
		Base:        http.DefaultTransport,
		HeaderValue: "test-value",
	}
	client := &http.Client{Transport: transport}

	req, _ := http.NewRequest("GET", server.URL, nil)
	resp, err := client.Do(req)
	require.NoError(t, err)
	defer func() { _ = resp.Body.Close() }()

	assert.Empty(t, req.Header.Get("Auth0-Client"))
}

// --- BuildChain ---

func TestBuildChain_DefaultHeaders(t *testing.T) {
	var capturedHeaders http.Header
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		capturedHeaders = r.Header.Clone()
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	chain, err := BuildChain(ChainConfig{})
	require.NoError(t, err)

	client := &http.Client{Transport: chain}
	req, _ := http.NewRequest("GET", server.URL, nil)
	resp, err := client.Do(req)
	require.NoError(t, err)
	defer func() { _ = resp.Body.Close() }()

	assert.Equal(t, "MyOrganization-Go/"+myorganization.SDKVersion, capturedHeaders.Get("User-Agent"))
	assert.NotEmpty(t, capturedHeaders.Get("Auth0-Client"))
}

func TestBuildChain_NoAuth0ClientInfo(t *testing.T) {
	var capturedHeaders http.Header
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		capturedHeaders = r.Header.Clone()
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	chain, err := BuildChain(ChainConfig{
		NoAuth0ClientInfo: true,
	})
	require.NoError(t, err)

	client := &http.Client{Transport: chain}
	req, _ := http.NewRequest("GET", server.URL, nil)
	resp, err := client.Do(req)
	require.NoError(t, err)
	defer func() { _ = resp.Body.Close() }()

	assert.Equal(t, "MyOrganization-Go/"+myorganization.SDKVersion, capturedHeaders.Get("User-Agent"))
	assert.Empty(t, capturedHeaders.Get("Auth0-Client"))
}

func TestBuildChain_WithCustomEnv(t *testing.T) {
	var capturedHeaders http.Header
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		capturedHeaders = r.Header.Clone()
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	chain, err := BuildChain(ChainConfig{
		Auth0ClientEnv: map[string]string{"myapp": "3.0.0"},
	})
	require.NoError(t, err)

	client := &http.Client{Transport: chain}
	req, _ := http.NewRequest("GET", server.URL, nil)
	resp, err := client.Do(req)
	require.NoError(t, err)
	defer func() { _ = resp.Body.Close() }()

	assert.NotEmpty(t, capturedHeaders.Get("Auth0-Client"))
}

func TestBuildChain_NilBase_UsesDefaultTransport(t *testing.T) {
	chain, err := BuildChain(ChainConfig{})
	require.NoError(t, err)
	assert.NotNil(t, chain)
}

// --- CloneHTTPClient ---

func TestCloneHTTPClient_Nil(t *testing.T) {
	client := CloneHTTPClient(nil)
	assert.NotNil(t, client)
}

func TestCloneHTTPClient_StandardClient(t *testing.T) {
	original := &http.Client{Timeout: 5 * time.Second}
	cloned := CloneHTTPClient(original)

	assert.NotNil(t, cloned)
	assert.Equal(t, original.Timeout, cloned.Timeout)

	// Verify it's a different pointer.
	original.Timeout = 10 * time.Second
	assert.NotEqual(t, original.Timeout, cloned.Timeout)
}

func TestCloneHTTPClient_CustomHTTPClient(t *testing.T) {
	custom := &mockHTTPClient{}
	client := CloneHTTPClient(custom)
	assert.NotNil(t, client)
	assert.NotNil(t, client.Transport)
}

type mockHTTPClient struct{}

func (m *mockHTTPClient) Do(_ *http.Request) (*http.Response, error) {
	return &http.Response{StatusCode: http.StatusOK}, nil
}
