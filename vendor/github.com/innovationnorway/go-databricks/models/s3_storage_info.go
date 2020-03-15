// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// S3StorageInfo s3 storage info
// swagger:model S3StorageInfo
type S3StorageInfo struct {

	// canned acl
	CannedACL string `json:"canned_acl,omitempty"`

	// destination
	Destination string `json:"destination,omitempty"`

	// enable encryption
	EnableEncryption bool `json:"enable_encryption,omitempty"`

	// encryption type
	EncryptionType string `json:"encryption_type,omitempty"`

	// endpoint
	Endpoint string `json:"endpoint,omitempty"`

	// kms key
	KmsKey string `json:"kms_key,omitempty"`

	// region
	Region string `json:"region,omitempty"`
}

// Validate validates this s3 storage info
func (m *S3StorageInfo) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *S3StorageInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *S3StorageInfo) UnmarshalBinary(b []byte) error {
	var res S3StorageInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
