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

// VpsOption Information about the options of a VPS Virtual Machine
type VpsOption struct {

	// Option The option name
	Option string `json:"option,omitempty"`

	// State The state of the option
	State string `json:"state,omitempty"`
}
