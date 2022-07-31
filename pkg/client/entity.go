package client

import "time"

// Type client type enum
type Type string

const (
	// TypeInternal for internal purposes only
	TypeInternal Type = "internal"
	// TypePublic OAuth 2.0 public client
	TypePublic Type = "public"
	// TypeConfidential OAuth 2.0 confidential client
	TypeConfidential Type = "confidential"
)

type Client struct {
	ClientId    string         `yaml:"client_id" json:"clientId"`
	Name        string         `yaml:"name" json:"name"`
	Description string         `yaml:"description" json:"description"`
	Type        Type           `yaml:"type" json:"type"`
	AuthConfig  Authentication `yaml:"auth_config" json:"authConfig"`
}

type Authentication struct {
	OAuth2Flows  OAuth2Flows  `yaml:"oauth2_flows" json:"oauth2_flows"`
	ClientSecret string       `yaml:"client_secret" json:"-"`
	RedirectUris []string     `yaml:"redirect_uris" json:"redirectUris"`
	Scopes       []string     `yaml:"scopes" json:"scopes"`
	TokensConfig TokensConfig `yaml:"tokens" json:"tokens"`
}

type OAuth2Flows struct {
	ClientCredentials OAuth2FlowCfg `yaml:"client_credentials" json:"clientCredentials"`
	AuthCodeGrant     OAuth2FlowCfg `yaml:"auth_code_grant" json:"authCodeGrant"`
	ResourceOwner     OAuth2FlowCfg `yaml:"resource_owner" json:"resourceOwner"`
}

type OAuth2FlowCfg struct {
	Enabled bool `yaml:"enabled" json:"enabled"`
}

type TokensConfig struct {
	Access  TokenConfig `yaml:"access" json:"access"`
	Refresh TokenConfig `yaml:"refresh" json:"refresh"`
	Id      TokenConfig `yaml:"id" json:"id"`
}

type TokenConfig struct {
	ExpiresIn time.Duration `yaml:"expires_in" json:"expiresIn"`
	SingleUse bool          `yaml:"single_use" json:"singleUse"`
}
