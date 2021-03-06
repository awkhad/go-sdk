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

// NichandleSubAccount Sub Account
type NichandleSubAccount struct {

	// CreationDate Creation date
	CreationDate *time.Time `json:"creationDate,omitempty"`

	// Description This sub-account description
	Description string `json:"description,omitempty"`

	// ID This sub-account id
	ID int64 `json:"id,omitempty"`
}
