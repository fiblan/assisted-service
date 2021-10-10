// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// InfraEnv infra env
//
// swagger:model infra-env
type InfraEnv struct {

	// A comma-separated list of NTP sources (name or IP) going to be added to all the hosts.
	AdditionalNtpSources string `json:"additional_ntp_sources,omitempty"`

	// If set, all hosts that register will be associated with the specified cluster.
	// Format: uuid
	ClusterID strfmt.UUID `json:"cluster_id,omitempty"`

	// The CPU architecture of the image (x86_64/arm64/etc).
	CPUArchitecture string `json:"cpu_architecture,omitempty"`

	// created at
	// Required: true
	// Format: date-time
	CreatedAt *strfmt.DateTime `json:"created_at" gorm:"type:timestamp with time zone"`

	// download url
	DownloadURL string `json:"download_url,omitempty"`

	// email domain
	EmailDomain string `json:"email_domain,omitempty"`

	// expires at
	// Format: date-time
	ExpiresAt strfmt.DateTime `json:"expires_at,omitempty" gorm:"type:timestamp with time zone"`

	// Image generator version.
	GeneratorVersion string `json:"generator_version,omitempty"`

	// Self link.
	// Required: true
	Href *string `json:"href"`

	// Unique identifier of the object.
	// Required: true
	// Format: uuid
	ID *strfmt.UUID `json:"id" gorm:"primary_key"`

	// Json formatted string containing the user overrides for the initial ignition config.
	IgnitionConfigOverride string `json:"ignition_config_override,omitempty"`

	// Indicates the type of this object.
	// Required: true
	// Enum: [InfraEnv]
	Kind *string `json:"kind"`

	// Name of the InfraEnv.
	// Required: true
	Name *string `json:"name"`

	// Version of the OpenShift cluster (used to infer the RHCOS version - temporary until generic logic implemented).
	OpenshiftVersion string `json:"openshift_version,omitempty"`

	// org id
	OrgID string `json:"org_id,omitempty"`

	// proxy
	Proxy *Proxy `json:"proxy,omitempty" gorm:"embedded;embedded_prefix:proxy_"`

	// True if the pull secret has been added to the cluster.
	PullSecretSet bool `json:"pull_secret_set,omitempty"`

	// size bytes
	// Minimum: 0
	SizeBytes *int64 `json:"size_bytes,omitempty"`

	// SSH public key for debugging the installation.
	SSHAuthorizedKey string `json:"ssh_authorized_key,omitempty"`

	// static network configuration string in the format expected by discovery ignition generation.
	StaticNetworkConfig string `json:"static_network_config,omitempty"`

	// type
	// Required: true
	Type ImageType `json:"type"`

	// The last time that this infraenv was updated.
	// Required: true
	// Format: date-time
	UpdatedAt *strfmt.DateTime `json:"updated_at" gorm:"type:timestamp with time zone"`

	// user name
	UserName string `json:"user_name,omitempty"`
}

// Validate validates this infra env
func (m *InfraEnv) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClusterID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExpiresAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHref(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKind(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProxy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSizeBytes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InfraEnv) validateClusterID(formats strfmt.Registry) error {

	if swag.IsZero(m.ClusterID) { // not required
		return nil
	}

	if err := validate.FormatOf("cluster_id", "body", "uuid", m.ClusterID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *InfraEnv) validateCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("created_at", "body", m.CreatedAt); err != nil {
		return err
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *InfraEnv) validateExpiresAt(formats strfmt.Registry) error {

	if swag.IsZero(m.ExpiresAt) { // not required
		return nil
	}

	if err := validate.FormatOf("expires_at", "body", "date-time", m.ExpiresAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *InfraEnv) validateHref(formats strfmt.Registry) error {

	if err := validate.Required("href", "body", m.Href); err != nil {
		return err
	}

	return nil
}

func (m *InfraEnv) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

var infraEnvTypeKindPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["InfraEnv"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		infraEnvTypeKindPropEnum = append(infraEnvTypeKindPropEnum, v)
	}
}

const (

	// InfraEnvKindInfraEnv captures enum value "InfraEnv"
	InfraEnvKindInfraEnv string = "InfraEnv"
)

// prop value enum
func (m *InfraEnv) validateKindEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, infraEnvTypeKindPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *InfraEnv) validateKind(formats strfmt.Registry) error {

	if err := validate.Required("kind", "body", m.Kind); err != nil {
		return err
	}

	// value enum
	if err := m.validateKindEnum("kind", "body", *m.Kind); err != nil {
		return err
	}

	return nil
}

func (m *InfraEnv) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *InfraEnv) validateProxy(formats strfmt.Registry) error {

	if swag.IsZero(m.Proxy) { // not required
		return nil
	}

	if m.Proxy != nil {
		if err := m.Proxy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("proxy")
			}
			return err
		}
	}

	return nil
}

func (m *InfraEnv) validateSizeBytes(formats strfmt.Registry) error {

	if swag.IsZero(m.SizeBytes) { // not required
		return nil
	}

	if err := validate.MinimumInt("size_bytes", "body", int64(*m.SizeBytes), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *InfraEnv) validateType(formats strfmt.Registry) error {

	if err := m.Type.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("type")
		}
		return err
	}

	return nil
}

func (m *InfraEnv) validateUpdatedAt(formats strfmt.Registry) error {

	if err := validate.Required("updated_at", "body", m.UpdatedAt); err != nil {
		return err
	}

	if err := validate.FormatOf("updated_at", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InfraEnv) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InfraEnv) UnmarshalBinary(b []byte) error {
	var res InfraEnv
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
