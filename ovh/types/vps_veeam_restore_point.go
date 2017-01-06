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

// VpsVeeamRestorePoint Informations about a VPS Veeam restore points
type VpsVeeamRestorePoint struct {

	// CreationTime The restore point's creation time
	CreationTime *time.Time `json:"creationTime,omitempty"`

	// ID The restore point's id
	ID int64 `json:"id,omitempty"`
}
