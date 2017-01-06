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

// BillingPaymentMeanValidation A validation required to add a payment mean
type BillingPaymentMeanValidation struct {
	URL string `json:"url,omitempty"`

	ValidationType string `json:"validationType,omitempty"`
}
