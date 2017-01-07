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

// TelephonyOfferTask Operation on a telephony offer
type TelephonyOfferTask struct {

	// Action Actual action that will be executed
	Action string `json:"action,omitempty"`

	// ExecutionDate Planned execution date
	ExecutionDate *time.Time `json:"executionDate,omitempty"`

	// Status Current status of the task
	Status string `json:"status,omitempty"`

	TaskID int64 `json:"taskId,omitempty"`

	// TType Type of operation that will be executed
	TType string `json:"type,omitempty"`
}