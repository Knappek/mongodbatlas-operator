package v1alpha1

import (
	ma "github.com/akshaykarle/go-mongodbatlas/mongodbatlas"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// MongoDBAtlasAlertConfigurationRequestBody defines the Request Body Parameters when creating/updating an alert configuration
type MongoDBAtlasAlertConfigurationRequestBody struct {
	EventTypeName   string             `json:"eventTypeName,omitempty"`
	Enabled         bool               `json:"enabled,omitempty"`
	Notifications   []ma.Notification  `json:"notifications,omitempty"`
	MetricThreshold ma.MetricThreshold `json:"metricThreshold,omitempty"`
	Matchers        []ma.Matcher       `json:"matchers,omitempty"`
}

// MongoDBAtlasAlertConfigurationSpec defines the desired state of MongoDBAtlasAlertConfiguration
// +k8s:openapi-gen=true
type MongoDBAtlasAlertConfigurationSpec struct {
	ProjectName                               string `json:"projectName,project"`
	MongoDBAtlasAlertConfigurationRequestBody `json:",inline"`
}

// MongoDBAtlasAlertConfigurationStatus defines the observed state of MongoDBAtlasAlertConfiguration
// +k8s:openapi-gen=true
type MongoDBAtlasAlertConfigurationStatus struct {
	ID                                        string `json:"id,omitempty"`
	GroupID                                   string `json:"groupID,omitempty"`
	MongoDBAtlasAlertConfigurationRequestBody `json:",inline"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MongoDBAtlasAlertConfiguration is the Schema for the mongodbatlasalertconfigurations API
// +k8s:openapi-gen=true
// +kubebuilder:printcolumn:name="ID",type="string",JSONPath=".status.id",description="The ID of the Alert Configuration"
// +kubebuilder:printcolumn:name="Project Name",type="string",JSONPath=".spec.projectName",description="The MongoDB Atlas Project to which the Alert Configuration is applied"
// +kubebuilder:printcolumn:name="Enabled",type="string",JSONPath=".status.enabled",description="Whether the Alert Configuration is enabled or disabled"
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=mongodbatlasalertconfigurations,shortName=maalertconfig
// +kubebuilder:categories=all,mongodbatlas
type MongoDBAtlasAlertConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MongoDBAtlasAlertConfigurationSpec   `json:"spec,omitempty"`
	Status MongoDBAtlasAlertConfigurationStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MongoDBAtlasAlertConfigurationList contains a list of MongoDBAtlasAlertConfiguration
type MongoDBAtlasAlertConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MongoDBAtlasAlertConfiguration `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MongoDBAtlasAlertConfiguration{}, &MongoDBAtlasAlertConfigurationList{})
}
