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

// DedicatedNas Storage nas
type DedicatedNas struct {

	// CanCreatePartition True, if partition creation is allowed on this nas
	CanCreatePartition bool `json:"canCreatePartition,omitempty"`

	// CustomName The name you give to the nas
	CustomName string `json:"customName,omitempty"`

	// Datacenter area of nas
	Datacenter string `json:"datacenter,omitempty"`

	// IP Access ip of nas
	IP string `json:"ip,omitempty"`

	// MountPath The storage mount path
	MountPath string `json:"mountPath,omitempty"`

	// ServiceName The storage service name
	ServiceName string `json:"serviceName,omitempty"`

	// ZpoolSize the size of the nas
	ZpoolSize int64 `json:"zpoolSize,omitempty"`
}
