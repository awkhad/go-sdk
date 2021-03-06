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

// TelephonyOvhPabxMenuPost ...
type TelephonyOvhPabxMenuPost struct {
	GreetSound int64 `json:"greetSound,omitempty"`

	GreetSoundTts int64 `json:"greetSoundTts,omitempty"`

	InvalidSound int64 `json:"invalidSound,omitempty"`

	InvalidSoundTts int64 `json:"invalidSoundTts,omitempty"`

	Name string `json:"name,omitempty"`
}
