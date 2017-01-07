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

// DedicatedCloudCommercialRange The commercial ranges actually available in your Dedicated Cloud
type DedicatedCloudCommercialRange struct {

	// AllowedHypervisorVersions The hypervisor versions compliant with this commercial Range
	AllowedHypervisorVersions []string `json:"allowedHypervisorVersions,omitempty"`

	// AllowedNetworkRoles The list of NetworkRoles allowed for one user in this commercial range
	AllowedNetworkRoles []string `json:"allowedNetworkRoles,omitempty"`

	// CommercialRangeName The name of this commercial range
	CommercialRangeName string `json:"commercialRangeName,omitempty"`

	// DedicatedCloudVersion The name of the dedicated Cloud version associated to this commercial range
	DedicatedCloudVersion string `json:"dedicatedCloudVersion,omitempty"`

	// TRange The range of this Datacenter in this Dedicated Cloud version
	TRange string `json:"range,omitempty"`
}