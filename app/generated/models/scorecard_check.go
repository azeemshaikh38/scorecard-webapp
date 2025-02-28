// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2021 OpenSSF Scorecard Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ScorecardCheck scorecard check
//
// swagger:model ScorecardCheck
type ScorecardCheck struct {

	// details
	Details []string `json:"details"`

	// score
	Score int64 `json:"score"`

	// reason
	Reason string `json:"reason,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// documentation
	Documentation *ScorecardCheckDocumentation `json:"documentation,omitempty"`
}

// Validate validates this scorecard check
func (m *ScorecardCheck) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDocumentation(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ScorecardCheck) validateDocumentation(formats strfmt.Registry) error {
	if swag.IsZero(m.Documentation) { // not required
		return nil
	}

	if m.Documentation != nil {
		if err := m.Documentation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("documentation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("documentation")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this scorecard check based on the context it is used
func (m *ScorecardCheck) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDocumentation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ScorecardCheck) contextValidateDocumentation(ctx context.Context, formats strfmt.Registry) error {

	if m.Documentation != nil {
		if err := m.Documentation.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("documentation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("documentation")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ScorecardCheck) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ScorecardCheck) UnmarshalBinary(b []byte) error {
	var res ScorecardCheck
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ScorecardCheckDocumentation scorecard check documentation
//
// swagger:model ScorecardCheckDocumentation
type ScorecardCheckDocumentation struct {

	// url
	URL string `json:"url,omitempty"`

	// short
	Short string `json:"short,omitempty"`
}

// Validate validates this scorecard check documentation
func (m *ScorecardCheckDocumentation) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this scorecard check documentation based on context it is used
func (m *ScorecardCheckDocumentation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ScorecardCheckDocumentation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ScorecardCheckDocumentation) UnmarshalBinary(b []byte) error {
	var res ScorecardCheckDocumentation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
