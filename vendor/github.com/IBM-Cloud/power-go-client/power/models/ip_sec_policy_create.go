// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// IPSecPolicyCreate IPSec Policy object used for creation
//
// swagger:model IPSecPolicyCreate
type IPSecPolicyCreate struct {

	// authentication
	Authentication IPSECPolicyAuthentication `json:"authentication,omitempty"`

	// Diffie-Hellman group
	// Example: 2
	// Required: true
	// Enum: [1 2 5 14 19 20 24]
	DhGroup *int64 `json:"dhGroup"`

	// connection encryption policy
	// Example: aes-256-cbc
	// Required: true
	// Enum: [aes-256-cbc aes-192-cbc aes-128-cbc aes-256-gcm aes-192-gcm aes-128-gcm 3des-cbc]
	Encryption *string `json:"encryption"`

	// key lifetime
	// Required: true
	KeyLifetime *KeyLifetime `json:"keyLifetime"`

	// IPSec Policy name
	// Example: ipSecPolicy2
	// Required: true
	// Max Length: 47
	// Min Length: 1
	Name *string `json:"name"`

	// Perfect Forward Secrecy
	// Example: true
	// Required: true
	Pfs *bool `json:"pfs"`
}

// Validate validates this IP sec policy create
func (m *IPSecPolicyCreate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthentication(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDhGroup(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEncryption(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKeyLifetime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePfs(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IPSecPolicyCreate) validateAuthentication(formats strfmt.Registry) error {
	if swag.IsZero(m.Authentication) { // not required
		return nil
	}

	if err := m.Authentication.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("authentication")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("authentication")
		}
		return err
	}

	return nil
}

var ipSecPolicyCreateTypeDhGroupPropEnum []interface{}

func init() {
	var res []int64
	if err := json.Unmarshal([]byte(`[1,2,5,14,19,20,24]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		ipSecPolicyCreateTypeDhGroupPropEnum = append(ipSecPolicyCreateTypeDhGroupPropEnum, v)
	}
}

// prop value enum
func (m *IPSecPolicyCreate) validateDhGroupEnum(path, location string, value int64) error {
	if err := validate.EnumCase(path, location, value, ipSecPolicyCreateTypeDhGroupPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *IPSecPolicyCreate) validateDhGroup(formats strfmt.Registry) error {

	if err := validate.Required("dhGroup", "body", m.DhGroup); err != nil {
		return err
	}

	// value enum
	if err := m.validateDhGroupEnum("dhGroup", "body", *m.DhGroup); err != nil {
		return err
	}

	return nil
}

var ipSecPolicyCreateTypeEncryptionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["aes-256-cbc","aes-192-cbc","aes-128-cbc","aes-256-gcm","aes-192-gcm","aes-128-gcm","3des-cbc"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		ipSecPolicyCreateTypeEncryptionPropEnum = append(ipSecPolicyCreateTypeEncryptionPropEnum, v)
	}
}

const (

	// IPSecPolicyCreateEncryptionAesDash256DashCbc captures enum value "aes-256-cbc"
	IPSecPolicyCreateEncryptionAesDash256DashCbc string = "aes-256-cbc"

	// IPSecPolicyCreateEncryptionAesDash192DashCbc captures enum value "aes-192-cbc"
	IPSecPolicyCreateEncryptionAesDash192DashCbc string = "aes-192-cbc"

	// IPSecPolicyCreateEncryptionAesDash128DashCbc captures enum value "aes-128-cbc"
	IPSecPolicyCreateEncryptionAesDash128DashCbc string = "aes-128-cbc"

	// IPSecPolicyCreateEncryptionAesDash256DashGcm captures enum value "aes-256-gcm"
	IPSecPolicyCreateEncryptionAesDash256DashGcm string = "aes-256-gcm"

	// IPSecPolicyCreateEncryptionAesDash192DashGcm captures enum value "aes-192-gcm"
	IPSecPolicyCreateEncryptionAesDash192DashGcm string = "aes-192-gcm"

	// IPSecPolicyCreateEncryptionAesDash128DashGcm captures enum value "aes-128-gcm"
	IPSecPolicyCreateEncryptionAesDash128DashGcm string = "aes-128-gcm"

	// IPSecPolicyCreateEncryptionNr3desDashCbc captures enum value "3des-cbc"
	IPSecPolicyCreateEncryptionNr3desDashCbc string = "3des-cbc"
)

// prop value enum
func (m *IPSecPolicyCreate) validateEncryptionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, ipSecPolicyCreateTypeEncryptionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *IPSecPolicyCreate) validateEncryption(formats strfmt.Registry) error {

	if err := validate.Required("encryption", "body", m.Encryption); err != nil {
		return err
	}

	// value enum
	if err := m.validateEncryptionEnum("encryption", "body", *m.Encryption); err != nil {
		return err
	}

	return nil
}

func (m *IPSecPolicyCreate) validateKeyLifetime(formats strfmt.Registry) error {

	if err := validate.Required("keyLifetime", "body", m.KeyLifetime); err != nil {
		return err
	}

	if err := validate.Required("keyLifetime", "body", m.KeyLifetime); err != nil {
		return err
	}

	if m.KeyLifetime != nil {
		if err := m.KeyLifetime.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("keyLifetime")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("keyLifetime")
			}
			return err
		}
	}

	return nil
}

func (m *IPSecPolicyCreate) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", *m.Name, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", *m.Name, 47); err != nil {
		return err
	}

	return nil
}

func (m *IPSecPolicyCreate) validatePfs(formats strfmt.Registry) error {

	if err := validate.Required("pfs", "body", m.Pfs); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this IP sec policy create based on the context it is used
func (m *IPSecPolicyCreate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAuthentication(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateKeyLifetime(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IPSecPolicyCreate) contextValidateAuthentication(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Authentication.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("authentication")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("authentication")
		}
		return err
	}

	return nil
}

func (m *IPSecPolicyCreate) contextValidateKeyLifetime(ctx context.Context, formats strfmt.Registry) error {

	if m.KeyLifetime != nil {
		if err := m.KeyLifetime.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("keyLifetime")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("keyLifetime")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IPSecPolicyCreate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IPSecPolicyCreate) UnmarshalBinary(b []byte) error {
	var res IPSecPolicyCreate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
