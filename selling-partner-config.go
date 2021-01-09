package amazonspapi

import (
	"errors"
	"os"
	"regexp"
)

type SellingPartnerConfig struct {
	Region       string
	RefreshToken string
	AccessToken  string

	RoleCredentials *RoleCredentialsConfig
	Options         *OptionsConfig
}

func (o SellingPartnerConfig) IsValid() (bool, error) {
	if o.RefreshToken == "" {
		return false, errors.New("refresh token is required")
	}

	if doesMatch, err := regexp.MatchString("^(eu|na|fe)$", o.Region); !doesMatch || err != nil {
		return false, errors.New("region should be one of eu, na, or fe")
	}

	return true, nil
}

type RoleCredentialsConfig struct {
	Id            string
	Secret        string
	SecurityToken string
}

type OptionsConfig struct {
	AutoRequestTokens    bool
	AutoRequestThrottled bool
	Debug                bool
}

type CredentialsConfig struct {
	AppClient *AppClientConfig
	AWSUser   *AWSUserConfig
}

type AppClientConfig struct {
	Id string
	Secret string
}

type AWSUserConfig struct {
	Id string
	Secret string
	Role string
}

func NewCredentialsConfig() (*CredentialsConfig, error) {
	config := &CredentialsConfig{}

	if appClient, err := NewAppClientConfig(); err != nil {
		return nil, err
	} else {
		config.AppClient = appClient
	}

	if awsUser, err := NewAWSUserConfig(); err != nil {
		return nil, err
	} else {
		config.AWSUser = awsUser
	}

	return config, nil
}

func NewAppClientConfig() (*AppClientConfig, error) {
	config := &AppClientConfig{}

	if val, ok := os.LookupEnv("SELLING_PARTNER_APP_CLIENT_ID"); ok && val != "" {
		config.Id = val
	} else {
		return nil, errors.New("SELLING_PARTNER_APP_CLIENT_ID env var missing")
	}

	if val, ok := os.LookupEnv("SELLING_PARTNER_APP_CLIENT_SECRET"); ok && val != "" {
		config.Secret = val
	} else {
		return nil, errors.New("SELLING_PARTNER_APP_CLIENT_SECRET env var missing")
	}

	return config, nil
}

func NewAWSUserConfig() (*AWSUserConfig, error) {
	config := &AWSUserConfig{}

	if val, ok := os.LookupEnv("AWS_ACCESS_KEY_ID"); ok && val != "" {
		config.Id = val
	} else {
		return nil, errors.New("AWS_ACCESS_KEY_ID env var missing")
	}

	if val, ok := os.LookupEnv("AWS_SECRET_ACCESS_KEY"); ok && val != "" {
		config.Secret = val
	} else {
		return nil, errors.New("AWS_SECRET_ACCESS_KEY env var missing")
	}

	if val, ok := os.LookupEnv("AWS_SELLING_PARTNER_ROLE"); ok && val != "" {
		config.Role = val
	} else {
		return nil, errors.New("AWS_SELLING_PARTNER_ROLE env var missing")
	}

	return config, nil
}

func NewOptionsConfig() *OptionsConfig {
	return &OptionsConfig{
		AutoRequestThrottled: true,
		AutoRequestTokens:    true,
		Debug:                false,
	}
}

func NewSellingPartnerConfig() *SellingPartnerConfig {
	return &SellingPartnerConfig{
		Options: NewOptionsConfig(),
	}
}
