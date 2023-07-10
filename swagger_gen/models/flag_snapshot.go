// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// FlagSnapshot flag snapshot
//
// swagger:model flagSnapshot
type FlagSnapshot struct {

	// flag
	// Required: true
	Flag *Flag `json:"flag"`

	// id
	// Required: true
	// Minimum: 1
	ID *int64 `json:"id"`

	// updated at
	// Required: true
	// Min Length: 1
	UpdatedAt *string `json:"updatedAt"`

	// updated by
	UpdatedBy string `json:"updatedBy,omitempty"`
}

// Validate validates this flag snapshot
func (m *FlagSnapshot) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFlag(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
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

func (m *FlagSnapshot) validateFlag(formats strfmt.Registry) error {

	if err := validate.Required("flag", "body", m.Flag); err != nil {
		return err
	}

	if m.Flag != nil {
		if err := m.Flag.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("flag")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("flag")
			}
			return err
		}
	}

	return nil
}

func (m *FlagSnapshot) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.MinimumInt("id", "body", *m.ID, 1, false); err != nil {
		return err
	}

	return nil
}

func (m *FlagSnapshot) validateUpdatedAt(formats strfmt.Registry) error {

	if err := validate.Required("updatedAt", "body", m.UpdatedAt); err != nil {
		return err
	}

	if err := validate.MinLength("updatedAt", "body", *m.UpdatedAt, 1); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this flag snapshot based on the context it is used
func (m *FlagSnapshot) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFlag(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FlagSnapshot) contextValidateFlag(ctx context.Context, formats strfmt.Registry) error {

	if m.Flag != nil {
		if err := m.Flag.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("flag")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("flag")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FlagSnapshot) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FlagSnapshot) UnmarshalBinary(b []byte) error {
	var res FlagSnapshot
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}