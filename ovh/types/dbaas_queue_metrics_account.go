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

// DBaasQueueMetricsAccount MetricsAccount
type DBaasQueueMetricsAccount struct {

	// Host OpenTSDB host url
	Host string `json:"host,omitempty"`

	// Token Token for OpenTSDB metrics access
	Token string `json:"token,omitempty"`
}