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

// TelephonyResetPhoneInfo Relevant informations of the phone reset
type TelephonyResetPhoneInfo struct {
	ResetCodeInfo *TelephonyResetPhoneCodeInfo `json:"resetCodeInfo,omitempty"`

	// ResetPhoneMethod Which way had been used to reset the phone
	ResetPhoneMethod string `json:"resetPhoneMethod,omitempty"`
}
