package config

import (
	"net/http"
	"os"

	dac "github.com/akshaykarle/go-http-digest-auth-client"
	ma "github.com/akshaykarle/go-mongodbatlas/mongodbatlas"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

type Config struct {
	AtlasUsername string
	AtlasAPIKey   string
}

// NewMongoDBAtlasClient returns a REST API client for MongoDB Atlas
func (c *Config) NewMongoDBAtlasClient() *ma.Client {
	t := dac.NewTransport(c.AtlasUsername, c.AtlasAPIKey)
	httpClient := &http.Client{Transport: &t}
	client := ma.NewClient(httpClient)
	return client
}

// GetKubernetesClient returns a Kubernetes Clientset in order to interact with Kubernetes resources
func GetKubernetesClient() (*kubernetes.Clientset, error) {
	var config *rest.Config
	var err error
	kubeconfig := os.Getenv("KUBECONFIG")
	if kubeconfig == "" {
		// creates the in-cluster config
		config, err = rest.InClusterConfig()
	} else {
		// creates out-of-cluster config
		config, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
	}
	if err != nil {
		panic(err.Error())
	}
	// creates the clientset
	return kubernetes.NewForConfig(config)
}
