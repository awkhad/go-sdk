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

// DedicatedServerIPOrderableDetails A structure describing informations about orderable IP address
type DedicatedServerIPOrderableDetails struct {

	// BlockSizes Orderable IP blocks sizes
	BlockSizes []int64 `json:"blockSizes,omitempty"`

	// Included Are those IP included with your offer
	Included bool `json:"included,omitempty"`

	// IPNumber Total number of IP that can be routed to this server
	IPNumber int64 `json:"ipNumber,omitempty"`

	// Number Total number of prefixes that can be routed to this server
	Number int64 `json:"number,omitempty"`

	// OptionRequired Which option is required to order this type of IP
	OptionRequired string `json:"optionRequired,omitempty"`

	// TType this IP type
	TType string `json:"type,omitempty"`
}
