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

// PackXdslAddressMoveMovePost ...
type PackXdslAddressMoveMovePost struct {
	Creation *PackXdslAddressMoveCreation `json:"creation,omitempty"`

	KeepCurrentNumber bool `json:"keepCurrentNumber,omitempty"`

	Landline *PackXdslAddressMoveLandline `json:"landline,omitempty"`

	MoveOutDate *time.Time `json:"moveOutDate,omitempty"`

	OfferCode string `json:"offerCode,omitempty"`

	Provider string `json:"provider,omitempty"`
}
