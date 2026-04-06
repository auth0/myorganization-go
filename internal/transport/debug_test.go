package transport

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDebugTransport_DisabledReturnsBase(t *testing.T) {
	base := http.DefaultTransport
	result := DebugTransport(base, false)
	assert.Equal(t, base, result)
}

func TestDebugTransport_EnabledReturnsWrapper(t *testing.T) {
	result := DebugTransport(http.DefaultTransport, true)
	_, ok := result.(RoundTripFunc)
	assert.True(t, ok, "expected RoundTripFunc wrapper when debug is enabled")
}

func TestDebugTransport_LogsRequestAndResponse(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"status":"ok"}`))
	}))
	defer server.Close()

	var buf bytes.Buffer
	log.SetOutput(&buf)
	log.SetFlags(0)
	defer log.SetOutput(os.Stderr)
	defer log.SetFlags(log.LstdFlags)

	dt := DebugTransport(http.DefaultTransport, true)
	client := &http.Client{Transport: dt}

	req, err := http.NewRequest("GET", server.URL+"/test-path", nil)
	require.NoError(t, err)
	req.Header.Set("X-Custom", "hello")

	resp, err := client.Do(req)
	require.NoError(t, err)
	defer func() { _ = resp.Body.Close() }()

	output := buf.String()

	// Request section.
	assert.Contains(t, output, "---[ REQUEST ]---")
	assert.Contains(t, output, "GET /test-path")
	assert.Contains(t, output, "X-Custom: hello")

	// Response section.
	assert.Contains(t, output, "---[ RESPONSE ]---")
	assert.Contains(t, output, "200 OK")
	assert.Contains(t, output, `{"status":"ok"}`)
}

func TestDebugTransport_RedactsSensitiveHeaders(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Set-Cookie", "session=secret123")
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	var buf bytes.Buffer
	log.SetOutput(&buf)
	log.SetFlags(0)
	defer log.SetOutput(os.Stderr)
	defer log.SetFlags(log.LstdFlags)

	dt := DebugTransport(http.DefaultTransport, true)
	client := &http.Client{Transport: dt}

	req, err := http.NewRequest("GET", server.URL, nil)
	require.NoError(t, err)
	req.Header.Set("Authorization", "Bearer super-secret-token")
	req.Header.Set("Cookie", "session=abc123")
	req.Header.Set("X-Api-Key", "my-api-key")

	resp, err := client.Do(req)
	require.NoError(t, err)
	defer func() { _ = resp.Body.Close() }()

	output := buf.String()

	// Sensitive values must NOT appear in output.
	assert.NotContains(t, output, "super-secret-token")
	assert.NotContains(t, output, "abc123")
	assert.NotContains(t, output, "my-api-key")
	assert.NotContains(t, output, "secret123")

	// Redacted placeholder must appear.
	assert.Contains(t, output, "[REDACTED]")
}

func TestDebugTransport_DoesNotModifyOriginalRequestHeaders(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	var buf bytes.Buffer
	log.SetOutput(&buf)
	log.SetFlags(0)
	defer log.SetOutput(os.Stderr)
	defer log.SetFlags(log.LstdFlags)

	dt := DebugTransport(http.DefaultTransport, true)
	client := &http.Client{Transport: dt}

	req, err := http.NewRequest("GET", server.URL, nil)
	require.NoError(t, err)
	req.Header.Set("Authorization", "Bearer my-token")

	resp, err := client.Do(req)
	require.NoError(t, err)
	defer func() { _ = resp.Body.Close() }()

	// Original header must be preserved.
	assert.Equal(t, "Bearer my-token", req.Header.Get("Authorization"))
}

func TestDebugTransport_ResponseBodyStillReadable(t *testing.T) {
	expectedBody := `{"data":"hello world"}`
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(expectedBody))
	}))
	defer server.Close()

	var buf bytes.Buffer
	log.SetOutput(&buf)
	log.SetFlags(0)
	defer log.SetOutput(os.Stderr)
	defer log.SetFlags(log.LstdFlags)

	dt := DebugTransport(http.DefaultTransport, true)
	client := &http.Client{Transport: dt}

	req, err := http.NewRequest("GET", server.URL, nil)
	require.NoError(t, err)

	resp, err := client.Do(req)
	require.NoError(t, err)
	defer func() { _ = resp.Body.Close() }()

	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	assert.Equal(t, expectedBody, string(body))
}

func TestDebugTransport_DoesNotModifyOriginalResponseHeaders(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Set-Cookie", "session=secret123")
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	var buf bytes.Buffer
	log.SetOutput(&buf)
	log.SetFlags(0)
	defer log.SetOutput(os.Stderr)
	defer log.SetFlags(log.LstdFlags)

	dt := DebugTransport(http.DefaultTransport, true)
	client := &http.Client{Transport: dt}

	req, err := http.NewRequest("GET", server.URL, nil)
	require.NoError(t, err)

	resp, err := client.Do(req)
	require.NoError(t, err)
	defer func() { _ = resp.Body.Close() }()

	// Original response header must be restored.
	assert.Equal(t, "session=secret123", resp.Header.Get("Set-Cookie"))
}

func TestDebugTransport_RequestAndResponseInSingleLogEntry(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"ok":true}`))
	}))
	defer server.Close()

	// Use a custom logger that records each log call separately.
	var entries []string
	log.SetOutput(logEntryWriter{fn: func(entry string) {
		entries = append(entries, entry)
	}})
	log.SetFlags(0)
	defer log.SetOutput(os.Stderr)
	defer log.SetFlags(log.LstdFlags)

	dt := DebugTransport(http.DefaultTransport, true)
	client := &http.Client{Transport: dt}

	req, err := http.NewRequest("GET", server.URL, nil)
	require.NoError(t, err)

	resp, err := client.Do(req)
	require.NoError(t, err)
	defer func() { _ = resp.Body.Close() }()

	// Both markers must appear in a single log entry (single Printf call).
	require.Len(t, entries, 1, "expected exactly one log entry for request+response pair")
	assert.Contains(t, entries[0], "---[ REQUEST ]---")
	assert.Contains(t, entries[0], "---[ RESPONSE ]---")
}

