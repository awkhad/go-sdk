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

// SaasCsp2BillingStatistics Billing statistics for the current period
type SaasCsp2BillingStatistics struct {

	// EndDate End of the billing period
	EndDate *time.Time `json:"endDate,omitempty"`

	// Lines List of lines associated to this statistics entity.
	Lines []*SaasCsp2BillingStatisticsLine `json:"lines,omitempty"`

	// StartDate Start of the billing period
	StartDate *time.Time `json:"startDate,omitempty"`
}