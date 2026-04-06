package telemetry

import (
	"encoding/base64"
	"encoding/json"
	"runtime"
	"testing"

	myorganization "github.com/auth0/myorganization-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBuildAuth0ClientInfo_Defaults(t *testing.T) {
	info := BuildAuth0ClientInfo(nil)
	assert.Equal(t, myorganization.SDKName, info.Name)
	assert.Equal(t, myorganization.SDKVersion, info.Version)
	assert.Equal(t, runtime.Version(), info.Env["go"])
}

func TestBuildAuth0ClientInfo_EmptyCustomEnv(t *testing.T) {
	info := BuildAuth0ClientInfo(map[string]string{})
	assert.Equal(t, myorganization.SDKName, info.Name)
	assert.Equal(t, runtime.Version(), info.Env["go"])
	assert.Len(t, info.Env, 1) // only "go"
}

func TestBuildAuth0ClientInfo_CustomEnv(t *testing.T) {
	custom := map[string]string{"myapp": "1.0.0", "mylib": "2.0.0"}
	info := BuildAuth0ClientInfo(custom)
	assert.Equal(t, "1.0.0", info.Env["myapp"])
	assert.Equal(t, "2.0.0", info.Env["mylib"])
	assert.Equal(t, runtime.Version(), info.Env["go"])
	assert.Len(t, info.Env, 3)
}

func TestBuildAuth0ClientInfo_CustomEnvDoesNotOverrideGo(t *testing.T) {
	custom := map[string]string{"go": "custom-go"}
	info := BuildAuth0ClientInfo(custom)
	// The "go" key is always set to the runtime version and cannot be overridden.
	assert.Equal(t, runtime.Version(), info.Env["go"])
}

func TestEncodeAuth0ClientInfo_RoundTrip(t *testing.T) {
	info := Auth0ClientInfo{
		Name:    "test-sdk",
		Version: "1.2.3",
		Env:     map[string]string{"go": "go1.21"},
	}
	encoded, err := EncodeAuth0ClientInfo(info)
	require.NoError(t, err)
	assert.NotEmpty(t, encoded)

	decoded, err := base64.StdEncoding.DecodeString(encoded)
	require.NoError(t, err)

	var roundTrip Auth0ClientInfo
	err = json.Unmarshal(decoded, &roundTrip)
	require.NoError(t, err)

	assert.Equal(t, info.Name, roundTrip.Name)
	assert.Equal(t, info.Version, roundTrip.Version)
	assert.Equal(t, info.Env, roundTrip.Env)
}

func TestEncodeAuth0ClientInfo_IsValidBase64(t *testing.T) {
	info := BuildAuth0ClientInfo(nil)
	encoded, err := EncodeAuth0ClientInfo(info)
	require.NoError(t, err)

	_, err = base64.StdEncoding.DecodeString(encoded)
	assert.NoError(t, err)
}

func TestEncodeAuth0ClientInfo_EmptyEnv(t *testing.T) {
	info := Auth0ClientInfo{
		Name:    "test",
		Version: "0.0.1",
	}
	encoded, err := EncodeAuth0ClientInfo(info)
	require.NoError(t, err)

	decoded, err := base64.StdEncoding.DecodeString(encoded)
	require.NoError(t, err)

	var roundTrip Auth0ClientInfo
	err = json.Unmarshal(decoded, &roundTrip)
	require.NoError(t, err)

	assert.Equal(t, "test", roundTrip.Name)
	assert.Nil(t, roundTrip.Env)
}
