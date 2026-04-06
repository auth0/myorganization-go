package core

import (
	"context"

	"golang.org/x/oauth2"
)

// ClientCredentialsOption configures OAuth2 client credentials authentication.
type ClientCredentialsOption struct {
	Context      context.Context
	ClientID     string
	ClientSecret string
	Audience     string
}

func (c *ClientCredentialsOption) applyRequestOptions(opts *RequestOptions) {
	opts.ClientID = c.ClientID
	opts.ClientSecret = c.ClientSecret
	opts.OAuthContext = c.Context
	opts.Audience = c.Audience
}

// PrivateKeyJwtOption configures Private Key JWT authentication.
type PrivateKeyJwtOption struct {
	Context          context.Context
	ClientID         string
	PrivateKeyPEM    string
	SigningAlgorithm string
	Audience         string
}

func (p *PrivateKeyJwtOption) applyRequestOptions(opts *RequestOptions) {
	opts.ClientID = p.ClientID
	opts.PrivateKeyPEM = p.PrivateKeyPEM
	opts.SigningAlgorithm = p.SigningAlgorithm
	opts.OAuthContext = p.Context
	opts.Audience = p.Audience
}

// TokenSourceOption configures a custom oauth2.TokenSource for authentication.
type TokenSourceOption struct {
	TokenSource oauth2.TokenSource
}

func (t *TokenSourceOption) applyRequestOptions(opts *RequestOptions) {
	opts.TokenSource = t.TokenSource
}

// NoAuth0ClientInfoOption disables the Auth0-Client telemetry header.
type NoAuth0ClientInfoOption struct{}

func (n *NoAuth0ClientInfoOption) applyRequestOptions(opts *RequestOptions) {
	opts.NoAuth0ClientInfo = true
}

// Auth0ClientEnvEntryOption adds a custom entry to the Auth0-Client env map.
type Auth0ClientEnvEntryOption struct {
	Key   string
	Value string
}

func (a *Auth0ClientEnvEntryOption) applyRequestOptions(opts *RequestOptions) {
	if opts.Auth0ClientEnv == nil {
		opts.Auth0ClientEnv = make(map[string]string)
	}
	opts.Auth0ClientEnv[a.Key] = a.Value
}

// OrganizationOption configures the organization name or ID sent as the
// "organization" parameter in OAuth2 token requests. It does not affect
// base URL or audience derivation.
type OrganizationOption struct {
	Organization string
}

func (o *OrganizationOption) applyRequestOptions(opts *RequestOptions) {
	opts.Organization = o.Organization
}

// DebugOption enables HTTP request/response debug logging.
type DebugOption struct {
	Debug bool
}

func (d *DebugOption) applyRequestOptions(opts *RequestOptions) {
	opts.Debug = d.Debug
}
