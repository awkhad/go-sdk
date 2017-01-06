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

// CloudProjectVolumeUsageDetail Volume usage
type CloudProjectVolumeUsageDetail struct {
	Price *OrderPrice `json:"price,omitempty"`

	VolumeCapacity *CloudProjectVolumeUsageDetailVolumeCapacity `json:"volumeCapacity,omitempty"`

	// VolumeID Volume id
	VolumeID string `json:"volumeId,omitempty"`

	// VolumeType Volume type
	VolumeType string `json:"volumeType,omitempty"`
}
