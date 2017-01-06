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

// VpsAutomatedBackup Backup your VPS
type VpsAutomatedBackup struct {

	// Schedule Scheduled time of your daily backup
	Schedule string `json:"schedule,omitempty"`

	// State Backup state
	State string `json:"state,omitempty"`
}
