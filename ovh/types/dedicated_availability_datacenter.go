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

// DedicatedAvailabilityDatacenter A structure describing the hardware availability for each datacenter
type DedicatedAvailabilityDatacenter struct {

	// Availability The availability
	Availability string `json:"availability,omitempty"`

	// Datacenter The code of the datacenter
	Datacenter string `json:"datacenter,omitempty"`
}
