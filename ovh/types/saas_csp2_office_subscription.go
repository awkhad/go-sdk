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

// SaasCsp2OfficeSubscription Office subscription
type SaasCsp2OfficeSubscription struct {

	// CreationDate Creation date
	CreationDate *time.Time `json:"creationDate,omitempty"`

	// ID Subscription's unique identifier
	ID int64 `json:"id,omitempty"`

	// LastUpdate Last update date
	LastUpdate *time.Time `json:"lastUpdate,omitempty"`

	// LicenseID License's type id
	LicenseID int64 `json:"licenseId,omitempty"`

	// Quantity Number of available licenses
	Quantity int64 `json:"quantity,omitempty"`

	// Status Subscription's status
	Status string `json:"status,omitempty"`

	// TaskPendingID Pending task's unique identifier
	TaskPendingID int64 `json:"taskPendingId,omitempty"`
}