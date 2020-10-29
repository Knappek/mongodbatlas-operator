package util

import (
	"net/http"
	"os"
	"strconv"
	"time"

	dac "github.com/akshaykarle/go-http-digest-auth-client"
	ma "github.com/akshaykarle/go-mongodbatlas/mongodbatlas"
)

// AtlasConfig stores Programmatic API Keys for authentication to Atlas API
type AtlasConfig struct {
	AtlasPublicKey  string
	AtlasPrivateKey string
}

// NewMongoDBAtlasClient returns a REST API client for MongoDB Atlas
func (c *AtlasConfig) newMongoDBAtlasClient() *ma.Client {
	t := dac.NewTransport(c.AtlasPublicKey, c.AtlasPrivateKey)
	httpClient := &http.Client{Transport: &t}
	client := ma.NewClient(httpClient)
	return client
}

// GetAtlasClient returns a MongoDB Atlas client
func GetAtlasClient() *ma.Client {
	// create MongoDB Atlas client
	privateKey, ok := os.LookupEnv("ATLAS_PRIVATE_KEY")
	if ok != true {
		panic("Error fetching private key: Env variable ATLAS_PRIVATE_KEY not set.")
	}
	publicKey, ok := os.LookupEnv("ATLAS_PUBLIC_KEY")
	if ok != true {
		panic("Error fetching public key: Env variable ATLAS_PUBLIC_KEY not set.")
	}
	atlasConfig := AtlasConfig{
		AtlasPublicKey:  publicKey,
		AtlasPrivateKey: privateKey,
	}
	return atlasConfig.newMongoDBAtlasClient()
}

// ReconciliationConfig let us customize reconcilitation
type ReconciliationConfig struct {
	Time time.Duration
}

// GetReconcilitationConfig gives us default values
func GetReconcilitationConfig() *ReconciliationConfig {
	// default reconciliation loop time is 2 minutes
	timeString := getenv("RECONCILIATION_TIME", "120")
	timeInt, _ := strconv.Atoi(timeString)
	reconciliationTime := time.Second * time.Duration(timeInt)
	return &ReconciliationConfig{
		Time: reconciliationTime,
	}
}

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}
