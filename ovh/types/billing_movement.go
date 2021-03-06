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

// BillingMovement Details about an OVH account
type BillingMovement struct {
	Amount *OrderPrice `json:"amount,omitempty"`

	Balance *OrderPrice `json:"balance,omitempty"`

	Date *time.Time `json:"date,omitempty"`

	Description string `json:"description,omitempty"`

	MovementID int64 `json:"movementId,omitempty"`

	Operation string `json:"operation,omitempty"`

	Order int64 `json:"order,omitempty"`

	PreviousBalance *OrderPrice `json:"previousBalance,omitempty"`
}
