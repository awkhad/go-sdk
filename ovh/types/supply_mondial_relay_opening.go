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

// SupplyMondialRelayOpening Day with schedule for mondial relay point opening
type SupplyMondialRelayOpening struct {

	// Friday Opening range
	Friday []*SupplyMondialRelayDayPeriod `json:"friday,omitempty"`

	// Monday Opening range
	Monday []*SupplyMondialRelayDayPeriod `json:"monday,omitempty"`

	// Saturday Opening range
	Saturday []*SupplyMondialRelayDayPeriod `json:"saturday,omitempty"`

	// Sunday Opening range
	Sunday []*SupplyMondialRelayDayPeriod `json:"sunday,omitempty"`

	// Thursday Opening range
	Thursday []*SupplyMondialRelayDayPeriod `json:"thursday,omitempty"`

	// Tuesday Opening range
	Tuesday []*SupplyMondialRelayDayPeriod `json:"tuesday,omitempty"`

	// Wednesday Opening range
	Wednesday []*SupplyMondialRelayDayPeriod `json:"wednesday,omitempty"`
}
