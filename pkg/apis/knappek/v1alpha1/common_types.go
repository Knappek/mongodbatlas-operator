package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
)

// AuthMongoDBAtlas defines the authentication struct for MongoDB Atlas API
type AuthMongoDBAtlas struct {
	PublicKey  string     `json:"publicKey"`
	PrivateKey PrivateKey `json:"privateKey"`
}

// PrivateKey defines the MongoDBAtlas Programmatic API Key reference
type PrivateKey struct {
	ValueFrom *PrivateKeySource `json:"valueFrom"`
}

// PrivateKeySource defines the MongoDBAtlas Programmatic API Key reference Kubernetes source
type PrivateKeySource struct {
	// Selects a key of a secret in the CR's namespace
	SecretKeyRef *corev1.SecretKeySelector `json:"secretKeyRef"`
}
