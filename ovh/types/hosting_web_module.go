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

// HostingWebModule Hosting modules installed
type HostingWebModule struct {

	// AdminFolder The admin folder, relative to the module's installation path
	AdminFolder string `json:"adminFolder,omitempty"`

	// AdminName Login for the admin account
	AdminName string `json:"adminName,omitempty"`

	// CreationDate Date of the installation of the module
	CreationDate *time.Time `json:"creationDate,omitempty"`

	// Dependencies The dependencies to which the module has access. A dependency can be a standard database (like MySQL or PostgreSQL) or a key-value store (like Redis or Memcached) for example
	Dependencies []*HostingWebModuleDependencyType `json:"dependencies,omitempty"`

	// ID Installation ID
	ID int64 `json:"id,omitempty"`

	// Language The language of the module
	Language string `json:"language,omitempty"`

	// LastUpdate Date of the last module's upgrade
	LastUpdate *time.Time `json:"lastUpdate,omitempty"`

	// ModuleID ID of the module associated with this installation
	ModuleID int64 `json:"moduleId,omitempty"`

	// Path Where the module is installed, relative to your home directory
	Path string `json:"path,omitempty"`

	// TargetURL The URL from where your module can be reached
	TargetURL string `json:"targetUrl,omitempty"`
}
