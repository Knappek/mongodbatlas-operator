package v1alpha1

import (
	"github.com/Knappek/mongodbatlas-operator/pkg/util"
	ma "github.com/akshaykarle/go-mongodbatlas/mongodbatlas"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.
// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/reference/generating-crd.html

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
// +kubebuilder:printcolumn:name="Project Name",type="string",JSONPath=".spec.projectName",description="The MongoDB Atlas Project where the cluster has been deployed"
// +kubebuilder:printcolumn:name="MongoDB_Version",type="string",JSONPath=".status.mongoDBVersion",description="The MongoDB version of the cluster"
// +kubebuilder:printcolumn:name="Status",type="string",JSONPath=".status.stateName",description="The status of the cluster"
// +kubebuilder:printcolumn:name="Region",type="string",JSONPath=".status.providerSettings.regionName",description="Physical location of your MongoDB cluster"
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:printcolumn:name="Provider",type="string",JSONPath=".status.providerSettings.providerName",description="Cloud service provider on which the servers are provisioned",priority="1"
// +kubebuilder:printcolumn:name="ContinuousBackups",type="boolean",JSONPath=".status.backupEnabled",description="Set to true to enable Atlas continuous backups for the cluster",priority="1"
// +kubebuilder:printcolumn:name="ProviderBackups",type="boolean",JSONPath=".status.providerBackupEnabled",description="Flag indicating if the cluster uses Cloud Provider Snapshots for backups",priority="1"
// +kubebuilder:printcolumn:name="SRV_Address",type="string",JSONPath=".status.srvAddress",description="Connection string (DNS SRV Record) for connecting to the Atlas cluster",priority="1"
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=mongodbatlasclusters,shortName=mac
// +kubebuilder:categories=all,mongodbatlas
type MongoDBAtlasCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MongoDBAtlasClusterSpec   `json:"spec,omitempty"`
	Status MongoDBAtlasClusterStatus `json:"status,omitempty"`
}

// MongoDBAtlasClusterRequestBody defines the Request Body Parameters when creating/updating a cluster
type MongoDBAtlasClusterRequestBody struct {
	MongoDBMajorVersion   string  `json:"mongoDBMajorVersion,omitempty"`
	DiskSizeGB            float64 `json:"diskSizeGB,omitempty"`
	BackupEnabled         bool    `json:"backupEnabled"`
	ProviderBackupEnabled bool    `json:"providerBackupEnabled"`
	// TODO: ReplicationSpec is deprecated, update to ReplicationSpecs.
	// This needs to be done in the Go clinet library first: https://github.com/akshaykarle/go-mongodbatlas
	ReplicationSpec  map[string]ma.ReplicationSpec `json:"replicationSpec,omitempty"`
	NumShards        int                           `json:"numShards,omitempty"`
	AutoScaling      ma.AutoScaling                `json:"autoScaling,omitempty"`
	ProviderSettings ma.ProviderSettings           `json:"providerSettings,omitempty"`
}

// IsMongoDBAtlasClusterToBeUpdated is used to compare spec.MongoDBAtlasClusterRequestBody with status.MongoDBAtlasClusterRequestBody
func IsMongoDBAtlasClusterToBeUpdated(m1 MongoDBAtlasClusterRequestBody, m2 MongoDBAtlasClusterRequestBody) bool {
	region := m1.ProviderSettings.RegionName
	if ok := util.IsNotEqual(m1.MongoDBMajorVersion, m2.MongoDBMajorVersion); ok {
		return true
	}
	if ok := util.IsNotEqual(m1.DiskSizeGB, m2.DiskSizeGB); ok {
		return true
	}
	if ok := util.IsNotEqual(m1.BackupEnabled, m2.BackupEnabled); ok {
		return true
	}
	if ok := util.IsNotEqual(m1.ReplicationSpec[region].Priority, m2.ReplicationSpec[region].Priority); ok {
		return true
	}
	if ok := util.IsNotEqual(m1.ReplicationSpec[region].ElectableNodes, m2.ReplicationSpec[region].ElectableNodes); ok {
		return true
	}
	if ok := util.IsNotEqual(m1.ReplicationSpec[region].ReadOnlyNodes, m2.ReplicationSpec[region].ReadOnlyNodes); ok {
		return true
	}
	if ok := util.IsNotEqual(m1.ReplicationSpec[region].AnalyticsNodes, m2.ReplicationSpec[region].AnalyticsNodes); ok {
		return true
	}
	if ok := util.IsNotEqual(m1.NumShards, m2.NumShards); ok {
		return true
	}
	if ok := util.IsNotEqual(m1.AutoScaling.DiskGBEnabled, m2.AutoScaling.DiskGBEnabled); ok {
		return true
	}
	if ok := util.IsNotEqual(m1.ProviderSettings.ProviderName, m2.ProviderSettings.ProviderName); ok {
		return true
	}
	if ok := util.IsNotEqual(m1.ProviderSettings.BackingProviderName, m2.ProviderSettings.BackingProviderName); ok {
		return true
	}
	if ok := util.IsNotEqual(m1.ProviderSettings.RegionName, m2.ProviderSettings.RegionName); ok {
		return true
	}
	if ok := util.IsNotEqual(m1.ProviderSettings.InstanceSizeName, m2.ProviderSettings.InstanceSizeName); ok {
		return true
	}
	if ok := util.IsNotEqual(m1.ProviderSettings.DiskIOPS, m2.ProviderSettings.DiskIOPS); ok {
		return true
	}
	if ok := util.IsNotEqual(m1.ProviderSettings.EncryptEBSVolume, m2.ProviderSettings.EncryptEBSVolume); ok {
		return true
	}
	return false
}



// MongoDBAtlasClusterSpec defines the desired state of MongoDBAtlasCluster
// +k8s:openapi-gen=true
type MongoDBAtlasClusterSpec struct {
	ProjectName                    string `json:"projectName,project"`
	MongoDBAtlasClusterRequestBody `json:",inline"`
}

// MongoDBAtlasClusterStatus defines the observed state of MongoDBAtlasCluster
// +k8s:openapi-gen=true
type MongoDBAtlasClusterStatus struct {
	ID                             string `json:"id,omitempty"`
	GroupID                        string `json:"groupID,omitempty"`
	Name                           string `json:"name,omitempty"`
	MongoDBAtlasClusterRequestBody `json:",inline"`
	MongoDBVersion                 string `json:"mongoDBVersion,omitempty"`
	MongoURI                       string `json:"mongoURI,omitempty"`
	MongoURIUpdated                string `json:"mongoURIUpdated,omitempty"`
	MongoURIWithOptions            string `json:"mongoURIWithOptions,omitempty"`
	SrvAddress                     string `json:"srvAddress,omitempty"`
	StateName                      string `json:"stateName,omitempty"`
	Paused                         bool   `json:"paused"`
}

func init() {
	SchemeBuilder.Register(&MongoDBAtlasCluster{}, &MongoDBAtlasClusterList{})
}
