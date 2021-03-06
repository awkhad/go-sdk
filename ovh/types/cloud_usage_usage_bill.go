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

// CloudUsageBill UsageBill
type CloudUsageBill struct {

	// BillID ID of the bill
	BillID string `json:"bill_id,omitempty"`

	// Credit Amount of credits used in this bill (not necessarily on part)
	Credit float64 `json:"credit,omitempty"`

	// Part Amount of the bill that accounts for services for the usage period, credits not taken into account
	Part float64 `json:"part,omitempty"`

	// PaymentType Payment type
	PaymentType string `json:"payment_type,omitempty"`

	// Total Total amount of the bill, credits not taken into account
	Total float64 `json:"total,omitempty"`
}
