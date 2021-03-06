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

// Contact Representation of a Contact
type Contact struct {
	Address *ContactAddress `json:"address,omitempty"`

	// BirthCity Birth city
	BirthCity string `json:"birthCity,omitempty"`

	// BirthCountry Birth Country
	BirthCountry string `json:"birthCountry,omitempty"`

	// BirthDay Birth date
	BirthDay *time.Time `json:"birthDay,omitempty"`

	// BirthZip Birth Zipcode
	BirthZip string `json:"birthZip,omitempty"`

	// CellPhone Cellphone number
	CellPhone string `json:"cellPhone,omitempty"`

	// CompanyNationalIDentificationNumber National identification number of your company
	CompanyNationalIDentificationNumber string `json:"companyNationalIdentificationNumber,omitempty"`

	// Email Email address
	Email string `json:"email,omitempty"`

	// Fax Fax number
	Fax string `json:"fax,omitempty"`

	// FirstName First name
	FirstName string `json:"firstName,omitempty"`

	// Gender Gender
	Gender string `json:"gender,omitempty"`

	// ID Contact Identifier
	ID int64 `json:"id,omitempty"`

	// Language Language
	Language string `json:"language,omitempty"`

	// LastName Last name
	LastName string `json:"lastName,omitempty"`

	// LegalForm Legal form of the contact
	LegalForm string `json:"legalForm,omitempty"`

	// NationalIDentificationNumber National identification number of the contact
	NationalIDentificationNumber string `json:"nationalIdentificationNumber,omitempty"`

	// Nationality Nationality
	Nationality string `json:"nationality,omitempty"`

	// OrganisationName Organisation name
	OrganisationName string `json:"organisationName,omitempty"`

	// OrganisationType Organisation type
	OrganisationType string `json:"organisationType,omitempty"`

	// Phone Telephone number
	Phone string `json:"phone,omitempty"`

	// SpareEmail Spare email address
	SpareEmail string `json:"spareEmail,omitempty"`

	// Vat VAT number
	Vat string `json:"vat,omitempty"`
}
