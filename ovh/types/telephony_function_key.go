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

// TelephonyFunctionKey Plug & Phone function key
type TelephonyFunctionKey struct {

	// TDefault The default function used by the key
	TDefault string `json:"default,omitempty"`

	// Function The function active on the key
	Function string `json:"function,omitempty"`

	// KeyNum The number of the function key
	KeyNum int64 `json:"keyNum,omitempty"`

	// Label The key label
	Label string `json:"label,omitempty"`

	// Parameter The function parameter
	Parameter string `json:"parameter,omitempty"`

	// TType The key type
	TType string `json:"type,omitempty"`
}
