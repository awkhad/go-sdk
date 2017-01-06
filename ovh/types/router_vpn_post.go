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

// RouterVpnPost ...
type RouterVpnPost struct {
	ClientIP string `json:"clientIp,omitempty"`

	ClientPrivNet string `json:"clientPrivNet,omitempty"`

	Psk string `json:"psk,omitempty"`

	ServerPrivNet string `json:"serverPrivNet,omitempty"`
}
