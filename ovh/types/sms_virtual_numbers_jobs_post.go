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

// SmsVirtualNumbersJobsPost ...
type SmsVirtualNumbersJobsPost struct {
	Charset string `json:"charset,omitempty"`

	Class string `json:"class,omitempty"`

	Coding string `json:"coding,omitempty"`

	DifferedPeriod int64 `json:"differedPeriod,omitempty"`

	Message string `json:"message,omitempty"`

	Priority string `json:"priority,omitempty"`

	Receivers []string `json:"receivers,omitempty"`

	ReceiversDocumentURL string `json:"receiversDocumentUrl,omitempty"`

	ReceiversSlotID string `json:"receiversSlotId,omitempty"`

	Tag string `json:"tag,omitempty"`

	ValidityPeriod int64 `json:"validityPeriod,omitempty"`
}
