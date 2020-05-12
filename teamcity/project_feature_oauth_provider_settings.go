package teamcity

import (
	"fmt"
	"strconv"
)

type ProjectFeatureOauthProviderSettingsOptions struct {
	DisplayName        string
	Endpoint           string
	FailOnError        bool
	ParameterNamespace string
	VaultNamespace     string
	ProviderType       string
	RoleId             string
	SecretId           string
	Url                string
}

// ProjectFeatureOauthProviderSettings represents the oauth provider settings feature for a project.
type ProjectFeatureOauthProviderSettings struct {
	id        string
	projectID string

	Options ProjectFeatureOauthProviderSettingsOptions
}

func NewProjectFeatureOauthProviderSettings(projectID string, options ProjectFeatureOauthProviderSettingsOptions) *ProjectFeatureOauthProviderSettings {
	return &ProjectFeatureOauthProviderSettings{
		projectID: projectID,
		Options:   options,
	}
}

// ID returns the ID of this project feature.
func (f *ProjectFeatureOauthProviderSettings) ID() string {
	return f.id
}

// SetID sets the ID of this project feature.
func (f *ProjectFeatureOauthProviderSettings) SetID(value string) {
	f.id = value
}

// Type represents the type of the project feature as a string.
func (f *ProjectFeatureOauthProviderSettings) Type() string {
	return "OAuthProvider"
}

// ProjectID represents the ID of the project the project feature is assigned to.
func (f *ProjectFeatureOauthProviderSettings) ProjectID() string {
	return f.projectID
}

// SetProjectID sets the ID of the project the project feature is assigned to.
func (f *ProjectFeatureOauthProviderSettings) SetProjectID(value string) {
	f.projectID = value
}

// Properties returns all properties for the oauth provider settings project feature.
func (f *ProjectFeatureOauthProviderSettings) Properties() *Properties {
	props := NewProperties(
		NewProperty("displayName", string(f.Options.DisplayName)),
		NewProperty("endpoint", string(f.Options.Endpoint)),
		NewProperty("fail-on-error", fmt.Sprintf("%t", f.Options.FailOnError)),
		NewProperty("namespace", string(f.Options.ParameterNamespace)),
		NewProperty("vault-namespace", string(f.Options.VaultNamespace)),
		NewProperty("providerType", string(f.Options.ProviderType)),
		NewProperty("role-id", string(f.Options.RoleId)),
		NewProperty("secure:secret-id", string(f.Options.SecretId)),
		NewProperty("url", string(f.Options.Url)),
	)

	return props
}

func loadProjectFeatureOauthProviderSettings(projectID string, feature projectFeatureJSON) (ProjectFeature, error) {
	settings := &ProjectFeatureOauthProviderSettings{
		id:        feature.ID,
		projectID: projectID,
		Options:   ProjectFeatureOauthProviderSettingsOptions{},
	}

	if encodedValue, ok := feature.Properties.GetOk("displayName"); ok {
		settings.Options.DisplayName = encodedValue
	}

	if encodedValue, ok := feature.Properties.GetOk("endpoint"); ok {
		settings.Options.Endpoint = encodedValue
	}

	if encodedValue, ok := feature.Properties.GetOk("fail-on-error"); ok {
		v, err := strconv.ParseBool(encodedValue)
		if err != nil {
			return nil, err
		}
		settings.Options.FailOnError = v
	}

	if encodedValue, ok := feature.Properties.GetOk("namespace"); ok {
		settings.Options.ParameterNamespace = encodedValue
	}

	if encodedValue, ok := feature.Properties.GetOk("vault-namespace"); ok {
		settings.Options.VaultNamespace = encodedValue
	}

	if encodedValue, ok := feature.Properties.GetOk("providerType"); ok {
		settings.Options.ProviderType = encodedValue
	}

	if encodedValue, ok := feature.Properties.GetOk("role-id"); ok {
		settings.Options.RoleId = encodedValue
	}

	settings.Options.SecretId = ""

	if encodedValue, ok := feature.Properties.GetOk("url"); ok {
		settings.Options.Url = encodedValue
	}
	return settings, nil
}
