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

// TelephonyConferenceHistory List past conferences on your number
type TelephonyConferenceHistory struct {

	// CountConnections The count of connections to the conference
	CountConnections int64 `json:"countConnections,omitempty"`

	// CountParticipants The count of unique participants of the conference
	CountParticipants int64 `json:"countParticipants,omitempty"`

	// DateBegin The date the conference began
	DateBegin *time.Time `json:"dateBegin,omitempty"`

	// DateEnd The date the conference end
	DateEnd *time.Time `json:"dateEnd,omitempty"`

	// Duration The duration of the conference in seconds
	Duration int64 `json:"duration,omitempty"`

	// Events The events of the conference (participants joining/left)
	Events []*TelephonyConferenceHistoryEvent `json:"events,omitempty"`

	// ID The id of the conference history
	ID int64 `json:"id,omitempty"`

	// RecordURL The audio record url if you set recording
	RecordURL string `json:"recordUrl,omitempty"`
}
