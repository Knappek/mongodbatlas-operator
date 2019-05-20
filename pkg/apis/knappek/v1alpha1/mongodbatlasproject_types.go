package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

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
type MongoDBAtlasProject struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MongoDBAtlasProjectSpec   `json:"spec,omitempty"`
	Status MongoDBAtlasProjectStatus `json:"status,omitempty"`
}

// MongoDBAtlasProjectSpec defines the desired state of MongoDBAtlasProject
// +k8s:openapi-gen=true
type MongoDBAtlasProjectSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html
	OrgID    MongoDBAtlasSecret `json:"orgId"`
	Username string             `json:"username"`
	APIKey   MongoDBAtlasSecret `json:"apiKey"`
}

// MongoDBAtlasProjectStatus defines the observed state of MongoDBAtlasProject
// +k8s:openapi-gen=true
type MongoDBAtlasProjectStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html
	ID           string `json:"id"`
	OrgID        string `json:"orgId"`
	Name         string `json:"name"`
	Status       string `json:"status"`
	ClusterCount int    `json:"clusterCount"`
}

// MongoDBAtlasSecret refers to Kubernetes secret objects required to interact with the MongoDB Atlas API
type MongoDBAtlasSecret struct {
	SecretName string `json:"secretName"`
	Key        string `json:"key"`
}

func init() {
	SchemeBuilder.Register(&MongoDBAtlasProject{}, &MongoDBAtlasProjectList{})
}
