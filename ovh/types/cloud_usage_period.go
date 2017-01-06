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

// CloudUsagePeriod Period
type CloudUsagePeriod struct {

	// From Usage from
	From *time.Time `json:"from,omitempty"`

	// To Usage to
	To *time.Time `json:"to,omitempty"`
}
