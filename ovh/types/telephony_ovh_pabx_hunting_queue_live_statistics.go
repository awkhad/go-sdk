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

// TelephonyOvhPabxHuntingQueueLiveStatistics Live statistics of the queue
type TelephonyOvhPabxHuntingQueueLiveStatistics struct {

	// CallsAnswered Total of calls answered
	CallsAnswered int64 `json:"callsAnswered,omitempty"`

	// CallsLost Total of calls lost
	CallsLost int64 `json:"callsLost,omitempty"`

	// CallsTotal Total of calls
	CallsTotal int64 `json:"callsTotal,omitempty"`

	// LastReset Last reset datetime of queue's statistics
	LastReset *time.Time `json:"lastReset,omitempty"`

	// TotalCallDuration Total call duration in seconds
	TotalCallDuration int64 `json:"totalCallDuration,omitempty"`

	// TotalWaitingDuration Total waiting duration in seconds
	TotalWaitingDuration int64 `json:"totalWaitingDuration,omitempty"`
}
