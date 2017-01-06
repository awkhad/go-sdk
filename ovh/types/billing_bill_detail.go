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

import (
	"time"
)

// BillingBillDetail Information about a Bill entry
type BillingBillDetail struct {
	BillDetailID string `json:"billDetailId,omitempty"`

	Description string `json:"description,omitempty"`

	Domain string `json:"domain,omitempty"`

	PeriodEnd *time.Time `json:"periodEnd,omitempty"`

	PeriodStart *time.Time `json:"periodStart,omitempty"`

	Quantity string `json:"quantity,omitempty"`

	TotalPrice *OrderPrice `json:"totalPrice,omitempty"`

	UnitPrice *OrderPrice `json:"unitPrice,omitempty"`
}