func TestDebugTransport_ErrorPath_LogsRequestOnly(t *testing.T) {
	// Create a server and immediately close it so connections are refused.
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	serverURL := server.URL
	server.Close()

	var buf bytes.Buffer
	log.SetOutput(&buf)
	log.SetFlags(0)
	defer log.SetOutput(os.Stderr)
	defer log.SetFlags(log.LstdFlags)

	dt := DebugTransport(http.DefaultTransport, true)
	client := &http.Client{Transport: dt}

	req, err := http.NewRequest("GET", serverURL+"/fail", nil)
	require.NoError(t, err)

	_, err = client.Do(req) //nolint:bodyclose
	require.Error(t, err)

	output := buf.String()

	// Request should still be logged even on error.
	assert.Contains(t, output, "---[ REQUEST ]---")
	assert.Contains(t, output, "GET /fail")

	// Response should NOT be logged since the round trip failed.
	assert.NotContains(t, output, "---[ RESPONSE ]---")
}

func TestDebugTransport_RequestWithBody(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	var buf bytes.Buffer
	log.SetOutput(&buf)
	log.SetFlags(0)
	defer log.SetOutput(os.Stderr)
	defer log.SetFlags(log.LstdFlags)

	dt := DebugTransport(http.DefaultTransport, true)
	client := &http.Client{Transport: dt}

	body := strings.NewReader(`{"name":"test"}`)
	req, err := http.NewRequest("POST", server.URL, body)
	require.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	require.NoError(t, err)
	defer func() { _ = resp.Body.Close() }()

	output := buf.String()
	assert.Contains(t, output, "POST")
	assert.Contains(t, output, `{"name":"test"}`)
}

// logEntryWriter captures each Write call (one per log.Printf) as a separate entry.
type logEntryWriter struct {
	fn func(string)
}

func (w logEntryWriter) Write(p []byte) (int, error) {
	w.fn(string(p))
	return len(p), nil
}
