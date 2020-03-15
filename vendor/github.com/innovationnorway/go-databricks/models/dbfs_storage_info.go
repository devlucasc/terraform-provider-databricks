// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DbfsStorageInfo dbfs storage info
// swagger:model DbfsStorageInfo
type DbfsStorageInfo struct {

	// destination
	Destination string `json:"destination,omitempty"`
}

// Validate validates this dbfs storage info
func (m *DbfsStorageInfo) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DbfsStorageInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DbfsStorageInfo) UnmarshalBinary(b []byte) error {
	var res DbfsStorageInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
