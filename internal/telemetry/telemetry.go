package telemetry

import (
	"encoding/base64"
	"encoding/json"
	"runtime"

	myorganization "github.com/auth0/myorganization-go"
)

// Auth0ClientInfo represents the Auth0-Client telemetry header payload.
type Auth0ClientInfo struct {
	Name    string            `json:"name"`
	Version string            `json:"version"`
	Env     map[string]string `json:"env,omitempty"`
}

// BuildAuth0ClientInfo constructs the telemetry info with default env entries
// and any user-provided custom entries merged in.
func BuildAuth0ClientInfo(customEnv map[string]string) Auth0ClientInfo {
	env := make(map[string]string, len(customEnv)+1)
	for k, v := range customEnv {
		env[k] = v
	}
	// Set "go" last so the runtime version cannot be overridden by custom env.
	env["go"] = runtime.Version()
	return Auth0ClientInfo{
		Name:    myorganization.SDKName,
		Version: myorganization.SDKVersion,
		Env:     env,
	}
}

// EncodeAuth0ClientInfo serializes the telemetry info to a base64-encoded JSON string
// suitable for use as the Auth0-Client header value.
func EncodeAuth0ClientInfo(info Auth0ClientInfo) (string, error) {
	data, err := json.Marshal(info)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(data), nil
}
