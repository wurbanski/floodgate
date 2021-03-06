/*
 * Spinnaker API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type ConstraintState struct {
	Status string `json:"status,omitempty"`
	Type_ string `json:"type,omitempty"`
	DeliveryConfigName string `json:"deliveryConfigName,omitempty"`
	JudgedBy string `json:"judgedBy,omitempty"`
	Attributes *interface{} `json:"attributes,omitempty"`
	ArtifactVersion string `json:"artifactVersion,omitempty"`
	EnvironmentName string `json:"environmentName,omitempty"`
	CreatedAt string `json:"createdAt,omitempty"`
	JudgedAt string `json:"judgedAt,omitempty"`
	Comment string `json:"comment,omitempty"`
}
