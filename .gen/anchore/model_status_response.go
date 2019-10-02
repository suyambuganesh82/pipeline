/*
 * Anchore Engine API Server
 *
 * This is the Anchore Engine API. Provides the primary external API for users of the service.
 *
 * API version: 0.1.12
 * Contact: nurmi@anchore.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package anchore

// System status response
type StatusResponse struct {
	Available bool `json:"available,omitempty"`
	Busy bool `json:"busy,omitempty"`
	Up bool `json:"up,omitempty"`
	Message string `json:"message,omitempty"`
	Version string `json:"version,omitempty"`
	DbVersion string `json:"db_version,omitempty"`
	Detail map[string]interface{} `json:"detail,omitempty"`
}
