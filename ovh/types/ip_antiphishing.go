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

// IPAntiphishing Phishing URLs hosted on your IP
type IPAntiphishing struct {

	// CreationDate Date of the event
	CreationDate *time.Time `json:"creationDate,omitempty"`

	// ID Internal ID of the phishing entry
	ID int64 `json:"id,omitempty"`

	// IPOnAntiphishing IP address hosting the phishing URL
	IPOnAntiphishing string `json:"ipOnAntiphishing,omitempty"`

	// State Current state of the phishing
	State string `json:"state,omitempty"`

	// URLPhishing Phishing URL
	URLPhishing string `json:"urlPhishing,omitempty"`
}
