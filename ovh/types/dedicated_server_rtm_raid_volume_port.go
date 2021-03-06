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

// DedicatedServerRtmRaidVolumePort Server raid volume port informations
type DedicatedServerRtmRaidVolumePort struct {
	Capacity *DedicatedServerRtmRaidVolumePortCapacity `json:"capacity,omitempty"`

	// Disk Port disk
	Disk string `json:"disk,omitempty"`

	// Model Port model name
	Model string `json:"model,omitempty"`

	// Port Raid volume port
	Port string `json:"port,omitempty"`

	// Serial Serial of this port
	Serial string `json:"serial,omitempty"`

	// Status Status of this port
	Status string `json:"status,omitempty"`

	// Syncprogress Raid port synchronization progress
	Syncprogress string `json:"syncprogress,omitempty"`
}
