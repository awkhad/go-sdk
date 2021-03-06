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

// PackXdslAddressMoveOffer An offer
type PackXdslAddressMoveOffer struct {
	Address *XdslEligibilityAddress `json:"address,omitempty"`

	// EstimatedDownload The estimated download synchronisation in kbps
	EstimatedDownload int64 `json:"estimatedDownload,omitempty"`

	// EstimatedUpload The estimated upload synchronisation in kbps
	EstimatedUpload int64 `json:"estimatedUpload,omitempty"`

	// LineSectionsLength Detailed information about the sections between the DSLAM and the telephone jack
	LineSectionsLength []*XdslLineSectionLength `json:"lineSectionsLength,omitempty"`

	// LineStatus The status of the landline
	LineStatus string `json:"lineStatus,omitempty"`

	MeetingSlots *XdslEligibilityMeetingSlots `json:"meetingSlots,omitempty"`

	// Nra The NRA of the landline
	Nra string `json:"nra,omitempty"`

	// OfferCode The code of the offer
	OfferCode string `json:"offerCode,omitempty"`

	Portability *XdslEligibilityPortability `json:"portability,omitempty"`

	Price *OrderPrice `json:"price,omitempty"`

	// Provider Status of the request
	Provider string `json:"provider,omitempty"`

	// Reseller Whether this is a reseller offer or not
	Reseller bool `json:"reseller,omitempty"`

	// SyncDownload The download synchronisation in kbps
	SyncDownload int64 `json:"syncDownload,omitempty"`

	// SyncUpload The upload synchronisation in kbps
	SyncUpload int64 `json:"syncUpload,omitempty"`

	// TType DSL technology
	TType string `json:"type,omitempty"`

	// Unbundling The unbundling type
	Unbundling string `json:"unbundling,omitempty"`
}
