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

// TelephonyLineOptions Line options
type TelephonyLineOptions struct {
	AbsentSubscriber bool `json:"absentSubscriber,omitempty"`

	// AnonymousCallRejection Reject anonymous calls
	AnonymousCallRejection bool `json:"anonymousCallRejection,omitempty"`

	// CallWaiting If the line receives a new call and the line is already in communication, this new call is dispatched once the current call is completed
	CallWaiting bool `json:"callWaiting,omitempty"`

	// Codecs Codecs preferences
	Codecs string `json:"codecs,omitempty"`

	// DefaultVoicemail The default voicemail of the line. It can be the voicemail of any line of your account.
	DefaultVoicemail string `json:"defaultVoicemail,omitempty"`

	DisplayNumber string `json:"displayNumber,omitempty"`

	DoNotDisturb bool `json:"doNotDisturb,omitempty"`

	// Domain The SIP domain of the line
	Domain string `json:"domain,omitempty"`

	// ForwardBackup Enable calls forward when the line is unavailable
	ForwardBackup bool `json:"forwardBackup,omitempty"`

	// ForwardBackupNature Nature of the forward when the line is unavailable
	ForwardBackupNature string `json:"forwardBackupNature,omitempty"`

	// ForwardBackupNumber Destination of the forward when the line is unavailable
	ForwardBackupNumber string `json:"forwardBackupNumber,omitempty"`

	// ForwardBusy Enable calls forward when the line is busy
	ForwardBusy bool `json:"forwardBusy,omitempty"`

	// ForwardBusyNature Nature of the forward when the line is busy
	ForwardBusyNature string `json:"forwardBusyNature,omitempty"`

	// ForwardBusyNumber Destination of the forward when the line is busy
	ForwardBusyNumber string `json:"forwardBusyNumber,omitempty"`

	// ForwardNoReply Enable calls forward on no-reply
	ForwardNoReply bool `json:"forwardNoReply,omitempty"`

	// ForwardNoReplyDelay Delay before forward on no-reply
	ForwardNoReplyDelay int64 `json:"forwardNoReplyDelay,omitempty"`

	// ForwardNoReplyNature Nature of the forward on no-reply
	ForwardNoReplyNature string `json:"forwardNoReplyNature,omitempty"`

	// ForwardNoReplyNumber Destination of the forward on no-reply
	ForwardNoReplyNumber string `json:"forwardNoReplyNumber,omitempty"`

	// ForwardUnconditional Enable unconditional calls forward
	ForwardUnconditional bool `json:"forwardUnconditional,omitempty"`

	// ForwardUnconditionalNature Nature of the unconditional forward
	ForwardUnconditionalNature string `json:"forwardUnconditionalNature,omitempty"`

	// ForwardUnconditionalNumber Destination of the unconditional forward
	ForwardUnconditionalNumber string `json:"forwardUnconditionalNumber,omitempty"`

	// IDentificationRestriction Do not display your number
	IDentificationRestriction bool `json:"identificationRestriction,omitempty"`

	// Intercom Intercom mode: takes automatically the call with the loudspeaker
	Intercom string `json:"intercom,omitempty"`

	// IPRestrictions The ip restrictions of your line
	IPRestrictions []string `json:"ipRestrictions,omitempty"`

	// Language Language of the line
	Language string `json:"language,omitempty"`

	// LockOutCall Disallow outgoing calls
	LockOutCall bool `json:"lockOutCall,omitempty"`

	// LockOutCallPassword Disallow outgoing calls password
	LockOutCallPassword string `json:"lockOutCallPassword,omitempty"`

	// RecordOutgoingCallsBeta Enable or disable record of outgoing calls
	RecordOutgoingCallsBeta bool `json:"recordOutgoingCallsBeta,omitempty"`

	// VoicemailExternalNumber Voicemail number to dial from any other line
	VoicemailExternalNumber string `json:"voicemailExternalNumber,omitempty"`

	// VoicemailInternalNumber Voicemail short number to dial from the line
	VoicemailInternalNumber string `json:"voicemailInternalNumber,omitempty"`
}
