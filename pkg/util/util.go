package util

import (
	"fmt"

	knappekv1alpha1 "github.com/Knappek/mongodbatlas-operator/pkg/apis/knappek/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// GetPrivateKey returns the Private Key of a Programmatic API Key based on a reference given in the yaml file.
// currently only reference to a Kubernetes secret possible, but this function can be extended to any kind of reference
// without the need to change anything in the operator code
func GetPrivateKey(c *kubernetes.Clientset, privateKey knappekv1alpha1.PrivateKey, namespace string) (string, error) {
	var privateKeyString string
	if privateKey.ValueFrom != nil && privateKey.ValueFrom.SecretKeyRef != nil {
		sec := privateKey.ValueFrom.SecretKeyRef
		privateKey, err := c.CoreV1().Secrets(namespace).Get(sec.Name, metav1.GetOptions{})
		if err != nil {
			return "", fmt.Errorf("Error fetching PrivateKey %v: %s", sec.Name, err)
		}
		privateKeyString = string(privateKey.Data[sec.Key])
	}
	return privateKeyString, nil
}
