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

// HostingWebCronPost ...
type HostingWebCronPost struct {
	Command string `json:"command,omitempty"`

	Description string `json:"description,omitempty"`

	Email string `json:"email,omitempty"`

	Frequency string `json:"frequency,omitempty"`

	Language string `json:"language,omitempty"`

	Status string `json:"status,omitempty"`
}
