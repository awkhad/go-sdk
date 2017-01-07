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

// XdslPendingAction Scheduled action before the next renewal of the service
type XdslPendingAction struct {
	Action string `json:"action,omitempty"`

	DateTodo *time.Time `json:"dateTodo,omitempty"`
}