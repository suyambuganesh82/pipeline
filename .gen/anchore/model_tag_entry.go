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
import (
	"time"
)

// A docker-pullable tag value as well as deconstructed components
type TagEntry struct {
	// The pullable string for the tag. E.g. \"docker.io/library/node:latest\"
	Pullstring string `json:"pullstring,omitempty"`
	// The registry hostname:port section of the pull string
	Registry string `json:"registry,omitempty"`
	// The repository section of the pull string
	Repository string `json:"repository,omitempty"`
	// The tag-only section of the pull string
	Tag string `json:"tag,omitempty"`
	// The timestamp at which the Anchore Engine detected this tag was mapped to the image digest. Does not necessarily indicate when the tag was actually pushed to the registry.
	DetectedAt time.Time `json:"detected_at,omitempty"`
}
