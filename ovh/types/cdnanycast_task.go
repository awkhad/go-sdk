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

// CDNanycastTask Task on a CDN
type CDNanycastTask struct {
	Comment string `json:"comment,omitempty"`

	Function string `json:"function,omitempty"`

	Status string `json:"status,omitempty"`

	TaskID int64 `json:"taskId,omitempty"`
}
