package v1alpha1

import (
	"github.com/Knappek/mongodbatlas-operator/pkg/util"
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

// IsMongoDBAtlasAlertConfigurationToBeUpdated is used to compare spec.MongoDBAtlasDatabaseUserRequestBody with status.MongoDBAtlasDatabaseUserRequestBody
func IsMongoDBAtlasAlertConfigurationToBeUpdated(m1 MongoDBAtlasAlertConfigurationRequestBody, m2 MongoDBAtlasAlertConfigurationRequestBody) bool {
	if ok := util.IsNotEqual(m1.EventTypeName, m2.EventTypeName); ok {
		return true
	}
	if ok := util.IsNotEqual(m1.Enabled, m2.Enabled); ok {
		return true
	}
	if ok := util.IsNotEqual(m1.MetricThreshold.MetricName, m2.MetricThreshold.MetricName); ok {
		return true
	}
	if ok := util.IsNotEqual(m1.MetricThreshold.Operator, m2.MetricThreshold.Operator); ok {
		return true
	}
	if ok := util.IsNotEqual(m1.MetricThreshold.Threshold, m2.MetricThreshold.Threshold); ok {
		return true
	}
	if ok := util.IsNotEqual(m1.MetricThreshold.Units, m2.MetricThreshold.Units); ok {
		return true
	}
	if ok := util.IsNotEqual(m1.MetricThreshold.Mode, m2.MetricThreshold.Mode); ok {
		return true
	}
	for idx, notification := range m1.Notifications {
		if ok := util.IsNotEqual(notification.TypeName, m2.Notifications[idx].TypeName); ok {
			return true
		}
		if ok := util.IsNotEqual(notification.IntervalMin, m2.Notifications[idx].IntervalMin); ok {
			return true
		}
		if ok := util.IsNotEqual(notification.DelayMin, m2.Notifications[idx].DelayMin); ok {
			return true
		}
		if ok := util.IsNotEqual(notification.EmailEnabled, m2.Notifications[idx].EmailEnabled); ok {
			return true
		}
		if ok := util.IsNotEqual(notification.SMSEnabled, m2.Notifications[idx].SMSEnabled); ok {
			return true
		}
		if ok := util.IsNotEqual(notification.Username, m2.Notifications[idx].Username); ok {
			return true
		}
		if ok := util.IsNotEqual(notification.TeamID, m2.Notifications[idx].TeamID); ok {
			return true
		}
		if ok := util.IsNotEqual(notification.EmailAddress, m2.Notifications[idx].EmailAddress); ok {
			return true
		}
		if ok := util.IsNotEqual(notification.MobileNumber, m2.Notifications[idx].MobileNumber); ok {
			return true
		}
		if ok := util.IsNotEqual(notification.NotificationToken, m2.Notifications[idx].NotificationToken); ok {
			return true
		}
		if ok := util.IsNotEqual(notification.RoomName, m2.Notifications[idx].RoomName); ok {
			return true
		}
		if ok := util.IsNotEqual(notification.ChannelName, m2.Notifications[idx].ChannelName); ok {
			return true
		}
		if ok := util.IsNotEqual(notification.APIToken, m2.Notifications[idx].APIToken); ok {
			return true
		}
		if ok := util.IsNotEqual(notification.OrgName, m2.Notifications[idx].OrgName); ok {
			return true
		}
		if ok := util.IsNotEqual(notification.FlowName, m2.Notifications[idx].FlowName); ok {
			return true
		}
		if ok := util.IsNotEqual(notification.FlowdockAPIToken, m2.Notifications[idx].FlowdockAPIToken); ok {
			return true
		}
		if ok := util.IsNotEqual(notification.ServiceKey, m2.Notifications[idx].ServiceKey); ok {
			return true
		}
		if ok := util.IsNotEqual(notification.VictorOpsAPIKey, m2.Notifications[idx].VictorOpsAPIKey); ok {
			return true
		}
		if ok := util.IsNotEqual(notification.VictorOpsRoutingKey, m2.Notifications[idx].VictorOpsRoutingKey); ok {
			return true
		}
		if ok := util.IsNotEqual(notification.OpsGenieAPIKey, m2.Notifications[idx].OpsGenieAPIKey); ok {
			return true
		}
	}
	for idx, matcher := range m1.Matchers {
		if ok := util.IsNotEqual(matcher.FieldName, m2.Matchers[idx].FieldName); ok {
			return true
		}
		if ok := util.IsNotEqual(matcher.Operator, m2.Matchers[idx].Operator); ok {
			return true
		}
		if ok := util.IsNotEqual(matcher.Value, m2.Matchers[idx].Value); ok {
			return true
		}
	}
	return false
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MongoDBAtlasAlertConfiguration is the Schema for the mongodbatlasalertconfigurations API
// +k8s:openapi-gen=true
// +kubebuilder:printcolumn:name="ID",type="string",JSONPath=".status.id",description="The ID of the Alert Configuration"
// +kubebuilder:printcolumn:name="Project Name",type="string",JSONPath=".spec.projectName",description="The MongoDB Atlas Project to which the Alert Configuration is applied"
// +kubebuilder:printcolumn:name="Enabled",type="string",JSONPath=".status.enabled",description="Whether the Alert Configuration is enabled or disabled"
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
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
