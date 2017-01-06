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

// XdslLine Information about the physical copper line
type XdslLine struct {
	ConcentrationPoint *XdslLandlineConcentrationPoint `json:"concentrationPoint,omitempty"`

	Deconsolidation string `json:"deconsolidation,omitempty"`

	// DirectDistribution True if the line is directly wired on the DSLAM
	DirectDistribution bool `json:"directDistribution,omitempty"`

	// Distance Distance in meters from the DSLAM
	Distance int64 `json:"distance,omitempty"`

	FaultRepairTime string `json:"faultRepairTime,omitempty"`

	// LineSectionsLength Detailed information about the sections between the DSLAM and the telephone jack
	LineSectionsLength []*XdslLineSectionLength `json:"lineSectionsLength,omitempty"`

	// Mitigation Mitigation of the line in dB
	Mitigation float64 `json:"mitigation,omitempty"`

	// Number The number of the line
	Number string `json:"number,omitempty"`

	// OriginalNumber The number used to place the order. Null if the same as the current number.
	OriginalNumber string `json:"originalNumber,omitempty"`

	// Portability Whether the line number has been ported to OVH, to be used with VoIP service
	Portability bool `json:"portability,omitempty"`

	// SyncDown The download synchronisation on the DSLAM in Kbps
	SyncDown float64 `json:"syncDown,omitempty"`

	// SyncUp The upload synchronisation on the DSLAM in Kbps
	SyncUp float64 `json:"syncUp,omitempty"`
}
