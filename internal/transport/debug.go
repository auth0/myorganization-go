package transport

import (
	"log"
	"net/http"
	"net/http/httputil"
)

// sensitiveHeaders is the set of HTTP headers whose values are redacted
// in debug output to avoid leaking credentials.
var sensitiveHeaders = []string{
	"Authorization",
	"Cookie",
	"Set-Cookie",
	"X-Api-Key",
	"Api-Key",
	"Auth-Token",
	"X-Auth-Token",
}

// RoundTripFunc is a function adapter that implements http.RoundTripper.
type RoundTripFunc func(*http.Request) (*http.Response, error)

// RoundTrip calls the underlying function.
func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}

// DebugTransport wraps base with request/response logging when debug is true.
// When debug is false, base is returned unchanged.
//
// The request is captured before the round trip (so the body is still
// available), and both request and response are emitted together in a
// single log.Printf call after the round trip completes. This keeps the
// pair atomic and prevents interleaving from concurrent goroutines.
func DebugTransport(base http.RoundTripper, debug bool) http.RoundTripper {
	if !debug {
		return base
	}
	return RoundTripFunc(func(req *http.Request) (*http.Response, error) {
		reqDump := captureRequest(req)

		resp, err := base.RoundTrip(req)
		if err != nil {
			log.Printf("\n---[ REQUEST ]---\n%s\n\n", reqDump)
			return resp, err
		}

		respDump := captureResponse(resp)
		log.Printf("\n---[ REQUEST ]---\n%s\n---[ RESPONSE ]---\n%s\n\n", reqDump, respDump)
		return resp, nil
	})
}

// captureRequest returns the wire representation of the request with
// sensitive headers redacted.
func captureRequest(req *http.Request) []byte {
	original := redactSensitiveHeaders(req.Header)
	defer restoreHeaders(req.Header, original)

	dump, err := httputil.DumpRequestOut(req, true)
	if err != nil {
		return []byte("error dumping request: " + err.Error())
	}
	return dump
}

// captureResponse returns the wire representation of the response with
// sensitive headers redacted.
func captureResponse(resp *http.Response) []byte {
	original := redactSensitiveHeaders(resp.Header)
	defer restoreHeaders(resp.Header, original)

	dump, err := httputil.DumpResponse(resp, true)
	if err != nil {
		return []byte("error dumping response: " + err.Error())
	}
	return dump
}

// redactSensitiveHeaders replaces sensitive header values with "[REDACTED]"
// and returns the original values so they can be restored.
func redactSensitiveHeaders(header http.Header) map[string][]string {
	original := make(map[string][]string)
	for _, key := range sensitiveHeaders {
		if vals := header.Values(key); len(vals) > 0 {
			original[key] = vals
			header.Set(key, "[REDACTED]")
		}
	}
	return original
}

// restoreHeaders restores the original header values after dumping.
func restoreHeaders(header http.Header, original map[string][]string) {
	for key, vals := range original {
		header.Del(key)
		for _, v := range vals {
			header.Add(key, v)
		}
	}
}
