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

// HpcspotConsumptionQuantity Quantity consumed (minutes, core minutes, licences, token)
type HpcspotConsumptionQuantity struct {
	Unit string `json:"unit,omitempty"`

	Value int64 `json:"value,omitempty"`
}
