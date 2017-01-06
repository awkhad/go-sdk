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

// OrderCartOffice365Post ...
type OrderCartOffice365Post struct {
	Duration string `json:"duration,omitempty"`

	PlanCode string `json:"planCode,omitempty"`

	PricingMode string `json:"pricingMode,omitempty"`

	Quantity int64 `json:"quantity,omitempty"`
}
