package v1alpha1

import (
	"github.com/Knappek/mongodbatlas-operator/pkg/util"
	ma "github.com/akshaykarle/go-mongodbatlas/mongodbatlas"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// MongoDBAtlasDatabaseUserRequestBody defines the Request Body Parameters when creating/updating a database user
type MongoDBAtlasDatabaseUserRequestBody struct {
	Password        string    `json:"password,omitempty"`
	DeleteAfterDate string    `json:"deleteAfterDate,omitempty"`
	DatabaseName    string    `json:"databaseName,omitempty"`
	Roles           []ma.Role `json:"roles,omitempty"`
}

// MongoDBAtlasDatabaseUserSpec defines the desired state of MongoDBAtlasDatabaseUser
// +k8s:openapi-gen=true
type MongoDBAtlasDatabaseUserSpec struct {
	ProjectName                         string `json:"projectName,project"`
	MongoDBAtlasDatabaseUserRequestBody `json:",inline"`
}

// MongoDBAtlasDatabaseUserStatus defines the observed state of MongoDBAtlasDatabaseUser
// +k8s:openapi-gen=true
type MongoDBAtlasDatabaseUserStatus struct {
	GroupID         string    `json:"groupID,omitempty"`
	Username        string    `json:"username,omitempty"`
	DeleteAfterDate string    `json:"deleteAfterDate,omitempty"`
	DatabaseName    string    `json:"databaseName,omitempty"`
	Roles           []ma.Role `json:"roles,omitempty"`
}

// IsMongoDBAtlasDatabaseUserToBeUpdated is used to compare spec.MongoDBAtlasDatabaseUserRequestBody with status
func IsMongoDBAtlasDatabaseUserToBeUpdated(m1 MongoDBAtlasDatabaseUserRequestBody, m2 MongoDBAtlasDatabaseUserStatus) bool {
	if ok := util.IsNotEqual(m1.DeleteAfterDate, m2.DeleteAfterDate); ok {
		return true
	}
	for idx, role := range m1.Roles {
		if ok := util.IsNotEqual(role.DatabaseName, m2.Roles[idx].DatabaseName); ok {
			return true
		}
		if ok := util.IsNotEqual(role.CollectionName, m2.Roles[idx].CollectionName); ok {
			return true
		}
		if ok := util.IsNotEqual(role.RoleName, m2.Roles[idx].RoleName); ok {
			return true
		}
	}
	return false
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MongoDBAtlasDatabaseUser is the Schema for the mongodbatlasdatabaseusers API
// +k8s:openapi-gen=true
// +kubebuilder:printcolumn:name="Project Name",type="string",JSONPath=".spec.projectName",description="The MongoDB Atlas Project to which the database user has access to"
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=mongodbatlasdatabaseusers,shortName=madbuser
// +kubebuilder:categories=all,mongodbatlas
type MongoDBAtlasDatabaseUser struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MongoDBAtlasDatabaseUserSpec   `json:"spec,omitempty"`
	Status MongoDBAtlasDatabaseUserStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MongoDBAtlasDatabaseUserList contains a list of MongoDBAtlasDatabaseUser
type MongoDBAtlasDatabaseUserList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MongoDBAtlasDatabaseUser `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MongoDBAtlasDatabaseUser{}, &MongoDBAtlasDatabaseUserList{})
}
