/*
 * Spinnaker API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"time"
)

type SpinnakerPluginRelease struct {
	Version string `json:"version,omitempty"`
	Requires string `json:"requires,omitempty"`
	Preferred bool `json:"preferred,omitempty"`
	Date time.Time `json:"date,omitempty"`
	Url string `json:"url,omitempty"`
	Sha512sum string `json:"sha512sum,omitempty"`
	RemoteExtensions []RemoteExtensionConfig `json:"remoteExtensions,omitempty"`
}
