package util

import (
	"fmt"

	knappekv1alpha1 "github.com/Knappek/mongodbatlas-operator/pkg/apis/knappek/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// GetAPIKey returns the API Key based on a reference given in the yaml file
// currently only reference to a Kubernetes secret possible, but this function can be extended to any kind of reference
// without the need to change anything in the operator code
func GetAPIKey(c *kubernetes.Clientset, cr *knappekv1alpha1.MongoDBAtlasProject) (string, error) {
	var apiKeyString string
	apiKey := cr.Spec.APIKey
	if apiKey.ValueFrom != nil && apiKey.ValueFrom.SecretKeyRef != nil {
		sec := apiKey.ValueFrom.SecretKeyRef
		apiKey, err := c.CoreV1().Secrets(cr.Namespace).Get(sec.Name, metav1.GetOptions{})
		if err != nil {
			return "", fmt.Errorf("Error fetching APIKey secret %v: %s", sec.Name, err)
		}
		apiKeyString = string(apiKey.Data[sec.Key])
	}
	return apiKeyString, nil
}

// GetOrgID returns the MongoDB Atlas OrgID/groupID based on a reference given in the yaml file
// currently only reference to a Kubernetes secret possible, but this function can be extended to any kind of reference
// without the need to change anything in the operator code
func GetOrgID(c *kubernetes.Clientset, cr *knappekv1alpha1.MongoDBAtlasProject) (string, error) {
	var orgIDString string
	orgID := cr.Spec.OrgID
	if orgID.ValueFrom != nil {
		if orgID.ValueFrom.SecretKeyRef != nil {
			sec := orgID.ValueFrom.SecretKeyRef
			orgID, err := c.CoreV1().Secrets(cr.Namespace).Get(sec.Name, metav1.GetOptions{})
			if err != nil {
				return "", fmt.Errorf("Error fetching APIKey secret %v: %s", sec.Name, err)
			}
			orgIDString = string(orgID.Data[sec.Key])
		}
		if orgID.ValueFrom.ConfigMapKeyRef != nil {
			conf := orgID.ValueFrom.ConfigMapKeyRef
			orgID, err := c.CoreV1().ConfigMaps(cr.Namespace).Get(conf.Name, metav1.GetOptions{})
			if err != nil {
				return "", fmt.Errorf("Error fetching APIKey secret %v: %s", conf.Name, err)
			}
			orgIDString = string(orgID.Data[conf.Key])
		}
	}
	return orgIDString, nil
}
