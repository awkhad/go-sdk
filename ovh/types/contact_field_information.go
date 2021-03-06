/*
 * OVH API - EU
 *
 * Build your own OVH world.
 *
 * OpenAPI spec version: 1.0.0
 * Contact: api-subscribe@ml.ovh.net
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package types

// ContactFieldInformation Extras informations about a field
type ContactFieldInformation struct {

	// FieldName Name of the field concerned by restrictions
	FieldName string `json:"fieldName,omitempty"`

	// Mandatory Indicates if the field is mandatory when editing
	Mandatory bool `json:"mandatory,omitempty"`

	// ReadOnly Indicates if the field can't be edited
	ReadOnly bool `json:"readOnly,omitempty"`
}
