package v1alpha1

import (
	ma "github.com/akshaykarle/go-mongodbatlas/mongodbatlas"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MongoDBAtlasClusterList contains a list of MongoDBAtlasCluster
type MongoDBAtlasClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MongoDBAtlasCluster `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MongoDBAtlasCluster is the Schema for the mongodbatlasclusters API
// +k8s:openapi-gen=true
type MongoDBAtlasCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MongoDBAtlasClusterSpec   `json:"spec,omitempty"`
	Status MongoDBAtlasClusterStatus `json:"status,omitempty"`
}

// MongoDBAtlasClusterSpec defines the desired state of MongoDBAtlasCluster
// +k8s:openapi-gen=true
type MongoDBAtlasClusterSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html
	MongoDBAtlasAuth      `json:",inline"`
	ProjectName           string                        `json:"projectName,project"`
	MongoDBVersion        string                        `json:"mongoDBVersion,omitempty"`
	MongoDBMajorVersion   string                        `json:"mongoDBMajorVersion,omitempty"`
	DiskSizeGB            float64                       `json:"diskSizeGB,omitempty"`
	BackupEnabled         bool                          `json:"backupEnabled"`
	ProviderBackupEnabled bool                          `json:"providerBackupEnabled"`
	ReplicationFactor     int                           `json:"replicationFactor,omitempty"`
	ReplicationSpec       map[string]ma.ReplicationSpec `json:"replicationSpec,omitempty"`
	NumShards             int                           `json:"numShards,omitempty"`
	AutoScaling           ma.AutoScaling                `json:"autoScaling,omitempty"`
	ProviderSettings      ma.ProviderSettings           `json:"providerSettings,omitempty"`
}

// MongoDBAtlasClusterStatus defines the observed state of MongoDBAtlasCluster
// +k8s:openapi-gen=true
type MongoDBAtlasClusterStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html
	// ID                    string                        `json:"id"`
	// GroupID               string                        `json:"groupId"`
	// Name                  string                        `json:"name"`
	// MongoDBVersion        string                        `json:"mongoDBVersion"`
	// MongoDBMajorVersion   string                        `json:"mongoDBMajorVersion"`
	// MongoURI              string                        `json:"mongoURI"`
	// MongoURIUpdated       string                        `json:"mongoURIUpdated"`
	// MongoURIWithOptions   string                        `json:"mongoURIWithOptions"`
	// SrvAddress            string                        `json:"srvAddress,omitempty"`
	// DiskSizeGB            float64                       `json:"diskSizeGB"`
	// BackupEnabled         bool                          `json:"backupEnabled"`
	// ProviderBackupEnabled bool                          `json:"providerBackupEnabled"`
	// StateName             string                        `json:"stateName"`
	// ReplicationFactor     int                           `json:"replicationFactor"`
	// ReplicationSpec       map[string]ma.ReplicationSpec `json:"replicationSpec"`
	// NumShards             int                           `json:"numShards"`
	// Paused                bool                          `json:"paused"`
	// AutoScaling           ma.AutoScaling                `json:"autoScaling"`
	// ProviderSettings      ma.ProviderSettings           `json:"providerSettings"`
	ID                    string                        `json:"id,omitempty"`
	GroupID               string                        `json:"groupId,omitempty"`
	Name                  string                        `json:"name,omitempty"`
	MongoDBVersion        string                        `json:"mongoDBVersion,omitempty"`
	MongoDBMajorVersion   string                        `json:"mongoDBMajorVersion,omitempty"`
	MongoURI              string                        `json:"mongoURI,omitempty"`
	MongoURIUpdated       string                        `json:"mongoURIUpdated,omitempty"`
	MongoURIWithOptions   string                        `json:"mongoURIWithOptions,omitempty"`
	SrvAddress            string                        `json:"srvAddress,omitempty"`
	DiskSizeGB            float64                       `json:"diskSizeGB,omitempty"`
	BackupEnabled         bool                          `json:"backupEnabled"`
	ProviderBackupEnabled bool                          `json:"providerBackupEnabled"`
	StateName             string                        `json:"stateName,omitempty"`
	ReplicationFactor     int                           `json:"replicationFactor,omitempty"`
	ReplicationSpec       map[string]ma.ReplicationSpec `json:"replicationSpec,omitempty"`
	NumShards             int                           `json:"numShards,omitempty"`
	Paused                bool                          `json:"paused"`
	AutoScaling           ma.AutoScaling                `json:"autoScaling,omitempty"`
	ProviderSettings      ma.ProviderSettings           `json:"providerSettings,omitempty"`
}

func init() {
	SchemeBuilder.Register(&MongoDBAtlasCluster{}, &MongoDBAtlasClusterList{})
}
