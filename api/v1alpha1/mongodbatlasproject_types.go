/*


Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// MongoDBAtlasProjectSpec defines the desired state of MongoDBAtlasProject
type MongoDBAtlasProjectSpec struct {
	// OrgID is the MongoDB Atlas Organization ID. More info: https://docs.atlas.mongodb.com/organizations-projects/#organizations
	OrgID string `json:"orgID"`
}

// MongoDBAtlasProjectStatus defines the observed state of MongoDBAtlasProject
type MongoDBAtlasProjectStatus struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	OrgID        string `json:"orgID"`
	Created      string `json:"created"`
	ClusterCount int    `json:"clusterCount"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// +kubebuilder:printcolumn:name="GroupID",type="string",JSONPath=".status.id",description="The ID of the Project"
// +kubebuilder:printcolumn:name="ClusterCount",type="integer",JSONPath=".status.clusterCount",description="The number of Clusters in the Project"
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:printcolumn:name="OrgID",type="string",JSONPath=".spec.orgID",description="The MongoDB Atlas Organization ID",priority=1
// +kubebuilder:resource:path=mongodbatlasprojects,shortName=map
// +kubebuilder:categories=all,mongodbatlas
// MongoDBAtlasProject is the Schema for the mongodbatlasprojects API
type MongoDBAtlasProject struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MongoDBAtlasProjectSpec   `json:"spec,omitempty"`
	Status MongoDBAtlasProjectStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MongoDBAtlasProjectList contains a list of MongoDBAtlasProject
type MongoDBAtlasProjectList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MongoDBAtlasProject `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MongoDBAtlasProject{}, &MongoDBAtlasProjectList{})
}
