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

// CloudImage Image
type CloudImage struct {

	// CreationDate Image creation date
	CreationDate string `json:"creationDate,omitempty"`

	// ID Image id
	ID string `json:"id,omitempty"`

	// MinDisk Minimum disks required to use image
	MinDisk int64 `json:"minDisk,omitempty"`

	// MinRAM Minimum RAM required to use image
	MinRAM int64 `json:"minRam,omitempty"`

	// Name Image name
	Name string `json:"name,omitempty"`

	// Region Image region
	Region string `json:"region,omitempty"`

	// Size Image size (in GiB)
	Size float64 `json:"size,omitempty"`

	// Status Image status
	Status string `json:"status,omitempty"`

	// TType Image type
	TType string `json:"type,omitempty"`

	// User User to connect with
	User string `json:"user,omitempty"`

	// Visibility Image visibility
	Visibility string `json:"visibility,omitempty"`
}
