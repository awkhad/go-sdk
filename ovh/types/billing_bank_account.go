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

import (
	"time"
)

// BillingBankAccount SEPA bank account info
type BillingBankAccount struct {
	Bic string `json:"bic,omitempty"`

	CreationDate *time.Time `json:"creationDate,omitempty"`

	DefaultPaymentMean bool `json:"defaultPaymentMean,omitempty"`

	// Description Custom description of this account
	Description string `json:"description,omitempty"`

	Iban string `json:"iban,omitempty"`

	ID int64 `json:"id,omitempty"`

	MandateSignatureDate *time.Time `json:"mandateSignatureDate,omitempty"`

	OwnerAddress string `json:"ownerAddress,omitempty"`

	OwnerName string `json:"ownerName,omitempty"`

	State string `json:"state,omitempty"`

	UniqueReference string `json:"uniqueReference,omitempty"`

	ValidationDocumentLink string `json:"validationDocumentLink,omitempty"`
}
