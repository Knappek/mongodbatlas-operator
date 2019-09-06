package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.
// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/reference/generating-crd.html

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MongoDBAtlasProjectList contains a list of MongoDBAtlasProject
type MongoDBAtlasProjectList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MongoDBAtlasProject `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MongoDBAtlasProject is the Schema for the mongodbatlasprojects API
// +k8s:openapi-gen=true
// +kubebuilder:printcolumn:name="OrgID",type="string",JSONPath=".spec.orgID",description="The MongoDB Atlas Organization ID"
// +kubebuilder:printcolumn:name="ClusterCount",type="integer",JSONPath=".status.clusterCount",description="The number of Clusters in the Project"
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:printcolumn:name="ProjectID",type="string",JSONPath=".status.id",description="The ID of the Project",priority="1"
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=mongodbatlasprojects,shortName=map
// +kubebuilder:categories=all,mongodbatlas
type MongoDBAtlasProject struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MongoDBAtlasProjectSpec   `json:"spec,omitempty"`
	Status MongoDBAtlasProjectStatus `json:"status,omitempty"`
}

// MongoDBAtlasProjectSpec defines the desired state of MongoDBAtlasProject
// +k8s:openapi-gen=true
type MongoDBAtlasProjectSpec struct {
	OrgID string `json:"orgID"`
}

// MongoDBAtlasProjectStatus defines the observed state of MongoDBAtlasProject
// +k8s:openapi-gen=true
type MongoDBAtlasProjectStatus struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	OrgID        string `json:"orgID"`
	Created      string `json:"created"`
	ClusterCount int    `json:"clusterCount"`
}

func init() {
	SchemeBuilder.Register(&MongoDBAtlasProject{}, &MongoDBAtlasProjectList{})
}
