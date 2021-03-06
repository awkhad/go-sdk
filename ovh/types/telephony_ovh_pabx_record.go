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

// TelephonyOvhPabxRecord The PABX records
type TelephonyOvhPabxRecord struct {

	// Agent The agent number of the recorded call
	Agent string `json:"agent,omitempty"`

	// CallEnd The end date of the recorded call
	CallEnd *time.Time `json:"callEnd,omitempty"`

	// CallStart The begin date of the recorded call
	CallStart *time.Time `json:"callStart,omitempty"`

	// CallerIDName The caller name of the recorded call
	CallerIDName string `json:"callerIdName,omitempty"`

	// CallerIDNumber The caller number of the recorded call
	CallerIDNumber string `json:"callerIdNumber,omitempty"`

	// DestinationNumber The destination number of the recorded call
	DestinationNumber string `json:"destinationNumber,omitempty"`

	// Duration The duration in seconds of the recorded call
	Duration int64 `json:"duration,omitempty"`

	// FileURL The record sound url
	FileURL string `json:"fileUrl,omitempty"`

	ID int64 `json:"id,omitempty"`
}
