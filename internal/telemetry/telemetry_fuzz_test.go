package telemetry

import (
	"encoding/base64"
	"encoding/json"
	"testing"
)

// FuzzEncodeAuth0ClientInfo verifies encoding never panics and always
// produces valid base64 that decodes to valid JSON.
func FuzzEncodeAuth0ClientInfo(f *testing.F) {
	f.Add("myorganization-go", "0.0.1", "go", "go1.21")
	f.Add("", "", "", "")
	f.Add("name-with-special-chars!@#$%", "v1.2.3-beta+build", "key with spaces", "value\nwith\nnewlines")

	f.Fuzz(func(t *testing.T, name, version, envKey, envValue string) {
		info := Auth0ClientInfo{
			Name:    name,
			Version: version,
		}
		if envKey != "" {
			info.Env = map[string]string{envKey: envValue}
		}

		encoded, err := EncodeAuth0ClientInfo(info)
		if err != nil {
			// json.Marshal can fail on certain inputs; that's acceptable.
			return
		}

		// Must be valid base64.
		decoded, err := base64.StdEncoding.DecodeString(encoded)
		if err != nil {
			t.Errorf("EncodeAuth0ClientInfo produced invalid base64: %v", err)
			return
		}

		// Must be valid JSON.
		var roundTrip Auth0ClientInfo
		if err := json.Unmarshal(decoded, &roundTrip); err != nil {
			t.Errorf("EncodeAuth0ClientInfo produced invalid JSON: %v", err)
		}
	})
}
