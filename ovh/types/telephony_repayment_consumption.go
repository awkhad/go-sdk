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

// TelephonyRepaymentConsumption Call which are repayable
type TelephonyRepaymentConsumption struct {

	// Called Called number of the call
	Called string `json:"called,omitempty"`

	// Calling Calling number of the call
	Calling string `json:"calling,omitempty"`

	ConsumptionID int64 `json:"consumptionId,omitempty"`

	// CreationDatetime the Datetime of the start of the call
	CreationDatetime *time.Time `json:"creationDatetime,omitempty"`

	Dialed string `json:"dialed,omitempty"`

	// Duration Duration of the call
	Duration int64 `json:"duration,omitempty"`

	// Price Price repayed with the call
	Price float64 `json:"price,omitempty"`
}
